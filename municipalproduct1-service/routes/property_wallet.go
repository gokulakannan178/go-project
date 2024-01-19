package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PropertyWalletRoutes : ""
func (route *Route) PropertyWalletRoutes(r *mux.Router) {
	r.Handle("/propertywallet", Adapt(http.HandlerFunc(route.Handler.SavePropertyWallet))).Methods("POST")
	r.Handle("/propertywallet", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyWallet))).Methods("GET")
	r.Handle("/propertywallet", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyWallet))).Methods("PUT")
	r.Handle("/propertywallet/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyWallet))).Methods("PUT")
	r.Handle("/propertywallet/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyWallet))).Methods("PUT")
	r.Handle("/propertywallet/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyWallet))).Methods("DELETE")
	r.Handle("/propertywallet/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyWallet))).Methods("POST")

}
