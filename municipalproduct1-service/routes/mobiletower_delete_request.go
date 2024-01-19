package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MobileTowerDeleteRequestRoutes : ""
func (route *Route) MobileTowerDeleteRequestRoutes(r *mux.Router) {
	// MobileTowerDelete Request Update
	r.Handle("/mobiletowerdeleterequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateMobileTowerDeleteRequest))).Methods("PUT")
	r.Handle("/mobiletowerdeleterequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptMobileTowerDeleteRequestUpdate))).Methods("PUT")
	r.Handle("/mobiletowerdeleterequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectMobileTowerDeleteRequestUpdate))).Methods("PUT")
	r.Handle("/mobiletowerdeleterequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterMobileTowerDeleteRequest))).Methods("POST")
	r.Handle("/mobiletowerdeleterequest", Adapt(http.HandlerFunc(route.Handler.GetSingleMobileTowerDeleteRequest))).Methods("GET")

	//
}
