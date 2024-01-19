package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DocumentTypeRoutes(r *mux.Router) {
	// DocumentType
	r.Handle("/documentType", Adapt(http.HandlerFunc(route.Handler.SaveDocumentType))).Methods("POST")
	r.Handle("/documentType", Adapt(http.HandlerFunc(route.Handler.GetSingleDocumentType))).Methods("GET")
	r.Handle("/documentType", Adapt(http.HandlerFunc(route.Handler.UpdateDocumentType))).Methods("PUT")
	r.Handle("/documentType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDocumentType))).Methods("PUT")
	r.Handle("/documentType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDocumentType))).Methods("PUT")
	r.Handle("/documentType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDocumentType))).Methods("DELETE")
	r.Handle("/documentType/filter", Adapt(http.HandlerFunc(route.Handler.FilterDocumentType))).Methods("POST")

}
