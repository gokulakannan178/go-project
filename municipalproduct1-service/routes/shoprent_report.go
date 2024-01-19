package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ShopRentReportRoute : ""
func (route *Route) ShopRentReportRoute(r *mux.Router) {
	r.Handle("/shoprent/report/daywisecollection", Adapt(http.HandlerFunc(route.Handler.DashboardDayWiseShoprentCollectionChart))).Methods("POST")
	r.Handle("/shoprent/report/daywisedemand", Adapt(http.HandlerFunc(route.Handler.DayWiseShoprentDemandChart))).Methods("POST")

	r.Handle("/shoprent/report/overalldemandreport", Adapt(http.HandlerFunc(route.Handler.ShopRentOverallDemandReport))).Methods("POST")

	r.Handle("/shoprent/daywiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardDayWiseShopRentCollectionReport))).Methods("POST")
	r.Handle("/shoprent/monthwiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardMonthWiseShopRentCollectionReport))).Methods("POST")
	r.Handle("/shoprent/yearwiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardYearWiseShopRentCollectionReport))).Methods("POST")

	r.Handle("/shoprent/daywiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardDayWiseShopRentDemandReport))).Methods("POST")
	r.Handle("/shoprent/monthwiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardMonthWiseShopRentDemandReport))).Methods("POST")
	r.Handle("/shoprent/yearwiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardYearWiseShopRentDemandReport))).Methods("POST")

	r.Handle("/shoprent/daywiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamDayWiseShopRentCollectionReport))).Methods("POST")
	r.Handle("/shoprent/monthwiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamMonthWiseShopRentCollectionReport))).Methods("POST")
	r.Handle("/shoprent/yearwiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamYearWiseShopRentCollectionReport))).Methods("POST")

}
