package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ShopRentDashboardRoutes(r *mux.Router) {
	// DashBoardShopRent
	r.Handle("/shoprentdashboard", Adapt(http.HandlerFunc(route.Handler.SaveShopRentDashboard))).Methods("POST")
	r.Handle("/shoprentdashboard", Adapt(http.HandlerFunc(route.Handler.GetSingleShopRentDashboard))).Methods("GET")
	r.Handle("/shoprentdashboard", Adapt(http.HandlerFunc(route.Handler.UpdateShopRentDashboard))).Methods("PUT")
	r.Handle("/shoprentdashboard/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableShopRentDashboard))).Methods("PUT")
	r.Handle("/shoprentdashboard/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableShopRentDashboard))).Methods("PUT")
	r.Handle("/shoprentdashboard/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteShopRentDashboard))).Methods("DELETE")
	r.Handle("/shoprentdashboard/filter", Adapt(http.HandlerFunc(route.Handler.FilterShopRentDashboard))).Methods("POST")

	r.Handle("/shoprent/dashboard/ddac", Adapt(http.HandlerFunc(route.Handler.DashboardShopRentDemandAndCollection))).Methods("POST")

}
