package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DashboardUserCountRoutes(r *mux.Router) {
	r.Handle("/dashboard/user/count", Adapt(http.HandlerFunc(route.Handler.DashboardUserCount))).Methods("POST")
	r.Handle("/user/report/daywisedemand", Adapt(http.HandlerFunc(route.Handler.DayWiseUserDemandChart))).Methods("POST")

}
