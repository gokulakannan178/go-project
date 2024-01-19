package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TicketCommentRoutes(r *mux.Router) {
	// TicketComment
	r.Handle("/ticketcomment", Adapt(http.HandlerFunc(route.Handler.SaveTicketComment))).Methods("POST")
	r.Handle("/ticketcomment", Adapt(http.HandlerFunc(route.Handler.GetSingleTicketComment))).Methods("GET")
	r.Handle("/ticketcomment", Adapt(http.HandlerFunc(route.Handler.UpdateTicketComment))).Methods("PUT")
	r.Handle("/ticketcomment/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTicketComment))).Methods("PUT")
	r.Handle("/ticketcomment/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTicketComment))).Methods("PUT")
	r.Handle("/ticketcomment/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTicketComment))).Methods("DELETE")
	r.Handle("/ticketcomment/filter", Adapt(http.HandlerFunc(route.Handler.FilterTicketComment))).Methods("POST")
}
