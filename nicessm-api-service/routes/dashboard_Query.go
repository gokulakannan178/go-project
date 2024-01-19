package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DashboardQueryCountRoutes(r *mux.Router) {
	r.Handle("/dashboard/query/count", Adapt(http.HandlerFunc(route.Handler.DashboardQueryCount))).Methods("POST")
	r.Handle("/query/report/daywisedemand", Adapt(http.HandlerFunc(route.Handler.DayWiseQueryDemandChart))).Methods("POST")

}
