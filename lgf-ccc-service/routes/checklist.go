package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ChecklistRoutes : ""
func (route *Route) ChecklistRoutes(r *mux.Router) {
	r.Handle("/checklist", Adapt(http.HandlerFunc(route.Handler.SaveChecklist))).Methods("POST")
	r.Handle("/checklist", Adapt(http.HandlerFunc(route.Handler.GetSingleChecklist))).Methods("GET")
	r.Handle("/checklist/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableChecklist))).Methods("PUT")
	r.Handle("/checklist/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableChecklist))).Methods("PUT")
	r.Handle("/checklist", Adapt(http.HandlerFunc(route.Handler.UpdateChecklist))).Methods("PUT")
	r.Handle("/checklist/filter", Adapt(http.HandlerFunc(route.Handler.FilterChecklist))).Methods("POST")
	r.Handle("/checklist/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteChecklist))).Methods("DELETE")
}
