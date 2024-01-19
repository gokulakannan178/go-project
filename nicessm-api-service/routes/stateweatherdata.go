package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) StateWeatherDataRoutes(r *mux.Router) {
	r.Handle("/stateWeatherData", Adapt(http.HandlerFunc(route.Handler.SaveStateWeatherData))).Methods("POST")
	r.Handle("/stateWeatherData", Adapt(http.HandlerFunc(route.Handler.GetSingleStateWeatherData))).Methods("GET")
	r.Handle("/stateWeatherData", Adapt(http.HandlerFunc(route.Handler.UpdateStateWeatherData))).Methods("PUT")
	r.Handle("/stateWeatherData/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableStateWeatherData))).Methods("PUT")
	r.Handle("/stateWeatherData/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableStateWeatherData))).Methods("PUT")
	r.Handle("/stateWeatherData/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteStateWeatherData))).Methods("DELETE")
	r.Handle("/stateWeatherData/filter", Adapt(http.HandlerFunc(route.Handler.FilterStateWeatherData))).Methods("POST")
	r.Handle("/stateWeatherData/state", Adapt(http.HandlerFunc(route.Handler.SaveStateWeatherDataWithOpenWebsite))).Methods("GET")
	r.Handle("/stateWeatherData/currentdate", Adapt(http.HandlerFunc(route.Handler.GetSingleStateWeatherDataWithCureentDate))).Methods("GET")
}
