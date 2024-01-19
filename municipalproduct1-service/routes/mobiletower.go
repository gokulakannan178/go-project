package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) MobileTowerRoutes(r *mux.Router) {

	// Mobile Tower Tax
	r.Handle("/mobiletowertax", Adapt(http.HandlerFunc(route.Handler.SaveMobileTowerTax))).Methods("POST")
	r.Handle("/mobiletowertax", Adapt(http.HandlerFunc(route.Handler.GetSingleMobileTowerTax))).Methods("GET")
	r.Handle("/mobiletowertax", Adapt(http.HandlerFunc(route.Handler.UpdateMobileTowerTax))).Methods("PUT")
	r.Handle("/mobiletowertax/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMobileTowerTax))).Methods("PUT")
	r.Handle("/mobiletowertax/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMobileTowerTax))).Methods("PUT")
	r.Handle("/mobiletowertax/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMobileTowerTax))).Methods("DELETE")
	r.Handle("/mobiletowertax/filter", Adapt(http.HandlerFunc(route.Handler.FilterMobileTowerTax))).Methods("POST")

	// Property Mobile Tower
	r.Handle("/propertymobiletower", Adapt(http.HandlerFunc(route.Handler.SavePropertyMobileTower))).Methods("POST")
	r.Handle("/propertymobiletower", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyMobileTower))).Methods("GET")
	r.Handle("/propertymobiletower", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyMobileTower))).Methods("PUT")
	r.Handle("/propertymobiletower/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyMobileTower))).Methods("PUT")
	r.Handle("/propertymobiletower/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyMobileTower))).Methods("PUT")
	r.Handle("/propertymobiletower/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyMobileTower))).Methods("DELETE")
	r.Handle("/propertymobiletower/status/reject", Adapt(http.HandlerFunc(route.Handler.RejectPropertyMobileTower))).Methods("DELETE")
	r.Handle("/propertymobiletower/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyMobileTower))).Methods("POST")
	r.Handle("/mobiletowerwithmobileno", Adapt(http.HandlerFunc(route.Handler.MobileTowerWithMobileNo))).Methods("POST")
	r.Handle("/mobiletower/penalty", Adapt(http.HandlerFunc(route.Handler.MobileTowerPenaltyUpdate))).Methods("PUT")

	// Basic Mobiletower Update
	r.Handle("/mobiletower/basicupdate", Adapt(http.HandlerFunc(route.Handler.BasicMobileTowerUpdate))).Methods("PUT")
	r.Handle("/mobiletower/basicupdate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptBasicMobileTowerUpdate))).Methods("PUT")
	r.Handle("/mobiletower/basicupdate/reject", Adapt(http.HandlerFunc(route.Handler.RejectBasicMobileTowerUpdate))).Methods("PUT")
	r.Handle("/mobiletower/basicupdate/filter", Adapt(http.HandlerFunc(route.Handler.FilterBasicMobileTowerUpdateLog))).Methods("POST")
	r.Handle("/mobiletower/basicupdate/getsingle", Adapt(http.HandlerFunc(route.Handler.GetSingleBasicMobileTowerUpdateLogV2))).Methods("GET")

}

func (route *Route) MobileTowerDemandRoutes(r *mux.Router) {
	r.Handle("/propertymobiletower/demand/recalculate", Adapt(http.HandlerFunc(route.Handler.ReCalcMobileTowerDemand))).Methods("GET")
	r.Handle("/propertymobiletower/demand", Adapt(http.HandlerFunc(route.Handler.CalcMobileTowerDemand))).Methods("GET")
}

func (route *Route) MobileTowerPaymentRoutes(r *mux.Router) {
	r.Handle("/propertymobiletower/payment/initiatepayment", Adapt(http.HandlerFunc(route.Handler.InitiateMobileTowerPayment))).Methods("POST")
	r.Handle("/propertymobiletower/payment/getpayment", Adapt(http.HandlerFunc(route.Handler.GetSingleMobileTowerPayment))).Methods("GET")
	r.Handle("/propertymobiletower/payment/makepayment", Adapt(http.HandlerFunc(route.Handler.MakeMobileTowerPayment))).Methods("POST")
	r.Handle("/propertymobiletower/payment/filter", Adapt(http.HandlerFunc(route.Handler.FilterMobileTowerPayment))).Methods("POST")
	r.Handle("/propertymobiletower/payment/verify", Adapt(http.HandlerFunc(route.Handler.VerifyMobileTowerPayment))).Methods("PUT")
	r.Handle("/propertymobiletower/payment/notverify", Adapt(http.HandlerFunc(route.Handler.NotVerifyMobileTowerPayment))).Methods("PUT")
	r.Handle("/propertymobiletower/payment/reject", Adapt(http.HandlerFunc(route.Handler.RejectMobileTowerPayment))).Methods("PUT")

	r.Handle("/mobiletower/payment/getpaymentreceipt", Adapt(http.HandlerFunc(route.Handler.GetMobileTowerPaymentReceiptsPDF))).Methods("GET")
	r.Handle("/mobiletower/basicupdate/getpayments", Adapt(http.HandlerFunc(route.Handler.BasicMobileTowerUpdateGetPaymentsToBeUpdated))).Methods("POST")

}

func (route *Route) MobileTowerRegistrationPaymentRoutes(r *mux.Router) {
	r.Handle("/propertymobiletower/payment/registration/initiatepayment", Adapt(http.HandlerFunc(route.Handler.InitiateMobileTowerRegisterPayment))).Methods("POST")
	r.Handle("/propertymobiletower/payment/registartion/makepayment", Adapt(http.HandlerFunc(route.Handler.MakeMobileTowerPaymentForRegistration))).Methods("POST")
	r.Handle("/propertymobiletower/payment/registration/verify", Adapt(http.HandlerFunc(route.Handler.VerifyMobileTowerRegistrationPayment))).Methods("PUT")
	r.Handle("/propertymobiletower/payment/registration/notverify", Adapt(http.HandlerFunc(route.Handler.NotVerifyMobileTowerRegistrationPayment))).Methods("PUT")
	r.Handle("/propertymobiletower/payment/registration/reject", Adapt(http.HandlerFunc(route.Handler.RejectMobileTowerRegistrationPayment))).Methods("PUT")

	r.Handle("/mobiletower/payment/registration/getpaymentreceipt", Adapt(http.HandlerFunc(route.Handler.GetMobileTowerRegistartionPaymentReceiptsPDF))).Methods("GET")

}

func (route *Route) MobileTowerLegacyRegistrationPaymentRoutes(r *mux.Router) {
	r.Handle("/propertymobiletower/payment/legacyregistration", Adapt(http.HandlerFunc(route.Handler.MakeMobileTowerLegacyRegistrationPayment))).Methods("POST")

}

func (route *Route) MobileTowerPaymentReportsRoutes(r *mux.Router) {}
