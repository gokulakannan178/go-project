package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ProductconfigsRoutes(r *mux.Router) {
	r.Handle("/save /productconfig", Adapt(http.HandlerFunc(route.Handler.SaveProductConfig))).Methods("POST")
	r.Handle("/Active /productconfig", Adapt(http.HandlerFunc(route.Handler.EnableProductConfig))).Methods("PUT")
	r.Handle("/productconfig/getactive", Adapt(http.HandlerFunc(route.Handler.GetactiveProductConfig))).Methods("GET")
}
