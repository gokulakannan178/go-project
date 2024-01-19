package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TradeLicenseDashboardRoutes(r *mux.Router) {
	// DashBoardTradeLicense
	r.Handle("/dashboard/tradelicense", Adapt(http.HandlerFunc(route.Handler.SaveTradeLicenseDashboard))).Methods("POST")
	r.Handle("/dashboard/tradelicense", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseDashboard))).Methods("GET")
	r.Handle("/dashboard/tradelicense", Adapt(http.HandlerFunc(route.Handler.UpdateTradeLicenseDashboard))).Methods("PUT")
	r.Handle("/dashboard/tradelicense/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTradeLicenseDashboard))).Methods("PUT")
	r.Handle("/dashboard/tradelicense/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTradeLicenseDashboard))).Methods("PUT")
	r.Handle("/dashboard/tradelicense/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDashBoardTradeLicense))).Methods("DELETE")
	r.Handle("/dashboard/tradelicense/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicenseDashboard))).Methods("POST")

	r.Handle("/tradelicense/dashboard/ddac", Adapt(http.HandlerFunc(route.Handler.DashboardTradeLicenseDemandAndCollection))).Methods("POST")
	r.Handle("/tradelicense/userwise/report", Adapt(http.HandlerFunc(route.Handler.UserwiseTradelicenseReport))).Methods("POST")

}
