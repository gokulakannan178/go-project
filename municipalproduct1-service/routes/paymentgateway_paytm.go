package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PaymentGatewayPaytmRoutes : ""
func (route *Route) PaymentGatewayPaytmRoutes(r *mux.Router) {
	r.Handle("/paymentgateway/patym/initiatetransaction", Adapt(http.HandlerFunc(route.Handler.PaytmPaymentInitiateTransaction))).Methods("POST")
	r.Handle("/paymentgateway/patym/qrcode/initiatetransaction", Adapt(http.HandlerFunc(route.Handler.PaytmtQrCodeInitTranscation))).Methods("POST")
	r.Handle("/paymentgateway/patym/createchecksum", Adapt(http.HandlerFunc(route.Handler.PaytmPaymentChecksum))).Methods("POST")
	r.Handle("/paymentgateway/default", Adapt(http.HandlerFunc(route.Handler.GetDefaultPaymentGateway))).Methods("GET")

}
