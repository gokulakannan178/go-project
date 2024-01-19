package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PopupNotificationRoutes(r *mux.Router) {
	//Popup Notification
	r.Handle("/popupnotification", Adapt(http.HandlerFunc(route.Handler.SavePopupNotification))).Methods("POST")
	r.Handle("/popupnotification", Adapt(http.HandlerFunc(route.Handler.GetSinglePopupNotification))).Methods("GET")
	r.Handle("/popupnotification", Adapt(http.HandlerFunc(route.Handler.UpdatePopupNotification))).Methods("PUT")
	r.Handle("/popupnotification/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePopupNotification))).Methods("PUT")
	r.Handle("/popupnotification/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePopupNotification))).Methods("PUT")
	r.Handle("/popupnotification/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePopupNotification))).Methods("DELETE")
	r.Handle("/popupnotification/filter", Adapt(http.HandlerFunc(route.Handler.FilterPopupNotification))).Methods("POST")
	r.Handle("/popupnotification/getdefault", Adapt(http.HandlerFunc(route.Handler.GetDefaultPopupNotification))).Methods("GET")
	r.Handle("/popupnotification/set", Adapt(http.HandlerFunc(route.Handler.SetPopupNotification))).Methods("PUT")
	r.Handle("/popupnotification/enable", Adapt(http.HandlerFunc(route.Handler.EnablePopupNotificationV2))).Methods("PUT")
	r.Handle("/popupnotification/disable", Adapt(http.HandlerFunc(route.Handler.DisablePopupNotificationV2))).Methods("PUT")
}
