package main

import (
	"time"
)

type Artist struct {
	Id			int			`json:"id"`				`redis:"id"`
	ArtistName	string		`json:"artistName"`		`redis:"artistName"`
	Concert		Concert 	`json:"concert"` 		`redis:"concert"`
}

type Concert struct {
	Artist Artist			`json:"artist"`			`redis:"artist"`
	TourName	string		`json:"tourName"`		`redis:"tourName"`
	Description	string		`json:"description"`	`redis:"description"`
	Venue		string		`json:"venue"`			`redis:"venue"`
	EventDate	time.time	`json:"eventDate"`		`redis:"eventDate"`	
	City		string		`json:"city"`			`redis:"city"`
	Country		string		`json:"country"`		`redis:"country"`

}

var Artists []Artist
var Concerts []Concert