package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cardil/knative-homedir/internal/service/fs"
)

// List is a controller that list files as JSON.
func List(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	listing, err := fs.List("~")
	if err != nil {
		handleErr(err, rw)
		return
	}
	b, err := json.Marshal(listing)
	if err != nil {
		handleErr(err, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(b)
	if err != nil {
		handleErr(err, rw)
		return
	}
}

func handleErr(err error, rw http.ResponseWriter) {
	log.Printf("ERROR: %v", err)
	rw.WriteHeader(http.StatusInternalServerError)
}
