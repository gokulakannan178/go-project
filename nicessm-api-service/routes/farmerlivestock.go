package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FarmerLiveStockRoutes : ""
func (route *Route) FarmerLiveStockRoutes(r *mux.Router) {
	r.Handle("/farmerLiveStock", Adapt(http.HandlerFunc(route.Handler.SaveFarmerLiveStock))).Methods("POST")
	r.Handle("/farmerLiveStock", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmerLiveStock))).Methods("GET")
	r.Handle("/farmerLiveStock", Adapt(http.HandlerFunc(route.Handler.UpdateFarmerLiveStock))).Methods("PUT")
	r.Handle("/farmerLiveStock/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFarmerLiveStock))).Methods("PUT")
	r.Handle("/farmerLiveStock/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFarmerLiveStock))).Methods("PUT")
	r.Handle("/farmerLiveStock/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFarmerLiveStock))).Methods("DELETE")
	r.Handle("/farmerLiveStock/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmerLiveStock))).Methods("POST")
}
