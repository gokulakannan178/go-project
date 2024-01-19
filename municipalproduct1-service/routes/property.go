package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PropertyRoutes : ""
func (route *Route) PropertyRoutes(r *mux.Router) {
	r.Handle("/property", Adapt(http.HandlerFunc(route.Handler.SaveProperty))).Methods("POST")
	r.Handle("/property/v2", Adapt(http.HandlerFunc(route.Handler.SavePropertyV2))).Methods("POST")
	r.Handle("/property", Adapt(http.HandlerFunc(route.Handler.GetSingleProperty))).Methods("GET")
	r.Handle("/property", Adapt(http.HandlerFunc(route.Handler.UpdateProperty))).Methods("PUT")
	r.Handle("/property/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableProperty))).Methods("PUT")
	r.Handle("/property/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableProperty))).Methods("PUT")
	r.Handle("/property/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteProperty))).Methods("DELETE")
	r.Handle("/property/filter", Adapt(http.HandlerFunc(route.Handler.FilterProperty))).Methods("POST")
	r.Handle("/property/activate", Adapt(http.HandlerFunc(route.Handler.ActivateProperty))).Methods("PUT")
	r.Handle("/property/reject", Adapt(http.HandlerFunc(route.Handler.RejectProperty))).Methods("PUT")
	r.Handle("/property/getdemand", Adapt(http.HandlerFunc(route.Handler.GetPropertyDemandCalc))).Methods("GET")
	r.Handle("/property/getalldemand", Adapt(http.HandlerFunc(route.Handler.GetAllPropertyDemandCalc))).Methods("GET")
	r.Handle("/property/getdemandwithstoredcalc", Adapt(http.HandlerFunc(route.Handler.GetPropertyDemandCalcWithStoredCalc))).Methods("GET")
	r.Handle("/property/holdingstatus/enable", Adapt(http.HandlerFunc(route.Handler.EnableHoldingProperty))).Methods("PUT")
	r.Handle("/property/holdingstatus/disable", Adapt(http.HandlerFunc(route.Handler.DisableHoldingProperty))).Methods("PUT")
	r.Handle("/property/checkoldholdingno", Adapt(http.HandlerFunc(route.Handler.CheckWardWisesOldHoldingNoOfProperty))).Methods("GET")

	r.Handle("/property/getdemand/v2", Adapt(http.HandlerFunc(route.Handler.GetPropertyDemandCalcV2))).Methods("GET")
	// r.Handle("/property/getdemand/pdf", Adapt(http.HandlerFunc(route.Handler.GetPropertyDemandCalcPDF))).Methods("GET")
	r.Handle("/property/getdemand/pdf", Adapt(http.HandlerFunc(route.Handler.GetPropertyDemandCalcPDFV2))).Methods("GET")
	r.Handle("/property/getdemand/pdf/v2", Adapt(http.HandlerFunc(route.Handler.GetPropertyDemandCalcPDFV2))).Methods("GET")
	// r.Handle("/property/getpaymentrecipt/pdf", Adapt(http.HandlerFunc(route.Handler.GetPaymentReceiptsPDF))).Methods("GET")
	r.Handle("/property/getpaymentrecipt/pdf", Adapt(http.HandlerFunc(route.Handler.GetPaymentReceiptsPDFV2))).Methods("GET")
	r.Handle("/pr", Adapt(http.HandlerFunc(route.Handler.GetPaymentReceiptsPDFV2))).Methods("GET")
	r.Handle("/property/getpaymentrecipt/pdfv2", Adapt(http.HandlerFunc(route.Handler.GetPaymentReceiptsPDFV2))).Methods("GET")

	r.Handle("/property/getdemand/fys", Adapt(http.HandlerFunc(route.Handler.GetPropertyDemandCalcForFYs))).Methods("POST")
	r.Handle("/property/getdemand/multiple", Adapt(http.HandlerFunc(route.Handler.GetMultiplePropertyDemandCalc))).Methods("POST")
	r.Handle("/property/getpaymentrecipt/pdfsave", Adapt(http.HandlerFunc(route.Handler.GETPaymentReceiptsPDFFilesaved))).Methods("POST")
	r.Handle("/property/savepaymentrecipt/pdfsave", Adapt(http.HandlerFunc(route.Handler.SavePaymentReceiptsPDFFilesaved))).Methods("POST")

	r.Handle("/property/getdemands", Adapt(http.HandlerFunc(route.Handler.DemandCalc))).Methods("POST")
	r.Handle("/property/dashboard/status", Adapt(http.HandlerFunc(route.Handler.DashboardPropertyStatus))).Methods("POST")
	r.Handle("/property/dashboard/collection", Adapt(http.HandlerFunc(route.Handler.DashboardTotalCollection))).Methods("POST")
	r.Handle("/property/dashboard/chart/collection", Adapt(http.HandlerFunc(route.Handler.DashboardTotalCollectionChart))).Methods("POST")
	r.Handle("/property/dashboard/collection/v2", Adapt(http.HandlerFunc(route.Handler.DashboardTotalCollectionOverview))).Methods("POST")
	r.Handle("/property/dashboard/chart/daywisecollection", Adapt(http.HandlerFunc(route.Handler.DashboardDayWiseCollectionChart))).Methods("POST")
	r.Handle("/property/dashboard/ddac", Adapt(http.HandlerFunc(route.Handler.DashboardDemandAndCollection))).Methods("POST")
	r.Handle("/property/dashboard/ddac/v2", Adapt(http.HandlerFunc(route.Handler.DashboardDemandAndCollectionV2))).Methods("POST")

	r.Handle("/property/report/wardwisecollection", Adapt(http.HandlerFunc(route.Handler.WardWiseCollectionReport))).Methods("POST")
	r.Handle("/property/report/tccollectionsummary", Adapt(http.HandlerFunc(route.Handler.TCCollectionSummaryReport))).Methods("POST")
	r.Handle("/property/demand", Adapt(http.HandlerFunc(route.Handler.PropertyDemand))).Methods("POST")
	r.Handle("/property/gistagging", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyGISTagging))).Methods("PUT")
	r.Handle("/property/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateProperty))).Methods("PUT")
	r.Handle("/property/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptBasicPropertyUpdate))).Methods("PUT")
	r.Handle("/property/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectBasicPropertyUpdate))).Methods("PUT")
	r.Handle("/property/basicupdate/filter", Adapt(http.HandlerFunc(route.Handler.FilterBasicPropertyUpdate))).Methods("POST")
	r.Handle("/property/basicupdate", Adapt(http.HandlerFunc(route.Handler.GetSingleBasicPropertyUpdateLog))).Methods("GET")
	r.Handle("/property/basicupdate/getpayments", Adapt(http.HandlerFunc(route.Handler.BasicPropertyUpdateGetPaymentsToBeUpdated))).Methods("POST")
	r.Handle("/property/parkpenalty/enable", Adapt(http.HandlerFunc(route.Handler.PropertyParkPenaltyEnable))).Methods("PUT")
	r.Handle("/property/parkpenalty/disable", Adapt(http.HandlerFunc(route.Handler.PropertyParkPenaltyDisable))).Methods("PUT")
	r.Handle("/property/location/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyLocation))).Methods("PUT")
	r.Handle("/property/picture/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyPicture))).Methods("PUT")

	//Payments
	r.Handle("/property/payment/initiate", Adapt(http.HandlerFunc(route.Handler.InitiatePropertyPayment))).Methods("POST")
	r.Handle("/property/payment/fortransaction", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyPaymentTxtID))).Methods("GET")
	r.Handle("/property/makepayment", Adapt(http.HandlerFunc(route.Handler.PropertyMakePayment))).Methods("POST")
	r.Handle("/property/getpayments", Adapt(http.HandlerFunc(route.Handler.GetAllPaymentsForProperty))).Methods("GET")
	r.Handle("/property/payment/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyPayment))).Methods("POST")
	r.Handle("/property/payment/generaterecipt", Adapt(http.HandlerFunc(route.Handler.GenerateReciptForAPayment))).Methods("GET")
	r.Handle("/property/payment/verifypayment", Adapt(http.HandlerFunc(route.Handler.VerifyPayment))).Methods("PUT")
	r.Handle("/property/payment/notverifypayment", Adapt(http.HandlerFunc(route.Handler.NotVerifiedPayment))).Methods("PUT")
	r.Handle("/property/report/wardwisedemand", Adapt(http.HandlerFunc(route.Handler.WardwiseDemand))).Methods("POST")
	r.Handle("/property/propertywisedemandandcollection", Adapt(http.HandlerFunc(route.Handler.PropertyWiseDemandandCollectionExcel))).Methods("POST")
	r.Handle("/property/report/wardwisecollectionv2", Adapt(http.HandlerFunc(route.Handler.WardwiseCollection))).Methods("POST")
	r.Handle("/property/report/arrerandcurrentcollection", Adapt(http.HandlerFunc(route.Handler.PropertyPaymentArrerAndCurrentCollection))).Methods("POST")

	//Report
	r.Handle("/property/chequereport", Adapt(http.HandlerFunc(route.Handler.ChequeReport))).Methods("POST")
	r.Handle("/property/counterreport", Adapt(http.HandlerFunc(route.Handler.CounterReport))).Methods("POST")
	// r.Handle("/property/teamwisecollection", Adapt(http.HandlerFunc(route.Handler.TeamWiseCollectionReport))).Methods("POST")
	r.Handle("/property/idwisereport", Adapt(http.HandlerFunc(route.Handler.PropertyIDWiseReport))).Methods("POST")
	r.Handle("/property/zoneandwardwisecollectionreport", Adapt(http.HandlerFunc(route.Handler.ZoneAndWardWiseCollection))).Methods("POST")
	//prev
	r.Handle("/property/payment/previousyrcollection", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyPreviousYrCollection))).Methods("PUT")
	//reject
	r.Handle("/property/payment/rejectpayment", Adapt(http.HandlerFunc(route.Handler.RejectPayment))).Methods("PUT")
	r.Handle("/property/payment/rejectpayment/byreceiptno", Adapt(http.HandlerFunc(route.Handler.RejectPaymentByReceiptNo))).Methods("PUT")
	r.Handle("/property/payment/bouncepayment", Adapt(http.HandlerFunc(route.Handler.BouncePayment))).Methods("PUT")

	// yearwise collection report for a property
	r.Handle("/property/payment/yearwise/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterPropertyMonthWiseCollectionReport))).Methods("POST")

	r.Handle("/property/payment/daterangewise/report", Adapt(http.HandlerFunc(route.Handler.DateRangeWisePropertyPaymentReport))).Methods("POST")
	r.Handle("/property/location/update/report", Adapt(http.HandlerFunc(route.Handler.PropertyUpdateLocationReport))).Methods("POST")
	r.Handle("/property/payment/datewise/report", Adapt(http.HandlerFunc(route.Handler.DateWisePropertyPaymentReport))).Methods("POST")

	// Online Payment
	r.Handle("/online/payment", Adapt(http.HandlerFunc(route.Handler.GetOnlinePayment))).Methods("GET", "POST")
	r.Handle("/online/payment/failed", Adapt(http.HandlerFunc(route.Handler.GetFailedOnlinePayment))).Methods("GET", "POST")

	// r.Handle("/online/payment", Adapt(http.HandlerFunc(route.Handler.PutOnlinePayment))).Methods("PUT")
	// r.Handle("/online/payment", Adapt(http.HandlerFunc(route.Handler.PostOnlinePayment))).Methods("POST")

	//Bulk Selected Transaction
	r.Handle("/property/collection/collect", Adapt(http.HandlerFunc(route.Handler.CollectedPropertyPayment))).Methods("PUT")
	r.Handle("/property/collection/reject", Adapt(http.HandlerFunc(route.Handler.RejectedPropertyPayment))).Methods("PUT")

	r.Handle("/property/collection/report", Adapt(http.HandlerFunc(route.Handler.FilterPropertyWiseCollectionReport))).Methods("POST")
	r.Handle("/property/demand/report", Adapt(http.HandlerFunc(route.Handler.FilterPropertyWiseDemandReport))).Methods("POST")
	r.Handle("/hdfc/pg", Adapt(http.HandlerFunc(route.Handler.PostReq)))
	r.Handle("/hdfc/pg/initiate", Adapt(http.HandlerFunc(route.Handler.GetReq))).Methods("GET")

	// property new uniqueId implementation
	r.Handle("/property/uniqueid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyUniqueID))).Methods("PUT")
	r.Handle("/owner/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdateOwnerPropertyID))).Methods("PUT")
	r.Handle("/floor/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdateFloorPropertyID))).Methods("PUT")
	r.Handle("/basicpropertyupdatelog/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdateBasicPropertyUpdateLogPropertyID))).Methods("PUT")
	r.Handle("/mobiletower/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdateMobileTowerPropertyID))).Methods("PUT")
	r.Handle("/overallpropertydemand/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdateOverAllPropertyDemandPropertyID))).Methods("PUT")
	r.Handle("/propertydemandlog/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyDemandLogPropertyID))).Methods("PUT")
	r.Handle("/propertyfydemandlog/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyFYDemandLogPropertyID))).Methods("PUT")
	r.Handle("/propertypaymentbasic/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyPaymentBasicPropertyID))).Methods("PUT")
	r.Handle("/propertypaymentfy/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyPaymentFYPropertyID))).Methods("PUT")
	r.Handle("/propertypaymentmodechange/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyPaymentModeChangePropertyID))).Methods("PUT")
	r.Handle("/propertypayment/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyPaymentPropertyID))).Methods("PUT")
	r.Handle("/propertyvisitlog/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyVisitLogPropertyID))).Methods("PUT")
	r.Handle("/reassessmentrequest/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyReassessmentRequestPropertyID))).Methods("PUT")
	r.Handle("/propertymutationrequest/propertyid/update", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyMutationRequestPropertyID))).Methods("PUT")

	//usercharge
	//r.Handle("/property/usercharge", Adapt(http.HandlerFunc(route.Handler.CreateUserChargeForProperty))).Methods("POST")
	//summary Update
	r.Handle("/property/payment/summary/update", Adapt(http.HandlerFunc(route.Handler.PropertyPaymentSummaryUpdate))).Methods("POST")

}
