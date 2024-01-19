package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//VehicleType : ""
func (route *Route) VehicleTypeRoutes(r *mux.Router) {
	r.Handle("/vehicletype", Adapt(http.HandlerFunc(route.Handler.SaveVehicleType))).Methods("POST")
	r.Handle("/vehicletype", Adapt(http.HandlerFunc(route.Handler.GetSingleVehicleType))).Methods("GET")
	r.Handle("/vehicletype", Adapt(http.HandlerFunc(route.Handler.UpdateVehicleType))).Methods("PUT")
	r.Handle("/vehicletype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVehicleType))).Methods("PUT")
	r.Handle("/vehicletype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVehicleType))).Methods("PUT")
	r.Handle("/vehicletype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVehicleType))).Methods("DELETE")
	r.Handle("/vehicletype/filter", Adapt(http.HandlerFunc(route.Handler.FilterVehicleType))).Methods("POST")

}
