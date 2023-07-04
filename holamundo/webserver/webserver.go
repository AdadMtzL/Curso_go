package webserver

import (
	"net/http"
)

func Webserver() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil) //No.Puerto
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}

