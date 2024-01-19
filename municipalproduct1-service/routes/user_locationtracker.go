package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserLocationTrackerRoutes(r *mux.Router) {
	// UserLocationTracker
	r.Handle("/userlocationtracker", Adapt(http.HandlerFunc(route.Handler.SaveUserLocationTracker))).Methods("POST")
	r.Handle("/userlocationtracker", Adapt(http.HandlerFunc(route.Handler.GetSingleUserLocationTracker))).Methods("GET")
	r.Handle("/userlocationtracker", Adapt(http.HandlerFunc(route.Handler.UpdateUserLocationTracker))).Methods("PUT")
	r.Handle("/userlocationtracker/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserLocationTracker))).Methods("PUT")
	r.Handle("/userlocationtracker/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserLocationTracker))).Methods("PUT")
	r.Handle("/userlocationtracker/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserLocationTracker))).Methods("DELETE")
	r.Handle("/userlocationtracker/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserLocationTracker))).Methods("POST")
	r.Handle("/userlocationtrackercoordinates", Adapt(http.HandlerFunc(route.Handler.UserLocationTrackerCoordinates))).Methods("POST")

}
