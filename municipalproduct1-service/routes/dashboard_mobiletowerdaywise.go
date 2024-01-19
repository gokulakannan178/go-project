package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) MobileTowerDashBoardDayWiseRoutes(r *mux.Router) {
	// MobileTower Day Wise
	r.Handle("/mobiletowerdaywise", Adapt(http.HandlerFunc(route.Handler.SaveMobileTowerDayWise))).Methods("POST")
	r.Handle("/mobiletowerdaywise", Adapt(http.HandlerFunc(route.Handler.GetSingleMobileTowerDayWise))).Methods("GET")
	r.Handle("/mobiletowerdaywise", Adapt(http.HandlerFunc(route.Handler.UpdateMobileTowerDayWise))).Methods("PUT")
	r.Handle("/mobiletowerdaywise/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMobileTowerDayWise))).Methods("PUT")
	r.Handle("/mobiletowerdaywise/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMobileTowerDayWise))).Methods("PUT")
	r.Handle("/mobiletowerdaywise/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMobileTowerDayWise))).Methods("DELETE")
	r.Handle("/mobiletowerdaywise/filter", Adapt(http.HandlerFunc(route.Handler.FilterMobileTowerDayWise))).Methods("POST")
}
