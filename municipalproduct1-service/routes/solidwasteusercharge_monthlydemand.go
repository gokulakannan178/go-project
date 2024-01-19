package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) SolidWasteUserChargeMonthlyDemandRoutes(r *mux.Router) {
	r.Handle("/solidwasteusercharge/monthly/demand", Adapt(http.HandlerFunc(route.Handler.GetSolidWasteUserChargeDemand))).Methods("GET")
	r.Handle("/solidwasteusercharge/monthly/payment/initiatepayment", Adapt(http.HandlerFunc(route.Handler.InitiateSolidWasteUserChargeMonthlyPayment))).Methods("POST")
	r.Handle("/solidwasteusercharge/monthly/payment/getpayment", Adapt(http.HandlerFunc(route.Handler.GetSingleSolidWasteUserChargePayment))).Methods("GET")
	r.Handle("/solidwasteusercharge/payment/makepayment", Adapt(http.HandlerFunc(route.Handler.MakeSolidWasteUserChargePayment))).Methods("POST")
	r.Handle("/solidwasteusercharge/payment/filter", Adapt(http.HandlerFunc(route.Handler.FilterSolidWasteUserChargePayment))).Methods("POST")
	r.Handle("/solidwasteusercharge/monthly/payment/getpaymentreceipt", Adapt(http.HandlerFunc(route.Handler.GetSolidWasteUserChargePaymentReceiptsPDF))).Methods("GET")
	r.Handle("/solidwasteusercharge/payment/verify", Adapt(http.HandlerFunc(route.Handler.VerifySolidWasteUserChargePayment))).Methods("PUT")
	r.Handle("/solidwasteusercharge/payment/notverify", Adapt(http.HandlerFunc(route.Handler.NotVerifySolidWasteUserChargePayment))).Methods("PUT")
	r.Handle("/solidwasteusercharge/payment/reject", Adapt(http.HandlerFunc(route.Handler.RejectSolidWasteUserChargePayment))).Methods("PUT")

}
