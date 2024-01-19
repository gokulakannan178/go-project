package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WeatherAlertNotInRangeRoutes(r *mux.Router) {
	r.Handle("/weatheralertnotinrange", Adapt(http.HandlerFunc(route.Handler.SaveWeatherAlertNotInRange))).Methods("POST")
	r.Handle("/weatheralertnotinrange", Adapt(http.HandlerFunc(route.Handler.GetSingleWeatherAlertNotInRange))).Methods("GET")
	r.Handle("/weatheralertnotinrange", Adapt(http.HandlerFunc(route.Handler.UpdateWeatherAlertNotInRange))).Methods("PUT")
	r.Handle("/weatheralertnotinrange/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWeatherAlertNotInRange))).Methods("PUT")
	r.Handle("/weatheralertnotinrange/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWeatherAlertNotInRange))).Methods("PUT")
	r.Handle("/weatheralertnotinrange/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWeatherAlertNotInRange))).Methods("DELETE")
	r.Handle("/weatheralertnotinrange/filter", Adapt(http.HandlerFunc(route.Handler.FilterWeatherAlertNotInRange))).Methods("POST")
}
