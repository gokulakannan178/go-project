package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) EstimatedPropertyDemandRoutes(r *mux.Router) {
	r.Handle("/estimatedpropertydemand", Adapt(http.HandlerFunc(route.Handler.SaveEstimatedPropertyDemand))).Methods("POST")
	r.Handle("/estimatedpropertydemand/getdemand", Adapt(http.HandlerFunc(route.Handler.GetEstimatedPropertyDemandCalc))).Methods("GET")

}
