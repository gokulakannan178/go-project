package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DashboardPropertyDayWiseRoutes(r *mux.Router) {
	// DashboardPropertyDayWise
	r.Handle("/propertydashboarddaywise", Adapt(http.HandlerFunc(route.Handler.SaveDashBoardPropertyDayWiseV2))).Methods("POST")
	r.Handle("/propertydashboarddaywise", Adapt(http.HandlerFunc(route.Handler.GetSingleDashBoardPropertyDayWiseV2))).Methods("GET")
	r.Handle("/propertydashboarddaywise", Adapt(http.HandlerFunc(route.Handler.UpdateDashBoardPropertyDayWiseV2))).Methods("PUT")
	r.Handle("/propertydashboarddaywise/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDashBoardPropertyDayWiseV2))).Methods("PUT")
	r.Handle("/propertydashboarddaywise/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDashBoardPropertyDayWiseV2))).Methods("PUT")
	r.Handle("/propertydashboarddaywise/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDashBoardPropertyDayWiseV2))).Methods("DELETE")
	r.Handle("/propertydashboarddaywise/filter", Adapt(http.HandlerFunc(route.Handler.FilterDashBoardPropertyDayWiseV2))).Methods("POST")
}
