package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ProductConfigurationRoutes : ""
func (route *Route) ContactUs(r *mux.Router) {
	r.Handle("/contactus/sendemail", Adapt(http.HandlerFunc(route.Handler.SendContactUs))).Methods("POST")
	r.Handle("/contactus", Adapt(http.HandlerFunc(route.Handler.SaveContact))).Methods("POST")
	r.Handle("/contactus", Adapt(http.HandlerFunc(route.Handler.GetSingleContact))).Methods("GET")
	r.Handle("/contactus", Adapt(http.HandlerFunc(route.Handler.UpdateContact))).Methods("PUT")
	r.Handle("/contactus/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableContact))).Methods("PUT")
	r.Handle("/contactus/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableContact))).Methods("PUT")
	r.Handle("/contactus/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteContact))).Methods("DELETE")
	r.Handle("/contactus/filter", Adapt(http.HandlerFunc(route.Handler.FilterContactUs))).Methods("POST")

}
