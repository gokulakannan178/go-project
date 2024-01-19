package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TradeLicensePart2Routes
func (route *Route) TradeLicenseDemandPart2Routes(r *mux.Router) {
	r.Handle("/tradelicense/getdemand/part2", Adapt(http.HandlerFunc(route.Handler.TradeLicenseDemandPart2))).Methods("GET")
}
