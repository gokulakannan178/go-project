package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) HospitalRoutes(r *mux.Router) {
	// Hospital
	r.Handle("/hospital", Adapt(http.HandlerFunc(route.Handler.SaveHospital))).Methods("POST")
	r.Handle("/hospital", Adapt(http.HandlerFunc(route.Handler.GetSingleHospital))).Methods("GET")
	r.Handle("/hospital", Adapt(http.HandlerFunc(route.Handler.UpdateHospital))).Methods("PUT")
	r.Handle("/hospital/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableHospital))).Methods("PUT")
	r.Handle("/hospital/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableHospital))).Methods("PUT")
	r.Handle("/hospital/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteHospital))).Methods("DELETE")
	r.Handle("/hospital/filter", Adapt(http.HandlerFunc(route.Handler.FilterHospital))).Methods("POST")
}
