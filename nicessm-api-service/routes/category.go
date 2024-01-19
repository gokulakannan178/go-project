package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CategoryRoutes(r *mux.Router) {
	r.Handle("/category", Adapt(http.HandlerFunc(route.Handler.SaveCategory))).Methods("POST")
	r.Handle("/category", Adapt(http.HandlerFunc(route.Handler.GetSingleCategory))).Methods("GET")
	r.Handle("/category", Adapt(http.HandlerFunc(route.Handler.UpdateCategory))).Methods("PUT")
	r.Handle("/category/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCategory))).Methods("PUT")
	r.Handle("/category/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCategory))).Methods("PUT")
	r.Handle("/category/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCategory))).Methods("DELETE")
	r.Handle("/category/filter", Adapt(http.HandlerFunc(route.Handler.FilterCategory))).Methods("POST")
}
