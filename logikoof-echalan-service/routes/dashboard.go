package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//DashboardRoutes : ""
func (route *Route) DashboardRoutes(r *mux.Router) {
	r.Handle("/dashboard/paymentwidget", Adapt(http.HandlerFunc(route.Handler.PaymentWidget))).Methods("POST")
	r.Handle("/dashboard/todaysoffencewidget", Adapt(http.HandlerFunc(route.Handler.TodaysOffenceWidget))).Methods("POST")
	r.Handle("/dashboard/topoffenceswidget", Adapt(http.HandlerFunc(route.Handler.TopOffencesWidget))).Methods("POST")

}
