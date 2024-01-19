package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)


//LeaseRentRoutes : ""
func (route *Route) LeaseRentRoutes(r *mux.Router) {
	r.Handle("/leaserent", Adapt(http.HandlerFunc(route.Handler.SaveLeaseRent))).Methods("POST")
	r.Handle("/leaserent", Adapt(http.HandlerFunc(route.Handler.GetSingleLeaseRent))).Methods("GET")
	r.Handle("/leaserent", Adapt(http.HandlerFunc(route.Handler.UpdateLeaseRent))).Methods("PUT")
	r.Handle("/leaserent/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLeaseRent))).Methods("PUT")
	r.Handle("/leaserent/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLeaseRent))).Methods("PUT")
	r.Handle("/leaserent/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLeaseRent))).Methods("DELETE")
	r.Handle("/leaserent/filter", Adapt(http.HandlerFunc(route.Handler.FilterLeaseRent))).Methods("POST")
}
