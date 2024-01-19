package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DistrictWeatherAlertNotInRangeRoutes(r *mux.Router) {
	r.Handle("/districtweatheralertnotinrange", Adapt(http.HandlerFunc(route.Handler.SaveDistrictWeatherAlertNotInRange))).Methods("POST")
	r.Handle("/districtweatheralertnotinrange", Adapt(http.HandlerFunc(route.Handler.GetSingleDistrictWeatherAlertNotInRange))).Methods("GET")
	r.Handle("/districtweatheralertnotinrange", Adapt(http.HandlerFunc(route.Handler.UpdateDistrictWeatherAlertNotInRange))).Methods("PUT")
	r.Handle("/districtweatheralertnotinrange/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDistrictWeatherAlertNotInRange))).Methods("PUT")
	r.Handle("/districtweatheralertnotinrange/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDistrictWeatherAlertNotInRange))).Methods("PUT")
	r.Handle("/districtweatheralertnotinrange/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDistrictWeatherAlertNotInRange))).Methods("DELETE")
	r.Handle("/districtweatheralertnotinrange/filter", Adapt(http.HandlerFunc(route.Handler.FilterDistrictWeatherAlertNotInRange))).Methods("POST")
}
