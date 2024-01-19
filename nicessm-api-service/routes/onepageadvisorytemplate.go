package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OnePageAdvisoryTemplateRoutes(r *mux.Router) {
	r.Handle("/onepageadvisorytemplate", Adapt(http.HandlerFunc(route.Handler.SaveOnePageAdvisoryTemplate))).Methods("POST")
	r.Handle("/onepageadvisorytemplate", Adapt(http.HandlerFunc(route.Handler.GetSingleOnePageAdvisoryTemplate))).Methods("GET")
	r.Handle("/onepageadvisorytemplate", Adapt(http.HandlerFunc(route.Handler.UpdateOnePageAdvisoryTemplate))).Methods("PUT")
	r.Handle("/onepageadvisorytemplate/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOnePageAdvisoryTemplate))).Methods("PUT")
	r.Handle("/onepageadvisorytemplate/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOnePageAdvisoryTemplate))).Methods("PUT")
	r.Handle("/onepageadvisorytemplate/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOnePageAdvisoryTemplate))).Methods("DELETE")
	r.Handle("/onepageadvisorytemplate/filter", Adapt(http.HandlerFunc(route.Handler.FilterOnePageAdvisoryTemplate))).Methods("POST")
}
