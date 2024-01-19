package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) RoadTypeRoutes(r *mux.Router) {
	r.Handle("/roadtype", Adapt(http.HandlerFunc(route.Handler.SaveRoadType))).Methods("POST")
	r.Handle("/roadtype", Adapt(http.HandlerFunc(route.Handler.GetSingleRoadType))).Methods("GET")
	r.Handle("/roadtype", Adapt(http.HandlerFunc(route.Handler.UpdateRoadType))).Methods("PUT")
	r.Handle("/roadtype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableRoadType))).Methods("PUT")
	r.Handle("/roadtype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableRoadType))).Methods("PUT")
	r.Handle("/roadtype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteRoadType))).Methods("DELETE")
	r.Handle("/roadtype/filter", Adapt(http.HandlerFunc(route.Handler.FilterRoadType))).Methods("POST")
}
