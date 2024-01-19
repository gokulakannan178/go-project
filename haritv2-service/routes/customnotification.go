package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// customnotificationRoutes : ""
func (route *Route) CustomNotificationRoutes(r *mux.Router) {
	// customnotification
	r.Handle("/customnotification", Adapt(http.HandlerFunc(route.Handler.SaveCustomNotification))).Methods("POST")
	r.Handle("/customnotification", Adapt(http.HandlerFunc(route.Handler.GetSingleCustomNotification))).Methods("GET")
	r.Handle("/customnotification", Adapt(http.HandlerFunc(route.Handler.UpdateCustomNotification))).Methods("PUT")
	r.Handle("/customnotification/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCustomNotification))).Methods("PUT")
	r.Handle("/customnotification/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCustomNotification))).Methods("PUT")
	r.Handle("/customnotification/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCustomNotification))).Methods("DELETE")
	r.Handle("/customnotification/filter", Adapt(http.HandlerFunc(route.Handler.FilterCustomNotification))).Methods("POST")
}
