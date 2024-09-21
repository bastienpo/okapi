package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"okapi-api/service"
)

func GetDockerContainers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request method"})
		return
	}

	containers, err := service.GetRunningDockerContainers()
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to fetch Docker containers"})
		return
	}

	json.NewEncoder(w).Encode(containers)
}
