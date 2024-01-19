package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeSalaryRoutes : ""
func (route *Route) EmployeeSalaryRoutes(r *mux.Router) {
	r.Handle("/employeesalary", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeSalary))).Methods("POST")
	r.Handle("/employeesalary", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeSalary))).Methods("GET")
	r.Handle("/employeesalary", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeSalary))).Methods("PUT")
	r.Handle("/employeesalary/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeSalary))).Methods("PUT")
	r.Handle("/employeesalary/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeSalary))).Methods("PUT")
	r.Handle("/employeesalary/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeSalary))).Methods("DELETE")
	r.Handle("/employeesalary/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeSalary))).Methods("POST")
	r.Handle("/employee/saveemployeesalary", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeSalaryWithEmployee))).Methods("POST")

}
