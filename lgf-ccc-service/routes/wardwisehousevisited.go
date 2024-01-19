package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//WardWiseHouseVisitedRoutes : ""
func (route *Route) WardWiseHouseVisitedRoutes(r *mux.Router) {
	r.Handle("/wardwisehousevisited", Adapt(http.HandlerFunc(route.Handler.SaveWardWiseHouseVisited))).Methods("POST")
	r.Handle("/wardwisehousevisited", Adapt(http.HandlerFunc(route.Handler.GetSingleWardWiseHouseVisited))).Methods("GET")
	r.Handle("/wardwisehousevisited", Adapt(http.HandlerFunc(route.Handler.UpdateWardWiseHouseVisited))).Methods("PUT")
	r.Handle("/wardwisehousevisited/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableWardWiseHouseVisited))).Methods("PUT")
	r.Handle("/wardwisehousevisited/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableWardWiseHouseVisited))).Methods("PUT")
	r.Handle("/wardwisehousevisited/filter", Adapt(http.HandlerFunc(route.Handler.FilterWardWiseHouseVisited))).Methods("POST")
	r.Handle("/wardwisehousevisited/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteWardWiseHouseVisited))).Methods("DELETE")
}
