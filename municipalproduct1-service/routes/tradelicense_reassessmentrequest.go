package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//TradeLicenseReassessmentRequestRoutes : ""
func (route *Route) TradeLicenseReassessmentRequestRoutes(r *mux.Router) {
	r.Handle("/tradelicensereassessmentrequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateTradeLicenseReassessmentRequest))).Methods("PUT")
	r.Handle("/tradelicensereassessmentrequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptTradeLicenseReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/tradelicensereassessmentrequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectTradeLicenseReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/tradelicensereassessmentrequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicenseReassessmentRequest))).Methods("POST")
	r.Handle("/tradelicensereassessmentrequest", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseReassessmentRequest))).Methods("GET")

}
