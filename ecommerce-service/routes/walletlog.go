package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WalletLogRoutes(r *mux.Router) {
	// Vendor
	r.Handle("/walletlog", Adapt(http.HandlerFunc(route.Handler.SaveWalletLog))).Methods("POST")
	r.Handle("/walletlog", Adapt(http.HandlerFunc(route.Handler.GetSingleWalletLog))).Methods("GET")
	r.Handle("/walletlog", Adapt(http.HandlerFunc(route.Handler.UpdateWalletLog))).Methods("PUT")
	r.Handle("/walletlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWalletLog))).Methods("PUT")
	r.Handle("/walletlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWalletLog))).Methods("PUT")
	r.Handle("/walletlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWalletLog))).Methods("DELETE")
	r.Handle("/walletlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterWalletLog))).Methods("POST")
}
