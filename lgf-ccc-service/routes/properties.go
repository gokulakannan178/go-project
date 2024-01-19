package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PropertiesRoutes : ""
func (route *Route) PropertiesRoutes(r *mux.Router) {
	r.Handle("/properties", Adapt(http.HandlerFunc(route.Handler.SaveProperties))).Methods("POST")
	r.Handle("/properties", Adapt(http.HandlerFunc(route.Handler.GetSingleProperties))).Methods("GET")
	r.Handle("/properties/get", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertiesWithHoldingNumber))).Methods("GET")
	r.Handle("/properties", Adapt(http.HandlerFunc(route.Handler.UpdateProperties))).Methods("PUT")
	r.Handle("/properties/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProperties))).Methods("PUT")
	r.Handle("/properties/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProperties))).Methods("PUT")
	r.Handle("/properties/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProperties))).Methods("DELETE")
	r.Handle("/properties/filter", Adapt(http.HandlerFunc(route.Handler.FilterProperties))).Methods("POST")

}
