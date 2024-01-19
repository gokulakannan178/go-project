package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ApplicantTypeRoutes : ""
func (route *Route) ApplicantTypeRoutes(r *mux.Router) {
	r.Handle("/applicanttype", Adapt(http.HandlerFunc(route.Handler.SaveApplicantType))).Methods("POST")
	r.Handle("/applicanttype", Adapt(http.HandlerFunc(route.Handler.GetSingleApplicantType))).Methods("GET")
	r.Handle("/applicanttype", Adapt(http.HandlerFunc(route.Handler.UpdateApplicantType))).Methods("PUT")
	r.Handle("/applicanttype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableApplicantType))).Methods("PUT")
	r.Handle("/applicanttype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableApplicantType))).Methods("PUT")
	r.Handle("/applicanttype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteApplicantType))).Methods("DELETE")
	r.Handle("/applicanttype/filter", Adapt(http.HandlerFunc(route.Handler.FilterApplicantType))).Methods("POST")
}

//EducationTypeRoutes : ""
func (route *Route) EducationTypeRoutes(r *mux.Router) {
	r.Handle("/educationtype", Adapt(http.HandlerFunc(route.Handler.SaveEducationType))).Methods("POST")
	r.Handle("/educationtype", Adapt(http.HandlerFunc(route.Handler.GetSingleEducationType))).Methods("GET")
	r.Handle("/educationtype", Adapt(http.HandlerFunc(route.Handler.UpdateEducationType))).Methods("PUT")
	r.Handle("/educationtype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableEducationType))).Methods("PUT")
	r.Handle("/educationtype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableEducationType))).Methods("PUT")
	r.Handle("/educationtype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteEducationType))).Methods("DELETE")
	r.Handle("/educationtype/filter", Adapt(http.HandlerFunc(route.Handler.FilterEducationType))).Methods("POST")
}
