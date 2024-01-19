package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TradeLicenseRoutes(r *mux.Router) {
	// TradeLicense
	r.Handle("/tradelicense", Adapt(http.HandlerFunc(route.Handler.SaveTradeLicense))).Methods("POST")
	r.Handle("/tradelicense", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicense))).Methods("GET")
	r.Handle("/tradelicense", Adapt(http.HandlerFunc(route.Handler.UpdateTradeLicense))).Methods("PUT")
	r.Handle("/tradelicense/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTradeLicense))).Methods("PUT")
	r.Handle("/tradelicense/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTradeLicense))).Methods("PUT")
	r.Handle("/tradelicense/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTradeLicense))).Methods("DELETE")
	r.Handle("/tradelicense/status/reject", Adapt(http.HandlerFunc(route.Handler.RejectedTradeLicense))).Methods("PUT")
	r.Handle("/tradelicense/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicense))).Methods("POST")

	// TradeLicenseV2
	r.Handle("/tradelicense/v2", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseV2))).Methods("GET")
	r.Handle("/tradelicense/pdf", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicensePDF))).Methods("GET")
	r.Handle("/tradelicense/certificate", Adapt(http.HandlerFunc(route.Handler.GetTradeLicensePaymentReceiptsPDFV2))).Methods("GET")

	// Basic TradeLicense Update
	r.Handle("/tradelicense/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicUpdateTradeLicense))).Methods("PUT")
	r.Handle("/tradelicense/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptBasicTradeLicenseUpdate))).Methods("PUT")
	r.Handle("/tradelicense/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectBasicTradeLicenseUpdate))).Methods("PUT")
	r.Handle("/tradelicense/basicupdate/filter", Adapt(http.HandlerFunc(route.Handler.FilterBasicTradeLicenseUpdateLog))).Methods("POST")
	r.Handle("/tradelicense/basicupdate/getsingle", Adapt(http.HandlerFunc(route.Handler.GetSingleBasicTradeLicenseUpdateLogV2))).Methods("GET")

	// Approve & Reject
	r.Handle("/tradelicense/approve", Adapt(http.HandlerFunc(route.Handler.ApproveTradeLicense))).Methods("PUT")
	r.Handle("/tradelicense/verify", Adapt(http.HandlerFunc(route.Handler.VerifyTradeLicense))).Methods("PUT")
	r.Handle("/tradelicense/notapprove", Adapt(http.HandlerFunc(route.Handler.NotApproveTradeLicense))).Methods("PUT")

	//online payment
	r.Handle("/tradelicense/online/payment", Adapt(http.HandlerFunc(route.Handler.GetOnlinePaymentTL))).Methods("GET", "POST")
	r.Handle("/tradelicense/online/payment/test", Adapt(http.HandlerFunc(route.Handler.GetOnlinePaymentTLtest))).Methods("GET", "POST")
	r.Handle("/tradelicense/online/payment/failed", Adapt(http.HandlerFunc(route.Handler.GetFailedOnlinePaymentTL))).Methods("GET", "POST")

}

func (route *Route) TradeLicenseDemandRoutes(r *mux.Router) {
	r.Handle("/tradelicense/demand", Adapt(http.HandlerFunc(route.Handler.CalcTradeLicenseDemand))).Methods("GET")

}

func (route *Route) TradeLicensePaymenRoutes(r *mux.Router) {
	r.Handle("/tradelicense/payment/initiatepayment", Adapt(http.HandlerFunc(route.Handler.InitiateTradeLicensePayment))).Methods("POST")
	r.Handle("/tradelicense/payment/getpayment", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicensePayment))).Methods("GET")
	r.Handle("/tradelicense/payment/getpaymentreceipt", Adapt(http.HandlerFunc(route.Handler.GetTradeLicensePaymentReceiptsPDF))).Methods("GET")
	r.Handle("/tradelicense/payment/makepayment", Adapt(http.HandlerFunc(route.Handler.MakeTradeLicensePayment))).Methods("POST")
	r.Handle("/tradelicense/payment/verify", Adapt(http.HandlerFunc(route.Handler.VerifyTradeLicensePayment))).Methods("PUT")
	r.Handle("/tradelicense/payment/notverify", Adapt(http.HandlerFunc(route.Handler.NotVerifyTradeLicensePayment))).Methods("PUT")
	r.Handle("/tradelicense/payment/reject", Adapt(http.HandlerFunc(route.Handler.RejectTradeLicensePayment))).Methods("PUT")
	r.Handle("/tradelicense/payment/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicensePayment))).Methods("POST")
	r.Handle("/tradelicense/basicupdate/getpayments", Adapt(http.HandlerFunc(route.Handler.BasicTradeLicenseUpdateGetPaymentsToBeUpdated))).Methods("POST")
	r.Handle("/tradelicense/payment/bouncepayment", Adapt(http.HandlerFunc(route.Handler.TradeLicenseBouncePayment))).Methods("PUT")

	r.Handle("/tradelicense/payment/daterangewise/report", Adapt(http.HandlerFunc(route.Handler.DateRangeWiseTradeLisencePaymentReport))).Methods("POST")
	r.Handle("/tradelicense/dashboard/saf/report", Adapt(http.HandlerFunc(route.Handler.GetTradeLicenseSAFDashboard))).Methods("POST")

}
