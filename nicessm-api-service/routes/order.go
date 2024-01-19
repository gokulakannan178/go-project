package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OrderRoutes(r *mux.Router) {
	r.Handle("/order", Adapt(http.HandlerFunc(route.Handler.SaveOrder))).Methods("POST")
	r.Handle("/order", Adapt(http.HandlerFunc(route.Handler.GetSingleOrder))).Methods("GET")
	r.Handle("/order", Adapt(http.HandlerFunc(route.Handler.UpdateOrder))).Methods("PUT")
	r.Handle("/order/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOrder))).Methods("PUT")
	r.Handle("/order/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOrder))).Methods("PUT")
	r.Handle("/order/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOrder))).Methods("DELETE")
	r.Handle("/order/filter", Adapt(http.HandlerFunc(route.Handler.FilterOrder))).Methods("POST")
	r.Handle("/order/create", Adapt(http.HandlerFunc(route.Handler.CreateOrder))).Methods("POST")
}
