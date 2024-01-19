package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MobileTowerReportRoute : ""
func (route *Route) MobileTowerReportRoute(r *mux.Router) {
	r.Handle("/mobiletower/report/daywisecollection", Adapt(http.HandlerFunc(route.Handler.DashboardDayWiseMobileTowerCollectionChart))).Methods("POST")
	r.Handle("/mobiletower/report/daywisedemand", Adapt(http.HandlerFunc(route.Handler.DayWiseMobileTowerDemandChart))).Methods("POST")
	r.Handle("/mobiletower/report/overalldemand", Adapt(http.HandlerFunc(route.Handler.MobileTowerOverallDemandReport))).Methods("POST")

	r.Handle("/mobiletower/daywiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardDayWiseMobileTowerCollectionReport))).Methods("POST")
	r.Handle("/mobiletower/monthwiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardMonthWiseMobileTowerCollectionReport))).Methods("POST")
	r.Handle("/mobiletower/yearwiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardYearWiseMobileTowerCollectionReport))).Methods("POST")

	r.Handle("/mobiletower/daywiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardDayWiseMobileTowerDemandReport))).Methods("POST")
	r.Handle("/mobiletower/monthwiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardMonthWiseMobileTowerDemandReport))).Methods("POST")
	r.Handle("/mobiletower/yearwiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardYearWiseMobileTowerDemandReport))).Methods("POST")

	r.Handle("/mobiletower/daywiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamDayWiseMobileTowerCollectionReport))).Methods("POST")
	r.Handle("/mobiletower/monthwiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamMonthWiseMobileTowerCollectionReport))).Methods("POST")
	r.Handle("/mobiletower/yearwiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamYearWiseMobileTowerCollectionReport))).Methods("POST")

}
