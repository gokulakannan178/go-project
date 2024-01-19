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
	r.Handle("/product/getdefaultproduct", Adapt(http.HandlerFunc(route.Handler.GetDefaultProduct))).Methods("GET")

}
