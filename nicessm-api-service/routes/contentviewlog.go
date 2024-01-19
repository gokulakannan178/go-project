package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ContentViewLogRoutes(r *mux.Router) {
	r.Handle("/contentviewlog", Adapt(http.HandlerFunc(route.Handler.SaveContentViewLog))).Methods("POST")
	r.Handle("/contentviewlog", Adapt(http.HandlerFunc(route.Handler.GetSingleContentViewLog))).Methods("GET")
	r.Handle("/contentviewlog", Adapt(http.HandlerFunc(route.Handler.UpdateContentViewLog))).Methods("PUT")
	r.Handle("/contentviewlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableContentViewLog))).Methods("PUT")
	r.Handle("/contentviewlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableContentViewLog))).Methods("PUT")
	r.Handle("/contentviewlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteContentViewLog))).Methods("DELETE")
	r.Handle("/contentviewlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterContentViewLog))).Methods("POST")
	r.Handle("/contentview/report/daywisedemand", Adapt(http.HandlerFunc(route.Handler.DayWiseContentViewChart))).Methods("POST")

}
