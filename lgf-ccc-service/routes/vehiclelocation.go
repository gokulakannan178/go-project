package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//VehicleLocationRoutes : ""
func (route *Route) VehicleLocationRoutes(r *mux.Router) {
	r.Handle("/vehiclelocation", Adapt(http.HandlerFunc(route.Handler.SaveVehicleLocation))).Methods("POST")
	r.Handle("/vehiclelocation", Adapt(http.HandlerFunc(route.Handler.GetSingleVehicleLocation))).Methods("GET")
	r.Handle("/vehiclelocation", Adapt(http.HandlerFunc(route.Handler.UpdateVehicleLocation))).Methods("PUT")
	r.Handle("/vehiclelocation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVehicleLocation))).Methods("PUT")
	r.Handle("/vehiclelocation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVehicleLocation))).Methods("PUT")
	r.Handle("/vehiclelocation/filter", Adapt(http.HandlerFunc(route.Handler.FilterVehicleLocation))).Methods("POST")
	r.Handle("/vehiclelocation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVehicleLocation))).Methods("DELETE")
}
