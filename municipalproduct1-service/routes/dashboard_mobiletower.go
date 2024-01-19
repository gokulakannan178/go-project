package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) MobileTowerDashboardRoutes(r *mux.Router) {
	// MobileTower
	r.Handle("/mobiletower", Adapt(http.HandlerFunc(route.Handler.SaveMobileTower))).Methods("POST")
	r.Handle("/mobiletower", Adapt(http.HandlerFunc(route.Handler.GetSingleMobileTower))).Methods("GET")
	r.Handle("/mobiletower", Adapt(http.HandlerFunc(route.Handler.UpdateMobileTower))).Methods("PUT")
	r.Handle("/mobiletower/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMobileTower))).Methods("PUT")
	r.Handle("/mobiletower/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMobileTower))).Methods("PUT")
	r.Handle("/mobiletower/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMobileTower))).Methods("DELETE")
	r.Handle("/mobiletower/filter", Adapt(http.HandlerFunc(route.Handler.FilterMobileTower))).Methods("POST")

	r.Handle("/mobiletower/dashboard/ddac", Adapt(http.HandlerFunc(route.Handler.DashboardMobileTowerDemandAndCollection))).Methods("POST")

}
