package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WeatherparameterRoutes(r *mux.Router) {
	r.Handle("/weatherparameter", Adapt(http.HandlerFunc(route.Handler.SaveWeatherParameter))).Methods("POST")
	r.Handle("/weatherparameter", Adapt(http.HandlerFunc(route.Handler.GetSingleWeatherParameter))).Methods("GET")
	r.Handle("/weatherparameter", Adapt(http.HandlerFunc(route.Handler.UpdateWeatherParameter))).Methods("PUT")
	r.Handle("/weatherparameter/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWeatherParameter))).Methods("PUT")
	r.Handle("/weatherparameter/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWeatherParameter))).Methods("PUT")
	r.Handle("/weatherparameter/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWeatherParameter))).Methods("DELETE")
	r.Handle("/weatherparameter/filter", Adapt(http.HandlerFunc(route.Handler.FilterWeatherParameter))).Methods("POST")
}
