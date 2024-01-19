package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PaymentModeRoutes(r *mux.Router) {
	// PaymentMode
	r.Handle("/paymentMode", Adapt(http.HandlerFunc(route.Handler.SavePaymentMode))).Methods("POST")
	r.Handle("/paymentMode", Adapt(http.HandlerFunc(route.Handler.GetSinglePaymentMode))).Methods("GET")
	r.Handle("/paymentMode", Adapt(http.HandlerFunc(route.Handler.UpdatePaymentMode))).Methods("PUT")
	r.Handle("/paymentMode/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePaymentMode))).Methods("PUT")
	r.Handle("/paymentMode/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePaymentMode))).Methods("PUT")
	r.Handle("/paymentMode/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePaymentMode))).Methods("DELETE")
	r.Handle("/paymentMode/filter", Adapt(http.HandlerFunc(route.Handler.FilterPaymentMode))).Methods("POST")
}
