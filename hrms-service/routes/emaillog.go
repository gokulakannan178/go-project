package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmailLogRoutes : ""
func (route *Route) EmailLogRoutes(r *mux.Router) {
	r.Handle("/emailLog", Adapt(http.HandlerFunc(route.Handler.SaveEmailLog))).Methods("POST")
	r.Handle("/emailLog", Adapt(http.HandlerFunc(route.Handler.GetSingleEmailLog))).Methods("GET")
	r.Handle("/emailLog", Adapt(http.HandlerFunc(route.Handler.UpdateEmailLog))).Methods("PUT")
	r.Handle("/emailLog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmailLog))).Methods("PUT")
	r.Handle("/emailLog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmailLog))).Methods("PUT")
	r.Handle("/emailLog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmailLog))).Methods("DELETE")
	r.Handle("/emailLog/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmailLog))).Methods("POST")
}
