package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PropertyDocumentRoutes(r *mux.Router) {
	r.Handle("/propertyDocument", Adapt(http.HandlerFunc(route.Handler.SavePropertyDocument))).Methods("POST")
	r.Handle("/propertyDocument", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyDocument))).Methods("GET")
	r.Handle("/propertyDocument", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyDocument))).Methods("PUT")
	r.Handle("/propertyDocument/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyDocument))).Methods("PUT")
	r.Handle("/propertyDocument/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyDocument))).Methods("PUT")
	r.Handle("/propertyDocument/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyDocument))).Methods("DELETE")
	r.Handle("/propertyDocument/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyDocument))).Methods("POST")
}
