package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//OurServiceRoutes : ""
func (route *Route) OurServiceRoutes(r *mux.Router) {
	r.Handle("/ourservice/{scenario}", Adapt(http.HandlerFunc(route.Handler.SaveOurService))).Methods("POST")
	r.Handle("/ourservice/{scenario}", Adapt(http.HandlerFunc(route.Handler.GetSingleOurService))).Methods("GET")
	r.Handle("/ourservice/{scenario}/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOurService))).Methods("PUT")
	r.Handle("/ourservice/{scenario}/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOurService))).Methods("PUT")
	r.Handle("/ourservice/{scenario}", Adapt(http.HandlerFunc(route.Handler.UpdateOurService))).Methods("PUT")
	r.Handle("/ourservice/{scenario}/filter", Adapt(http.HandlerFunc(route.Handler.FilterOurService))).Methods("POST")
	r.Handle("/ourservice/{scenario}/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOurService))).Methods("DELETE")
}
