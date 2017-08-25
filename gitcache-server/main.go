package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &httpHandler{})
	http.ListenAndServe(":80", nil)
}

type httpHandler struct {
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("req:", req.URL)

	str := "req " + req.URL.Path
	w.Write([]byte(str))
}
