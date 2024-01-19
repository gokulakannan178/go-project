package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//UserReportRoutes
func (route *Route) UserReportRoutes(r *mux.Router) {
	r.Handle("/userreport/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserReport))).Methods("POST")
	r.Handle("/duplicateuserreport/filter", Adapt(http.HandlerFunc(route.Handler.FilterDuplicateUserReport))).Methods("POST")

}
