package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p4ndabk/ibge-integration/internal/ibge"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", HealthCheckHandlerRequest).Methods("GET")
	router.HandleFunc("/api/city/coordinates/{ibge_code}", ibge.CheckCoordinateRequest).Methods("GET")

	router.HandleFunc("/api/city", ibge.CityIndexRequest).Methods("GET")
	router.HandleFunc("/api/city/{city_id}", ibge.CityShowRequest).Methods("GET")

	router.HandleFunc("/api/solar_efficiencie/{city_id}", ibge.SolarEfficiencieByCodeRequest).Methods("GET")

	fmt.Println("Server is running on port 4001")
	log.Fatal(http.ListenAndServe(":4001", router))
}

func HealthCheckHandlerRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
