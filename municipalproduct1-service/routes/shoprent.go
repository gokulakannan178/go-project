package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ShopRentRoutes : ""
func (route *Route) ShopRentRoutes(r *mux.Router) {
	r.Handle("/shoprent", Adapt(http.HandlerFunc(route.Handler.SaveShopRent))).Methods("POST")
	r.Handle("/shoprent", Adapt(http.HandlerFunc(route.Handler.GetSingleShopRent))).Methods("GET")
	r.Handle("/shoprent", Adapt(http.HandlerFunc(route.Handler.UpdateShopRent))).Methods("PUT")
	r.Handle("/shoprent/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableShopRent))).Methods("PUT")
	r.Handle("/shoprent/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableShopRent))).Methods("PUT")
	r.Handle("/shoprent/status/reject", Adapt(http.HandlerFunc(route.Handler.RejectedShopRent))).Methods("PUT")
	r.Handle("/shoprent/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteShopRent))).Methods("DELETE")
	r.Handle("/shoprent/filter", Adapt(http.HandlerFunc(route.Handler.FilterShopRent))).Methods("POST")

	r.Handle("/shoprent/report/wardwisecollection", Adapt(http.HandlerFunc(route.Handler.WardWiseShoprentReport))).Methods("POST")

	//reject
	r.Handle("/shoprent/payment/bouncepayment", Adapt(http.HandlerFunc(route.Handler.ShoprentBouncePayment))).Methods("PUT")

}
func (route *Route) ShopRentCalcRoutes(r *mux.Router) {
	r.Handle("/shoprent/calc/demandandcollection", Adapt(http.HandlerFunc(route.Handler.UpdateShopRentDemandAndCollections))).Methods("PUT")

}

func (route *Route) ShopRentDemandRoutes(r *mux.Router) {
	r.Handle("/shoprent/demand", Adapt(http.HandlerFunc(route.Handler.CalcShopRentDemand))).Methods("GET")

}

func (route *Route) ShopRentPaymenRoutes(r *mux.Router) {
	r.Handle("/shoprent/payment/initiatepayment", Adapt(http.HandlerFunc(route.Handler.InitiateShopRentPayment))).Methods("POST")
	r.Handle("/shoprent/legacypayment/initiatepayment", Adapt(http.HandlerFunc(route.Handler.InitiateShopRentMonthlyLegacyPayment))).Methods("POST")
	r.Handle("/shoprent/payment/getpayment", Adapt(http.HandlerFunc(route.Handler.GetSingleShopRentPayment))).Methods("GET")
	r.Handle("/shoprent/payment/getpaymentreceipt", Adapt(http.HandlerFunc(route.Handler.GetShopRentPaymentReceiptsPDF))).Methods("GET")
	r.Handle("/shoprent/payment/makepayment", Adapt(http.HandlerFunc(route.Handler.MakeShopRentPayment))).Methods("POST")
	r.Handle("/shoprent/legacypayment/makepayment", Adapt(http.HandlerFunc(route.Handler.MakeShopRentMonthlyLegacyPayment))).Methods("POST")
	r.Handle("/shoprent/payment/verify", Adapt(http.HandlerFunc(route.Handler.VerifyShopRentPayment))).Methods("PUT")
	r.Handle("/shoprent/payment/notverify", Adapt(http.HandlerFunc(route.Handler.NotVerifyShopRentPayment))).Methods("PUT")
	r.Handle("/shoprent/payment/reject", Adapt(http.HandlerFunc(route.Handler.RejectShopRentPayment))).Methods("PUT")
	r.Handle("/shoprent/payment/filter", Adapt(http.HandlerFunc(route.Handler.FilterShopRentPayment))).Methods("POST")
	r.Handle("/shoprent/basicupdate/getpayments", Adapt(http.HandlerFunc(route.Handler.BasicShopRentUpdateGetPaymentsToBeUpdated))).Methods("POST")

}
