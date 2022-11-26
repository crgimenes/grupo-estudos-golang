package main

import (
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet { // GET show login form
		data := map[string]interface{}{}
		page := "assets/login.html"

		err := printTemplate(page, r, w, data)
		if err != nil {
			log.Println(err)
		}

		return
	}

	if r.Method == http.MethodPost { // POST validate login form
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		email := r.FormValue("email")
		passwd := r.FormValue("password")

		// *FAKE* LOGIN
		if email == "user@example.com" && passwd == "1234" {
			userID := 1
			sessionData, sid := sc.Create(userID)
			// set session cookie
			sc.Save(w, sid, sessionData)

			http.Redirect(w, r, "/", http.StatusFound)

			return
		}

		data := map[string]interface{}{}
		page := "assets/login.html"

		err = printTemplate(page, r, w, data)
		if err != nil {
			log.Println(err)
		}

		return
	}

	http.Redirect(w, r, "/error", http.StatusFound)
}
