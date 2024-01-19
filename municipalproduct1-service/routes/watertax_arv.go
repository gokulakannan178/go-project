package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) WaterTaxArvRoutes(r *mux.Router) {
	// WaterTaxArv
	r.Handle("/watertaxarv", Adapt(http.HandlerFunc(route.Handler.SaveWaterTaxArv))).Methods("POST")
	r.Handle("/watertaxarv", Adapt(http.HandlerFunc(route.Handler.GetSingleWaterTaxArv))).Methods("GET")
	r.Handle("/watertaxarv", Adapt(http.HandlerFunc(route.Handler.UpdateWaterTaxArv))).Methods("PUT")
	r.Handle("/watertaxarv/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWaterTaxArv))).Methods("PUT")
	r.Handle("/watertaxarv/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWaterTaxArv))).Methods("PUT")
	r.Handle("/watertaxarv/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWaterTaxArv))).Methods("DELETE")
	r.Handle("/watertaxarv/filter", Adapt(http.HandlerFunc(route.Handler.FilterWaterTaxArv))).Methods("POST")
}
