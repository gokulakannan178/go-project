package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PropertyDeleteRequestRoutes : ""
func (route *Route) PropertyDeleteRequestRoutes(r *mux.Router) {
	// PropertyDelete Request Update
	r.Handle("/propertydeleterequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdatePropertyDeleteRequest))).Methods("PUT")
	r.Handle("/propertydeleterequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptPropertyDeleteRequestUpdate))).Methods("PUT")
	r.Handle("/propertydeleterequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectPropertyDeleteRequestUpdate))).Methods("PUT")
	r.Handle("/propertydeleterequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyDeleteRequest))).Methods("POST")
	r.Handle("/propertydeleterequest", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyDeleteRequest))).Methods("GET")

	//
}
