package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//TradeLicenseRebateRoutes : ""
func (route *Route) TradeLicenseRebateRoutes(r *mux.Router) {
	r.Handle("/tradelicenserebate", Adapt(http.HandlerFunc(route.Handler.SaveTradeLicenseRebate))).Methods("POST")
	r.Handle("/tradelicenserebate", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseRebate))).Methods("GET")
	r.Handle("/tradelicenserebate", Adapt(http.HandlerFunc(route.Handler.UpdateTradeLicenseRebate))).Methods("PUT")
	r.Handle("/tradelicenserebate/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTradeLicenseRebate))).Methods("PUT")
	r.Handle("/tradelicenserebate/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTradeLicenseRebate))).Methods("PUT")
	r.Handle("/tradelicenserebate/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTradeLicenseRebate))).Methods("DELETE")
	r.Handle("/tradelicenserebate/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicenseRebate))).Methods("POST")
}
