package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/munjeli/testAPI/service"
)

func CreateCount(w http.ResponseWriter, r *http.Request) {
	if err := validate(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	c, err := service.CountWordsInText(string(b))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func validate(r *http.Request) error {
	l := r.ContentLength
	if l == -1 {
		return fmt.Errorf("cannot read Content-Length for request")
	}
	mb := (l / 1024) / 1024
	if mb > 2 {
		return fmt.Errorf("only 2 MB allowed: request body is %v mb", mb)
	}
	ct := r.Header.Get("Content-Type")
	if ct != "text/plain" {
		return fmt.Errorf("the Content-Type header must be text/plain - received: %v", ct)
	}
	return nil
}
