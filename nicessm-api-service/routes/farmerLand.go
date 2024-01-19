package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) FarmerLandRoutes(r *mux.Router) {
	r.Handle("/farmerland", Adapt(http.HandlerFunc(route.Handler.SaveFarmerLand))).Methods("POST")
	r.Handle("/farmerland", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmerLand))).Methods("GET")
	r.Handle("/farmerland", Adapt(http.HandlerFunc(route.Handler.UpdateFarmerLand))).Methods("PUT")
	r.Handle("/farmerland/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFarmerLand))).Methods("PUT")
	r.Handle("/farmerland/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFarmerLand))).Methods("PUT")
	r.Handle("/farmerland/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFarmerLand))).Methods("DELETE")
	r.Handle("/farmerland/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmerLand))).Methods("POST")
}
