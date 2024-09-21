package main

import (
	"log"
	"net/http"
	"okapi-api/handler"
)

func main() {
	http.HandleFunc("/containers", handler.GetDockerContainers)

	log.Println("API is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
