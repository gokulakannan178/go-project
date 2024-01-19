package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserLocationLogRoutes(r *mux.Router) {
	// UserLocationLog
	r.Handle("/userlocationlog", Adapt(http.HandlerFunc(route.Handler.SaveUserLocationLog))).Methods("POST")
	r.Handle("/userlocationlog", Adapt(http.HandlerFunc(route.Handler.GetSingleUserLocationLog))).Methods("GET")
	r.Handle("/userlocationlog", Adapt(http.HandlerFunc(route.Handler.UpdateUserLocationLog))).Methods("PUT")
	r.Handle("/userlocationlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserLocationLog))).Methods("PUT")
	r.Handle("/userlocationlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserLocationLog))).Methods("PUT")
	r.Handle("/userlocationlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserLocationLog))).Methods("DELETE")
	r.Handle("/userlocationlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserLocationLog))).Methods("POST")
}
