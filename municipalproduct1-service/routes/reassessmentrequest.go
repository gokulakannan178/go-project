package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ConstructionTypeRoutes : ""
func (route *Route) ReassessmentRequestRoutes(r *mux.Router) {
	// Reassessment Request Update
	r.Handle("/reassessmentrequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateReassessmentRequest))).Methods("PUT")
	r.Handle("/reassessmentrequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/reassessmentrequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/reassessmentrequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterReassessmentRequest))).Methods("POST")
	r.Handle("/reassessmentrequest", Adapt(http.HandlerFunc(route.Handler.GetSingleReassessmentRequest))).Methods("GET")

}
