package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//TaskRoutes : ""
func (route *Route) TaskRoutes(r *mux.Router) {
	r.Handle("/task", Adapt(http.HandlerFunc(route.Handler.SaveTask))).Methods("POST")
	r.Handle("/task", Adapt(http.HandlerFunc(route.Handler.GetSingleTask))).Methods("GET")
	r.Handle("/task/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTask))).Methods("PUT")
	r.Handle("/task/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTask))).Methods("PUT")
	r.Handle("/task/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTask))).Methods("DELETE")
	r.Handle("/task", Adapt(http.HandlerFunc(route.Handler.UpdateTask))).Methods("PUT")
	r.Handle("/task/filter", Adapt(http.HandlerFunc(route.Handler.FilterTask))).Methods("POST")
	r.Handle("/task/addteammember", Adapt(http.HandlerFunc(route.Handler.SaveTaskTeamMember))).Methods("POST")
	r.Handle("/task/status/removeteammember", Adapt(http.HandlerFunc(route.Handler.DisableTaskTeamMember))).Methods("PUT")

	// TaskMessage : ""
	r.Handle("/taskmessage", Adapt(http.HandlerFunc(route.Handler.SaveTaskMessage))).Methods("POST")
	r.Handle("/taskmessage/filter", Adapt(http.HandlerFunc(route.Handler.FilterTaskMessage))).Methods("POST")

}
