package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Probationary : ""
func (route *Route) ProbationaryRoutes(r *mux.Router) {
	r.Handle("/probationary", Adapt(http.HandlerFunc(route.Handler.SaveProbationary))).Methods("POST")
	r.Handle("/probationary", Adapt(http.HandlerFunc(route.Handler.GetSingleProbationary))).Methods("GET")
	r.Handle("/probationary", Adapt(http.HandlerFunc(route.Handler.UpdateProbationary))).Methods("PUT")
	r.Handle("/probationary/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProbationary))).Methods("PUT")
	r.Handle("/probationary/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProbationary))).Methods("PUT")
	r.Handle("/probationary/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProbationary))).Methods("DELETE")
	r.Handle("/probationary/filter", Adapt(http.HandlerFunc(route.Handler.FilterProbationary))).Methods("POST")

}
