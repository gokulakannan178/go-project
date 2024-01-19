package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//PreregistrationRoutes : ""
func (route *Route) PreregistrationRoutes(r *mux.Router) {
	r.Handle("/preregistration", Adapt(http.HandlerFunc(route.Handler.SavePreregistration))).Methods("POST")
	r.Handle("/preregistration/submit", Adapt(http.HandlerFunc(route.Handler.SubmitPreregistration))).Methods("POST")
	r.Handle("/preregistration/reapply", Adapt(http.HandlerFunc(route.Handler.ReapplyPreregistration))).Methods("POST")
	r.Handle("/preregistration/mobile", Adapt(http.HandlerFunc(route.Handler.GetSinglePreregistration))).Methods("GET")
	r.Handle("/preregistration/uniqueId", Adapt(http.HandlerFunc(route.Handler.GetSinglePreregistrationWithUniqueID))).Methods("GET")
	r.Handle("/preregistration", Adapt(http.HandlerFunc(route.Handler.UpdatePreregistration))).Methods("PUT")
	r.Handle("/preregistration/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePreregistration))).Methods("PUT")
	r.Handle("/preregistration/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePreregistration))).Methods("PUT")
	r.Handle("/preregistration/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePreregistration))).Methods("DELETE")
	r.Handle("/preregistration/filter", Adapt(http.HandlerFunc(route.Handler.FilterPreregistration))).Methods("POST")
	r.Handle("/preregistration/validate/mobile", Adapt(http.HandlerFunc(route.Handler.ValidateMobileNumber))).Methods("GET")
	r.Handle("/preregistration/statuschange", Adapt(http.HandlerFunc(route.Handler.PreregistrationStatusChange))).Methods("POST")
	r.Handle("/preregistration/payment", Adapt(http.HandlerFunc(route.Handler.PaymentPreregistration))).Methods("POST")
	r.Handle("/preregistration/notice/paymentpending", Adapt(http.HandlerFunc(route.Handler.PaymentPendingNoticeForPreRegistration))).Methods("POST")

}

// OTPRoutes : ""
func (route *Route) OTPRoutes(r *mux.Router) {
	r.Handle("/send/otp", Adapt(http.HandlerFunc(route.Handler.ApplicantLoginSendOTP))).Methods("GET")
	r.Handle("/validate/otp", Adapt(http.HandlerFunc(route.Handler.ApplicantLoginValidateOTP))).Methods("POST")
}

//ApplicantRoutes : ""
func (route *Route) ApplicantRoutes(r *mux.Router) {
	r.Handle("/applicant", Adapt(http.HandlerFunc(route.Handler.SaveApplicant))).Methods("POST")
	r.Handle("/applicant", Adapt(http.HandlerFunc(route.Handler.GetSingleApplicant))).Methods("GET")
	r.Handle("/applicant", Adapt(http.HandlerFunc(route.Handler.UpdateApplicant))).Methods("PUT")
	r.Handle("/applicant/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableApplicant))).Methods("PUT")
	r.Handle("/applicant/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableApplicant))).Methods("PUT")
	r.Handle("/applicant/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteApplicant))).Methods("DELETE")
	r.Handle("/applicant/status/blacklist", Adapt(http.HandlerFunc(route.Handler.BlacklistApplicant))).Methods("PUT")
	r.Handle("/applicant/status/lisencecancel", Adapt(http.HandlerFunc(route.Handler.LicenseCancelApplicant))).Methods("PUT")
	r.Handle("/applicant/status/reactivate", Adapt(http.HandlerFunc(route.Handler.ReActivateApplicant))).Methods("PUT")
	r.Handle("/applicant/filter", Adapt(http.HandlerFunc(route.Handler.FilterApplicant))).Methods("POST")
}
