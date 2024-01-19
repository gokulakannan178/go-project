package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CitizenGrievanceRoutes(r *mux.Router) {
	// Citizen Grievance
	r.Handle("/citizengrievance", Adapt(http.HandlerFunc(route.Handler.SaveCitizenGrievance))).Methods("POST")
	r.Handle("/citizengrievance", Adapt(http.HandlerFunc(route.Handler.GetSingleCitizenGrievance))).Methods("GET")
	r.Handle("/citizengrievance", Adapt(http.HandlerFunc(route.Handler.UpdateCitizenGrievance))).Methods("PUT")
	r.Handle("/citizengrievance/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCitizenGrievance))).Methods("PUT")
	r.Handle("/citizengrievance/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCitizenGrievance))).Methods("PUT")
	r.Handle("/citizengrievance/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCitizenGrievance))).Methods("DELETE")
	r.Handle("/citizengrievance/status/complete", Adapt(http.HandlerFunc(route.Handler.CompletedCitizenGrievance))).Methods("PUT")
	r.Handle("/citizengrievance/status/rejected", Adapt(http.HandlerFunc(route.Handler.RejectedCitizenGrievance))).Methods("PUT")
	r.Handle("/citizengrievance/filter", Adapt(http.HandlerFunc(route.Handler.FilterCitizenGrievance))).Methods("POST")
	r.Handle("/citizengrievance/solution", Adapt(http.HandlerFunc(route.Handler.UpdateCitizenGrievanceSolution))).Methods("PUT")

}
