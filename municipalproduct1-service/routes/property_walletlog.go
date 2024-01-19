package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PropertyWalletLogRoutes : ""
func (route *Route) PropertyWalletLogRoutes(r *mux.Router) {
	r.Handle("/propertywalletlog", Adapt(http.HandlerFunc(route.Handler.SavePropertyWalletLog))).Methods("POST")
	r.Handle("/propertywalletlog", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyWalletLog))).Methods("GET")
	r.Handle("/propertywalletlog", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyWalletLog))).Methods("PUT")
	r.Handle("/propertywalletlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyWalletLog))).Methods("PUT")
	r.Handle("/propertywalletlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyWalletLog))).Methods("PUT")
	r.Handle("/propertywalletlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyWalletLog))).Methods("DELETE")
	r.Handle("/propertywalletlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyWalletLog))).Methods("POST")

}
