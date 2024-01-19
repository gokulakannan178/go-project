package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TradeLicensePaymentsPart2Routes : ""
func (route *Route) TradeLicensePaymentsPart2Routes(r *mux.Router) {
	r.Handle("/tradelicense/payment/initiate/part2", Adapt(http.HandlerFunc(route.Handler.InitiateTradeLicensePaymentPart2))).Methods("POST")
	r.Handle("/tradelicense/payment/getpayment/part2", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicensePaymentPart2))).Methods("GET")
	r.Handle("/tradelicense/payment/makepayment/part2", Adapt(http.HandlerFunc(route.Handler.MakeTradeLicensePaymentPart2))).Methods("POST")
	r.Handle("/tradelicense/payment/filter/part2", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicensePaymentPart2))).Methods("POST")
	r.Handle("/tradelicense/payment/getpaymentreceipt/part2", Adapt(http.HandlerFunc(route.Handler.GetTradeLicensePaymentReceiptsPDFPart2))).Methods("GET")
	r.Handle("/tradelicense/certificate/part2", Adapt(http.HandlerFunc(route.Handler.GetTradeLicensePaymentReceiptsPDFV2Part2))).Methods("GET")
	r.Handle("/tradelicense/payment/part2/verify", Adapt(http.HandlerFunc(route.Handler.VerifyTradeLicensePaymentPart2))).Methods("PUT")
	r.Handle("/tradelicense/payment/part2/notverify", Adapt(http.HandlerFunc(route.Handler.NotVerifyTradeLicensePaymentPart2))).Methods("PUT")
	r.Handle("/tradelicense/payment/part2/reject", Adapt(http.HandlerFunc(route.Handler.RejectTradeLicensePaymentPart2))).Methods("PUT")
	r.Handle("/tradelicense/v2/part2", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseV2Part2))).Methods("GET")
	r.Handle("/tradelicense/basicupdate/getpayments/part2", Adapt(http.HandlerFunc(route.Handler.BasicTradeLicenseUpdateGetPaymentsToBeUpdatedPart2))).Methods("POST")

}
