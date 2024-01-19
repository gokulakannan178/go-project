package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ReportRoutes : ""
func (route *Route) ReportRoutes(r *mux.Router) {

	r.Handle("/report/legacy/dumphistory/daywise/quantity", Adapt(http.HandlerFunc(route.Handler.DayWiseDumphistoryCount))).Methods("POST")
	r.Handle("/report/legacy/dumphistory/monthwise/quantity", Adapt(http.HandlerFunc(route.Handler.MonthWiseDumphistoryCount))).Methods("POST")
	r.Handle("/report/swm/circle/housevisited", Adapt(http.HandlerFunc(route.Handler.CircleWiseHouseVisitedCount))).Methods("POST")
	r.Handle("/report/swm/ward/housevisited", Adapt(http.HandlerFunc(route.Handler.DayWiseWardHouseVisitedCount))).Methods("POST")
	//	r.Handle("/report/swm/qrscan/housevisited", Adapt(http.HandlerFunc(route.Handler.DayWiseWardHouseVisitedCount))).Methods("POST")
}
