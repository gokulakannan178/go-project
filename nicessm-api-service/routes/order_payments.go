package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OrderPaymentsRoutes(r *mux.Router) {
	r.Handle("/orderpayments", Adapt(http.HandlerFunc(route.Handler.SaveOrderPayment))).Methods("POST")
	r.Handle("/orderpayments", Adapt(http.HandlerFunc(route.Handler.GetSingleOrderPayment))).Methods("GET")
	r.Handle("/orderpayments", Adapt(http.HandlerFunc(route.Handler.UpdateOrderPayment))).Methods("PUT")
	r.Handle("/orderpayments/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOrderPayment))).Methods("PUT")
	r.Handle("/orderpayments/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOrderPayment))).Methods("PUT")
	r.Handle("/orderpayments/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOrderPayment))).Methods("DELETE")
	r.Handle("/orderpayments/filter", Adapt(http.HandlerFunc(route.Handler.FilterOrderPayment))).Methods("POST")
}
