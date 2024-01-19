package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//TelegramLogRoutes : ""
func (route *Route) TelegramLogRoutes(r *mux.Router) {
	r.Handle("/telegramLog", Adapt(http.HandlerFunc(route.Handler.SaveTelegramLog))).Methods("POST")
	r.Handle("/telegramLog", Adapt(http.HandlerFunc(route.Handler.GetSingleTelegramLog))).Methods("GET")
	r.Handle("/telegramLog", Adapt(http.HandlerFunc(route.Handler.UpdateTelegramLog))).Methods("PUT")
	r.Handle("/telegramLog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTelegramLog))).Methods("PUT")
	r.Handle("/telegramLog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTelegramLog))).Methods("PUT")
	r.Handle("/telegramLog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTelegramLog))).Methods("DELETE")
	r.Handle("/telegramLog/filter", Adapt(http.HandlerFunc(route.Handler.FilterTelegramLog))).Methods("POST")
}
