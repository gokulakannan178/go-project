package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) BillClaimRoutes(r *mux.Router) {
	// BillClaim
	r.Handle("/billClaim", Adapt(http.HandlerFunc(route.Handler.SaveBillClaim))).Methods("POST")
	r.Handle("/billClaim", Adapt(http.HandlerFunc(route.Handler.GetSingleBillClaim))).Methods("GET")
	r.Handle("/billClaim", Adapt(http.HandlerFunc(route.Handler.UpdateBillClaim))).Methods("PUT")
	r.Handle("/billClaim/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBillClaim))).Methods("PUT")
	r.Handle("/billClaim/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBillClaim))).Methods("PUT")
	r.Handle("/billClaim/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBillClaim))).Methods("DELETE")
	r.Handle("/billClaim/filter", Adapt(http.HandlerFunc(route.Handler.FilterBillClaim))).Methods("POST")
	r.Handle("/billClaim/approved", Adapt(http.HandlerFunc(route.Handler.ApprovedBillClaim))).Methods("PUT")
	r.Handle("/billClaim/rejected", Adapt(http.HandlerFunc(route.Handler.RejectedBillClaim))).Methods("PUT")

}
