package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) EmployeeReportRoutes(r *mux.Router) {
	// Employee
	r.Handle("/employeereport", Adapt(http.HandlerFunc(route.Handler.EmployeeReport))).Methods("POST")

}
