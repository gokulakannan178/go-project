package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) RoleRoutes(r *mux.Router) {
	// Role
	r.Handle("/role", Adapt(http.HandlerFunc(route.Handler.SaveRole))).Methods("POST")
	r.Handle("/role", Adapt(http.HandlerFunc(route.Handler.GetSingleRole))).Methods("GET")
	r.Handle("/role", Adapt(http.HandlerFunc(route.Handler.UpdateRole))).Methods("PUT")
	r.Handle("/role/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableRole))).Methods("PUT")
	r.Handle("/role/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableRole))).Methods("PUT")
	r.Handle("/role/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteRole))).Methods("DELETE")
	r.Handle("/role/filter", Adapt(http.HandlerFunc(route.Handler.FilterRole))).Methods("POST")
}
