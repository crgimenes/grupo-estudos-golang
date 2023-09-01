package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewFilesystemStore(
	os.TempDir(),
	[]byte("chave com 32 caracteres ........"))

const sessionName = "session-name"

func clearSession(w http.ResponseWriter, session string) {
	cookie := &http.Cookie{
		Name:   session,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		clearSession(w, sessionName)
		session.Options.MaxAge = -1
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["foo"] = "bar"
	session.Values[42] = "The answer to life, the universe and everything"

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a := fmt.Sprintf("42 is %q", session.Values[42])

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(a))
}

func getSessionHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		clearSession(w, sessionName)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a := fmt.Sprintf("42 is %q", session.Values[42])

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(a))
}

func getNonExistingValueHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		clearSession(w, sessionName)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a, ok := session.Values["non"]
	if !ok {

		http.Error(w, "value not set", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	switch a.(type) {
	case string:
		w.Write([]byte(a.(string)))
	default:
		http.Error(w, "value is not type string", http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/get", getSessionHandler)
	r.HandleFunc("/non", getNonExistingValueHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listen at port :8000")
	log.Println("Temp Dir:", os.TempDir())

	log.Fatal(srv.ListenAndServe())

}
