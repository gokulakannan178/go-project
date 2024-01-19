package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OnePageAdvisoryRoutes(r *mux.Router) {
	r.Handle("/onepageadvisory", Adapt(http.HandlerFunc(route.Handler.SaveOnePageAdvisory))).Methods("POST")
	r.Handle("/onepageadvisory", Adapt(http.HandlerFunc(route.Handler.GetSingleOnePageAdvisory))).Methods("GET")
	r.Handle("/onepageadvisory", Adapt(http.HandlerFunc(route.Handler.UpdateOnePageAdvisory))).Methods("PUT")
	r.Handle("/onepageadvisory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOnePageAdvisory))).Methods("PUT")
	r.Handle("/onepageadvisory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOnePageAdvisory))).Methods("PUT")
	r.Handle("/onepageadvisory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOnePageAdvisory))).Methods("DELETE")
	r.Handle("/onepageadvisory/filter", Adapt(http.HandlerFunc(route.Handler.FilterOnePageAdvisory))).Methods("POST")
}
