package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p4ndabk/ibge-integration/internal/controller"
	"github.com/p4ndabk/ibge-integration/internal/middleware"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/api/health", middleware.SetContentTypeMiddleware(http.HandlerFunc(controller.HealthCheckHandlerRequest))).Methods("GET")
	router.HandleFunc("/api/city/coordinates/{ibge_code}", controller.CheckCoordinateRequest).Methods("GET")

	router.Handle("/api/city", middleware.SetContentTypeMiddleware(http.HandlerFunc(controller.AllCityRequest))).Methods("GET")
	router.Handle("/api/city/{city_id}", middleware.SetContentTypeMiddleware(http.HandlerFunc(controller.CityRequest))).Methods("GET")

	router.Handle("/api/solar-efficiencie/{city_id}", middleware.SetContentTypeMiddleware(http.HandlerFunc(controller.SolarEfficiencieByCodeRequest))).Methods("GET")

	fmt.Println("Server is running on port 4001")
	log.Fatal(http.ListenAndServe(":4001", router))
}
