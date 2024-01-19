package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WeatherAlertTypeRoutes(r *mux.Router) {
	r.Handle("/weatheralerttype", Adapt(http.HandlerFunc(route.Handler.SaveWeatherAlertType))).Methods("POST")
	r.Handle("/weatheralerttype", Adapt(http.HandlerFunc(route.Handler.GetSingleWeatherAlertType))).Methods("GET")
	r.Handle("/weatheralerttype", Adapt(http.HandlerFunc(route.Handler.UpdateWeatherAlertType))).Methods("PUT")
	r.Handle("/weatheralerttype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWeatherAlertType))).Methods("PUT")
	r.Handle("/weatheralerttype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWeatherAlertType))).Methods("PUT")
	r.Handle("/weatheralerttype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWeatherAlertType))).Methods("DELETE")
	r.Handle("/weatheralerttype/filter", Adapt(http.HandlerFunc(route.Handler.FilterWeatherAlertType))).Methods("POST")
}
