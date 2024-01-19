package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ProductCategoryRoutes(r *mux.Router) {
	// Product Category
	r.Handle("/productcategory", Adapt(http.HandlerFunc(route.Handler.SaveProductCategory))).Methods("POST")
	r.Handle("/productcategory", Adapt(http.HandlerFunc(route.Handler.GetSingleProductCategory))).Methods("GET")
	r.Handle("/productcategory", Adapt(http.HandlerFunc(route.Handler.UpdateProductCategory))).Methods("PUT")
	r.Handle("/productcategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProductCategory))).Methods("PUT")
	r.Handle("/productcategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProductCategory))).Methods("PUT")
	r.Handle("/productcategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProductCategory))).Methods("DELETE")
	r.Handle("/productcategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterProductCategory))).Methods("POST")
	r.Handle("/productcategory/getdefaultproductcategory", Adapt(http.HandlerFunc(route.Handler.GetDefaultProductCategory))).Methods("GET")

}
