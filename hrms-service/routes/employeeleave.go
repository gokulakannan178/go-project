package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeLeaveRoutes : ""
func (route *Route) EmployeeLeaveRoutes(r *mux.Router) {
	r.Handle("/employeeleave", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeLeave))).Methods("POST")
	r.Handle("/employeeleave", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeLeave))).Methods("GET")
	r.Handle("/employeeleave", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeLeave))).Methods("PUT")
	r.Handle("/employeeleave/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeLeave))).Methods("PUT")
	r.Handle("/employeeleave/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeLeave))).Methods("PUT")
	r.Handle("/employeeleave/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeLeave))).Methods("DELETE")
	r.Handle("/employeeleave/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeLeave))).Methods("POST")
	r.Handle("/employeeleave/count", Adapt(http.HandlerFunc(route.Handler.GetEmployeeLeaveCount))).Methods("POST")
	r.Handle("/employeeleave/updatevalue", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeLeaveFromTimeOff))).Methods("PUT")
	r.Handle("/employeeleave/list", Adapt(http.HandlerFunc(route.Handler.EmployeeLeaveList))).Methods("POST")
	r.Handle("/employeeleave/list/v2", Adapt(http.HandlerFunc(route.Handler.EmployeeLeaveListV2))).Methods("POST")
	r.Handle("/allemployeeleave/list", Adapt(http.HandlerFunc(route.Handler.GetAllEmployeeLeaveList))).Methods("POST")
	r.Handle("/employeeleave/updateleave", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeLeaveWithEmployeeId))).Methods("PUT")

}
