package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ProductVariantRoutes(r *mux.Router) {
	// ProductVariant
	r.Handle("/productvariant", Adapt(http.HandlerFunc(route.Handler.SaveProductVariant))).Methods("POST")
	r.Handle("/productvariant", Adapt(http.HandlerFunc(route.Handler.GetSingleProductVariant))).Methods("GET")
	r.Handle("/productvariant", Adapt(http.HandlerFunc(route.Handler.UpdateProductVariant))).Methods("PUT")
	r.Handle("/productvariant/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProductVariant))).Methods("PUT")
	r.Handle("/productvariant/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProductVariant))).Methods("PUT")
	r.Handle("/productvariant/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProductVariant))).Methods("DELETE")
	r.Handle("/productvariant/filter", Adapt(http.HandlerFunc(route.Handler.FilterProductVariant))).Methods("POST")
	r.Handle("/productvariant/register", Adapt(http.HandlerFunc(route.Handler.ProductVariantRegister))).Methods("POST")
	r.Handle("/productvariant/getmyinventory", Adapt(http.HandlerFunc(route.Handler.GetMyInventory))).Methods("GET")
}
