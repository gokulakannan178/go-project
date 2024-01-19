package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SalaryConfigLogRoutes : ""
func (route *Route) SalaryConfigLogRoutes(r *mux.Router) {
	r.Handle("/salaryconfiglog", Adapt(http.HandlerFunc(route.Handler.SaveSalaryConfigLog))).Methods("POST")
	r.Handle("/salaryconfiglog", Adapt(http.HandlerFunc(route.Handler.GetSingleSalaryConfigLog))).Methods("GET")
	r.Handle("/salaryconfiglog", Adapt(http.HandlerFunc(route.Handler.UpdateSalaryConfigLog))).Methods("PUT")
	r.Handle("/salaryconfiglog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSalaryConfigLog))).Methods("PUT")
	r.Handle("/salaryconfiglog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSalaryConfigLog))).Methods("PUT")
	r.Handle("/salaryconfiglog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSalaryConfigLog))).Methods("DELETE")
	r.Handle("/salaryconfiglog/filter", Adapt(http.HandlerFunc(route.Handler.FilterSalaryConfigLog))).Methods("POST")

}
