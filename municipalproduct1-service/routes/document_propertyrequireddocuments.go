package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PropertyRequiredDocumentRoutes(r *mux.Router) {
	r.Handle("/propertyrequireddocument", Adapt(http.HandlerFunc(route.Handler.SavePropertyRequiredDocument))).Methods("POST")
	r.Handle("/propertyrequireddocument", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyRequiredDocument))).Methods("GET")
	r.Handle("/propertyrequireddocument", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyRequiredDocument))).Methods("PUT")
	r.Handle("/propertyrequireddocument/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyRequiredDocument))).Methods("PUT")
	r.Handle("/propertyrequireddocument/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyRequiredDocument))).Methods("PUT")
	r.Handle("/propertyrequireddocument/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyRequiredDocument))).Methods("DELETE")
	r.Handle("/propertyrequireddocument/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyRequiredDocument))).Methods("POST")
}
