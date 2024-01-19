package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewsLikeRoutes : ""
func (route *Route) NewsLikeRoutes(r *mux.Router) {
	r.Handle("/newslike", Adapt(http.HandlerFunc(route.Handler.SaveNewsLike))).Methods("POST")
	r.Handle("/newslike", Adapt(http.HandlerFunc(route.Handler.GetSingleNewsLike))).Methods("GET")
	r.Handle("/newslike", Adapt(http.HandlerFunc(route.Handler.UpdateNewsLike))).Methods("PUT")
	r.Handle("/newslike/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableNewsLike))).Methods("PUT")
	r.Handle("/newslike/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableNewsLike))).Methods("PUT")
	r.Handle("/newslike/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteNewsLike))).Methods("DELETE")
	r.Handle("/newslike/filter", Adapt(http.HandlerFunc(route.Handler.FilterNewsLike))).Methods("POST")

}
