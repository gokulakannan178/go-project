package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserChargeDashboardRoutes(r *mux.Router) {
	// DashBoardUserCharge
	r.Handle("/usercharge", Adapt(http.HandlerFunc(route.Handler.SaveUserChargeDashboard))).Methods("POST")
	r.Handle("/usercharge", Adapt(http.HandlerFunc(route.Handler.GetSingleUserChargeDashboard))).Methods("GET")
	r.Handle("/usercharge", Adapt(http.HandlerFunc(route.Handler.UpdateUserChargeDashboard))).Methods("PUT")
	r.Handle("/usercharge/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserChargeDashboard))).Methods("PUT")
	r.Handle("/usercharge/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserChargeDashboard))).Methods("PUT")
	r.Handle("/usercharge/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDashBoardUserCharge))).Methods("DELETE")
	r.Handle("/usercharge/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserChargeDashboard))).Methods("POST")

	r.Handle("/usercharge/dashboard/saf/report", Adapt(http.HandlerFunc(route.Handler.GetUserChargeSAFDashboard))).Methods("POST")
	r.Handle("/usercharge/userwise/report", Adapt(http.HandlerFunc(route.Handler.UserwiseUserChargeReport))).Methods("POST")

}
