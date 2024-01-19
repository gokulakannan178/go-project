package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WorkScheduleRoutes(r *mux.Router) {
	// WorkSchedule
	r.Handle("/workSchedule", Adapt(http.HandlerFunc(route.Handler.SaveWorkSchedule))).Methods("POST")
	r.Handle("/workSchedule", Adapt(http.HandlerFunc(route.Handler.GetSingleWorkSchedule))).Methods("GET")
	r.Handle("/workSchedule", Adapt(http.HandlerFunc(route.Handler.UpdateWorkSchedule))).Methods("PUT")
	r.Handle("/workSchedule/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWorkSchedule))).Methods("PUT")
	r.Handle("/workSchedule/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWorkSchedule))).Methods("PUT")
	r.Handle("/workSchedule/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWorkSchedule))).Methods("DELETE")
	r.Handle("/workSchedule/filter", Adapt(http.HandlerFunc(route.Handler.FilterWorkSchedule))).Methods("POST")

}
