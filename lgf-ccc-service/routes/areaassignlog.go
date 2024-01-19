package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//AreaAssignLog : ""
func (route *Route) AreaAssignLogRoutes(r *mux.Router) {
	r.Handle("/areaassignlog", Adapt(http.HandlerFunc(route.Handler.SaveAreaAssignLog))).Methods("POST")
	r.Handle("/areaassignlog", Adapt(http.HandlerFunc(route.Handler.GetSingleAreaAssignLog))).Methods("GET")
	r.Handle("/areaassignlog", Adapt(http.HandlerFunc(route.Handler.UpdateAreaAssignLog))).Methods("PUT")
	r.Handle("/areaassignlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAreaAssignLog))).Methods("PUT")
	r.Handle("/areaassignlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAreaAssignLog))).Methods("PUT")
	r.Handle("/areaassignlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAreaAssignLog))).Methods("DELETE")
	r.Handle("/areaassignlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterAreaAssignLog))).Methods("POST")
}
