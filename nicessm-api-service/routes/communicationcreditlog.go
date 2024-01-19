package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CommunicationCreditLogRoutes : ""
func (route *Route) CommunicationCreditLogRoutes(r *mux.Router) {
	r.Handle("/communicationcreditlog", Adapt(http.HandlerFunc(route.Handler.SaveCommunicationCreditLog))).Methods("POST")
	r.Handle("/communicationcreditlog", Adapt(http.HandlerFunc(route.Handler.GetSingleCommunicationCreditLog))).Methods("GET")
	r.Handle("/communicationcreditlog", Adapt(http.HandlerFunc(route.Handler.UpdateCommunicationCreditLog))).Methods("PUT")
	r.Handle("/communicationcreditlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommunicationCreditLog))).Methods("PUT")
	r.Handle("/communicationcreditlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommunicationCreditLog))).Methods("PUT")
	r.Handle("/communicationcreditlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommunicationCreditLog))).Methods("DELETE")
	r.Handle("/communicationcreditlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommunicationCreditLog))).Methods("POST")
	r.Handle("/communicationcreditlog/addcredit", Adapt(http.HandlerFunc(route.Handler.UpdateCommunicationCreditLogWithPostCredit))).Methods("PUT")
}
