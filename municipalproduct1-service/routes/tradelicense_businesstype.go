package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TradeLicenseBusinessTypeRoutes(r *mux.Router) {
	// TradeLicenseBusinessType
	r.Handle("/tradelicensebusinesstype", Adapt(http.HandlerFunc(route.Handler.SaveTradeLicenseBusinessType))).Methods("POST")
	r.Handle("/tradelicensebusinesstype", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseBusinessType))).Methods("GET")
	r.Handle("/tradelicensebusinesstype", Adapt(http.HandlerFunc(route.Handler.UpdateTradeLicenseBusinessType))).Methods("PUT")
	r.Handle("/tradelicensebusinesstype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTradeLicenseBusinessType))).Methods("PUT")
	r.Handle("/tradelicensebusinesstype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTradeLicenseBusinessType))).Methods("PUT")
	r.Handle("/tradelicensebusinesstype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTradeLicenseBusinessType))).Methods("DELETE")
	r.Handle("/tradelicensebusinesstype/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicenseBusinessType))).Methods("POST")
}
