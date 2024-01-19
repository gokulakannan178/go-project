package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Cron Routes: ""
func (rout *Route) CronRoutes(r *mux.Router) {
	r.Handle("/cron/inventorymonthlyupdate", Adapt(http.HandlerFunc(rout.Handler.NotifyForULBInventoryUpdate))).Methods("GET")
}
