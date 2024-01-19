package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DashboardFarmerCountRoutes(r *mux.Router) {
	r.Handle("/dashboard/farmer/count", Adapt(http.HandlerFunc(route.Handler.DashboardFarmerCount))).Methods("POST")
	r.Handle("/farmer/report/daywisedemand", Adapt(http.HandlerFunc(route.Handler.DayWiseFarmerDemandChart))).Methods("POST")

}
