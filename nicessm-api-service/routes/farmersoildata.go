package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FarmerSoilDataRoutes : ""
func (route *Route) FarmerSoilDataRoutes(r *mux.Router) {
	r.Handle("/farmerSoilData", Adapt(http.HandlerFunc(route.Handler.SaveFarmerSoilData))).Methods("POST")
	r.Handle("/farmerSoilData", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmerSoilData))).Methods("GET")
	r.Handle("/farmerSoilData", Adapt(http.HandlerFunc(route.Handler.UpdateFarmerSoilData))).Methods("PUT")
	r.Handle("/farmerSoilData/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFarmerSoilData))).Methods("PUT")
	r.Handle("/farmerSoilData/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFarmerSoilData))).Methods("PUT")
	r.Handle("/farmerSoilData/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFarmerSoilData))).Methods("DELETE")
	r.Handle("/farmerSoilData/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmerSoilData))).Methods("POST")
}
