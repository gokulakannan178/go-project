package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TradelicensePaymentModeChangeRoutes : ""
func (route *Route) TradelicensePaymentModeChangeRoutes(r *mux.Router) {
	r.Handle("/tradelicensepayment/modechange/update", Adapt(http.HandlerFunc(route.Handler.SaveTradelicensePaymentModeChange))).Methods("PUT")
	r.Handle("/tradelicensepayment/modechange", Adapt(http.HandlerFunc(route.Handler.GetSingleTradelicensePaymentModeChange))).Methods("GET")
	r.Handle("/tradelicensepayment/modechange/update/accept", Adapt(http.HandlerFunc(route.Handler.AcceptTradelicensePaymentModeChange))).Methods("PUT")
	r.Handle("/tradelicensepayment/modechange/update/reject", Adapt(http.HandlerFunc(route.Handler.RejectTradelicensePaymentModeChange))).Methods("PUT")
	r.Handle("/tradelicensepayment/modechange/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradelicensePaymentModeChange))).Methods("POST")
}
