package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserChargeDayWiseDashboardRoutes(r *mux.Router) {
	// DashBoardUserCharge
	r.Handle("/userchargedaywise", Adapt(http.HandlerFunc(route.Handler.SaveUserChargeDayWiseDashboard))).Methods("POST")
	r.Handle("/userchargedaywise", Adapt(http.HandlerFunc(route.Handler.GetSingleUserChargeDayWiseDashboard))).Methods("GET")
	r.Handle("/userchargedaywise", Adapt(http.HandlerFunc(route.Handler.UpdateUserChargeDayWiseDashboard))).Methods("PUT")
	r.Handle("/userchargedaywise/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserChargeDayWiseDashboard))).Methods("PUT")
	r.Handle("/userchargedaywise/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserChargeDayWiseDashboard))).Methods("PUT")
	r.Handle("/userchargedaywise/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDashBoardUserCharge))).Methods("DELETE")
	r.Handle("/userchargedaywise/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserChargeDayWiseDashboard))).Methods("POST")
}
