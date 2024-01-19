package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//VehicleRoutes : ""
func (route *Route) VehicleRoutes(r *mux.Router) {
	r.Handle("/vehicle", Adapt(http.HandlerFunc(route.Handler.SaveVehicle))).Methods("POST")
	r.Handle("/vehicle", Adapt(http.HandlerFunc(route.Handler.GetSingleVehicle))).Methods("GET")
	r.Handle("/vehicle", Adapt(http.HandlerFunc(route.Handler.UpdateVehicle))).Methods("PUT")
	r.Handle("/vehicle/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVehicle))).Methods("PUT")
	r.Handle("/vehicle/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVehicle))).Methods("PUT")
	r.Handle("/vehicle/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVehicle))).Methods("DELETE")
	r.Handle("/vehicle/filter", Adapt(http.HandlerFunc(route.Handler.FilterVehicle))).Methods("POST")
}

//VehicleChallanRoutes : ""
func (route *Route) VehicleChallanRoutes(r *mux.Router) {
	r.Handle("/vehiclechallan", Adapt(http.HandlerFunc(route.Handler.SaveVehicleChallan))).Methods("POST")
	r.Handle("/vehiclechallan", Adapt(http.HandlerFunc(route.Handler.GetSingleVehicleChallan))).Methods("GET")
	r.Handle("/vehiclechallan", Adapt(http.HandlerFunc(route.Handler.UpdateVehicleChallan))).Methods("PUT")
	r.Handle("/vehiclechallan/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVehicleChallan))).Methods("PUT")
	r.Handle("/vehiclechallan/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVehicleChallan))).Methods("PUT")
	r.Handle("/vehiclechallan/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVehicleChallan))).Methods("DELETE")
	r.Handle("/vehiclechallan/filter", Adapt(http.HandlerFunc(route.Handler.FilterVehicleChallan))).Methods("POST")
}

//OffenceTypeRoutes : ""
func (route *Route) OffenceTypeRoutes(r *mux.Router) {
	r.Handle("/offencetype", Adapt(http.HandlerFunc(route.Handler.SaveOffenceType))).Methods("POST")
	r.Handle("/offencetype", Adapt(http.HandlerFunc(route.Handler.GetSingleOffenceType))).Methods("GET")
	r.Handle("/offencetype", Adapt(http.HandlerFunc(route.Handler.UpdateOffenceType))).Methods("PUT")
	r.Handle("/offencetype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOffenceType))).Methods("PUT")
	r.Handle("/offencetype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOffenceType))).Methods("PUT")
	r.Handle("/offencetype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOffenceType))).Methods("DELETE")
	r.Handle("/offencetype/filter", Adapt(http.HandlerFunc(route.Handler.FilterOffenceType))).Methods("POST")
}
