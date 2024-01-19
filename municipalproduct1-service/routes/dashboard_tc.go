package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ProductConfigurationRoutes : ""
func (route *Route) TcDashboardRoutes(r *mux.Router) {
	r.Handle("/tcdashboard", Adapt(http.HandlerFunc(route.Handler.TcDashboard))).Methods("POST")
	r.Handle("/pmdashboard", Adapt(http.HandlerFunc(route.Handler.PmDashboard))).Methods("POST")

}
