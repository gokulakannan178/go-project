package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//VehicleLog : ""
func (route *Route) VehicleLogRoutes(r *mux.Router) {
	r.Handle("/vehiclelog", Adapt(http.HandlerFunc(route.Handler.SaveVehicleLog))).Methods("POST")
	r.Handle("/vehiclelog", Adapt(http.HandlerFunc(route.Handler.GetSingleVehicleLog))).Methods("GET")
	r.Handle("/vehiclelog", Adapt(http.HandlerFunc(route.Handler.UpdateVehicleLog))).Methods("PUT")
	r.Handle("/vehiclelog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVehicleLog))).Methods("PUT")
	r.Handle("/vehiclelog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVehicleLog))).Methods("PUT")
	r.Handle("/vehiclelog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVehicleLog))).Methods("DELETE")
	r.Handle("/vehiclelog/filter", Adapt(http.HandlerFunc(route.Handler.FilterVehicleLog))).Methods("POST")
}
