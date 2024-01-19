package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PaymentRoutes(r *mux.Router) {

	r.Handle("/payment/filter", Adapt(http.HandlerFunc(route.Handler.FilterPayment))).Methods("POST")

}
