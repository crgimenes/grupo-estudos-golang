package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"session_example/session"
	"time"
)

var (
	sc *session.Control

	//go:embed assets
	assets embed.FS
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	sid, sd, ok := sc.Get(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// renew session
	sc.Save(w, sid, sd)

	data := struct {
		ID int
	}{
		ID: sd.UserID,
	}
	page := "assets/home.html"

	err := printTemplate(page, r, w, data)
	if err != nil {
		log.Println(err)
	}
}

func printTemplate(page string, r *http.Request, w http.ResponseWriter, data any) error {
	tpl, err := template.ParseFS(assets, page, "assets/head.html", "assets/foot.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	return tpl.Execute(w, data)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	page := "assets/error.html"

	sid, _, ok := sc.Get(r)
	if ok {
		sc.Delete(w, sid)
	}

	err := printTemplate(page, r, w, data)
	if err != nil {
		log.Println(err)
	}

	return
}

func main() {
	sc = session.New()

	go func() {
		for {
			time.Sleep(5 * time.Minute)
			sc.RemoveExpired()
		}
	}()

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/", homeHandler)

	fs := http.FileServer(http.FS(assets))

	http.Handle("/assets/", fs)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
