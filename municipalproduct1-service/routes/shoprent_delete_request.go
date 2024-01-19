package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ShopRentDeleteRequestRoutes : ""
func (route *Route) ShopRentDeleteRequestRoutes(r *mux.Router) {
	// ShopRentDelete Request Update
	r.Handle("/shoprentdeleterequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateShopRentDeleteRequest))).Methods("PUT")
	r.Handle("/shoprentdeleterequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptShopRentDeleteRequestUpdate))).Methods("PUT")
	r.Handle("/shoprentdeleterequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectShopRentDeleteRequestUpdate))).Methods("PUT")
	r.Handle("/shoprentdeleterequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterShopRentDeleteRequest))).Methods("POST")
	r.Handle("/shoprentdeleterequest", Adapt(http.HandlerFunc(route.Handler.GetSingleShopRentDeleteRequest))).Methods("GET")

	//
}
