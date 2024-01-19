package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//WardWiseDumpHistoryRoutes : ""
func (route *Route) WardWiseDumpHistoryRoutes(r *mux.Router) {
	r.Handle("/wardwisedumphistory", Adapt(http.HandlerFunc(route.Handler.SaveWardWiseDumpHistory))).Methods("POST")
	r.Handle("/wardwisedumphistory", Adapt(http.HandlerFunc(route.Handler.GetSingleWardWiseDumpHistory))).Methods("GET")
	r.Handle("/wardwisedumphistory", Adapt(http.HandlerFunc(route.Handler.UpdateWardWiseDumpHistory))).Methods("PUT")
	r.Handle("/wardwisedumphistory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWardWiseDumpHistory))).Methods("PUT")
	r.Handle("/wardwisedumphistory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWardWiseDumpHistory))).Methods("PUT")
	r.Handle("/wardwisedumphistory/filter", Adapt(http.HandlerFunc(route.Handler.FilterWardWiseDumpHistory))).Methods("POST")
	r.Handle("/wardwisedumphistory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWardWiseDumpHistory))).Methods("DELETE")
}
