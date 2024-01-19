package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//BranchRoutes : ""
func (route *Route) BranchRoutes(r *mux.Router) {
	r.Handle("/Branch", Adapt(http.HandlerFunc(route.Handler.SaveBranch))).Methods("POST")
	r.Handle("/Branch", Adapt(http.HandlerFunc(route.Handler.GetSingleBranch))).Methods("GET")
	r.Handle("/Branch/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBranch))).Methods("PUT")
	r.Handle("/Branch/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBranch))).Methods("PUT")
	r.Handle("/Branch", Adapt(http.HandlerFunc(route.Handler.UpdateBranch))).Methods("PUT")
	r.Handle("/Branch/filter", Adapt(http.HandlerFunc(route.Handler.FilterBranch))).Methods("POST")
	r.Handle("/Branch/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBranch))).Methods("DELETE")
}
