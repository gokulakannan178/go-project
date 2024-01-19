package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PropertyVisitLogRoutes : ""
func (route *Route) PropertyVisitLog(r *mux.Router) {

	// PropertyVisitLog
	r.Handle("/propertyvisitlog", Adapt(http.HandlerFunc(route.Handler.SavePropertyVisitLog))).Methods("POST")
	r.Handle("/propertyvisitlog", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyVisitLog))).Methods("GET")
	r.Handle("/propertyvisitlog", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyVisitLog))).Methods("PUT")
	r.Handle("/propertyvisitlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyVisitLog))).Methods("PUT")
	r.Handle("/propertyvisitlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyVisitLog))).Methods("PUT")
	r.Handle("/propertyvisitlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyVisitLog))).Methods("DELETE")
	r.Handle("/propertyvisitlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyVisitLog))).Methods("POST")
}
