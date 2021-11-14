package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/avelino/slugify"
	"github.com/gosidekick/goconfig"
)

type config struct {
	InputFolder  string `json:"i" cfg:"i" cfgDefault:"./"`
	OutputFolder string `json:"o" cfg:"o"`
}

var (
	cfg = config{}
)

func execHelper(path, name string, arg ...string) (out []byte, err error) {
	var (
		outbuf bytes.Buffer
		errbuf bytes.Buffer
	)
	stdout := bufio.NewWriter(&outbuf)
	stderr := bufio.NewWriter(&errbuf)

	cmd := exec.Command(name, arg...)
	cmd.Dir = path
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err = cmd.Run()
	if err != nil {
		return
	}
	if errbuf.Len() > 0 {
		err = errors.New(errbuf.String())
		return
	}
	out = outbuf.Bytes()
	return
}

func fileExists(path string) (ret bool) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Println(err)
		return
	}
	ret = true
	return
}

func visit(path string, f os.FileInfo, perr error) error {
	if perr != nil {
		return perr
	}
	if !f.IsDir() {
		return nil
	}
	fe := fileExists(path + "/README.md")
	if !fe {
		return nil
	}

	/*
		if f.Name() == "vendor" {
			return filepath.SkipDir
		}
	*/
	pathAbs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	pathOutputAbs, err := filepath.Abs(cfg.OutputFolder)
	if err != nil {
		return err
	}

	filePath := filepath.Join(pathAbs, "/README.md")

	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if body[0] == '#' {
		s := strings.Split(string(body), "\n\n")
		name := s[0][2:]
		title := fmt.Sprintf("title = \"%s\"\n", name)

		outputFileName := slugify.Slugify(name)

		out, err := execHelper(
			path,
			"git",
			"log",
			"--reverse",
			"--format=%cI",
			"--name-only",
			"--diff-filter=A",
			filePath)
		if err != nil {
			return err
		}
		dateSplit := strings.Split(string(out), "\n")
		date := fmt.Sprintf("date = \"%s\"\n", dateSplit[0])

		description := fmt.Sprintf("description = \"%s\"\n", strings.TrimSpace(s[1]))
		tags := "tags = [\"golang\"]\n"
		metadata := "+++\n" + title + date + description + tags + "+++\n\n"

		metadata += string(body)

		// coloca arquivos no rodapé

		metadata += "\n### Arquivos desse post\n\n"
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, local := range files {
			name := fmt.Sprintf(
				"https://github.com/go-br/estudos/blob/master/exemplos/%s/%s",
				f.Name(),
				local.Name())
			metadata += fmt.Sprintf("- [%s/%s](%s)\n", f.Name(), local.Name(), name)

			fmt.Println(name)
		}

		outputFileName = filepath.Join(pathOutputAbs, outputFileName+".md")
		fmt.Println(outputFileName)

		err = ioutil.WriteFile(outputFileName, []byte(metadata), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := goconfig.Parse(&cfg)
	if err != nil {
		log.Println(err)
		return
	}

	if cfg.OutputFolder == "" {
		log.Println("error output folder not indicated")
		return
	}

	err = filepath.Walk(cfg.InputFolder, visit)
	if err == nil || err == io.EOF {
		return
	}
	log.Println(err)
}
