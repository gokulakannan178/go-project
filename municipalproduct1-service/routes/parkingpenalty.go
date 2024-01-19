package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ParkingPenaltyRoutes : ""
func (route *Route) ParkingPenaltyRoutes(r *mux.Router) {
	r.Handle("/parkingpenalty", Adapt(http.HandlerFunc(route.Handler.SaveParkingPenalty))).Methods("POST")
	r.Handle("/parkingpenalty", Adapt(http.HandlerFunc(route.Handler.GetSingleParkingPenalty))).Methods("GET")
	r.Handle("/parkingpenalty", Adapt(http.HandlerFunc(route.Handler.UpdateParkingPenalty))).Methods("PUT")
	r.Handle("/parkingpenalty/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableParkingPenalty))).Methods("PUT")
	r.Handle("/parkingpenalty/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableParkingPenalty))).Methods("PUT")
	r.Handle("/parkingpenalty/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteParkingPenalty))).Methods("DELETE")
	r.Handle("/parkingpenalty/filter", Adapt(http.HandlerFunc(route.Handler.FilterParkingPenalty))).Methods("POST")
}
