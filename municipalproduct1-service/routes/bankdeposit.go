package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CreatedBankDeposit(r *mux.Router) {
	r.Handle("/bankdeposit", Adapt(http.HandlerFunc(route.Handler.CreatedBankDeposit))).Methods("POST")
	r.Handle("/bankdeposit/verify", Adapt(http.HandlerFunc(route.Handler.VerifyBankDeposit))).Methods("POST")
	r.Handle("/bankdeposit/notverify", Adapt(http.HandlerFunc(route.Handler.NotVerifyBankDeposit))).Methods("POST")
	r.Handle("/bankdepositfilter", Adapt(http.HandlerFunc(route.Handler.BankDepositFilter))).Methods("POST")
	r.Handle("/bankdeposit", Adapt(http.HandlerFunc(route.Handler.GetSingleBankDeposit))).Methods("GET")
	r.Handle("/bankdepositcollectionsubmissionrequest", Adapt(http.HandlerFunc(route.Handler.CollectionSubmissionRequest))).Methods("POST")
	r.Handle("/bankdepositcollectionsubmissionrequest/filter", Adapt(http.HandlerFunc(route.Handler.CollectionSubmissionRequestFilter))).Methods("POST")

}
