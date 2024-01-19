package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TicketUserRoutes(r *mux.Router) {
	// TicketUser
	r.Handle("/ticketuser", Adapt(http.HandlerFunc(route.Handler.SaveTicketUser))).Methods("POST")
	r.Handle("/ticketuser", Adapt(http.HandlerFunc(route.Handler.GetSingleTicketUser))).Methods("GET")
	r.Handle("/ticketuser", Adapt(http.HandlerFunc(route.Handler.UpdateTicketUser))).Methods("PUT")
	r.Handle("/ticketuser/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTicketUser))).Methods("PUT")
	r.Handle("/ticketuser/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTicketUser))).Methods("PUT")
	r.Handle("/ticketuser/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTicketUser))).Methods("DELETE")
	r.Handle("/ticketuser/filter", Adapt(http.HandlerFunc(route.Handler.FilterTicketUser))).Methods("POST")
}
