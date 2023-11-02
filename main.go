package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p4ndabk/ibge-integration/ibge"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/api/city/coordinates/{ibge_code}", ibge.CheckCoordinates).Methods("GET")

	fmt.Println("Server is running on port 4001")
	log.Fatal(http.ListenAndServe(":4001", router))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
