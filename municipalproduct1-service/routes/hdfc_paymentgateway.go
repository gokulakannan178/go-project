package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HDFCPaymentGatewayRoutes : ""
func (route *Route) HDFCPaymentGatewayRoutes(r *mux.Router) {
	r.Handle("/hdfc/paymentgateway", Adapt(http.HandlerFunc(route.Handler.SaveHDFCPaymentGateway))).Methods("POST")
	r.Handle("/hdfc/paymentgateway/merchantid", Adapt(http.HandlerFunc(route.Handler.GetSingleMerchantIDHDFCPaymentGateway))).Methods("GET")
	r.Handle("/hdfc/paymentgateway", Adapt(http.HandlerFunc(route.Handler.GetSingleDefaultHDFCPaymentGateway))).Methods("GET")
	r.Handle("/hdfc/paymentgateway", Adapt(http.HandlerFunc(route.Handler.GetSingleDefaultHDFCPaymentGateway))).Methods("POST")

}
