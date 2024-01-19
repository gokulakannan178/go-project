package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DocumentMuxMasterRoutes(r *mux.Router) {
	// DocumentMuxMaster
	r.Handle("/documentMuxMaster", Adapt(http.HandlerFunc(route.Handler.SaveDocumentMuxMaster))).Methods("POST")
	r.Handle("/documentMuxMaster", Adapt(http.HandlerFunc(route.Handler.GetSingleDocumentMuxMaster))).Methods("GET")
	r.Handle("/documentMuxMaster", Adapt(http.HandlerFunc(route.Handler.UpdateDocumentMuxMaster))).Methods("PUT")
	r.Handle("/documentMuxMaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDocumentMuxMaster))).Methods("PUT")
	r.Handle("/documentMuxMaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDocumentMuxMaster))).Methods("PUT")
	r.Handle("/documentMuxMaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDocumentMuxMaster))).Methods("DELETE")
	r.Handle("/documentMuxMaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterDocumentMuxMaster))).Methods("POST")

}
