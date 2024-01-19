package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ShopRentMonthlyDemandRoutes(r *mux.Router) {
	r.Handle("/shoprent/monthly/demand", Adapt(http.HandlerFunc(route.Handler.CalcShopRentMonthlyDemand))).Methods("GET")
	r.Handle("/shoprent/monthly/payment/initiatepayment", Adapt(http.HandlerFunc(route.Handler.InitiateShopRentMonthlyPayment))).Methods("POST")
	r.Handle("/shoprent/monthly/payment/getpaymentreceipt", Adapt(http.HandlerFunc(route.Handler.GetShopRentMonthlyPaymentReceiptsPDF))).Methods("GET")
	r.Handle("/shoprent/monthly/payment/getpayment", Adapt(http.HandlerFunc(route.Handler.GetShopRentMonthlyPayment))).Methods("GET")
	r.Handle("/shoprent/monthly/payment/filter", Adapt(http.HandlerFunc(route.Handler.FilterShopRentMonthlyPayment))).Methods("POST")

}
