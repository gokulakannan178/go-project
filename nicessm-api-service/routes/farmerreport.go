package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FarmerReportRoutes
func (route *Route) FarmerReportRoutes(r *mux.Router) {
	r.Handle("/farmerreport/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmerReport))).Methods("POST")
	r.Handle("/farmer/report", Adapt(http.HandlerFunc(route.Handler.FilterFarmerReport2))).Methods("POST")
	r.Handle("/duplicatefarmerreport/filter", Adapt(http.HandlerFunc(route.Handler.FilterDuplicateFarmer))).Methods("POST")

}
