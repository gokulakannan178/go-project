package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SmsLogRoutes : ""
func (route *Route) SmsLogRoutes(r *mux.Router) {
	r.Handle("/smsLog", Adapt(http.HandlerFunc(route.Handler.SaveSmsLog))).Methods("POST")
	r.Handle("/smsLog", Adapt(http.HandlerFunc(route.Handler.GetSingleSmsLog))).Methods("GET")
	r.Handle("/smsLog", Adapt(http.HandlerFunc(route.Handler.UpdateSmsLog))).Methods("PUT")
	r.Handle("/smsLog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableSmsLog))).Methods("PUT")
	r.Handle("/smsLog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableSmsLog))).Methods("PUT")
	r.Handle("/smsLog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteSmsLog))).Methods("DELETE")
	r.Handle("/smsLog/filter", Adapt(http.HandlerFunc(route.Handler.FilterSmsLog))).Methods("POST")
}
