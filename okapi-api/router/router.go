package router

import (
	"okapi-api/handler"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/containers", handler.GetDockerContainers).Methods("GET")
	router.HandleFunc("/containers/{id}/exec", handler.ExecPythonInContainer).Methods("POST")
	return router
}
