package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CircleWiseDumpHistoryRoutes : ""
func (route *Route) CircleWiseDumpHistoryRoutes(r *mux.Router) {
	r.Handle("/circlewisedumphistory", Adapt(http.HandlerFunc(route.Handler.SaveCircleWiseDumpHistory))).Methods("POST")
	r.Handle("/circlewisedumphistory", Adapt(http.HandlerFunc(route.Handler.GetSingleCircleWiseDumpHistory))).Methods("GET")
	r.Handle("/circlewisedumphistory", Adapt(http.HandlerFunc(route.Handler.UpdateCircleWiseDumpHistory))).Methods("PUT")
	r.Handle("/circlewisedumphistory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCircleWiseDumpHistory))).Methods("PUT")
	r.Handle("/circlewisedumphistory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCircleWiseDumpHistory))).Methods("PUT")
	r.Handle("/circlewisedumphistory/filter", Adapt(http.HandlerFunc(route.Handler.FilterCircleWiseDumpHistory))).Methods("POST")
	r.Handle("/circlewisedumphistory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCircleWiseDumpHistory))).Methods("DELETE")
}
