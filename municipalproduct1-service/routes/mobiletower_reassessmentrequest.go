package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//MobileTowerReassessmentRequestRoutes : ""
func (route *Route) MobileTowerReassessmentRequestRoutes(r *mux.Router) {
	r.Handle("/mobiletowerreassessmentrequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateMobileTowerReassessmentRequest))).Methods("PUT")
	r.Handle("/mobiletowerreassessmentrequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptMobileTowerReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/mobiletowerreassessmentrequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectMobileTowerReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/mobiletowerreassessmentrequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterMobileTowerReassessmentRequest))).Methods("POST")
	r.Handle("/mobiletowerreassessmentrequest", Adapt(http.HandlerFunc(route.Handler.GetSingleMobileTowerReassessmentRequest))).Methods("GET")

}
