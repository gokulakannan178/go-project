package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) EmergencyContactRoutes(r *mux.Router) {
	// EmergencyContact
	r.Handle("/emergencyContact", Adapt(http.HandlerFunc(route.Handler.SaveEmergencyContact))).Methods("POST")
	r.Handle("/emergencyContact", Adapt(http.HandlerFunc(route.Handler.GetSingleEmergencyContact))).Methods("GET")
	r.Handle("/emergencyContact", Adapt(http.HandlerFunc(route.Handler.UpdateEmergencyContact))).Methods("PUT")
	r.Handle("/emergencyContact/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmergencyContact))).Methods("PUT")
	r.Handle("/emergencyContact/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmergencyContact))).Methods("PUT")
	r.Handle("/emergencyContact/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmergencyContact))).Methods("DELETE")
	r.Handle("/emergencyContact/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmergencyContact))).Methods("POST")

}
