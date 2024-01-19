package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CitizenGraviansLogRoutes(r *mux.Router) {
	// CitizenGraviansLog
	r.Handle("/citizengravianslog", Adapt(http.HandlerFunc(route.Handler.SaveCitizenGraviansLog))).Methods("POST")
	r.Handle("/citizengravianslog", Adapt(http.HandlerFunc(route.Handler.GetSingleCitizenGraviansLog))).Methods("GET")
	r.Handle("/citizengravianslog", Adapt(http.HandlerFunc(route.Handler.UpdateCitizenGraviansLog))).Methods("PUT")
	r.Handle("/citizengravianslog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCitizenGraviansLog))).Methods("PUT")
	r.Handle("/citizengravianslog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCitizenGraviansLog))).Methods("PUT")
	r.Handle("/citizengravianslog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCitizenGraviansLog))).Methods("DELETE")
	r.Handle("/citizengravianslog/filter", Adapt(http.HandlerFunc(route.Handler.FilterCitizenGraviansLog))).Methods("POST")
}
