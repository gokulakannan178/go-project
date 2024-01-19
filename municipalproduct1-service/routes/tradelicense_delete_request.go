package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TradeLicenseDeleteRequestRoutes : ""
func (route *Route) TradeLicenseDeleteRequestRoutes(r *mux.Router) {
	// TradeLicenseDelete Request Update
	r.Handle("/tradeLicensedeleterequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateTradeLicenseDeleteRequest))).Methods("PUT")
	r.Handle("/tradeLicensedeleterequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptTradeLicenseDeleteRequestUpdate))).Methods("PUT")
	r.Handle("/tradeLicensedeleterequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectTradeLicenseDeleteRequestUpdate))).Methods("PUT")
	r.Handle("/tradeLicensedeleterequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicenseDeleteRequest))).Methods("POST")
	r.Handle("/tradeLicensedeleterequest", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseDeleteRequest))).Methods("GET")

	//
}
