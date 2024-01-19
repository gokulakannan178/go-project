package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ShoprentPaymentModeChangeRoutes : ""
func (route *Route) ShoprentPaymentModeChangeRoutes(r *mux.Router) {
	r.Handle("/shoprentpayment/modechange/update", Adapt(http.HandlerFunc(route.Handler.SaveShoprentPaymentModeChange))).Methods("PUT")
	r.Handle("/shoprentpayment/modechange", Adapt(http.HandlerFunc(route.Handler.GetSingleShoprentPaymentModeChange))).Methods("GET")
	r.Handle("/shoprentpayment/modechange/update/accept", Adapt(http.HandlerFunc(route.Handler.AcceptShoprentPaymentModeChange))).Methods("PUT")
	r.Handle("/shoprentpayment/modechange/update/reject", Adapt(http.HandlerFunc(route.Handler.RejectShoprentPaymentModeChange))).Methods("PUT")
	r.Handle("/shoprentpayment/modechange/filter", Adapt(http.HandlerFunc(route.Handler.FilterShoprentPaymentModeChange))).Methods("POST")
}
