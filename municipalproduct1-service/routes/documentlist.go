package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DocumentListRoutes(r *mux.Router) {
	r.Handle("/documentList", Adapt(http.HandlerFunc(route.Handler.SaveDocumentList))).Methods("POST")
	r.Handle("/documentList", Adapt(http.HandlerFunc(route.Handler.GetSingleDocumentList))).Methods("GET")
	r.Handle("/documentList", Adapt(http.HandlerFunc(route.Handler.UpdateDocumentList))).Methods("PUT")
	r.Handle("/documentList/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDocumentList))).Methods("PUT")
	r.Handle("/documentList/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDocumentList))).Methods("PUT")
	r.Handle("/documentList/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDocumentList))).Methods("DELETE")
	r.Handle("/documentList/filter", Adapt(http.HandlerFunc(route.Handler.FilterDocumentList))).Methods("POST")
}
