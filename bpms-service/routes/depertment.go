package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//DepartmentRoutes : ""
func (route *Route) DepartmentRoutes(r *mux.Router) {
	r.Handle("/department", Adapt(http.HandlerFunc(route.Handler.SaveDepartment))).Methods("POST")
	r.Handle("/department", Adapt(http.HandlerFunc(route.Handler.GetSingleDepartment))).Methods("GET")
	r.Handle("/department", Adapt(http.HandlerFunc(route.Handler.UpdateDepartment))).Methods("PUT")
	r.Handle("/department/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDepartment))).Methods("PUT")
	r.Handle("/department/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDepartment))).Methods("PUT")
	r.Handle("/department/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDepartment))).Methods("DELETE")
	r.Handle("/department/filter", Adapt(http.HandlerFunc(route.Handler.FilterDepartment))).Methods("POST")
}

//DeptChecklistRoutes : ""
func (route *Route) DeptChecklistRoutes(r *mux.Router) {
	r.Handle("/deptchecklist", Adapt(http.HandlerFunc(route.Handler.SaveDeptChecklist))).Methods("POST")
	r.Handle("/deptchecklist", Adapt(http.HandlerFunc(route.Handler.GetSingleDeptChecklist))).Methods("GET")
	r.Handle("/deptchecklist", Adapt(http.HandlerFunc(route.Handler.UpdateDeptChecklist))).Methods("PUT")
	r.Handle("/deptchecklist/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDeptChecklist))).Methods("PUT")
	r.Handle("/deptchecklist/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDeptChecklist))).Methods("PUT")
	r.Handle("/deptchecklist/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDeptChecklist))).Methods("DELETE")
	r.Handle("/deptchecklist/filter", Adapt(http.HandlerFunc(route.Handler.FilterDeptChecklist))).Methods("POST")
}

//DepartmentTypeRoutes : ""
func (route *Route) DepartmentTypeRoutes(r *mux.Router) {
	r.Handle("/departmenttype", Adapt(http.HandlerFunc(route.Handler.SaveDepartmentType))).Methods("POST")
	r.Handle("/departmenttype", Adapt(http.HandlerFunc(route.Handler.GetSingleDepartmentType))).Methods("GET")
	r.Handle("/departmenttype", Adapt(http.HandlerFunc(route.Handler.UpdateDepartmentType))).Methods("PUT")
	r.Handle("/departmenttype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDepartmentType))).Methods("PUT")
	r.Handle("/departmenttype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDepartmentType))).Methods("PUT")
	r.Handle("/departmenttype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDepartmentType))).Methods("DELETE")
	r.Handle("/departmenttype/filter", Adapt(http.HandlerFunc(route.Handler.FilterDepartmentType))).Methods("POST")
}
