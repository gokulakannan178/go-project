package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DistrictWeatherAlertMasterRoutes(r *mux.Router) {
	r.Handle("/districtweatheralertmaster", Adapt(http.HandlerFunc(route.Handler.SaveDistrictWeatherAlertMaster))).Methods("POST")
	r.Handle("/districtweatheralertmaster", Adapt(http.HandlerFunc(route.Handler.GetSingleDistrictWeatherAlertMaster))).Methods("GET")
	r.Handle("/districtweatheralertmaster", Adapt(http.HandlerFunc(route.Handler.UpdateDistrictWeatherAlertMaster))).Methods("PUT")
	r.Handle("/updatedistrictweatheralertmaster/min", Adapt(http.HandlerFunc(route.Handler.UpdateDistrictWeatherAlertMasterUpsertwithMin))).Methods("PUT")
	r.Handle("/updatedistrictweatheralertmaster/max", Adapt(http.HandlerFunc(route.Handler.UpdateDistrictWeatherAlertMasterUpsertwithMax))).Methods("PUT")
	r.Handle("/districtweatheralertmaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDistrictWeatherAlertMaster))).Methods("PUT")
	r.Handle("/districtweatheralertmaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDistrictWeatherAlertMaster))).Methods("PUT")
	r.Handle("/districtweatheralertmaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDistrictWeatherAlertMaster))).Methods("DELETE")
	r.Handle("/districtweatheralertmaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterDistrictWeatherAlertMaster))).Methods("POST")
	r.Handle("/districtweatheralertmaster/block/weatheralert", Adapt(http.HandlerFunc(route.Handler.GetDistrictWeatherAlertMaster))).Methods("POST")
}
