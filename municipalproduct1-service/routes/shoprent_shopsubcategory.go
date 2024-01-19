package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ShopRentShopSubCategoryRoutes : ""
func (route *Route) ShopRentShopSubCategoryRoutes(r *mux.Router) {

	r.Handle("/shoprentshopsubcategory", Adapt(http.HandlerFunc(route.Handler.SaveShopRentShopSubCategory))).Methods("POST")
	r.Handle("/shoprentshopsubcategory", Adapt(http.HandlerFunc(route.Handler.GetSingleShopRentShopSubCategory))).Methods("GET")
	r.Handle("/shoprentshopsubcategory", Adapt(http.HandlerFunc(route.Handler.UpdateShopRentShopSubCategory))).Methods("PUT")
	r.Handle("/shoprentshopsubcategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableShopRentShopSubCategory))).Methods("PUT")
	r.Handle("/shoprentshopsubcategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableShopRentShopSubCategory))).Methods("PUT")
	r.Handle("/shoprentshopsubcategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteShopRentShopSubCategory))).Methods("DELETE")
	r.Handle("/shoprentshopsubcategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterShopRentShopSubCategory))).Methods("POST")
}
