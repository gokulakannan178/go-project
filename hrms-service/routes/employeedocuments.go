package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) EmployeeDocumentsRoutes(r *mux.Router) {
	// EmployeeDocuments
	r.Handle("/employeeDocuments", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeDocuments))).Methods("POST")
	r.Handle("/employeeDocuments", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeDocuments))).Methods("GET")
	r.Handle("/employeeDocuments", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeDocuments))).Methods("PUT")
	r.Handle("/employeeDocuments/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeDocuments))).Methods("PUT")
	r.Handle("/employeeDocuments/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeDocuments))).Methods("PUT")
	r.Handle("/employeeDocuments/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeDocuments))).Methods("DELETE")
	r.Handle("/employeeDocuments/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeDocuments))).Methods("POST")
	r.Handle("/employeeDocuments/list", Adapt(http.HandlerFunc(route.Handler.EmployeeDocumentsList))).Methods("POST")
	r.Handle("/employeeDocuments/updateDocument", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeDocumentsWithUpsert))).Methods("PUT")
	r.Handle("/employeeDocuments/removeDocument", Adapt(http.HandlerFunc(route.Handler.RemoveEmployeeDocuments))).Methods("PUT")

}
