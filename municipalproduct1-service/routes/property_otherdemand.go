package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PropertyOtherDemandRoutes(r *mux.Router) {
	// PropertyOtherDemand
	r.Handle("/propertyotherdemand", Adapt(http.HandlerFunc(route.Handler.SavePropertyOtherDemand))).Methods("POST")
	r.Handle("/propertyotherdemand", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyOtherDemand))).Methods("GET")
	r.Handle("/propertyotherdemand", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyOtherDemand))).Methods("PUT")
	r.Handle("/propertyotherdemand/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyOtherDemand))).Methods("PUT")
	r.Handle("/propertyotherdemand/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyOtherDemand))).Methods("PUT")
	r.Handle("/propertyotherdemand/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyOtherDemand))).Methods("DELETE")
	r.Handle("/propertyotherdemand/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyOtherDemand))).Methods("POST")

	// Property Other Demand Payments
	r.Handle("/propertyotherdemand/payment/initiate", Adapt(http.HandlerFunc(route.Handler.InitiatePropertyOtherDemandPayment))).Methods("POST")
	r.Handle("/propertyotherdemand/payment/fortransaction", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyOtherDemandPaymentTxtID))).Methods("GET")
	r.Handle("/propertyotherdemand/makepayment", Adapt(http.HandlerFunc(route.Handler.PropertyOtherDemandMakePayment))).Methods("POST")
	r.Handle("/propertyotherdemand/payment/verifypayment", Adapt(http.HandlerFunc(route.Handler.PropertyOtherDemandVerifyPayment))).Methods("PUT")
	r.Handle("/propertyotherdemand/payment/notverifypayment", Adapt(http.HandlerFunc(route.Handler.PropertyOtherDemandNotVerifiedPayment))).Methods("PUT")
	r.Handle("/propertyotherdemand/payment/rejectpayment", Adapt(http.HandlerFunc(route.Handler.PropertyOtherDemandRejectPayment))).Methods("PUT")
	r.Handle("/propertyotherdemand/payment/generaterecipt", Adapt(http.HandlerFunc(route.Handler.GetPropertyOtherDemandPaymentReceiptsPDF))).Methods("GET")

}
