package apis

import (
	"github.com/gorilla/mux"
)

func GetHTTPRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/simplify", simplifyPost).Methods("POST")
	router.HandleFunc("/simplify", simplifyGet).Methods("GET")

	return router
}
