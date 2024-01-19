package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PropertyReportRoutes : ""
func (route *Route) PropertyReportRoutes(r *mux.Router) {

	r.Handle("/property/report/ovealldemand", Adapt(http.HandlerFunc(route.Handler.PropertyOverallDemandReport))).Methods("POST")
	// overall demand
	r.Handle("/property/report/daywisedemand", Adapt(http.HandlerFunc(route.Handler.DayWisePropertyDemandReport))).Methods("POST")

	r.Handle("/property/yearwise/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterYearWisePropertyDemandReport))).Methods("POST")
	r.Handle("/property/yearwise/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterYearWisePropertyCollectionReport))).Methods("POST")

	// ward wise collection
	r.Handle("/property/daywiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardDayWisePropertyCollectionReport))).Methods("POST")
	r.Handle("/property/monthwiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardMonthWisePropertyCollectionReport))).Methods("POST")
	r.Handle("/property/yearwiseward/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardYearWisePropertyCollectionReport))).Methods("POST")
	// ward wise demand
	r.Handle("/property/daywiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardDayWisePropertyDemandReport))).Methods("POST")
	r.Handle("/property/monthwiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardMonthWisePropertyDemandReport))).Methods("POST")
	r.Handle("/property/yearwiseward/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterWardYearWisePropertyDemandReport))).Methods("POST")
	// team wise collection
	r.Handle("/property/daywiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamDayWisePropertyCollectionReport))).Methods("POST")
	r.Handle("/property/monthwiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamMonthWisePropertyCollectionReport))).Methods("POST")
	r.Handle("/property/yearwiseteam/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterTeamYearWisePropertyCollectionReport))).Methods("POST")
	//ward wise collection And Demand
	r.Handle("/property/wardwise/demandcollection/report", Adapt(http.HandlerFunc(route.Handler.FilterWardWisePropertyDemandAndCollectionReport))).Methods("POST")
	// User Wise Property Collection Report
	r.Handle("/property/userwise/collection/report", Adapt(http.HandlerFunc(route.Handler.UserWisePropertyCollectionReport))).Methods("POST")
	// FilterPropertyArrearAndCurrentCollectionReport
	r.Handle("/property/arrearandcurrent/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterPropertyArrearAndCurrentCollectionReport))).Methods("POST")
	r.Handle("/property/counterreport/v2", Adapt(http.HandlerFunc(route.Handler.FilterCounterReportV2))).Methods("POST")
	r.Handle("/property/userwise/paymentcollection", Adapt(http.HandlerFunc(route.Handler.FilterUserWisePropertyCollectionReport))).Methods("POST")
	r.Handle("/property/collection/v2", Adapt(http.HandlerFunc(route.Handler.FilterHoldingWiseCollectionReport))).Methods("POST")
	// r.Handle("property/reassessmentrequest/report", Adapt(http.HandlerFunc(route.Handler.GetSingleReassessmentRequest))).Methods("GET")
	r.Handle("/property/dcreport", Adapt(http.HandlerFunc(route.Handler.FilterPropertyDemandAndCollectionReport))).Methods("POST")
	r.Handle("/property/counterreport/v3", Adapt(http.HandlerFunc(route.Handler.FilterPaymentCOllection))).Methods("POST")
	r.Handle("/property/counterreport/spare/v3", Adapt(http.HandlerFunc(route.Handler.FilterPaymentCOllectionSpare))).Methods("POST")
	r.Handle("/property/collection/v3", Adapt(http.HandlerFunc(route.Handler.FilterHoldingWiseCollectionReportJSONV2))).Methods("POST")
	//r.Handle("/property/dcreport/v2", Adapt(http.HandlerFunc(route.Handler.FilterPropertyDemandAndCollectionReportV2))).Methods("POST")
	r.Handle("/property/propertywisedemandandcollection/v2", Adapt(http.HandlerFunc(route.Handler.PropertyWiseDemandandCollectionExcelV2))).Methods("POST")
	r.Handle("/property/demandcollectionbalance/report", Adapt(http.HandlerFunc(route.Handler.PropertyWiseDemandCollectionandBalanceReport))).Methods("POST")
}
