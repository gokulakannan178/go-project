package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PropertyVisitLogRemarkTypeRoutes : ""
func (route *Route) PropertyVisitLogRemarkType(r *mux.Router) {

	// PropertyVisitLogRemarkType
	r.Handle("/propertyvisitlogremarktype", Adapt(http.HandlerFunc(route.Handler.SavePropertyVisitLogRemarkType))).Methods("POST")
	r.Handle("/propertyvisitlogremarktype", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyVisitLogRemarkType))).Methods("GET")
	r.Handle("/propertyvisitlogremarktype", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyVisitLogRemarkType))).Methods("PUT")
	r.Handle("/propertyvisitlogremarktype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyVisitLogRemarkType))).Methods("PUT")
	r.Handle("/propertyvisitlogremarktype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyVisitLogRemarkType))).Methods("PUT")
	r.Handle("/propertyvisitlogremarktype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyVisitLogRemarkType))).Methods("DELETE")
	r.Handle("/propertyvisitlogremarktype/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyVisitLogRemarkType))).Methods("POST")
}
