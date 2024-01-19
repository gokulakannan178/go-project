package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//EmployeeDeductionMasterRoutes : ""
func (route *Route) EmployeeDeductionMasterRoutes(r *mux.Router) {
	r.Handle("/employeedeductionmaster", Adapt(http.HandlerFunc(route.Handler.SaveEmployeeDeductionMaster))).Methods("POST")
	r.Handle("/employeedeductionmaster", Adapt(http.HandlerFunc(route.Handler.GetSingleEmployeeDeductionMaster))).Methods("GET")
	r.Handle("/employeedeductionmaster", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeDeductionMaster))).Methods("PUT")
	r.Handle("/employeedeductionmaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEmployeeDeductionMaster))).Methods("PUT")
	r.Handle("/employeedeductionmaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEmployeeDeductionMaster))).Methods("PUT")
	r.Handle("/employeedeductionmaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEmployeeDeductionMaster))).Methods("DELETE")
	r.Handle("/employeedeductionmaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterEmployeeDeductionMaster))).Methods("POST")

}
