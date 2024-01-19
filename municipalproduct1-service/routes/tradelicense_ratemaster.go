package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TradeLicenseRateMasterRoutes(r *mux.Router) {
	// TradeLicenseRateMaster
	r.Handle("/tradelicenseratemaster", Adapt(http.HandlerFunc(route.Handler.SaveTradeLicenseRateMaster))).Methods("POST")
	r.Handle("/tradelicenseratemaster", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseRateMaster))).Methods("GET")
	r.Handle("/tradelicenseratemaster", Adapt(http.HandlerFunc(route.Handler.UpdateTradeLicenseRateMaster))).Methods("PUT")
	r.Handle("/tradelicenseratemaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTradeLicenseRateMaster))).Methods("PUT")
	r.Handle("/tradelicenseratemaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTradeLicenseRateMaster))).Methods("PUT")
	r.Handle("/tradelicenseratemaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTradeLicenseRateMaster))).Methods("DELETE")
	r.Handle("/tradelicenseratemaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicenseRateMaster))).Methods("POST")
}
