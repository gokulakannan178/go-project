package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// EmployeeAssetsRoutes : ""
func (route *Route) EmployeeAssetsRoutes(r *mux.Router) {
	r.Handle("/employeeassets", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeAssets))).Methods("POST")
	r.Handle("/employeeassets", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeAssets))).Methods("GET")
	r.Handle("/employeeassets", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeAssets))).Methods("PUT")
	r.Handle("/employeeassets/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeAssets))).Methods("PUT")
	r.Handle("/employeeassets/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeAssets))).Methods("PUT")
	r.Handle("/employeeassets/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeAssets))).Methods("DELETE")
	r.Handle("/employeeassets/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeAssets))).Methods("POST")

}
