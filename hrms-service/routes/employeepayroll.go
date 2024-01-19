package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeePayrollRoutes : ""
func (route *Route) EmployeePayrollRoutes(r *mux.Router) {
	r.Handle("/employeepayroll", Adapt(http.HandlerFunc(route.Handler.SaveEmployeePayroll))).Methods("POST")
	r.Handle("/employeepayroll", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeePayroll))).Methods("GET")
	r.Handle("/employeepayroll", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeePayroll))).Methods("PUT")
	r.Handle("/employeepayroll/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeePayroll))).Methods("PUT")
	r.Handle("/employeepayroll/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeePayroll))).Methods("PUT")
	r.Handle("/employeepayroll/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeePayroll))).Methods("DELETE")
	r.Handle("/employeepayroll/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeePayroll))).Methods("POST")
	r.Handle("/employeepayroll/updatepayroll", Adapt(http.HandlerFunc(route.Handler.EmployeeUpdatePayrollWithNewPayroll))).Methods("PUT")
	r.Handle("/employeepayroll/withearningdeduction", Adapt(http.HandlerFunc(route.Handler.SaveEmployeePayrollWithEaringDeduction))).Methods("POST")

}
