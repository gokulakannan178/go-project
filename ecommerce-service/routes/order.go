package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OrderRoutes(r *mux.Router) {
	// Order
	r.Handle("/order/initiate", Adapt(http.HandlerFunc(route.Handler.InitiateOrder))).Methods("POST")
	r.Handle("/order", Adapt(http.HandlerFunc(route.Handler.GetSingleOrder))).Methods("GET")
	r.Handle("/order/filter", Adapt(http.HandlerFunc(route.Handler.FilterOrder))).Methods("POST")

}
