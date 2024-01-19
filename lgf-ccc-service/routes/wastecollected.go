package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// WasteCollectedRoutes
func (route *Route) WasteCollectedRoutes(r *mux.Router) {
	r.Handle("/wastecollected", Adapt(http.HandlerFunc(route.Handler.SaveWasteCollected))).Methods("POST")
	r.Handle("/wastecollected", Adapt(http.HandlerFunc(route.Handler.GetSingleWasteCollected))).Methods("GET")
	r.Handle("/wastecollected", Adapt(http.HandlerFunc(route.Handler.UpdateWasteCollected))).Methods("PUT")
	r.Handle("/wastecollected/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWasteCollected))).Methods("PUT")
	r.Handle("/wastecollected/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWasteCollected))).Methods("PUT")
	r.Handle("/wastecollected/status/completed", Adapt(http.HandlerFunc(route.Handler.WasteCollectedCompleted))).Methods("PUT")
	r.Handle("/wastecollected/status/pending", Adapt(http.HandlerFunc(route.Handler.WasteCollectedPending))).Methods("PUT")
	r.Handle("/wastecollected/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWasteCollected))).Methods("DELETE")
	r.Handle("/wastecollected/filter", Adapt(http.HandlerFunc(route.Handler.FilterWasteCollected))).Methods("POST")
	//r.Handle("/wasteCollected/assign", Adapt(http.HandlerFunc(route.Handler.WasteCollectedAssign))).Methods("POST")
	//r.Handle("/wasteCollected/revoke", Adapt(http.HandlerFunc(route.Handler.RevokeWasteCollected))).Methods("PUT")

}
