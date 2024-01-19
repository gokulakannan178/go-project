package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DocumentPolicyRoutes
func (route *Route) DocumentPolicyRoutes(r *mux.Router) {
	r.Handle("/documentpolicy", Adapt(http.HandlerFunc(route.Handler.SaveDocumentPolicy))).Methods("POST")
	r.Handle("/documentpolicy", Adapt(http.HandlerFunc(route.Handler.GetSingleDocumentPolicy))).Methods("GET")
	r.Handle("/documentpolicy", Adapt(http.HandlerFunc(route.Handler.UpdateDocumentPolicy))).Methods("PUT")
	r.Handle("/documentpolicy/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDocumentPolicy))).Methods("PUT")
	r.Handle("/documentpolicy/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDocumentPolicy))).Methods("PUT")
	r.Handle("/documentpolicy/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDocumentPolicy))).Methods("DELETE")
	r.Handle("/documentpolicy/filter", Adapt(http.HandlerFunc(route.Handler.FilterDocumentPolicy))).Methods("POST")

}
