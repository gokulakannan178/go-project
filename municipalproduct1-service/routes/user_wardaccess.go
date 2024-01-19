package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserWardAccessRoutes(r *mux.Router) {
	// UserWardAccess
	r.Handle("/userwardaccess", Adapt(http.HandlerFunc(route.Handler.SaveUserWardAccess))).Methods("POST")
	r.Handle("/userwardaccess", Adapt(http.HandlerFunc(route.Handler.GetSingleUserWardAccess))).Methods("GET")
	r.Handle("/userwardaccess", Adapt(http.HandlerFunc(route.Handler.UpdateUserWardAccess))).Methods("PUT")
	r.Handle("/userwardaccess/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserWardAccess))).Methods("PUT")
	r.Handle("/userwardaccess/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserWardAccess))).Methods("PUT")
	r.Handle("/userwardaccess/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserWardAccess))).Methods("DELETE")
	r.Handle("/userwardaccess/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserWardAccess))).Methods("POST")
}
