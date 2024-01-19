package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ProductRoutes(r *mux.Router) {
	// Product
	r.Handle("/product", Adapt(http.HandlerFunc(route.Handler.SaveProduct))).Methods("POST")
	r.Handle("/product", Adapt(http.HandlerFunc(route.Handler.GetSingleProduct))).Methods("GET")
	r.Handle("/product", Adapt(http.HandlerFunc(route.Handler.UpdateProduct))).Methods("PUT")
	r.Handle("/product/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProduct))).Methods("PUT")
	r.Handle("/product/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProduct))).Methods("PUT")
	r.Handle("/product/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProduct))).Methods("DELETE")
	r.Handle("/product/filter", Adapt(http.HandlerFunc(route.Handler.FilterProduct))).Methods("POST")
}

func (route *Route) ProductRegisterRoutes(r *mux.Router) {

	r.Handle("/productregister", Adapt(http.HandlerFunc(route.Handler.SaveRegisterProduct))).Methods("POST")
}
