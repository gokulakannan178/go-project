package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NotificationLogRoutes : ""
func (route *Route) NotificationLogRoutes(r *mux.Router) {
	r.Handle("/notificationLog", Adapt(http.HandlerFunc(route.Handler.SaveNotificationLog))).Methods("POST")
	r.Handle("/notificationLog", Adapt(http.HandlerFunc(route.Handler.GetSingleNotificationLog))).Methods("GET")
	r.Handle("/notificationLog", Adapt(http.HandlerFunc(route.Handler.UpdateNotificationLog))).Methods("PUT")
	r.Handle("/notificationLog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableNotificationLog))).Methods("PUT")
	r.Handle("/notificationLog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableNotificationLog))).Methods("PUT")
	r.Handle("/notificationLog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteNotificationLog))).Methods("DELETE")
	r.Handle("/notificationLog/filter", Adapt(http.HandlerFunc(route.Handler.FilterNotificationLog))).Methods("POST")
}
