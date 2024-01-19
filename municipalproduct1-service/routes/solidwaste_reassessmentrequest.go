package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SolidWasteReassessmentRequestRoutes : ""
func (route *Route) SolidWasteReassessmentRequestRoutes(r *mux.Router) {
	r.Handle("/solidwastereassessmentrequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateSolidWasteReassessmentRequest))).Methods("PUT")
	r.Handle("/solidwastereassessmentrequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptSolidWasteReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/solidwastereassessmentrequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectSolidWasteReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/solidwastereassessmentrequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterSolidWasteReassessmentRequest))).Methods("POST")
	r.Handle("/solidwastereassessmentrequest", Adapt(http.HandlerFunc(route.Handler.GetSingleSolidWasteReassessmentRequest))).Methods("GET")
	r.Handle("/solidwaste/basicupdate/getpayments", Adapt(http.HandlerFunc(route.Handler.BasicSolidWasteUpdateGetPaymentsToBeUpdated))).Methods("POST")

}
