package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ShopRentRateMasterRoutes : ""
func (route *Route) ShopRentRateMasterRoutes(r *mux.Router) {
	r.Handle("/shoprentratemaster", Adapt(http.HandlerFunc(route.Handler.SaveShopRentRateMaster))).Methods("POST")
	r.Handle("/shoprentratemaster", Adapt(http.HandlerFunc(route.Handler.GetSingleShopRentRateMaster))).Methods("GET")
	r.Handle("/shoprentratemaster", Adapt(http.HandlerFunc(route.Handler.UpdateShopRentRateMaster))).Methods("PUT")
	r.Handle("/shoprentratemaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableShopRentRateMaster))).Methods("PUT")
	r.Handle("/shoprentratemaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableShopRentRateMaster))).Methods("PUT")
	r.Handle("/shoprentratemaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteShopRentRateMaster))).Methods("DELETE")
	r.Handle("/shoprentratemaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterShopRentRateMaster))).Methods("POST")
}
