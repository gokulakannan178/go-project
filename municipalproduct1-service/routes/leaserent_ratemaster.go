package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)


//LeaseRentRateMasterRoutes : ""
func (route *Route) LeaseRentRateMasterRoutes(r *mux.Router) {
	r.Handle("/leaserentratemaster", Adapt(http.HandlerFunc(route.Handler.SaveLeaseRentRateMaster))).Methods("POST")
	r.Handle("/leaserentratemaster", Adapt(http.HandlerFunc(route.Handler.GetSingleLeaseRentRateMaster))).Methods("GET")
	r.Handle("/leaserentratemaster", Adapt(http.HandlerFunc(route.Handler.UpdateLeaseRentRateMaster))).Methods("PUT")
	r.Handle("/leaserentratemaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLeaseRentRateMaster))).Methods("PUT")
	r.Handle("/leaserentratemaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLeaseRentRateMaster))).Methods("PUT")
	r.Handle("/leaserentratemaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLeaseRentRateMaster))).Methods("DELETE")
	r.Handle("/leaserentratemaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterLeaseRentRateMaster))).Methods("POST")
}
