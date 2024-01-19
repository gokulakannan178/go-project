package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserAclRoutes(r *mux.Router) {
	r.Handle("/useracl", Adapt(http.HandlerFunc(route.Handler.SaveUserAcl))).Methods("POST")
	r.Handle("/useracl", Adapt(http.HandlerFunc(route.Handler.GetSingleUserAcl))).Methods("GET")
	r.Handle("/useracl", Adapt(http.HandlerFunc(route.Handler.UpdateUserAcl))).Methods("PUT")
	r.Handle("/useracl/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserAcl))).Methods("PUT")
	r.Handle("/useracl/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserAcl))).Methods("PUT")
	r.Handle("/useracl/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserAcl))).Methods("DELETE")
	r.Handle("/useracl/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserAcl))).Methods("POST")
	r.Handle("/useracl/username", Adapt(http.HandlerFunc(route.Handler.SaveUserAclWithUpsert))).Methods("POST")
	r.Handle("/useracl/username", Adapt(http.HandlerFunc(route.Handler.GetSingleUserAclWithUserName))).Methods("GET")

}
