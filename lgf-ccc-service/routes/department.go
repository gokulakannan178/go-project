package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DepartmentRoutes : ""
func (route *Route) DepartmentRoutes(r *mux.Router) {
	r.Handle("/Department", Adapt(http.HandlerFunc(route.Handler.SaveDepartment))).Methods("POST")
	r.Handle("/Department", Adapt(http.HandlerFunc(route.Handler.GetSingleDepartment))).Methods("GET")
	r.Handle("/Department", Adapt(http.HandlerFunc(route.Handler.UpdateDepartment))).Methods("PUT")
	r.Handle("/Department/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDepartment))).Methods("PUT")
	r.Handle("/Department/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDepartment))).Methods("PUT")
	r.Handle("/Department/filter", Adapt(http.HandlerFunc(route.Handler.FilterDepartment))).Methods("POST")
	r.Handle("/Department/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDepartment))).Methods("DELETE")
}
