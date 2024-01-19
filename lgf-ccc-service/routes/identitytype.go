package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//IdentityTypeRoutes : ""
func (route *Route) IdentityTypeRoutes(r *mux.Router) {
	r.Handle("/identitytype", Adapt(http.HandlerFunc(route.Handler.SaveIdentityType))).Methods("POST")
	r.Handle("/identitytype", Adapt(http.HandlerFunc(route.Handler.GetSingleIdentityType))).Methods("GET")
	r.Handle("/identitytype", Adapt(http.HandlerFunc(route.Handler.UpdateIdentityType))).Methods("PUT")
	r.Handle("/identitytype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableIdentityType))).Methods("PUT")
	r.Handle("/identitytype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableIdentityType))).Methods("PUT")
	r.Handle("/identitytype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteIdentityType))).Methods("DELETE")
	r.Handle("/identitytype/filter", Adapt(http.HandlerFunc(route.Handler.FilterIdentityType))).Methods("POST")

}
