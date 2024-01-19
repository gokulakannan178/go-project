package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ULBRoutes : ""
func (route *Route) ULBRoutes(r *mux.Router) {
	r.Handle("/ulb", Adapt(http.HandlerFunc(route.Handler.SaveULB))).Methods("POST")
	r.Handle("/ulb", Adapt(http.HandlerFunc(route.Handler.GetSingleULB))).Methods("GET")
	r.Handle("/ulb", Adapt(http.HandlerFunc(route.Handler.UpdateULB))).Methods("PUT")
	r.Handle("/ulb/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableULB))).Methods("PUT")
	r.Handle("/ulb/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableULB))).Methods("PUT")
	r.Handle("/ulb/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteULB))).Methods("DELETE")
	r.Handle("/ulb/filter", Adapt(http.HandlerFunc(route.Handler.FilterULB))).Methods("POST")
	r.Handle("/ulb/updatelocation", Adapt(http.HandlerFunc(route.Handler.UpdateUlbLocation))).Methods("PUT")
	r.Handle("/ulb/nearby", Adapt(http.HandlerFunc(route.Handler.ULBNearBy))).Methods("POST")
	r.Handle("/ulb/nearby/{state}", Adapt(http.HandlerFunc(route.Handler.UlbInTheStateV2))).Methods("GET")
	r.Handle("/ulb/nearby/status", Adapt(http.HandlerFunc(route.Handler.UlbInTheStateV3))).Methods("GET")
	r.Handle("/ulb/compost/{state}", Adapt(http.HandlerFunc(route.Handler.UlbCompostInTheState))).Methods("GET")
	r.Handle("/ulb/nearby/district", Adapt(http.HandlerFunc(route.Handler.ULBNearBy))).Methods("POST")
	r.Handle("/ulb/addtestcert", Adapt(http.HandlerFunc(route.Handler.AddULBTestCert))).Methods("PUT")
	r.Handle("/ulb/testcertificate/apply", Adapt(http.HandlerFunc(route.Handler.ApplyForTestCert))).Methods("PUT")
	r.Handle("/ulb/testcertificate/reapply", Adapt(http.HandlerFunc(route.Handler.ReApplyForTestCert))).Methods("PUT")
	r.Handle("/ulb/testcertificate/accept", Adapt(http.HandlerFunc(route.Handler.AcceptTestCert))).Methods("PUT")
	r.Handle("/ulb/testcertificate/reject", Adapt(http.HandlerFunc(route.Handler.RejectTestCert))).Methods("PUT")
	r.Handle("/ulb/testcert/status", Adapt(http.HandlerFunc(route.Handler.ULBTestCertStatus))).Methods("PUT")
	// report
	r.Handle("/ulb/report/master/v2", Adapt(http.HandlerFunc(route.Handler.ULBMasterReportV2))).Methods("POST")

	// ULBLogin
	// r.Handle("/user/ulb/auth", Adapt(http.HandlerFunc(route.Handler.ULBLogin))).Methods("POST")
	r.Handle("/ulb/mobile/uniqueness", Adapt(http.HandlerFunc(route.Handler.ULBMobileUniqueness))).Methods("POST")

}
