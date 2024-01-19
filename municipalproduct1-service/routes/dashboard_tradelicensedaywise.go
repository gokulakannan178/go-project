package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TradeLicenseDashboardDayWiseRoutes(r *mux.Router) {
	// DashBoardTradeLicenseDayWise
	r.Handle("/tradelicensedaywise", Adapt(http.HandlerFunc(route.Handler.SaveTradeLicenseDashboardDayWise))).Methods("POST")
	r.Handle("/tradelicensedaywise", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseDashboardDayWise))).Methods("GET")
	r.Handle("/tradelicensedaywise", Adapt(http.HandlerFunc(route.Handler.UpdateTradeLicenseDashboardDayWise))).Methods("PUT")
	r.Handle("/tradelicensedaywise/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTradeLicenseDashboardDayWise))).Methods("PUT")
	r.Handle("/tradelicensedaywise/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTradeLicenseDashboardDayWise))).Methods("PUT")
	r.Handle("/tradelicensedaywise/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDashBoardTradeLicenseDayWise))).Methods("DELETE")
	r.Handle("/tradelicensedaywise/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicenseDashboardDayWise))).Methods("POST")
}
