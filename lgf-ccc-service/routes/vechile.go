package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// VechileRoutes
func (route *Route) VechileRoutes(r *mux.Router) {
	r.Handle("/vechile", Adapt(http.HandlerFunc(route.Handler.SaveVechile))).Methods("POST")
	r.Handle("/vechile", Adapt(http.HandlerFunc(route.Handler.GetSingleVechile))).Methods("GET")
	r.Handle("/vechile", Adapt(http.HandlerFunc(route.Handler.UpdateVechile))).Methods("PUT")
	r.Handle("/vechile/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVechile))).Methods("PUT")
	r.Handle("/vechile/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVechile))).Methods("PUT")
	r.Handle("/vechile/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVechile))).Methods("DELETE")
	r.Handle("/vechile/filter", Adapt(http.HandlerFunc(route.Handler.FilterVechile))).Methods("POST")
	r.Handle("/vechile/assign", Adapt(http.HandlerFunc(route.Handler.VechileAssign))).Methods("POST")

}
