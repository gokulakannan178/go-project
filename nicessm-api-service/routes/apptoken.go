package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ApptokenRoutes(r *mux.Router) {
	//Apptoken
	r.Handle("/apptoken", Adapt(http.HandlerFunc(route.Handler.SaveApptoken))).Methods("POST")
	r.Handle("/apptoken", Adapt(http.HandlerFunc(route.Handler.GetSingleApptoken))).Methods("GET")
	r.Handle("/apptoken", Adapt(http.HandlerFunc(route.Handler.UpdateApptoken))).Methods("PUT")
	r.Handle("/apptoken/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableApptoken))).Methods("PUT")
	r.Handle("/apptoken/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableApptoken))).Methods("PUT")
	r.Handle("/apptoken/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteApptoken))).Methods("DELETE")
	r.Handle("/apptoken/filter", Adapt(http.HandlerFunc(route.Handler.FilterApptoken))).Methods("GET")
	r.Handle("/apptoken/login", Adapt(http.HandlerFunc(route.Handler.LoginRegistration))).Methods("POST")
	r.Handle("/apptoken/logout", Adapt(http.HandlerFunc(route.Handler.LogoutApptoken))).Methods("DELETE")
	r.Handle("/apptoken/uniquecheck", Adapt(http.HandlerFunc(route.Handler.GetSingleApptokenWithUniqueCheck))).Methods("GET")

	//	r.Handle("/notification/notify", Adapt(http.HandlerFunc(route.Handler.SendAppTokenNotification))).Methods("POST")

}
