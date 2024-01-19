package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserZoneAccessRoutes(r *mux.Router) {
	// UserZoneAccess
	r.Handle("/userzoneaccess", Adapt(http.HandlerFunc(route.Handler.SaveUserZoneAccess))).Methods("POST")
	r.Handle("/userzoneaccess", Adapt(http.HandlerFunc(route.Handler.GetSingleUserZoneAccess))).Methods("GET")
	r.Handle("/userzoneaccess", Adapt(http.HandlerFunc(route.Handler.UpdateUserZoneAccess))).Methods("PUT")
	r.Handle("/userzoneaccess/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserZoneAccess))).Methods("PUT")
	r.Handle("/userzoneaccess/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserZoneAccess))).Methods("PUT")
	r.Handle("/userzoneaccess/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserZoneAccess))).Methods("DELETE")
	r.Handle("/userzoneaccess/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserZoneAccess))).Methods("POST")
}
