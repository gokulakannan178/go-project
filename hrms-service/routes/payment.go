package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PaymentRoutes : ""
func (route *Route) PaymentRoutes(r *mux.Router) {
	r.Handle("/payment", Adapt(http.HandlerFunc(route.Handler.SavePayment))).Methods("POST")
	r.Handle("/payment", Adapt(http.HandlerFunc(route.Handler.GetSinglePayment))).Methods("GET")
	r.Handle("/payment", Adapt(http.HandlerFunc(route.Handler.UpdatePayment))).Methods("PUT")
	r.Handle("/payment/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePayment))).Methods("PUT")
	r.Handle("/payment/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePayment))).Methods("PUT")
	r.Handle("/payment/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePayment))).Methods("DELETE")
	r.Handle("/payment/filter", Adapt(http.HandlerFunc(route.Handler.FilterPayment))).Methods("POST")
	r.Handle("/payment/pending", Adapt(http.HandlerFunc(route.Handler.PaymentPending))).Methods("POST")
	r.Handle("/payment/accept", Adapt(http.HandlerFunc(route.Handler.PaymentAccept))).Methods("PUT")
	r.Handle("/payment/reject", Adapt(http.HandlerFunc(route.Handler.PaymentReject))).Methods("PUT")
	r.Handle("/payment/received", Adapt(http.HandlerFunc(route.Handler.PaymentReceived))).Methods("PUT")

}
