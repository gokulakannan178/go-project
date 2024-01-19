package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//StateRoutes : ""
func (route *Route) StateRoutes(r *mux.Router) {
	r.Handle("/state", Adapt(http.HandlerFunc(route.Handler.SaveState))).Methods("POST")
	r.Handle("/state", Adapt(http.HandlerFunc(route.Handler.GetSingleState))).Methods("GET")
	r.Handle("/state", Adapt(http.HandlerFunc(route.Handler.UpdateState))).Methods("PUT")
	r.Handle("/state/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableState))).Methods("PUT")
	r.Handle("/state/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableState))).Methods("PUT")
	r.Handle("/state/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteState))).Methods("DELETE")
	r.Handle("/state/filter", Adapt(http.HandlerFunc(route.Handler.FilterState))).Methods("POST")
}

//DistrictRoutes : ""
func (route *Route) DistrictRoutes(r *mux.Router) {
	r.Handle("/district", Adapt(http.HandlerFunc(route.Handler.SaveDistrict))).Methods("POST")
	r.Handle("/district", Adapt(http.HandlerFunc(route.Handler.GetSingleDistrict))).Methods("GET")
	r.Handle("/district", Adapt(http.HandlerFunc(route.Handler.UpdateDistrict))).Methods("PUT")
	r.Handle("/district/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDistrict))).Methods("PUT")
	r.Handle("/district/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDistrict))).Methods("PUT")
	r.Handle("/district/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDistrict))).Methods("DELETE")
	r.Handle("/district/filter", Adapt(http.HandlerFunc(route.Handler.FilterDistrict))).Methods("POST")
}

//VillageRoutes : ""
func (route *Route) VillageRoutes(r *mux.Router) {
	r.Handle("/village", Adapt(http.HandlerFunc(route.Handler.SaveVillage))).Methods("POST")
	r.Handle("/village", Adapt(http.HandlerFunc(route.Handler.GetSingleVillage))).Methods("GET")
	r.Handle("/village", Adapt(http.HandlerFunc(route.Handler.UpdateVillage))).Methods("PUT")
	r.Handle("/village/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVillage))).Methods("PUT")
	r.Handle("/village/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVillage))).Methods("PUT")
	r.Handle("/village/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVillage))).Methods("DELETE")
	r.Handle("/village/filter", Adapt(http.HandlerFunc(route.Handler.FilterVillage))).Methods("POST")
}

//ZoneRoutes : ""
func (route *Route) ZoneRoutes(r *mux.Router) {
	r.Handle("/zone", Adapt(http.HandlerFunc(route.Handler.SaveZone))).Methods("POST")
	r.Handle("/zone", Adapt(http.HandlerFunc(route.Handler.GetSingleZone))).Methods("GET")
	r.Handle("/zone", Adapt(http.HandlerFunc(route.Handler.UpdateZone))).Methods("PUT")
	r.Handle("/zone/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableZone))).Methods("PUT")
	r.Handle("/zone/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableZone))).Methods("PUT")
	r.Handle("/zone/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteZone))).Methods("DELETE")
	r.Handle("/zone/filter", Adapt(http.HandlerFunc(route.Handler.FilterZone))).Methods("POST")
}

//WardRoutes : ""
func (route *Route) WardRoutes(r *mux.Router) {
	r.Handle("/ward", Adapt(http.HandlerFunc(route.Handler.SaveWard))).Methods("POST")
	r.Handle("/ward", Adapt(http.HandlerFunc(route.Handler.GetSingleWard))).Methods("GET")
	r.Handle("/ward", Adapt(http.HandlerFunc(route.Handler.UpdateWard))).Methods("PUT")
	r.Handle("/ward/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWard))).Methods("PUT")
	r.Handle("/ward/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWard))).Methods("PUT")
	r.Handle("/ward/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWard))).Methods("DELETE")
	r.Handle("/ward/filter", Adapt(http.HandlerFunc(route.Handler.FilterWard))).Methods("POST")
}
