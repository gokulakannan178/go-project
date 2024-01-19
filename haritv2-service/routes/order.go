package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OrderRoutes(r *mux.Router) {

	r.Handle("/order/initiate", Adapt(http.HandlerFunc(route.Handler.InitiateOrder))).Methods("POST")
	r.Handle("/order", Adapt(http.HandlerFunc(route.Handler.GetSingleOrder))).Methods("GET")
	r.Handle("/order/filter", Adapt(http.HandlerFunc(route.Handler.FilterOrder))).Methods("POST")
	r.Handle("/order/placeorder", Adapt(http.HandlerFunc(route.Handler.PlaceOrder))).Methods("PUT")
	r.Handle("/order/cancel", Adapt(http.HandlerFunc(route.Handler.OrderCancel))).Methods("PUT")
	r.Handle("/order/initiateandplaceorder", Adapt(http.HandlerFunc(route.Handler.InitiateAndPlaceOrder))).Methods("POST")
	r.Handle("/order/status/rejected", Adapt(http.HandlerFunc(route.Handler.RejectedOrder))).Methods("PUT")

}
