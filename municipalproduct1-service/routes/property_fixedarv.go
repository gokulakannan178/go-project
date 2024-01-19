package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PropertyFixedArvRoutes(r *mux.Router) {
	// PropertyFixedArv
	r.Handle("/propertyfixedarv", Adapt(http.HandlerFunc(route.Handler.SavePropertyFixedArv))).Methods("POST")
	r.Handle("/propertyfixedarv", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyFixedArv))).Methods("GET")
	r.Handle("/propertyfixedarv", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyFixedArv))).Methods("PUT")
	r.Handle("/propertyfixedarv/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyFixedArv))).Methods("PUT")
	r.Handle("/propertyfixedarv/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyFixedArv))).Methods("PUT")
	r.Handle("/propertyfixedarv/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyFixedArv))).Methods("DELETE")
	r.Handle("/propertyfixedarv/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyFixedArv))).Methods("POST")
	r.Handle("/propertyfixedarv/accept", Adapt(http.HandlerFunc(route.Handler.AcceptPropertyFixedArv))).Methods("PUT")
	r.Handle("/propertyfixedarv/reject", Adapt(http.HandlerFunc(route.Handler.RejectPropertyFixedArv))).Methods("PUT")
	r.Handle("/propertyfixedarv/multipleaccept", Adapt(http.HandlerFunc(route.Handler.AcceptMultiplePropertyFixedArv))).Methods("PUT")

}
