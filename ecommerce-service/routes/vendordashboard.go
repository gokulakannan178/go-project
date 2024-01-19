package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) VendorDashboardRoutes(r *mux.Router) {
	// VendorDashboardRoutes
	r.Handle("/dashboard/vendor/mobile", Adapt(http.HandlerFunc(route.Handler.VendorDashboard))).Methods("POST")
}
