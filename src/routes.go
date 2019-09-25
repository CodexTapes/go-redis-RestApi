package main

import (
	mux "github.com/gorilla/mux"
)

type Route struct {
	Method			string
	Path			string
	HandlerFunc 	mux.HandlerFunc
}

type Routes	[]Route

var routes = Routes {
	Route{
		"GET",
		"/artists",
		getAllArtistsConcerts,
	},
	Route{
		"GET",
		"/artists/{artistname}",
		getArtistConcerts
	},
	Route{
		"POST",
		"/artists",
		createArtistConcerts
	},
	Route{
		"PUT",
		"/artists/{artistname}",
		updateArtistConcerts
	},
	Route{
		"DELETE",
		"/artists/{artistname}",
		deleteArtistConcerts
	},
}