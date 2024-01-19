package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NotifyRoutes : ""
func (route *Route) NotifyRoutes(r *mux.Router) {
	// notify routes
	r.Handle("/notify/updatelocation", Adapt(http.HandlerFunc(route.Handler.NotifyForUpdateLocation))).Methods("POST")
	r.Handle("/notify/updatelocation", Adapt(http.HandlerFunc(route.Handler.NotifyForUpdateProfile))).Methods("POST")
	r.Handle("/notify/ulbinventoryupdate", Adapt(http.HandlerFunc(route.Handler.NotifyForULBInventoryUpdate))).Methods("POST")
	r.Handle("/notify/ulbinventoryupdatev2", Adapt(http.HandlerFunc(route.Handler.NotifyForULBInventoryUpdateV2))).Methods("POST")

}
