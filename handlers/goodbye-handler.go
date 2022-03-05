package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

// Depency Injection
func NewGoodByeHandler(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) GetAllDataGoodBye(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("GoodBye")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oops"))
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "%s", data)
}
