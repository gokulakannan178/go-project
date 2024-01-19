package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//DashboardRoutes : ""
func (route *Route) DashboardRoutes(r *mux.Router) {

	r.Handle("/dashboard/count", Adapt(http.HandlerFunc(route.Handler.GetCollectionCount))).Methods("POST")
	r.Handle("/dashboard/housevisited/count", Adapt(http.HandlerFunc(route.Handler.GetHousevisitedCount))).Methods("POST")
	r.Handle("/dashboard/vehicle/count", Adapt(http.HandlerFunc(route.Handler.GetvehicleCount))).Methods("POST")
	r.Handle("/dashboard/dumbsite/count", Adapt(http.HandlerFunc(route.Handler.GetDumbSiteCount))).Methods("POST")
	r.Handle("/dashboard/usertype/count", Adapt(http.HandlerFunc(route.Handler.GetUsertypeCount))).Methods("POST")
	r.Handle("/dashboard/property/count", Adapt(http.HandlerFunc(route.Handler.GetPropertyCount))).Methods("POST")
	r.Handle("/dashboard/garbagge/count", Adapt(http.HandlerFunc(route.Handler.GetGarbaggeCount))).Methods("POST")
}
