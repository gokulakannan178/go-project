package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TicketRoutes(r *mux.Router) {
	// Ticket
	r.Handle("/ticket", Adapt(http.HandlerFunc(route.Handler.SaveTicket))).Methods("POST")
	r.Handle("/ticket", Adapt(http.HandlerFunc(route.Handler.GetSingleTicket))).Methods("GET")
	r.Handle("/ticket", Adapt(http.HandlerFunc(route.Handler.UpdateTicket))).Methods("PUT")
	r.Handle("/ticket/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTicket))).Methods("PUT")
	r.Handle("/ticket/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTicket))).Methods("PUT")
	r.Handle("/ticket/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTicket))).Methods("DELETE")
	r.Handle("/ticket/filter", Adapt(http.HandlerFunc(route.Handler.FilterTicket))).Methods("POST")
}
