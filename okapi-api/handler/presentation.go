package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"okapi-api/service"

	"github.com/gorilla/mux"
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

func ExecPythonInContainer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	containerId := mux.Vars(r)["id"]

	var cmd struct {
		Command string `json:"command"`
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request method"})
		return
	}

	output, err := service.ExecPythonInContainer(containerId, cmd.Command)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to execute Python script in container"})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
	json.NewEncoder(w).Encode(map[string]string{"message": "Python script executed successfully"})
}
