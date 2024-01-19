package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ContentCountLogRoutes(r *mux.Router) {
	r.Handle("/contentcountlog", Adapt(http.HandlerFunc(route.Handler.SaveContentCountLog))).Methods("POST")
	r.Handle("/contentcountlog", Adapt(http.HandlerFunc(route.Handler.GetSingleContentCountLog))).Methods("GET")
	r.Handle("/contentcountlog", Adapt(http.HandlerFunc(route.Handler.UpdateContentCountLog))).Methods("PUT")
	r.Handle("/contentcountlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableContentCountLog))).Methods("PUT")
	r.Handle("/contentcountlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableContentCountLog))).Methods("PUT")
	r.Handle("/contentcountlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteContentCountLog))).Methods("DELETE")
	r.Handle("/contentcountlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterContentCountLog))).Methods("POST")
}
