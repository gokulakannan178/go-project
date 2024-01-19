package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DashboardPropertyRoutes(r *mux.Router) {
	// DashboardProperty
	r.Handle("/propertydashboard", Adapt(http.HandlerFunc(route.Handler.SaveDashBoardProperty))).Methods("POST")
	r.Handle("/propertydashboard", Adapt(http.HandlerFunc(route.Handler.GetSingleDashBoardProperty))).Methods("GET")
	r.Handle("/propertydashboard", Adapt(http.HandlerFunc(route.Handler.UpdateDashBoardProperty))).Methods("PUT")
	r.Handle("/propertydashboard/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDashBoardProperty))).Methods("PUT")
	r.Handle("/propertydashboard/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDashBoardProperty))).Methods("PUT")
	r.Handle("/propertydashboard/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDashBoardProperty))).Methods("DELETE")
	r.Handle("/propertydashboard/filter", Adapt(http.HandlerFunc(route.Handler.FilterDashBoardProperty))).Methods("POST")
	r.Handle("/getoverallpropertydashBoard", Adapt(http.HandlerFunc(route.Handler.GetOverAllPropertyDashBoard))).Methods("GET")
}
