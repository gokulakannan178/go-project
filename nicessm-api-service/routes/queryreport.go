package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//QueryReportRoutes
func (route *Route) QueryReportRoutes(r *mux.Router) {
	r.Handle("/queryreport/filter", Adapt(http.HandlerFunc(route.Handler.FilterQueryReport))).Methods("POST")

}
