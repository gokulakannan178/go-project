package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OrganisationConfigRoutes(r *mux.Router) {
	r.Handle("/organisationConfig", Adapt(http.HandlerFunc(route.Handler.SaveOrganisationConfig))).Methods("POST")
	r.Handle("/organisationConfig", Adapt(http.HandlerFunc(route.Handler.GetSingleOrganisationConfig))).Methods("GET")
	r.Handle("/organisationConfig", Adapt(http.HandlerFunc(route.Handler.UpdateOrganisationConfig))).Methods("PUT")
	r.Handle("/organisationConfig/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOrganisationConfig))).Methods("PUT")
	r.Handle("/organisationConfig/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOrganisationConfig))).Methods("PUT")
	r.Handle("/organisationConfig/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOrganisationConfig))).Methods("DELETE")
	r.Handle("/organisationConfig/filter", Adapt(http.HandlerFunc(route.Handler.FilterOrganisationConfig))).Methods("POST")
	r.Handle("/organisationConfig/setdefault", Adapt(http.HandlerFunc(route.Handler.SetdefaultOrganisationConfig))).Methods("PUT")
	r.Handle("/organisationconfig/getdefault", Adapt(http.HandlerFunc(route.Handler.GetactiveOrganisationConfig))).Methods("GET")
}
