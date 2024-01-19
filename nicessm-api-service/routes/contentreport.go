package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ContentReportRoutes
func (route *Route) ContentReportRoutes(r *mux.Router) {
	r.Handle("/contentreport/filter", Adapt(http.HandlerFunc(route.Handler.FilterContentReport))).Methods("POST")
	r.Handle("/duplicatecontentreport/filter", Adapt(http.HandlerFunc(route.Handler.FilterDuplicateContentReport))).Methods("GET")

}
