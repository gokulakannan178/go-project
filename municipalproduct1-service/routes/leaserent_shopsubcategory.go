package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// LeaseRentShopSubCategoryRoutes : ""
func (route *Route) LeaseRentShopSubCategoryRoutes(r *mux.Router) {

	r.Handle("/leaserentshopsubcategory", Adapt(http.HandlerFunc(route.Handler.SaveLeaseRentShopSubCategory))).Methods("POST")
	r.Handle("/leaserentshopsubcategory", Adapt(http.HandlerFunc(route.Handler.GetSingleLeaseRentShopSubCategory))).Methods("GET")
	r.Handle("/leaserentshopsubcategory", Adapt(http.HandlerFunc(route.Handler.UpdateLeaseRentShopSubCategory))).Methods("PUT")
	r.Handle("/leaserentshopsubcategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLeaseRentShopSubCategory))).Methods("PUT")
	r.Handle("/leaserentshopsubcategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLeaseRentShopSubCategory))).Methods("PUT")
	r.Handle("/leaserentshopsubcategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLeaseRentShopSubCategory))).Methods("DELETE")
	r.Handle("/leaserentshopsubcategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterLeaseRentShopSubCategory))).Methods("POST")
}
