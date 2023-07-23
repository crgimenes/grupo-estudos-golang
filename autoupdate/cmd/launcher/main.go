package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"autoupdate/version"
)

const (
	downloadURL = "https://cloud.crg.eti.br/autoupdate"
)

var (
	cfg        = make(map[string]interface{})
	helperQuit = make(chan bool)
	mutex      sync.Mutex
	stdin      io.WriteCloser
	stdout     io.ReadCloser
	stderr     io.ReadCloser
)

func write(w io.Writer, msg string) {
	mutex.Lock()
	w.Write([]byte(msg))
	mutex.Unlock()
}

func loadCfg() error {
	d := filepath.Join(*localPath, "config.json")
	c, err := os.ReadFile(d)
	if err != nil {
		return nil
	}
	log.Println("config.json:", string(c))
	return json.Unmarshal(c, &cfg)
}

func parseVersion(version string) (int, error) {
	a := strings.Split(version, ".")
	if len(a) != 3 {
		return 0, errors.New("helper_version format error")
	}
	major, err := strconv.Atoi(a[0])
	if err != nil {
		return 0, err
	}
	minor, err := strconv.Atoi(a[1])
	if err != nil {
		return 0, err
	}
	revision, err := strconv.Atoi(a[2])
	if err != nil {
		return 0, err
	}
	aux := fmt.Sprintf("%03d%03d%04d", major, minor, revision)
	return strconv.Atoi(aux)
}

func helperVersion() (int, error) {
	v := "0.0.0"
	i, ok := cfg["helper_version"]
	if ok {
		v, ok = i.(string)
		if !ok {
			return 0, errors.New("helper_version format error")
		}
	}
	return parseVersion(v)
}

func serverHelperVersion() (int, string, error) {
	vs := &struct {
		HelperVersion string `json:"helper_version,omitempty"`
	}{}
	c := &http.Client{Timeout: 5 * time.Second}
	r, err := c.Get(downloadURL + "/version.json")
	if err != nil {
		return 0, "", err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(vs)
	if err != nil {
		return 0, "", err
	}
	i, err := parseVersion(vs.HelperVersion)
	return i, vs.HelperVersion, err
}

func pwd() string {
	d, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return d
}

func downloadNewVersion(dir string) error {
	f := filepath.Join(dir, version.HelperZipName)
	log.Println("donload:", f)
	o, err := os.Create(f)
	if err != nil {
		return err
	}
	defer o.Close()
	log.Println("download new version from:",
		downloadURL+"/"+
			version.HelperZipName)
	r, err := http.Get(downloadURL + "/" + version.HelperZipName)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", r.Status)
	}
	_, err = io.Copy(o, r.Body)
	return err
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fpath,
			filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
		if err != nil {
			return err
		}
		outFile, err := os.OpenFile(
			fpath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode())
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func execHelper() error {
	log.Println("starting execHelper")
	cmd := exec.Command(
		filepath.Join(*localPath, version.HelperName),
		"-path",
		*localPath,
	)
	stderr, err := cmd.StderrPipe()
	if nil != err {
		log.Println("error obtaining stderr:", err.Error())
		return err
	}
	stdin, err = cmd.StdinPipe()
	if nil != err {
		log.Println("error obtaining stdin:", err.Error())
		return err
	}
	stdout, err = cmd.StdoutPipe()
	if nil != err {
		log.Println("error obtaining stdout:", err.Error())
		return err
	}
	reader := bufio.NewReader(stdout)
	readerError := bufio.NewReader(stderr)

	go func() {
		scanner := bufio.NewScanner(readerError)
		for scanner.Scan() {
			s := scanner.Text()
			log.Println("helper stderr:", s)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			s := scanner.Text()
			log.Println("helper stdout:", s)
			switch s {
			case "ready":
				write(stdin, "ok\n")
			case "ping":
				write(stdin, "pong\n")
			}
		}
	}()

	/*
		go func() {
			<-time.After(5 * time.Second)
			write(stdin, "quit\n")
		}()
	*/
	err = cmd.Start()
	if err != nil {
		e := fmt.Errorf(
			"error starting program: %s, %s",
			cmd.Path,
			err.Error())
		helperQuit <- true
		log.Println(e)
		return e
	}
	cmd.Wait()
	log.Println("helper exited sending helperQuit message")
	helperQuit <- true
	log.Println("exiting execHelper")
	return nil
}

var localPath *string

func main() {

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc
		fmt.Printf("\r\nfreeing up resources...\r\n")
		if stdin != nil {
			write(stdin, "quit\n")
			<-helperQuit
		}
		fmt.Printf("have a nice day!\r\n")
		os.Exit(0)
	}()

	localPath = flag.String("path", pwd(), "path to config.json")
	flag.Parse()

	log.Println("using path:", *localPath)

	err := loadCfg()
	if err != nil {
		log.Println("error load config", err)
		return
	}
	bt, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	waitUpdate := false
	// chk if there is a new version
	go func() {
		for {
			if waitUpdate {
				<-time.After(10 * time.Second)
			}
			waitUpdate = true
			hv, err := helperVersion()
			if err != nil {
				log.Println("error parssing helper version ", err)
				continue
			}
			hsv, hvs, err := serverHelperVersion()
			if err != nil {
				log.Println("error parssing server helper version", err)
				continue
			}
			if hsv > hv { // server version is gt local version
				log.Println("new version avaliable")
				d := *localPath
				err = downloadNewVersion(d)
				if err != nil {
					log.Println("error downloading helper", err)
					continue
				}
				helperZipFile := filepath.Join(d, version.HelperZipName)
				d = filepath.Join(d, "tmp_helper")
				err = unzip(helperZipFile, d)
				if err != nil {
					log.Println("error unzip:", err)
					continue
				}
				err = os.Remove(helperZipFile)
				if err != nil {
					log.Println("error remove:", helperZipFile, err)
					continue
				}
				log.Println("A new version is available for upgrade.")
				write(stdin, "quit\n")
				log.Println("wait for helper to quit")
				<-helperQuit // wait for helper to quit
				err = moveNewVersion()
				if err != nil {
					log.Println("error moving new version", err)
					continue
				}
				err = os.Remove(d)
				if err != nil {
					log.Println("error removing tmp_helper", err)
				}
				//write json to file
				log.Println("updating config.json")
				cfg["helper_version"] = hvs
				d = *localPath
				d = filepath.Join(d, "config.json")
				bt, err = json.MarshalIndent(cfg, "", "\t")
				if err != nil {
					log.Println("error MarshalIdent", err)
					continue
				}
				log.Println(string(bt))
				err = os.WriteFile(d, bt, 0644)
				if err != nil {
					log.Println("error write to config.file", err)
					continue
				}
				log.Println("config.json updated")

				go func() {
					err := execHelper()
					if err != nil {
						log.Println("error starting helper", err)
					}
				}()
				fmt.Println("upgrade successfully.")
			}
		}
	}()

	err = execHelper()
	if err != nil {
		log.Println("error starting helper", err)
	}
	c := make(chan bool)
	<-c
}

func moveNewVersion() error {
	dir := *localPath
	log.Println("remove from", dir)
	os.Remove(filepath.Join(dir, version.HelperName))

	o := filepath.Join(dir, "tmp_helper", version.HelperName)
	d := filepath.Join(dir, version.HelperName)
	log.Println("move", o, d)
	os.Rename(o, d)

	return nil
}
