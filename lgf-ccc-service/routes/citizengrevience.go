package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CitizenGrevience : ""
func (route *Route) CitizenGrevienceRoutes(r *mux.Router) {
	r.Handle("/CitizenGrevience", Adapt(http.HandlerFunc(route.Handler.SaveCitizenGrevience))).Methods("POST")
	r.Handle("/CitizenGrevience", Adapt(http.HandlerFunc(route.Handler.GetSingleCitizenGrevience))).Methods("GET")
	r.Handle("/CitizenGrevience", Adapt(http.HandlerFunc(route.Handler.UpdateCitizenGrevience))).Methods("PUT")
	r.Handle("/CitizenGrevience/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCitizenGrevience))).Methods("PUT")
	r.Handle("/CitizenGrevience/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCitizenGrevience))).Methods("PUT")
	r.Handle("/CitizenGrevience/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCitizenGrevience))).Methods("DELETE")
	r.Handle("/CitizenGrevience/filter", Adapt(http.HandlerFunc(route.Handler.FilterCitizenGrevience))).Methods("POST")
	r.Handle("/CitizenGrevience/citizen/property", Adapt(http.HandlerFunc(route.Handler.UpdateCitizenProperty))).Methods("PUT")

}
