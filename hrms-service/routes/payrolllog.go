package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PayrollLogRoutes(r *mux.Router) {
	// PayrollLog
	r.Handle("/payrollLog", Adapt(http.HandlerFunc(route.Handler.SavePayrollLog))).Methods("POST")
	r.Handle("/payrollLog", Adapt(http.HandlerFunc(route.Handler.GetSinglePayrollLog))).Methods("GET")
	r.Handle("/payrollLog", Adapt(http.HandlerFunc(route.Handler.UpdatePayrollLog))).Methods("PUT")
	r.Handle("/payrollLog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePayrollLog))).Methods("PUT")
	r.Handle("/payrollLog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePayrollLog))).Methods("PUT")
	r.Handle("/payrollLog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePayrollLog))).Methods("DELETE")
	r.Handle("/payrollLog/filter", Adapt(http.HandlerFunc(route.Handler.FilterPayrollLog))).Methods("POST")
	r.Handle("/payroll/employeesalaryslip", Adapt(http.HandlerFunc(route.Handler.GetEmployeeSalarySlip))).Methods("POST")
	r.Handle("/payroll/list", Adapt(http.HandlerFunc(route.Handler.PayrollLogList))).Methods("POST")

}
