package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ShopRentDashboardDayWiseRoutes(r *mux.Router) {
	// DashBoardShopRentDayWise
	r.Handle("/shoprentdaywise", Adapt(http.HandlerFunc(route.Handler.SaveShopRentDashboardDayWise))).Methods("POST")
	r.Handle("/shoprentdaywise", Adapt(http.HandlerFunc(route.Handler.GetSingleShopRentDashboardDayWise))).Methods("GET")
	r.Handle("/shoprentdaywise", Adapt(http.HandlerFunc(route.Handler.UpdateShopRentDashboardDayWise))).Methods("PUT")
	r.Handle("/shoprentdaywise/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableShopRentDashboardDayWise))).Methods("PUT")
	r.Handle("/shoprentdaywise/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableShopRentDashboardDayWise))).Methods("PUT")
	r.Handle("/shoprentdaywise/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteShopRentDashboardDayWise))).Methods("DELETE")
	r.Handle("/shoprentdaywise/filter", Adapt(http.HandlerFunc(route.Handler.FilterShopRentDashboardDayWise))).Methods("POST")
}
