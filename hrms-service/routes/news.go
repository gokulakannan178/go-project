package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewsRoutes : ""
func (route *Route) NewsRoutes(r *mux.Router) {
	r.Handle("/news", Adapt(http.HandlerFunc(route.Handler.SaveNews))).Methods("POST")
	r.Handle("/news", Adapt(http.HandlerFunc(route.Handler.GetSingleNews))).Methods("GET")
	r.Handle("/news", Adapt(http.HandlerFunc(route.Handler.UpdateNews))).Methods("PUT")
	r.Handle("/news/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableNews))).Methods("PUT")
	r.Handle("/news/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableNews))).Methods("PUT")
	r.Handle("/news/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteNews))).Methods("DELETE")
	r.Handle("/news/status/published", Adapt(http.HandlerFunc(route.Handler.PublishedNews))).Methods("POST")
	r.Handle("/news/filter", Adapt(http.HandlerFunc(route.Handler.FilterNews))).Methods("POST")

}
