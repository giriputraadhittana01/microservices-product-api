package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHelloHandler(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) GetAllDataHello(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oops"))
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "%s", data)
}
