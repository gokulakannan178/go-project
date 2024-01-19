package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) BillClaimLogRoutes(r *mux.Router) {
	// BillClaimLog
	r.Handle("/billClaimLog", Adapt(http.HandlerFunc(route.Handler.SaveBillClaimLog))).Methods("POST")
	r.Handle("/billClaimLog", Adapt(http.HandlerFunc(route.Handler.GetSingleBillClaimLog))).Methods("GET")
	r.Handle("/billClaimLog", Adapt(http.HandlerFunc(route.Handler.UpdateBillClaimLog))).Methods("PUT")
	r.Handle("/billClaimLog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBillClaimLog))).Methods("PUT")
	r.Handle("/billClaimLog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBillClaimLog))).Methods("PUT")
	r.Handle("/billClaimLog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBillClaimLog))).Methods("DELETE")
	r.Handle("/billClaimLog/filter", Adapt(http.HandlerFunc(route.Handler.FilterBillClaimLog))).Methods("POST")

}
