package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// BillclaimLevelsRoutes
func (route *Route) BillclaimLevelsRoutes(r *mux.Router) {
	r.Handle("/billclaimlevels", Adapt(http.HandlerFunc(route.Handler.SaveBillclaimLevels))).Methods("POST")
	r.Handle("/billclaimlevels", Adapt(http.HandlerFunc(route.Handler.GetSingleBillclaimLevels))).Methods("GET")
	r.Handle("/billclaimlevels", Adapt(http.HandlerFunc(route.Handler.UpdateBillclaimLevels))).Methods("PUT")
	r.Handle("/billclaimlevels/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBillclaimLevels))).Methods("PUT")
	r.Handle("/billclaimlevels/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBillclaimLevels))).Methods("PUT")
	r.Handle("/billclaimlevels/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBillclaimLevels))).Methods("DELETE")
	r.Handle("/billclaimlevels/filter", Adapt(http.HandlerFunc(route.Handler.FilterBillclaimLevels))).Methods("POST")
	r.Handle("/billclaimlevels/status/approved", Adapt(http.HandlerFunc(route.Handler.ApprovedBillclaimLevels))).Methods("PUT")

}
