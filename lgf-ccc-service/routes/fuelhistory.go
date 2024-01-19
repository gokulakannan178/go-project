package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FuelHistory : ""
func (route *Route) FuelHistoryRoutes(r *mux.Router) {
	r.Handle("/fuelhistory", Adapt(http.HandlerFunc(route.Handler.SaveFuelHistory))).Methods("POST")
	r.Handle("/fuelhistory", Adapt(http.HandlerFunc(route.Handler.GetSingleFuelHistory))).Methods("GET")
	r.Handle("/fuelhistory", Adapt(http.HandlerFunc(route.Handler.UpdateFuelHistory))).Methods("PUT")
	r.Handle("/fuelhistory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFuelHistory))).Methods("PUT")
	r.Handle("/fuelhistory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFuelHistory))).Methods("PUT")
	r.Handle("/fuelhistory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFuelHistory))).Methods("DELETE")
	r.Handle("/fuelhistory/filter", Adapt(http.HandlerFunc(route.Handler.FilterFuelHistory))).Methods("POST")
}
