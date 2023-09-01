package main

import "net/http"

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sid, _, ok := sc.Get(r)
	if ok {
		sc.Delete(w, sid)
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
