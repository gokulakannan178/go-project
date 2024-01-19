package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// LeaseRentShopCategoryRoutes : ""
func (route *Route) LeaseRentShopCategoryRoutes(r *mux.Router) {

	r.Handle("/leaserentshopcategory", Adapt(http.HandlerFunc(route.Handler.SaveLeaseRentShopCategory))).Methods("POST")
	r.Handle("/leaserentshopcategory", Adapt(http.HandlerFunc(route.Handler.GetSingleLeaseRentShopCategory))).Methods("GET")
	r.Handle("/leaserentshopcategory", Adapt(http.HandlerFunc(route.Handler.UpdateLeaseRentShopCategory))).Methods("PUT")
	r.Handle("/leaserentshopcategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLeaseRentShopCategory))).Methods("PUT")
	r.Handle("/leaserentshopcategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLeaseRentShopCategory))).Methods("PUT")
	r.Handle("/leaserentshopcategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLeaseRentShopCategory))).Methods("DELETE")
	r.Handle("/leaserentshopcategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterLeaseRentShopCategory))).Methods("POST")
}
