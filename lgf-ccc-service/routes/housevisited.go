package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// housevisited : ""
func (route *Route) HouseVisitedRoutes(r *mux.Router) {
	r.Handle("/housevisited", Adapt(http.HandlerFunc(route.Handler.SaveHouseVisited))).Methods("POST")
	r.Handle("/housevisited", Adapt(http.HandlerFunc(route.Handler.GetSingleHouseVisited))).Methods("GET")
	r.Handle("/housevisited", Adapt(http.HandlerFunc(route.Handler.UpdateHouseVisited))).Methods("PUT")
	r.Handle("/housevisited/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableHouseVisited))).Methods("PUT")
	r.Handle("/housevisited/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableHouseVisited))).Methods("PUT")
	r.Handle("/housevisited/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteHouseVisited))).Methods("DELETE")
	r.Handle("/housevisited/filter", Adapt(http.HandlerFunc(route.Handler.FilterHouseVisited))).Methods("POST")
	r.Handle("/housevisited/status/Collected", Adapt(http.HandlerFunc(route.Handler.CollectedHouseVisited))).Methods("PUT")

}
