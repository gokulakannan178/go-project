package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WalletRoutes(r *mux.Router) {
	// Vendor
	r.Handle("/wallet", Adapt(http.HandlerFunc(route.Handler.SaveWallet))).Methods("POST")
	r.Handle("/wallet", Adapt(http.HandlerFunc(route.Handler.GetSingleWallet))).Methods("GET")
	r.Handle("/wallet", Adapt(http.HandlerFunc(route.Handler.UpdateWallet))).Methods("PUT")
	r.Handle("/wallet/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWallet))).Methods("PUT")
	r.Handle("/wallet/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWallet))).Methods("PUT")
	r.Handle("/wallet/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWallet))).Methods("DELETE")
	r.Handle("/wallet/filter", Adapt(http.HandlerFunc(route.Handler.FilterWallet))).Methods("POST")
}
