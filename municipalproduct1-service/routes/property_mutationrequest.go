package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PropertyMutationRequestRoutes : ""
func (route *Route) PropertyMutationRequestRoutes(r *mux.Router) {
	// PropertyMutation Request Update
	r.Handle("/propertymutationrequest/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdatePropertyMutationRequest))).Methods("PUT")
	r.Handle("/propertymutationrequest/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptPropertyMutationRequestUpdate))).Methods("PUT")
	r.Handle("/propertymutationrequest/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectPropertyMutationRequestUpdate))).Methods("PUT")
	r.Handle("/propertymutationrequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyMutationRequest))).Methods("POST")
	r.Handle("/propertymutationrequest", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyMutationRequest))).Methods("GET")

	r.Handle("/mutatedproperty", Adapt(http.HandlerFunc(route.Handler.SaveMutatedProperty))).Methods("POST")
	r.Handle("/mutatedproperty", Adapt(http.HandlerFunc(route.Handler.GetSingleMutatedProperty))).Methods("GET")
	r.Handle("/mutatedproperty/filter", Adapt(http.HandlerFunc(route.Handler.FilterMutatedProperty))).Methods("POST")
	r.Handle("/mutatedproperty/remainingarea", Adapt(http.HandlerFunc(route.Handler.RemainingAreaOfMutatedProperty))).Methods("GET")

	//
}
