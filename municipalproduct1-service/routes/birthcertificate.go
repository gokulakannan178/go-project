package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) BirthCertificateRoutes(r *mux.Router) {
	// BirthCertificate
	r.Handle("/birthcertificate", Adapt(http.HandlerFunc(route.Handler.SaveBirthCertificate))).Methods("POST")
	r.Handle("/birthcertificate", Adapt(http.HandlerFunc(route.Handler.GetSingleBirthCertificate))).Methods("GET")
	r.Handle("/birthcertificate", Adapt(http.HandlerFunc(route.Handler.UpdateBirthCertificate))).Methods("PUT")
	r.Handle("/birthcertificate/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBirthCertificate))).Methods("PUT")
	r.Handle("/birthcertificate/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBirthCertificate))).Methods("PUT")
	r.Handle("/birthcertificate/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBirthCertificate))).Methods("DELETE")
	r.Handle("/birthcertificate/filter", Adapt(http.HandlerFunc(route.Handler.FilterBirthCertificate))).Methods("POST")
	r.Handle("/birthcertificate/approve", Adapt(http.HandlerFunc(route.Handler.ApproveBirthCertificate))).Methods("PUT")
	r.Handle("/birthcertificate/reject", Adapt(http.HandlerFunc(route.Handler.RejectBirthCertificate))).Methods("PUT")
}
