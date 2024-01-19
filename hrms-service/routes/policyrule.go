package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PolicyRule : ""
func (route *Route) PolicyRuleRoutes(r *mux.Router) {
	r.Handle("/policyrule", Adapt(http.HandlerFunc(route.Handler.SavePolicyRule))).Methods("POST")
	r.Handle("/policyrule", Adapt(http.HandlerFunc(route.Handler.GetSinglePolicyRule))).Methods("GET")
	r.Handle("/policyrule", Adapt(http.HandlerFunc(route.Handler.UpdatePolicyRule))).Methods("PUT")
	r.Handle("/policyrule/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePolicyRule))).Methods("PUT")
	r.Handle("/policyrule/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePolicyRule))).Methods("PUT")
	r.Handle("/policyrule/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePolicyRule))).Methods("DELETE")
	r.Handle("/policyrule/filter", Adapt(http.HandlerFunc(route.Handler.FilterPolicyRule))).Methods("POST")

}
