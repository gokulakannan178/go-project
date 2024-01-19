package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MySurvey : ""
func (route *Route) MySurveyRoutes(r *mux.Router) {
	r.Handle("/mysurvey", Adapt(http.HandlerFunc(route.Handler.SaveMySurvey))).Methods("POST")
	r.Handle("/mysurvey", Adapt(http.HandlerFunc(route.Handler.GetSingleMySurvey))).Methods("GET")
	r.Handle("/mysurvey", Adapt(http.HandlerFunc(route.Handler.UpdateMySurvey))).Methods("PUT")
	r.Handle("/mysurvey/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMySurvey))).Methods("PUT")
	r.Handle("/mysurvey/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMySurvey))).Methods("PUT")
	r.Handle("/mysurvey/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMySurvey))).Methods("DELETE")
	r.Handle("/mysurvey/filter", Adapt(http.HandlerFunc(route.Handler.FilterMySurvey))).Methods("POST")
	r.Handle("/mysurvey/citizen/property", Adapt(http.HandlerFunc(route.Handler.UpdateCitizenProperty))).Methods("PUT")

}
