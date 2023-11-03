package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/nuveo/log"
)

func closer(body io.Closer) {
	err := body.Close()
	if err != nil {
		log.Errorln(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorln(err)
	}
	defer closer(r.Body)
	fmt.Printf("client enviou: %q\n", string(b))
	w.Write([]byte("ok\n"))
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
