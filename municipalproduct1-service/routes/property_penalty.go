package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PropertyPenaltyRoutes : ""
func (route *Route) PropertyPenaltyRoutes(r *mux.Router) {

	// Property Penalty
	r.Handle("/propertypenalty", Adapt(http.HandlerFunc(route.Handler.SavePropertyPenalty))).Methods("POST")
	r.Handle("/propertypenalty", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyPenalty))).Methods("GET")
	r.Handle("/propertypenalty", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyPenalty))).Methods("PUT")
	r.Handle("/propertypenalty/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyPenalty))).Methods("PUT")
	r.Handle("/propertypenalty/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyPenalty))).Methods("PUT")
	r.Handle("/propertypenalty/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyPenalty))).Methods("DELETE")
	r.Handle("/propertypenalty/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyPenalty))).Methods("POST")
}
