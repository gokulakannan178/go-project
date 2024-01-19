package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ConstructionTypeRoutes : ""
func (route *Route) ShoprentReassessmentRequestRoutes(r *mux.Router) {
	// Reassessment Request Update
	r.Handle("/shoprentreassessmentrequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateShoprentReassessmentRequest))).Methods("PUT")
	r.Handle("/shoprentreassessmentrequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptShoprentReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/shoprentreassessmentrequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectShoprentReassessmentRequestUpdate))).Methods("PUT")
	r.Handle("/shoprentreassessmentrequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterShoprentReassessmentRequest))).Methods("POST")
	r.Handle("/shoprentreassessmentrequest", Adapt(http.HandlerFunc(route.Handler.GetSingleShoprentReassessmentRequest))).Methods("GET")

}
