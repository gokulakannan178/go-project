package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OrderPaymentRoutes(r *mux.Router) {
	// OrderPayment
	r.Handle("/orderPayment", Adapt(http.HandlerFunc(route.Handler.SaveOrderPayment))).Methods("POST")
	r.Handle("/orderPayment", Adapt(http.HandlerFunc(route.Handler.GetSingleOrderPayment))).Methods("GET")
	r.Handle("/orderPayment", Adapt(http.HandlerFunc(route.Handler.UpdateOrderPayment))).Methods("PUT")
	r.Handle("/orderPayment/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOrderPayment))).Methods("PUT")
	r.Handle("/orderPayment/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOrderPayment))).Methods("PUT")
	r.Handle("/orderPayment/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOrderPayment))).Methods("DELETE")
	r.Handle("/orderPayment/filter", Adapt(http.HandlerFunc(route.Handler.FilterOrderPayment))).Methods("POST")
	r.Handle("/orderPayment/makepayment", Adapt(http.HandlerFunc(route.Handler.MakePayment))).Methods("POST")
}
