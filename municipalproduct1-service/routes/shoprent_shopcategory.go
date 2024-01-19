package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ShopRentShopCategoryRoutes : ""
func (route *Route) ShopRentShopCategoryRoutes(r *mux.Router) {

	r.Handle("/shoprentshopcategory", Adapt(http.HandlerFunc(route.Handler.SaveShopRentShopCategory))).Methods("POST")
	r.Handle("/shoprentshopcategory", Adapt(http.HandlerFunc(route.Handler.GetSingleShopRentShopCategory))).Methods("GET")
	r.Handle("/shoprentshopcategory", Adapt(http.HandlerFunc(route.Handler.UpdateShopRentShopCategory))).Methods("PUT")
	r.Handle("/shoprentshopcategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableShopRentShopCategory))).Methods("PUT")
	r.Handle("/shoprentshopcategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableShopRentShopCategory))).Methods("PUT")
	r.Handle("/shoprentshopcategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteShopRentShopCategory))).Methods("DELETE")
	r.Handle("/shoprentshopcategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterShopRentShopCategory))).Methods("POST")
}
