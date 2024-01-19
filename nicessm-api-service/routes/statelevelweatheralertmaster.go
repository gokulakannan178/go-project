package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) StateWeatherAlertMasterRoutes(r *mux.Router) {
	r.Handle("/stateweatheralertmaster", Adapt(http.HandlerFunc(route.Handler.SaveStateWeatherAlertMaster))).Methods("POST")
	r.Handle("/stateweatheralertmaster", Adapt(http.HandlerFunc(route.Handler.GetSingleStateWeatherAlertMaster))).Methods("GET")
	r.Handle("/stateweatheralertmaster", Adapt(http.HandlerFunc(route.Handler.UpdateStateWeatherAlertMaster))).Methods("PUT")
	r.Handle("/updatestateweatheralertmaster/min", Adapt(http.HandlerFunc(route.Handler.UpdateStateWeatherAlertMasterUpsertwithMin))).Methods("PUT")
	r.Handle("/updatestateweatheralertmaster/max", Adapt(http.HandlerFunc(route.Handler.UpdateStateWeatherAlertMasterUpsertwithMax))).Methods("PUT")
	r.Handle("/stateweatheralertmaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableStateWeatherAlertMaster))).Methods("PUT")
	r.Handle("/stateweatheralertmaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableStateWeatherAlertMaster))).Methods("PUT")
	r.Handle("/stateweatheralertmaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteStateWeatherAlertMaster))).Methods("DELETE")
	r.Handle("/stateweatheralertmaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterStateWeatherAlertMaster))).Methods("POST")
	r.Handle("/stateweatheralertmaster/block/weatheralert", Adapt(http.HandlerFunc(route.Handler.GetStateWeatherAlertMaster))).Methods("POST")
}
