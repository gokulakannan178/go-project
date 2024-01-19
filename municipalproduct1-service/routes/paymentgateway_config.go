package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PaymentGatewayConfigurationRoutes : ""
func (route *Route) PaymentGatewayConfigurationRoutes(r *mux.Router) {
	r.Handle("/paymentgateway/config", Adapt(http.HandlerFunc(route.Handler.SavePaymentGateway))).Methods("POST")
	r.Handle("/paymentgateway/config", Adapt(http.HandlerFunc(route.Handler.GetSinglePaymentGateway))).Methods("GET")
	r.Handle("/paymentgateway/config", Adapt(http.HandlerFunc(route.Handler.UpdatePaymentGateway))).Methods("PUT")
	r.Handle("/paymentgateway/config/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePaymentGateway))).Methods("PUT")
	r.Handle("/paymentgateway/config/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePaymentGateway))).Methods("PUT")
	r.Handle("/paymentgateway/config/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePaymentGateway))).Methods("DELETE")
	r.Handle("/paymentgateway/config/filter", Adapt(http.HandlerFunc(route.Handler.FilterPaymentGateway))).Methods("POST")

}
