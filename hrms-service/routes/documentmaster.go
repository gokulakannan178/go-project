package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DocumentMaterRoutes
func (route *Route) DocumentMaterRoutes(r *mux.Router) {
	r.Handle("/documentmaster", Adapt(http.HandlerFunc(route.Handler.SaveDocumentMaster))).Methods("POST")
	r.Handle("/documentmaster", Adapt(http.HandlerFunc(route.Handler.GetSingleDocumentMaster))).Methods("GET")
	r.Handle("/documentmaster", Adapt(http.HandlerFunc(route.Handler.UpdateDocumentMaster))).Methods("PUT")
	r.Handle("/documentmaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDocumentMaster))).Methods("PUT")
	r.Handle("/documentmaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDocumentMaster))).Methods("PUT")
	r.Handle("/documentmaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDocumentMaster))).Methods("DELETE")
	r.Handle("/documentmaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterDocumentMaster))).Methods("POST")

}
