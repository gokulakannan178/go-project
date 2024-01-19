package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DistrictWeatherAlertRoutes(r *mux.Router) {
	r.Handle("/districtweatheralert", Adapt(http.HandlerFunc(route.Handler.SaveDistrictWeatherAlert))).Methods("POST")
	r.Handle("/districtweatheralert", Adapt(http.HandlerFunc(route.Handler.GetSingleDistrictWeatherAlert))).Methods("GET")
	r.Handle("/districtweatheralert", Adapt(http.HandlerFunc(route.Handler.UpdateDistrictWeatherAlert))).Methods("PUT")
	r.Handle("/districtweatheralert/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDistrictWeatherAlert))).Methods("PUT")
	r.Handle("/districtweatheralert/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDistrictWeatherAlert))).Methods("PUT")
	r.Handle("/districtweatheralert/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDistrictWeatherAlert))).Methods("DELETE")
	r.Handle("/districtweatheralert/filter", Adapt(http.HandlerFunc(route.Handler.FilterDistrictWeatherAlert))).Methods("POST")
	r.Handle("/districtweatheralert/updateweatheralertmaster", Adapt(http.HandlerFunc(route.Handler.UpdateDistrictWeatherAlertMasterwithWeatheralert))).Methods("PUT")
}
