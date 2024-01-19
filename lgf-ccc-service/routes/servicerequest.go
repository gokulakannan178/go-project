package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ServiceRequest : ""
func (route *Route) ServiceRequestRoutes(r *mux.Router) {
	r.Handle("/servicerequest", Adapt(http.HandlerFunc(route.Handler.SaveServiceRequest))).Methods("POST")
	r.Handle("/servicerequest", Adapt(http.HandlerFunc(route.Handler.GetSingleServiceRequest))).Methods("GET")
	r.Handle("/servicerequest", Adapt(http.HandlerFunc(route.Handler.UpdateServiceRequest))).Methods("PUT")
	r.Handle("/servicerequest/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableServiceRequest))).Methods("PUT")
	r.Handle("/servicerequest/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableServiceRequest))).Methods("PUT")
	r.Handle("/servicerequest/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteServiceRequest))).Methods("DELETE")
	r.Handle("/servicerequest/status/init", Adapt(http.HandlerFunc(route.Handler.InitServiceRequest))).Methods("PUT")
	r.Handle("/servicerequest/status/pending", Adapt(http.HandlerFunc(route.Handler.PendingServiceRequest))).Methods("PUT")
	r.Handle("/servicerequest/status/inprogress", Adapt(http.HandlerFunc(route.Handler.InProgressServiceRequest))).Methods("PUT")
	r.Handle("/servicerequest/status/completed", Adapt(http.HandlerFunc(route.Handler.CompletedServiceRequest))).Methods("PUT")
	r.Handle("/servicerequest/filter", Adapt(http.HandlerFunc(route.Handler.FilterServiceRequest))).Methods("POST")
	r.Handle("/servicerequest/getdetailservice", Adapt(http.HandlerFunc(route.Handler.GetDetailServiceRequest))).Methods("GET")
	r.Handle("/servicerequest/assign", Adapt(http.HandlerFunc(route.Handler.AssignServiceRequest))).Methods("PUT")

}
