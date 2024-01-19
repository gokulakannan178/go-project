package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TradeLicenseReportRoute : ""
func (route *Route) SaveStoredDemandRoute(r *mux.Router) {
	r.Handle("/save/storeddemand", Adapt(http.HandlerFunc(route.Handler.SaveStoredDemand))).Methods("GET")
	r.Handle("/property/getdemandv3", Adapt(http.HandlerFunc(route.Handler.GetDemandV3))).Methods("GET")
}
