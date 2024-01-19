package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) EmployeeRoutes(r *mux.Router) {
	// Employee
	r.Handle("/employee", Adapt(http.HandlerFunc(route.Handler.SaveEmployee))).Methods("POST")
	r.Handle("/employee", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployee))).Methods("GET")
	r.Handle("/employee", Adapt(http.HandlerFunc(route.Handler.UpdateEmployee))).Methods("PUT")
	r.Handle("/employee/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployee))).Methods("PUT")
	r.Handle("/employee/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployee))).Methods("PUT")
	r.Handle("/employee/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployee))).Methods("DELETE")
	r.Handle("/employee/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployee))).Methods("POST")
	r.Handle("/employee/reject", Adapt(http.HandlerFunc(route.Handler.EmployeeReject))).Methods("PUT")
	r.Handle("/employee/onboarding", Adapt(http.HandlerFunc(route.Handler.EmployeeOnboarding))).Methods("PUT")
	r.Handle("/employee/probationary", Adapt(http.HandlerFunc(route.Handler.EmployeeProbationary))).Methods("PUT")
	r.Handle("/employee/activeemployee", Adapt(http.HandlerFunc(route.Handler.EmployeeActive))).Methods("PUT")
	r.Handle("/employee/bench", Adapt(http.HandlerFunc(route.Handler.EmployeeBench))).Methods("PUT")
	r.Handle("/employee/notice", Adapt(http.HandlerFunc(route.Handler.EmployeeNotice))).Methods("PUT")
	r.Handle("/employee/offboard", Adapt(http.HandlerFunc(route.Handler.EmployeeOffboard))).Methods("PUT")
	r.Handle("/employee/relieve", Adapt(http.HandlerFunc(route.Handler.EmployeeRelieve))).Methods("PUT")
	r.Handle("/employee/updateBioData", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeBioData))).Methods("PUT")
	r.Handle("/employee/updateprofileimage", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeProfileImage))).Methods("PUT")
	r.Handle("/employee/updateEmergencyContact", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeEmergencyContact))).Methods("PUT")
	r.Handle("/employee/updatePersonalInformation", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeePersonalInformation))).Methods("PUT")
	r.Handle("/attendance/employeedaywisereport", Adapt(http.HandlerFunc(route.Handler.EmployeeDayWiseAttendanceReport))).Methods("POST")
	r.Handle("/employee/upload", Adapt(http.HandlerFunc(route.Handler.EmployeeUpload))).Methods("POST")
	r.Handle("/employee/upload/v2", Adapt(http.HandlerFunc(route.Handler.EmployeeUploadV2))).Methods("POST")
	r.Handle("/dashboard/employee/count", Adapt(http.HandlerFunc(route.Handler.DashboardEmployeeCount))).Methods("POST")
	r.Handle("/employee/orgchart", Adapt(http.HandlerFunc(route.Handler.GetEmployeeChild))).Methods("GET")
	r.Handle("/employee/getorgchart", Adapt(http.HandlerFunc(route.Handler.GetOrgChart))).Methods("GET")
	r.Handle("/employee/linemanageremployee", Adapt(http.HandlerFunc(route.Handler.GetLineManagerEmployee))).Methods("GET")
	r.Handle("/employee/linemanagercheck", Adapt(http.HandlerFunc(route.Handler.GetEmployeeLinemanagerCheck))).Methods("GET")
	r.Handle("/employee/updateloginId", Adapt(http.HandlerFunc(route.Handler.EmployeeUpdateLoginId))).Methods("POST")

}

// func (route *Route) EmployeeAuthRoutes(r *mux.Router) {
// 	r.Handle("/employee/auth", Adapt(http.HandlerFunc(route.Handler.EmployeeLogin))).Methods("POST")
// }
