package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//VehicleInsurance : ""
func (route *Route) VehicleInsuranceRoutes(r *mux.Router) {
	r.Handle("/vehicleinsurance", Adapt(http.HandlerFunc(route.Handler.SaveVehicleInsurance))).Methods("POST")
	r.Handle("/vehicleinsurance", Adapt(http.HandlerFunc(route.Handler.GetSingleVehicleInsurance))).Methods("GET")
	r.Handle("/vehicleinsurance", Adapt(http.HandlerFunc(route.Handler.UpdateVehicleInsurance))).Methods("PUT")
	r.Handle("/vehicleinsurance/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVehicleInsurance))).Methods("PUT")
	r.Handle("/vehicleinsurance/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVehicleInsurance))).Methods("PUT")
	r.Handle("/vehicleinsurance/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVehicleInsurance))).Methods("DELETE")
	r.Handle("/vehicleinsurance/filter", Adapt(http.HandlerFunc(route.Handler.FilterVehicleInsurance))).Methods("POST")
}
