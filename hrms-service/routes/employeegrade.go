package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GradeRoutes
func (route *Route) GradeRoutes(r *mux.Router) {
	r.Handle("/grade", Adapt(http.HandlerFunc(route.Handler.SaveGrade))).Methods("POST")
	r.Handle("/grade", Adapt(http.HandlerFunc(route.Handler.GetSingleGrade))).Methods("GET")
	r.Handle("/grade", Adapt(http.HandlerFunc(route.Handler.UpdateGrade))).Methods("PUT")
	r.Handle("/grade/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableGrade))).Methods("PUT")
	r.Handle("/grade/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableGrade))).Methods("PUT")
	r.Handle("/grade/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteGrade))).Methods("DELETE")
	r.Handle("/grade/filter", Adapt(http.HandlerFunc(route.Handler.FilterGrade))).Methods("POST")

}
