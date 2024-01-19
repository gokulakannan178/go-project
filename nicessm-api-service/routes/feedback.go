package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) FeedBackRoutes(r *mux.Router) {
	r.Handle("/feedback", Adapt(http.HandlerFunc(route.Handler.SaveFeedBack))).Methods("POST")
	r.Handle("/feedback", Adapt(http.HandlerFunc(route.Handler.GetSingleFeedBack))).Methods("GET")
	r.Handle("/feedback", Adapt(http.HandlerFunc(route.Handler.UpdateFeedBack))).Methods("PUT")
	r.Handle("/feedback/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFeedBack))).Methods("PUT")
	r.Handle("/feedback/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFeedBack))).Methods("PUT")
	r.Handle("/feedback/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFeedBack))).Methods("DELETE")
	r.Handle("/feedback/filter", Adapt(http.HandlerFunc(route.Handler.FilterFeedBack))).Methods("POST")
	r.Handle("/consolidatedFeedBack", Adapt(http.HandlerFunc(route.Handler.ConsolidatedFeedBack))).Methods("POST")

}
