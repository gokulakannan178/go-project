package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// UserchargePaymentModeChangeRoutes : ""
func (route *Route) UserchargePaymentModeChangeRoutes(r *mux.Router) {
	r.Handle("/userchargepayment/modechange/update", Adapt(http.HandlerFunc(route.Handler.SaveUserchargePaymentModeChange))).Methods("PUT")
	r.Handle("/userchargepayment/modechange", Adapt(http.HandlerFunc(route.Handler.GetSingleUserchargePaymentModeChange))).Methods("GET")
	r.Handle("/userchargepayment/modechange/update/accept", Adapt(http.HandlerFunc(route.Handler.AcceptUserchargePaymentModeChange))).Methods("PUT")
	r.Handle("/userchargepayment/modechange/update/reject", Adapt(http.HandlerFunc(route.Handler.RejectUserchargePaymentModeChange))).Methods("PUT")
	r.Handle("/userchargepayment/modechange/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserchargePaymentModeChange))).Methods("POST")
}
