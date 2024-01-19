package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//DashboardRoutes : ""
func (route *Route) DashboardRoutes(r *mux.Router) {

	r.Handle("/dashboard/count", Adapt(http.HandlerFunc(route.Handler.GetCollectionCount))).Methods("GET")
}
