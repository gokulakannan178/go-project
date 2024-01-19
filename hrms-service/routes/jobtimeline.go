package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) JobTimelineRoutes(r *mux.Router) {
	// JobTimeline
	r.Handle("/jobTimeline", Adapt(http.HandlerFunc(route.Handler.SaveJobTimeline))).Methods("POST")
	r.Handle("/jobTimeline", Adapt(http.HandlerFunc(route.Handler.GetSingleJobTimeline))).Methods("GET")
	r.Handle("/jobTimeline", Adapt(http.HandlerFunc(route.Handler.UpdateJobTimeline))).Methods("PUT")
	r.Handle("/jobTimeline/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableJobTimeline))).Methods("PUT")
	r.Handle("/jobTimeline/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableJobTimeline))).Methods("PUT")
	r.Handle("/jobTimeline/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteJobTimeline))).Methods("DELETE")
	r.Handle("/jobTimeline/filter", Adapt(http.HandlerFunc(route.Handler.FilterJobTimeline))).Methods("POST")

}
