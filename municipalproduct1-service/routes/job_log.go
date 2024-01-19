package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) JobLogRoutes(r *mux.Router) {
	// JobLog
	r.Handle("/joblog", Adapt(http.HandlerFunc(route.Handler.SaveJobLog))).Methods("POST")
	r.Handle("/joblog", Adapt(http.HandlerFunc(route.Handler.GetSingleJobLog))).Methods("GET")
	r.Handle("/joblog", Adapt(http.HandlerFunc(route.Handler.UpdateJobLog))).Methods("PUT")
	r.Handle("/joblog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableJobLog))).Methods("PUT")
	r.Handle("/joblog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableJobLog))).Methods("PUT")
	r.Handle("/joblog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteJobLog))).Methods("DELETE")
	r.Handle("/joblog/filter", Adapt(http.HandlerFunc(route.Handler.FilterJobLog))).Methods("POST")
}
