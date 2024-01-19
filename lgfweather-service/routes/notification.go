package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NotificationRoutes : ""
func (route *Route) NotificationRoutes(r *mux.Router) {
	// Notification
	r.Handle("/notification", Adapt(http.HandlerFunc(route.Handler.SaveNotification))).Methods("POST")
	r.Handle("/notification", Adapt(http.HandlerFunc(route.Handler.GetSingleNotification))).Methods("GET")
	r.Handle("/notification", Adapt(http.HandlerFunc(route.Handler.UpdateNotification))).Methods("PUT")
	r.Handle("/notification/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableNotification))).Methods("PUT")
	r.Handle("/notification/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableNotification))).Methods("PUT")
	r.Handle("/notification/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteNotification))).Methods("DELETE")
	r.Handle("/notification/filter", Adapt(http.HandlerFunc(route.Handler.FilterNotification))).Methods("POST")

}
