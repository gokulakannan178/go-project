package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CommunicationCreditRoutes : ""
func (route *Route) CommunicationCreditRoutes(r *mux.Router) {
	r.Handle("/communicationCredit", Adapt(http.HandlerFunc(route.Handler.SaveCommunicationCredit))).Methods("POST")
	r.Handle("/communicationCredit", Adapt(http.HandlerFunc(route.Handler.GetSingleCommunicationCredit))).Methods("GET")
	r.Handle("/communicationCredit", Adapt(http.HandlerFunc(route.Handler.UpdateCommunicationCredit))).Methods("PUT")
	r.Handle("/communicationCredit/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommunicationCredit))).Methods("PUT")
	r.Handle("/communicationCredit/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommunicationCredit))).Methods("PUT")
	r.Handle("/communicationCredit/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommunicationCredit))).Methods("DELETE")
	r.Handle("/communicationCredit/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommunicationCredit))).Methods("POST")
}
