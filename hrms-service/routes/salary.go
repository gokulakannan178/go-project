package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SalaryRoutes : ""
func (route *Route) SalaryRoutes(r *mux.Router) {
	r.Handle("/salary", Adapt(http.HandlerFunc(route.Handler.SaveSalary))).Methods("POST")
	r.Handle("/salary", Adapt(http.HandlerFunc(route.Handler.GetSingleSalary))).Methods("GET")
	r.Handle("/salary", Adapt(http.HandlerFunc(route.Handler.UpdateSalary))).Methods("PUT")
	r.Handle("/salary/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSalary))).Methods("PUT")
	r.Handle("/salary/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSalary))).Methods("PUT")
	r.Handle("/salary/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSalary))).Methods("DELETE")
	r.Handle("/salary/filter", Adapt(http.HandlerFunc(route.Handler.FilterSalary))).Methods("POST")

}
