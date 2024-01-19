package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewsCommentRoutes : ""
func (route *Route) NewsCommentRoutes(r *mux.Router) {
	r.Handle("/newscomment", Adapt(http.HandlerFunc(route.Handler.SaveNewsComment))).Methods("POST")
	r.Handle("/newscomment", Adapt(http.HandlerFunc(route.Handler.GetSingleNewsComment))).Methods("GET")
	r.Handle("/newscomment", Adapt(http.HandlerFunc(route.Handler.UpdateNewsComment))).Methods("PUT")
	r.Handle("/newscomment/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableNewsComment))).Methods("PUT")
	r.Handle("/newscomment/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableNewsComment))).Methods("PUT")
	r.Handle("/newscomment/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteNewsComment))).Methods("DELETE")
	r.Handle("/newscomment/filter", Adapt(http.HandlerFunc(route.Handler.FilterNewsComment))).Methods("POST")

}
