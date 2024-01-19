package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) MajorUpdateRoutes(r *mux.Router) {
	// MajorUpdate
	r.Handle("/majorupdate", Adapt(http.HandlerFunc(route.Handler.SaveMajorUpdate))).Methods("POST")
	r.Handle("/majorupdate", Adapt(http.HandlerFunc(route.Handler.GetSingleMajorUpdate))).Methods("GET")
	r.Handle("/majorupdate", Adapt(http.HandlerFunc(route.Handler.UpdateMajorUpdate))).Methods("PUT")
	r.Handle("/majorupdate/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMajorUpdate))).Methods("PUT")
	r.Handle("/majorupdate/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMajorUpdate))).Methods("PUT")
	r.Handle("/majorupdate/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMajorUpdate))).Methods("DELETE")
	r.Handle("/majorupdate/filter", Adapt(http.HandlerFunc(route.Handler.FilterMajorUpdate))).Methods("POST")
}
