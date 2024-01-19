package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PropertyTypeRoutes : ""
func (route *Route) PropertyTypeRoutes(r *mux.Router) {
	r.Handle("/propertytype", Adapt(http.HandlerFunc(route.Handler.SavePropertyType))).Methods("POST")
	r.Handle("/propertytype", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyType))).Methods("GET")
	r.Handle("/propertytype", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyType))).Methods("PUT")
	r.Handle("/propertytype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyType))).Methods("PUT")
	r.Handle("/propertytype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyType))).Methods("PUT")
	r.Handle("/propertytype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyType))).Methods("DELETE")
	r.Handle("/propertytype/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyType))).Methods("POST")

}
