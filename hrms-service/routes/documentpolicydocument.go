package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DocumentPolicyDocumentsRoutes
func (route *Route) DocumentPolicyDocumentsRoutes(r *mux.Router) {
	r.Handle("/documentpolicydocuments", Adapt(http.HandlerFunc(route.Handler.SaveDocumentPolicyDocuments))).Methods("POST")
	r.Handle("/documentpolicydocuments", Adapt(http.HandlerFunc(route.Handler.GetSingleDocumentPolicyDocuments))).Methods("GET")
	r.Handle("/documentpolicydocuments", Adapt(http.HandlerFunc(route.Handler.UpdateDocumentPolicyDocuments))).Methods("PUT")
	r.Handle("/documentpolicydocuments/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDocumentPolicyDocuments))).Methods("PUT")
	r.Handle("/documentpolicydocuments/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDocumentPolicyDocuments))).Methods("PUT")
	r.Handle("/documentpolicydocuments/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDocumentPolicyDocuments))).Methods("DELETE")
	r.Handle("/documentpolicydocuments/filter", Adapt(http.HandlerFunc(route.Handler.FilterDocumentPolicyDocuments))).Methods("POST")

}
