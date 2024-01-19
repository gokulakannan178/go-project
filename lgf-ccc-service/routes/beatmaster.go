package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) BeatMasterRoutes(r *mux.Router) {
	// BeatMaster
	r.Handle("/beatmaster", Adapt(http.HandlerFunc(route.Handler.SaveBeatMaster))).Methods("POST")
	r.Handle("/beatmaster", Adapt(http.HandlerFunc(route.Handler.GetSingleBeatMaster))).Methods("GET")
	r.Handle("/beatmaster", Adapt(http.HandlerFunc(route.Handler.UpdateBeatMaster))).Methods("PUT")
	r.Handle("/beatmaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBeatMaster))).Methods("PUT")
	r.Handle("/beatmaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBeatMaster))).Methods("PUT")
	r.Handle("/beatmaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBeatMaster))).Methods("DELETE")
	r.Handle("/beatmaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterBeatMaster))).Methods("POST")
	r.Handle("/beatmaster/vehicle/assign", Adapt(http.HandlerFunc(route.Handler.VehicleAssignForBeatMaster))).Methods("PUT")

}
