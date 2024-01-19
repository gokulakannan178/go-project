package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ProductVariantTypeRoutes(r *mux.Router) {
	// ProductVariantType
	r.Handle("/productVariantType", Adapt(http.HandlerFunc(route.Handler.SaveProductVariantType))).Methods("POST")
	r.Handle("/productVariantType", Adapt(http.HandlerFunc(route.Handler.GetSingleProductVariantType))).Methods("GET")
	r.Handle("/productVariantType", Adapt(http.HandlerFunc(route.Handler.UpdateProductVariantType))).Methods("PUT")
	r.Handle("/productVariantType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProductVariantType))).Methods("PUT")
	r.Handle("/productVariantType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProductVariantType))).Methods("PUT")
	r.Handle("/productVariantType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProductVariantType))).Methods("DELETE")
	r.Handle("/productVariantType/filter", Adapt(http.HandlerFunc(route.Handler.FilterProductVariantType))).Methods("POST")
}
