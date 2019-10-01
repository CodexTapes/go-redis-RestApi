package routes

import (
	"net/http"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/artists",
		GetAllArtistsConcerts,
	},
	Route{
		"GET",
		"/artist/{artistname}",
		GetArtistConcerts,
	},
	Route{
		"POST",
		"/artist",
		CreateArtistConcerts,
	},
	Route{
		"PUT",
		"/artist/{artistname}",
		UpdateArtistConcerts,
	},
	Route{
		"DELETE",
		"/artist/{artistname}",
		DeleteArtistConcerts,
	},
}
