package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) LeaseDayWiseRoutes(r *mux.Router) {
	// Lease Day Wise
	r.Handle("/leasedaywise", Adapt(http.HandlerFunc(route.Handler.SaveLeaseDayWise))).Methods("POST")
	r.Handle("/leasedaywise", Adapt(http.HandlerFunc(route.Handler.GetSingleLeaseDayWise))).Methods("GET")
	r.Handle("/leasedaywise", Adapt(http.HandlerFunc(route.Handler.UpdateLeaseDayWise))).Methods("PUT")
	r.Handle("/leasedaywise/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLeaseDayWise))).Methods("PUT")
	r.Handle("/leasedaywise/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLeaseDayWise))).Methods("PUT")
	r.Handle("/leasedaywise/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLeaseDayWise))).Methods("DELETE")
	r.Handle("/leasedaywise/filter", Adapt(http.HandlerFunc(route.Handler.FilterLeaseDayWise))).Methods("POST")
}
