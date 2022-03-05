package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetAllDataHy(rw http.ResponseWriter, r *http.Request) {
	var l *log.Logger = log.New(os.Stdout, "hy-api", log.LstdFlags)
	l.Println("Hy")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oops"))
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "%s", data)
}
