package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeAttendanceCalendarRoutes : ""
func (route *Route) EmployeeAttendanceCalendarRoutes(r *mux.Router) {
	r.Handle("/employeeattendancecalendar", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeAttendanceCalendar))).Methods("POST")
	r.Handle("/employeeattendancecalendar", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeAttendanceCalendar))).Methods("GET")
	r.Handle("/employeeattendancecalendar", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeAttendanceCalendar))).Methods("PUT")
	r.Handle("/employeeattendancecalendar/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeAttendanceCalendar))).Methods("PUT")
	r.Handle("/employeeattendancecalendar/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeAttendanceCalendar))).Methods("PUT")
	r.Handle("/employeeattendancecalendar/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeAttendanceCalendar))).Methods("DELETE")
	r.Handle("/employeeattendancecalendar/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeAttendanceCalendar))).Methods("POST")
	r.Handle("/employeeattendancecalendar/currentmonth", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeAttendanceCalendarWithCurrentMonth))).Methods("GET")

}
