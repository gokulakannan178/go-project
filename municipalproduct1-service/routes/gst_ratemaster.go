package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GSTRateMasterRoutes : ""
func (route *Route) GSTRateMasterRoutes(r *mux.Router) {

	r.Handle("/gstratemaster", Adapt(http.HandlerFunc(route.Handler.SaveGSTRateMaster))).Methods("POST")
	r.Handle("/gstratemaster", Adapt(http.HandlerFunc(route.Handler.GetSingleGSTRateMaster))).Methods("GET")
	r.Handle("/gstratemaster", Adapt(http.HandlerFunc(route.Handler.UpdateGSTRateMaster))).Methods("PUT")
	r.Handle("/gstratemaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableGSTRateMaster))).Methods("PUT")
	r.Handle("/gstratemaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableGSTRateMaster))).Methods("PUT")
	r.Handle("/gstratemaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteGSTRateMaster))).Methods("DELETE")
	r.Handle("/gstratemaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterGSTRateMaster))).Methods("POST")
}
