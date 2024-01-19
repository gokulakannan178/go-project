package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// BillclaimConfigRoutes
func (route *Route) BillclaimConfigRoutes(r *mux.Router) {
	r.Handle("/billclaimconfig", Adapt(http.HandlerFunc(route.Handler.SaveBillclaimConfig))).Methods("POST")
	r.Handle("/billclaimconfig", Adapt(http.HandlerFunc(route.Handler.GetSingleBillclaimConfig))).Methods("GET")
	r.Handle("/billclaimconfig", Adapt(http.HandlerFunc(route.Handler.UpdateBillclaimConfig))).Methods("PUT")
	r.Handle("/billclaimconfig/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBillclaimConfig))).Methods("PUT")
	r.Handle("/billclaimconfig/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBillclaimConfig))).Methods("PUT")
	r.Handle("/billclaimconfig/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBillclaimConfig))).Methods("DELETE")
	r.Handle("/billclaimconfig/filter", Adapt(http.HandlerFunc(route.Handler.FilterBillclaimConfig))).Methods("POST")

}
