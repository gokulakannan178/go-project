package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//VehicleTripRoutes : ""
func (route *Route) VehicleTripRoutes(r *mux.Router) {
	r.Handle("/vehicletrip", Adapt(http.HandlerFunc(route.Handler.SaveVehicleTrip))).Methods("POST")
	r.Handle("/vehicletrip", Adapt(http.HandlerFunc(route.Handler.GetSingleVehicleTrip))).Methods("GET")
	r.Handle("/vehicletrip", Adapt(http.HandlerFunc(route.Handler.UpdateVehicleTrip))).Methods("PUT")
	r.Handle("/vehicletrip/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVehicleTrip))).Methods("PUT")
	r.Handle("/vehicletrip/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVehicleTrip))).Methods("PUT")
	r.Handle("/vehicletrip/filter", Adapt(http.HandlerFunc(route.Handler.FilterVehicleTrip))).Methods("POST")
	r.Handle("/vehicletrip/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVehicleTrip))).Methods("DELETE")
}
