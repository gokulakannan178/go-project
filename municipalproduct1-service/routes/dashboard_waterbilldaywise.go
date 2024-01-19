package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WaterBillDayWiseDashboardRoutes(r *mux.Router) {
	// DashBoardWaterBill
	r.Handle("/waterbilldaywise", Adapt(http.HandlerFunc(route.Handler.SaveWaterBillDayWiseDashboard))).Methods("POST")
	r.Handle("/waterbilldaywise", Adapt(http.HandlerFunc(route.Handler.GetSingleWaterBillDayWiseDashboard))).Methods("GET")
	r.Handle("/waterbilldaywise", Adapt(http.HandlerFunc(route.Handler.UpdateWaterBillDayWiseDashboard))).Methods("PUT")
	r.Handle("/waterbilldaywise/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWaterBillDayWiseDashboard))).Methods("PUT")
	r.Handle("/waterbilldaywise/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWaterBillDayWiseDashboard))).Methods("PUT")
	r.Handle("/waterbilldaywise/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWaterBillDayWiseDashboard))).Methods("DELETE")
	r.Handle("/waterbilldaywise/filter", Adapt(http.HandlerFunc(route.Handler.FilterWaterBillDayWiseDashboard))).Methods("POST")
}
