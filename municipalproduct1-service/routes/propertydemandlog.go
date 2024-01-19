package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PropertyDemandLogRoutes : ""
func (route *Route) PropertyDemandLogRoutes(r *mux.Router) {
	r.Handle("/property/overallproperty/demand", Adapt(http.HandlerFunc(route.Handler.SaveOverAllPropertyDemand))).Methods("GET")
	r.Handle("/property/overallproperty/demand/view", Adapt(http.HandlerFunc(route.Handler.GetOverAllPropertyDemand))).Methods("GET")
}
