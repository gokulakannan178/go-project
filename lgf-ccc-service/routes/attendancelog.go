package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//AttendanceLogRoutes : ""
func (route *Route) AttendanceLogRoutes(r *mux.Router) {
	r.Handle("/attendancelog", Adapt(http.HandlerFunc(route.Handler.SaveAttendanceLog))).Methods("POST")
	r.Handle("/attendancelog", Adapt(http.HandlerFunc(route.Handler.GetSingleAttendanceLog))).Methods("GET")
	r.Handle("/attendancelog/employeeid", Adapt(http.HandlerFunc(route.Handler.GetSingleAttendanceLoglast))).Methods("GET")
	r.Handle("/attendancelog/employee/todaystatus", Adapt(http.HandlerFunc(route.Handler.AttendanceEmployeeTodayStatus))).Methods("GET")
	r.Handle("/attendancelog", Adapt(http.HandlerFunc(route.Handler.UpdateAttendanceLog))).Methods("PUT")
	r.Handle("/attendancelog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAttendanceLog))).Methods("PUT")
	r.Handle("/attendancelog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAttendanceLog))).Methods("PUT")
	r.Handle("/attendancelog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAttendanceLog))).Methods("DELETE")
	r.Handle("/attendancelog/filter", Adapt(http.HandlerFunc(route.Handler.FilterAttendanceLog))).Methods("POST")
	r.Handle("/attendancelog/employee/todaylogs", Adapt(http.HandlerFunc(route.Handler.AttendanceEmployeeTodayLogs))).Methods("GET")
}
