package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CircleWiseHouseVisitedRoutes : ""
func (route *Route) CircleWiseHouseVisitedRoutes(r *mux.Router) {
	r.Handle("/circlewisehousevisited", Adapt(http.HandlerFunc(route.Handler.SaveCircleWiseHouseVisited))).Methods("POST")
	r.Handle("/circlewisehousevisited", Adapt(http.HandlerFunc(route.Handler.GetSingleCircleWiseHouseVisited))).Methods("GET")
	r.Handle("/circlewisehousevisited", Adapt(http.HandlerFunc(route.Handler.UpdateCircleWiseHouseVisited))).Methods("PUT")
	r.Handle("/circlewisehousevisited/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCircleWiseHouseVisited))).Methods("PUT")
	r.Handle("/circlewisehousevisited/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCircleWiseHouseVisited))).Methods("PUT")
	r.Handle("/circlewisehousevisited/filter", Adapt(http.HandlerFunc(route.Handler.FilterCircleWiseHouseVisited))).Methods("POST")
	r.Handle("/circlewisehousevisited/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCircleWiseHouseVisited))).Methods("DELETE")
}
