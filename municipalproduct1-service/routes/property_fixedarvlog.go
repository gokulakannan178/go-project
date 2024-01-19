package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PropertyFixedArvLogRoutes(r *mux.Router) {
	// PropertyFixedArvLog
	r.Handle("/propertyfixedarvlog", Adapt(http.HandlerFunc(route.Handler.SavePropertyFixedArvLog))).Methods("POST")
	r.Handle("/propertyfixedarvlog", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyFixedArvLog))).Methods("GET")
	r.Handle("/propertyfixedarvlog", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyFixedArvLog))).Methods("PUT")
	r.Handle("/propertyfixedarvlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyFixedArvLog))).Methods("PUT")
	r.Handle("/propertyfixedarvlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyFixedArvLog))).Methods("PUT")
	r.Handle("/propertyfixedarvlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyFixedArvLog))).Methods("DELETE")
	r.Handle("/propertyfixedarvlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyFixedArvLog))).Methods("POST")
}
