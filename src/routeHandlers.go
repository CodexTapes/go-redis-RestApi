package main

import (
	"fmt"
	"encoding/json"

	mux "github.com/gorilla/mux"
)

func setHeaders() {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}


func (a *Artist) getAllArtistsConcerts(w http.ResponseWriter, req *http.Request) {

}

func (a *Artist) getArtistConcerts(w http.ResponseWriter, req *http.Request) {
	
}

func (a *Artist) createArtistConcert(w http.ResponseWriter, req *http.Request) {
	
}

func (a *Artist) updateArtistConcerts(w http.ResponseWriter, req *http.Request) {
	
}

func (a *Artist) deleteArtistConcerts(w http.ResponseWriter, req *http.Request) {
	
}
