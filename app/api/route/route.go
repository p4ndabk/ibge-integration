package route

import (
	"github.com/gorilla/mux"
	"github.com/p4ndabk/ibge-integration/app/api/controller"
	"github.com/p4ndabk/ibge-integration/app/api/middleware"
	"net/http"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/api/health", middleware.SetContentTypeMiddleware(http.HandlerFunc(controller.HealthCheckHandlerRequest))).Methods("GET")
	router.Handle("/api/city/coordinates/{ibge_code}", middleware.SetContentTypeMiddleware(http.HandlerFunc(controller.CheckCoordinateRequest))).Methods("GET")

	router.Handle("/api/city", middleware.SetContentTypeMiddleware(http.HandlerFunc(controller.AllCityRequest))).Methods("GET")
	router.Handle("/api/city/{city_id}", middleware.SetContentTypeMiddleware(http.HandlerFunc(controller.CityRequest))).Methods("GET")

	router.Handle("/api/solar-efficiencie/{city_id}", middleware.SetContentTypeMiddleware(http.HandlerFunc(controller.SolarEfficiencieByCodeRequest))).Methods("GET")

	return router
}
