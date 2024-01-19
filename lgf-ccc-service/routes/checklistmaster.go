package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ChecklistRoutes : ""
func (route *Route) ChecklistMasterRoutes(r *mux.Router) {
	r.Handle("/checklistMaster", Adapt(http.HandlerFunc(route.Handler.SaveChecklist))).Methods("POST")
	r.Handle("/checklistMaster", Adapt(http.HandlerFunc(route.Handler.GetSingleChecklist))).Methods("GET")
	r.Handle("/checklistMaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableChecklist))).Methods("PUT")
	r.Handle("/checklistMaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableChecklist))).Methods("PUT")
	r.Handle("/checklistMaster", Adapt(http.HandlerFunc(route.Handler.UpdateChecklist))).Methods("PUT")
	r.Handle("/checklistMaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterChecklist))).Methods("POST")
	r.Handle("/checklistMaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteChecklist))).Methods("DELETE")
}
