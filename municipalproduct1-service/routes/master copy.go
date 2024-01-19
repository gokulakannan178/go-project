package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//HonorifficRoutes : ""
func (route *Route) HonorifficRoutes(r *mux.Router) {
	r.Handle("/honoriffic", Adapt(http.HandlerFunc(route.Handler.SaveHonoriffic))).Methods("POST")
	r.Handle("/honoriffic", Adapt(http.HandlerFunc(route.Handler.GetSingleHonoriffic))).Methods("GET")
	r.Handle("/honoriffic", Adapt(http.HandlerFunc(route.Handler.UpdateHonoriffic))).Methods("PUT")
	r.Handle("/honoriffic/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableHonoriffic))).Methods("PUT")
	r.Handle("/honoriffic/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableHonoriffic))).Methods("PUT")
	r.Handle("/honoriffic/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteHonoriffic))).Methods("DELETE")
	r.Handle("/honoriffic/filter", Adapt(http.HandlerFunc(route.Handler.FilterHonoriffic))).Methods("POST")
}

//RelationRoutes : ""
func (route *Route) RelationRoutes(r *mux.Router) {
	r.Handle("/relation", Adapt(http.HandlerFunc(route.Handler.SaveRelation))).Methods("POST")
	r.Handle("/relation", Adapt(http.HandlerFunc(route.Handler.GetSingleRelation))).Methods("GET")
	r.Handle("/relation", Adapt(http.HandlerFunc(route.Handler.UpdateRelation))).Methods("PUT")
	r.Handle("/relation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableRelation))).Methods("PUT")
	r.Handle("/relation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableRelation))).Methods("PUT")
	r.Handle("/relation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteRelation))).Methods("DELETE")
	r.Handle("/relation/filter", Adapt(http.HandlerFunc(route.Handler.FilterRelation))).Methods("POST")
}
