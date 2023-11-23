package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p4ndabk/ibge-integration/internal/controller"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", controller.HealthCheckHandlerRequest).Methods("GET")
	router.HandleFunc("/api/city/coordinates/{ibge_code}", controller.CheckCoordinateRequest).Methods("GET")

	router.HandleFunc("/api/city", controller.AllCityRequest).Methods("GET")
	router.HandleFunc("/api/city/{city_id}", controller.CityRequest).Methods("GET")

	router.HandleFunc("/api/solar-efficiencie/{city_id}", controller.SolarEfficiencieByCodeRequest).Methods("GET")

	fmt.Println("Server is running on port 4001")
	log.Fatal(http.ListenAndServe(":4001", router))
}
