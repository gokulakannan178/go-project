package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WaterTaxConnectionTypeRoutes(r *mux.Router) {
	// WaterTaxConnectionType
	r.Handle("/watertaxconnectiontype", Adapt(http.HandlerFunc(route.Handler.SaveWaterTaxConnectionType))).Methods("POST")
	r.Handle("/watertaxconnectiontype", Adapt(http.HandlerFunc(route.Handler.GetSingleWaterTaxConnectionType))).Methods("GET")
	r.Handle("/watertaxconnectiontype", Adapt(http.HandlerFunc(route.Handler.UpdateWaterTaxConnectionType))).Methods("PUT")
	r.Handle("/watertaxconnectiontype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWaterTaxConnectionType))).Methods("PUT")
	r.Handle("/watertaxconnectiontype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWaterTaxConnectionType))).Methods("PUT")
	r.Handle("/watertaxconnectiontype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWaterTaxConnectionType))).Methods("DELETE")
	r.Handle("/watertaxconnectiontype/filter", Adapt(http.HandlerFunc(route.Handler.FilterWaterTaxConnectionType))).Methods("POST")
}
