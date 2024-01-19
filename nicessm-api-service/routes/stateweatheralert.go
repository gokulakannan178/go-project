package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) StateWeatherAlertRoutes(r *mux.Router) {
	r.Handle("/weatheralert", Adapt(http.HandlerFunc(route.Handler.SaveStateWeatherAlert))).Methods("POST")
	r.Handle("/weatheralert", Adapt(http.HandlerFunc(route.Handler.GetSingleStateWeatherAlert))).Methods("GET")
	r.Handle("/weatheralert", Adapt(http.HandlerFunc(route.Handler.UpdateStateWeatherAlert))).Methods("PUT")
	r.Handle("/weatheralert/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableStateWeatherAlert))).Methods("PUT")
	r.Handle("/weatheralert/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableStateWeatherAlert))).Methods("PUT")
	r.Handle("/weatheralert/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteStateWeatherAlert))).Methods("DELETE")
	r.Handle("/weatheralert/filter", Adapt(http.HandlerFunc(route.Handler.FilterStateWeatherAlert))).Methods("POST")
	r.Handle("/weatheralert/updateweatheralertmaster", Adapt(http.HandlerFunc(route.Handler.UpdateStateWeatherAlertMatser))).Methods("PUT")
}
