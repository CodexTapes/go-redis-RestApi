package main

import (
	"go-redis-restApi/src/router"
	"log"
	"net/http"
)

func main() {

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
