package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserChargeMonthlyDemandRoutes(r *mux.Router) {
	r.Handle("/usercharge/monthly/payment/initiatepayment", Adapt(http.HandlerFunc(route.Handler.InitiateUserChargeMonthlyPayment))).Methods("POST")
	r.Handle("/usercharge/monthly/payment/getpaymentreceipt", Adapt(http.HandlerFunc(route.Handler.GetUserChargeMonthlyPaymentReceiptsPDF))).Methods("GET")
	r.Handle("/usercharge/monthly/payment/getpayment", Adapt(http.HandlerFunc(route.Handler.GetUserChargeMonthlyPayment))).Methods("GET")
	r.Handle("/usercharge/monthly/payment/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserChargeMonthlyPayment))).Methods("POST")
	r.Handle("/usercharge/property/makepayment", Adapt(http.HandlerFunc(route.Handler.MakeUserChargePayment))).Methods("POST")
	r.Handle("/usercharge/payment/bouncepayment", Adapt(http.HandlerFunc(route.Handler.UserChargeBouncePayment))).Methods("PUT")
	r.Handle("/usercharge/payment/verifypayment", Adapt(http.HandlerFunc(route.Handler.UserChargVerifyPayment))).Methods("PUT")
	r.Handle("/usercharge/payment/notverifypayment", Adapt(http.HandlerFunc(route.Handler.UserChargNotVerifiedPayment))).Methods("PUT")
	r.Handle("/usercharge/payment/rejectpayment", Adapt(http.HandlerFunc(route.Handler.UserChargRejectPayment))).Methods("PUT")
	r.Handle("/usercharge/payment/fortransaction", Adapt(http.HandlerFunc(route.Handler.GetSingleUserChargePaymentTxtID))).Methods("GET")

	//dashboard
	r.Handle("/usercharge/payment/daterangewise/report", Adapt(http.HandlerFunc(route.Handler.DateRangeWiseUserchargePaymentReport))).Methods("POST")

}
