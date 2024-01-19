package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//LeaveMaster : ""
func (route *Route) LeaveMasterRoutes(r *mux.Router) {
	r.Handle("/leavemaster", Adapt(http.HandlerFunc(route.Handler.SaveLeaveMaster))).Methods("POST")
	r.Handle("/leavemaster", Adapt(http.HandlerFunc(route.Handler.GetSingleLeaveMaster))).Methods("GET")
	r.Handle("/leavemaster", Adapt(http.HandlerFunc(route.Handler.UpdateLeaveMaster))).Methods("PUT")
	r.Handle("/leavemaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLeaveMaster))).Methods("PUT")
	r.Handle("/leavemaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLeaveMaster))).Methods("PUT")
	r.Handle("/leavemaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLeaveMaster))).Methods("DELETE")
	r.Handle("/leavemaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterLeaveMaster))).Methods("POST")

}
