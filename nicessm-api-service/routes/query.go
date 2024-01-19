package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) QueryRoutes(r *mux.Router) {
	r.Handle("/query", Adapt(http.HandlerFunc(route.Handler.SaveQuery))).Methods("POST")
	r.Handle("/query", Adapt(http.HandlerFunc(route.Handler.GetSingleQuery))).Methods("GET")
	r.Handle("/query", Adapt(http.HandlerFunc(route.Handler.UpdateQuery))).Methods("PUT")
	r.Handle("/query/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableQuery))).Methods("PUT")
	r.Handle("/query/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableQuery))).Methods("PUT")
	r.Handle("/query/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteQuery))).Methods("DELETE")
	r.Handle("/query/filter", Adapt(http.HandlerFunc(route.Handler.FilterQuery))).Methods("POST")
	r.Handle("/query/assinguser", Adapt(http.HandlerFunc(route.Handler.AssignuserQuery))).Methods("PUT")
	r.Handle("/query/resolveuser", Adapt(http.HandlerFunc(route.Handler.ResolveuserQuery))).Methods("PUT")

}
