package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//AttendanceRoutes : ""
func (route *Route) AttendanceRoutes(r *mux.Router) {
	r.Handle("/attendance", Adapt(http.HandlerFunc(route.Handler.SaveAttendance))).Methods("POST")
	r.Handle("/attendance/employeeedit", Adapt(http.HandlerFunc(route.Handler.SaveAttendanceWithEditEmployee))).Methods("POST")
	r.Handle("/attendance", Adapt(http.HandlerFunc(route.Handler.GetSingleAttendance))).Methods("GET")
	r.Handle("/attendance", Adapt(http.HandlerFunc(route.Handler.UpdateAttendance))).Methods("PUT")
	r.Handle("/attendance/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAttendance))).Methods("PUT")
	r.Handle("/attendance/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAttendance))).Methods("PUT")
	r.Handle("/attendance/filter", Adapt(http.HandlerFunc(route.Handler.FilterAttendance))).Methods("POST")
	r.Handle("/attendance/clockin", Adapt(http.HandlerFunc(route.Handler.ClockinAttendance))).Methods("POST")
	r.Handle("/attendance/clockout", Adapt(http.HandlerFunc(route.Handler.ClockoutAttendance))).Methods("PUT")
	r.Handle("/attendance/employeeleave/request", Adapt(http.HandlerFunc(route.Handler.AttendanceEmployeeLeaveRequest))).Methods("POST")
	//r.Handle("/attendance/employeeleave/approve", Adapt(http.HandlerFunc(route.Handler.AttendanceEmployeeLeaveApprove))).Methods("PUT")
	//	r.Handle("/attendance/employeetodaystatus", Adapt(http.HandlerFunc(route.Handler.EmployeeAttendanceTodayStatus))).Methods("GET")
	//r.Handle("/attendance/employeeleave/reject", Adapt(http.HandlerFunc(route.Handler.AttendanceEmployeeLeaveReject))).Methods("PUT")
	r.Handle("/attendance/employee/daywisereport", Adapt(http.HandlerFunc(route.Handler.DayWiseAttendanceReport))).Methods("POST")
	r.Handle("/attendancelog/employee/statistics", Adapt(http.HandlerFunc(route.Handler.AttendanceEmployeeStatistics))).Methods("GET")
	r.Handle("/attendance/todayemployees/leave", Adapt(http.HandlerFunc(route.Handler.TodayEmployessLeave))).Methods("GET")
	r.Handle("/attendance/approved", Adapt(http.HandlerFunc(route.Handler.EmployeeAttendanceApprove))).Methods("PUT")
	r.Handle("/attendance/approved/allemployee", Adapt(http.HandlerFunc(route.Handler.AllEmployeeAttendanceApprove))).Methods("PUT")
	r.Handle("/attendance/rejected/allemployee", Adapt(http.HandlerFunc(route.Handler.AllEmployeeAttendanceReject))).Methods("PUT")
	r.Handle("/attendance/rejected", Adapt(http.HandlerFunc(route.Handler.EmployeeAttendanceRejected))).Methods("PUT")
	r.Handle("/attendance/todaytimeoff", Adapt(http.HandlerFunc(route.Handler.GetTodayEmployeeTimeOff))).Methods("GET")
	r.Handle("/attendance/todayuplannedleave", Adapt(http.HandlerFunc(route.Handler.GetTodayEmployeeUplannedLeave))).Methods("GET")
	r.Handle("/attendance/todayPunchin", Adapt(http.HandlerFunc(route.Handler.GetTodayEmployeePunchin))).Methods("GET")
	r.Handle("/attendance/todayabsent", Adapt(http.HandlerFunc(route.Handler.GetTodayEmployeeAbsent))).Methods("POST")

}
