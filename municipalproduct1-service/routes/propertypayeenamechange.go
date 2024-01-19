package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PropertyPayeeNameChangeRoutes(r *mux.Router) {
	r.Handle("/propertypayeenamechange", Adapt(http.HandlerFunc(route.Handler.SavePropertyPayeeNameChange))).Methods("POST")
	r.Handle("/propertypayeenamechange", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyPayeeNameChange))).Methods("GET")
	r.Handle("/propertypayeenamechange", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyPayeeNameChange))).Methods("PUT")
	r.Handle("/propertypayeenamechange/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyPayeeNameChange))).Methods("PUT")
	r.Handle("/propertypayeenamechange/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyPayeeNameChange))).Methods("PUT")
	r.Handle("/propertypayeenamechange/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyPayeeNameChange))).Methods("DELETE")
	r.Handle("/propertypayeenamechange/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyPayeeNameChange))).Methods("POST")

	r.Handle("/propertypayeenamechange/approve", Adapt(http.HandlerFunc(route.Handler.ApprovePropertyPayeeNameChange))).Methods("PUT")
	r.Handle("/propertypayeenamechange/rejected", Adapt(http.HandlerFunc(route.Handler.NotApprovePropertyPayeeNameChange))).Methods("PUT")

}
