package handlers

import (
	"fmt"
	"encoding/json"

	mux "github.com/gorilla/mux"
	"go-redis-restApi/src/models"
)

func setHeaders(requestStatus) {
	w.Header().Set("Content-Type", "application/json")
	// pass in status code as an argument
	w.WriteHeader(http.StatusOK)
}

/***************************							
*	RestFul Route Handlers	*
 ***************************/							

 //  Retrieve all concert events for all artists in RedisDB  - /artists
func (a *Artist) GetAllArtistsConcerts(w http.ResponseWriter, req *http.Request, _ mux.params) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	artists := allArtists()

	if err := json.NewEncoder(w).Encode(artists); err != nil {
		panic(err)
	}

}

// Retrieve concert event for artist with artistname - /artist/{artistname}
func (a *Artist) GetArtistConcerts(w http.ResponseWriter, req *http.Request, reqParams mux.Params) {
	
	artist, err := strconv.Atoi(reqParams.ByName("artistName")); if err != nil {
		panic(err)
	}

	artistConcert := findArtistConcert(artist)

	if err := json.NewEncoder(w).Encode(artistConcert); err != nil {
		panic(err)
	}
}


// Create a concert event for an artist- /artist
func (a *Artist) CreateArtistConcert(w http.ResponseWriter, req *http.Request, _ mux.Params) {
	
	var artist Artist

	// Read request body
	requestBody, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	// Convert request body json to Artist struct
	if err := json.Unmarshall(requestBody, &artist); err != nil {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	    w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//  Send artist struct - converted/unmarshalled from json in the request body to artist struct to Redis
	func createConcert(artist) 
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// Update concert event for artist - /artist/{artistname}
func (a *Artist) UpdateArtistConcerts(w http.ResponseWriter, req *http.Request) {
	
}

// Delete concert event for artist - /artist/{artistname}
func (a *Artist) DeleteArtistConcerts(w http.ResponseWriter, req *http.Request) {
	
}
