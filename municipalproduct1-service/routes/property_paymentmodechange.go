package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PropertyPaymentModeChangeRoutes : ""
func (route *Route) PropertyPaymentModeChangeRoutes(r *mux.Router) {
	r.Handle("/propertypayment/modechange/update", Adapt(http.HandlerFunc(route.Handler.SavePropertyPaymentModeChange))).Methods("PUT")
	r.Handle("/propertypayment/modechange", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyPaymentModeChange))).Methods("GET")
	r.Handle("/propertypayment/modechange/update/accept", Adapt(http.HandlerFunc(route.Handler.AcceptPropertyPaymentModeChange))).Methods("PUT")
	r.Handle("/propertypayment/modechange/update/reject", Adapt(http.HandlerFunc(route.Handler.RejectPropertyPaymentModeChange))).Methods("PUT")
	r.Handle("/propertypayment/modechange/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyPaymentModeChange))).Methods("POST")
}
