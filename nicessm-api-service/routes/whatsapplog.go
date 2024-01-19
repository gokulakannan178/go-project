package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//WhatsappLogRoutes : ""
func (route *Route) WhatsappLogRoutes(r *mux.Router) {
	r.Handle("/whatsappLog", Adapt(http.HandlerFunc(route.Handler.SaveWhatsappLog))).Methods("POST")
	r.Handle("/whatsappLog", Adapt(http.HandlerFunc(route.Handler.GetSingleWhatsappLog))).Methods("GET")
	r.Handle("/whatsappLog", Adapt(http.HandlerFunc(route.Handler.UpdateWhatsappLog))).Methods("PUT")
	r.Handle("/whatsappLog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWhatsappLog))).Methods("PUT")
	r.Handle("/whatsappLog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWhatsappLog))).Methods("PUT")
	r.Handle("/whatsappLog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWhatsappLog))).Methods("DELETE")
	r.Handle("/whatsappLog/filter", Adapt(http.HandlerFunc(route.Handler.FilterWhatsappLog))).Methods("POST")
}
