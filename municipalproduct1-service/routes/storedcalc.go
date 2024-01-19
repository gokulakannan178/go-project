package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) StoredCalcRoutes(r *mux.Router) {
	r.Handle("/storedcalc", Adapt(http.HandlerFunc(route.Handler.SaveStoredCalc))).Methods("POST")
	r.Handle("/storedcalc", Adapt(http.HandlerFunc(route.Handler.GetSingleStoredCalc))).Methods("GET")
	r.Handle("/storedcalc", Adapt(http.HandlerFunc(route.Handler.UpdateStoredCalc))).Methods("PUT")
	r.Handle("/storedcalc/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableStoredCalc))).Methods("PUT")
	r.Handle("/storedcalc/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableStoredCalc))).Methods("PUT")
	r.Handle("/storedcalc/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteStoredCalc))).Methods("DELETE")
	r.Handle("/storedcalc/filter", Adapt(http.HandlerFunc(route.Handler.FilterStoredCalc))).Methods("POST")
}
