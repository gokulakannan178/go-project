package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PropertyPartPaymentPaymentRoutes(r *mux.Router) {

	r.Handle("/property/partpayment/makepayment", Adapt(http.HandlerFunc(route.Handler.SavePropertyPartPayment))).Methods("POST")
	r.Handle("/property/partpayment/additionalpayment", Adapt(http.HandlerFunc(route.Handler.SavePropertyPartPaymentAdditional))).Methods("POST")

	r.Handle("/property/partpayment/validate", Adapt(http.HandlerFunc(route.Handler.ValidatePartPayments))).Methods("POST")
	r.Handle("/property/partpayment/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyPartPayment))).Methods("POST")
	r.Handle("/property/partpayment/payment/verify", Adapt(http.HandlerFunc(route.Handler.VerifyPropertyPartPayment))).Methods("PUT")
	r.Handle("/property/partpayment/payment/notverify", Adapt(http.HandlerFunc(route.Handler.NotVerifyPropertyPartPayment))).Methods("PUT")
	r.Handle("/property/partpayment/payment/reject", Adapt(http.HandlerFunc(route.Handler.RejectPropertyPartPayment))).Methods("PUT")

	r.Handle("/payment/partpayment", Adapt(http.HandlerFunc(route.Handler.GetPropertyPaymentsWithPartPayments))).Methods("GET")

}
