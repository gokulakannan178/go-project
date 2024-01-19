package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TradeLicenseReportRoute : ""
func (route *Route) TradeLicenseReportRoute(r *mux.Router) {
	r.Handle("/tradelicense/report/daywisecollection", Adapt(http.HandlerFunc(route.Handler.DashboardDayWiseTradelicenseCollectionChart))).Methods("POST")
	r.Handle("/tradelicense/report/daywisedemand", Adapt(http.HandlerFunc(route.Handler.DayWiseTradelicenseDemandChart))).Methods("POST")

	r.Handle("/tradelicense/report/overalldemandreport", Adapt(http.HandlerFunc(route.Handler.TradeLicenseOverallDemandReport))).Methods("POST")

	r.Handle("/tradelicense/daywiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardDayWiseTradeLicenseCollectionReport))).Methods("POST")
	r.Handle("/tradelicense/monthwiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardMonthWiseTradeLicenseCollectionReport))).Methods("POST")
	r.Handle("/tradelicense/yearwiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardYearWiseTradeLicenseCollectionReport))).Methods("POST")

	r.Handle("/tradelicense/daywiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardDayWiseTradeLicenseDemandReport))).Methods("POST")
	r.Handle("/tradelicense/monthwiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardMonthWiseTradeLicenseDemandReport))).Methods("POST")
	r.Handle("/tradelicense/yearwiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardYearWiseTradeLicenseDemandReport))).Methods("POST")

	r.Handle("/tradelicense/daywiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamDayWiseTradeLicenseCollectionReport))).Methods("POST")
	r.Handle("/tradelicense/monthwiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamMonthWiseTradeLicenseCollectionReport))).Methods("POST")
	r.Handle("/tradelicense/yearwiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamYearWiseTradeLicenseCollectionReport))).Methods("POST")

}
