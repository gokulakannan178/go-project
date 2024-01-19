package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SelfConsumptionRoutes(r *mux.Router) {
	// SelfConsumption
	r.Handle("/selfConsumption", Adapt(http.HandlerFunc(route.Handler.SaveSelfConsumption))).Methods("POST")
	r.Handle("/selfConsumption", Adapt(http.HandlerFunc(route.Handler.GetSingleSelfConsumption))).Methods("GET")
	r.Handle("/selfConsumption", Adapt(http.HandlerFunc(route.Handler.UpdateSelfConsumption))).Methods("PUT")
	r.Handle("/selfConsumption/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSelfConsumption))).Methods("PUT")
	r.Handle("/selfConsumption/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSelfConsumption))).Methods("PUT")
	r.Handle("/selfConsumption/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSelfConsumption))).Methods("DELETE")
	r.Handle("/selfConsumption/filter", Adapt(http.HandlerFunc(route.Handler.FilterSelfConsumption))).Methods("POST")
	r.Handle("/selfConsumption/decrease", Adapt(http.HandlerFunc(route.Handler.DecreaseInventoryForULBandFPO))).Methods("PUT")

}
