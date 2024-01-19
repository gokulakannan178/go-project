package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WaterBillDashboardRoutes(r *mux.Router) {
	// DashBoardWaterBill
	r.Handle("/waterbill", Adapt(http.HandlerFunc(route.Handler.SaveWaterBillDashboard))).Methods("POST")
	r.Handle("/waterbill", Adapt(http.HandlerFunc(route.Handler.GetSingleWaterBillDashboard))).Methods("GET")
	r.Handle("/waterbill", Adapt(http.HandlerFunc(route.Handler.UpdateWaterBillDashboard))).Methods("PUT")
	r.Handle("/waterbill/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWaterBillDashboard))).Methods("PUT")
	r.Handle("/waterbill/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWaterBillDashboard))).Methods("PUT")
	r.Handle("/waterbill/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDashBoardWaterBill))).Methods("DELETE")
	r.Handle("/waterbill/filter", Adapt(http.HandlerFunc(route.Handler.FilterWaterBillDashboard))).Methods("POST")
}
