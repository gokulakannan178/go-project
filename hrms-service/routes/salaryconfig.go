package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SalaryConfigRoutes : ""
func (route *Route) SalaryConfigRoutes(r *mux.Router) {
	r.Handle("/salaryconfig", Adapt(http.HandlerFunc(route.Handler.SaveSalaryConfig))).Methods("POST")
	r.Handle("/salaryconfig", Adapt(http.HandlerFunc(route.Handler.GetSingleSalaryConfig))).Methods("GET")
	r.Handle("/salaryconfig", Adapt(http.HandlerFunc(route.Handler.UpdateSalaryConfig))).Methods("PUT")
	r.Handle("/salaryconfig/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSalaryConfig))).Methods("PUT")
	r.Handle("/salaryconfig/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSalaryConfig))).Methods("PUT")
	r.Handle("/salaryconfig/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSalaryConfig))).Methods("DELETE")
	r.Handle("/salaryconfig/filter", Adapt(http.HandlerFunc(route.Handler.FilterSalaryConfig))).Methods("POST")
	r.Handle("/savesalaryconfig/employeetype", Adapt(http.HandlerFunc(route.Handler.SaveSalaryConfigWithEmployeeType))).Methods("POST")

}
