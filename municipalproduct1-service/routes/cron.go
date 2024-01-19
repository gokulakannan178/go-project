package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CronRoutes(r *mux.Router) {
	r.Handle("/cron/dailylog", Adapt(http.HandlerFunc(route.Handler.TodaysLog))).Methods("POST")
	r.Handle("/asset/dashboard2", Adapt(http.HandlerFunc(route.Handler.AssetAPIForInterview))).Methods("GET")
}
