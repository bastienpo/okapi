package main

import (
	"log"
	"net/http"
	"okapi-api/router"
)

func main() {
	r := router.InitRoutes()
    port := ":8080"

	log.Printf("API is running on port %s", port)
    if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("API fail to start : %v", err)
	}
}
