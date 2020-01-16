package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a http.Handler that respond to request to the / endpoint
type Hello struct {
	l *log.Logger
}

// NewHello returns a pointer to handlers.Hello
func NewHello(l *log.Logger) *Hello {
	return &Hello{l: l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s\n", d)
}
