package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//DealerRoutes : ""
func (route *Route) DealerRoutes(r *mux.Router) {
	r.Handle("/dealer", Adapt(http.HandlerFunc(route.Handler.SaveDealer))).Methods("POST")
	r.Handle("/dealer", Adapt(http.HandlerFunc(route.Handler.GetSingleDealer))).Methods("GET")
	r.Handle("/dealer", Adapt(http.HandlerFunc(route.Handler.UpdateDealer))).Methods("PUT")
	r.Handle("/dealer/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDealer))).Methods("PUT")
	r.Handle("/dealer/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDealer))).Methods("PUT")
	r.Handle("/dealer/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDealer))).Methods("DELETE")
	r.Handle("/dealer/filter", Adapt(http.HandlerFunc(route.Handler.FilterDealer))).Methods("POST")
	r.Handle("/dealer/registration/uniquecheck", Adapt(http.HandlerFunc(route.Handler.DealerUniquenessCheckRegistration))).Methods("GET")
	r.Handle("/dealer/nearby", Adapt(http.HandlerFunc(route.Handler.DealerNearBy))).Methods("POST")
	r.Handle("/dealer/certification/apply", Adapt(http.HandlerFunc(route.Handler.DealerCertificationApply))).Methods("PUT")
	r.Handle("/dealer/certification/approve", Adapt(http.HandlerFunc(route.Handler.DealerCertificationApprove))).Methods("PUT")
	r.Handle("/dealer/certification/Reject", Adapt(http.HandlerFunc(route.Handler.DealerCertificationReject))).Methods("PUT")

}

//ConsumerRegistrationRoutes : ""
func (route *Route) DealerRegistrationRoutes(r *mux.Router) {
	r.Handle("/dealer/registration/generateotp", Adapt(http.HandlerFunc(route.Handler.DealerregistrationGenerateOTP))).Methods("POST")
	r.Handle("/dealer/registration/validateotp", Adapt(http.HandlerFunc(route.Handler.DealerregistrationValidateOTP))).Methods("POST")
}
