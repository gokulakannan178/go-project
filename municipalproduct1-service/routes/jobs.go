package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) JobRoutes(r *mux.Router) {
	// Job
	r.Handle("/job", Adapt(http.HandlerFunc(route.Handler.SaveJob))).Methods("POST")
	r.Handle("/job", Adapt(http.HandlerFunc(route.Handler.GetSingleJob))).Methods("GET")
	r.Handle("/job", Adapt(http.HandlerFunc(route.Handler.UpdateJob))).Methods("PUT")
	r.Handle("/job/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableJob))).Methods("PUT")
	r.Handle("/job/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableJob))).Methods("PUT")
	r.Handle("/job/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteJob))).Methods("DELETE")
	r.Handle("/job/filter", Adapt(http.HandlerFunc(route.Handler.FilterJob))).Methods("POST")
	// ExecuteJob

	r.Handle("/job/execute", Adapt(http.HandlerFunc(route.Handler.ExecuteJob))).Methods("POST")

}
