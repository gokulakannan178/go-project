package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//OrganisationRoutes : ""
func (route *Route) OrganisationRoutes(r *mux.Router) {
	r.Handle("/organisation", Adapt(http.HandlerFunc(route.Handler.SaveOrganisation))).Methods("POST")
	r.Handle("/organisation", Adapt(http.HandlerFunc(route.Handler.GetSingleOrganisation))).Methods("GET")
	r.Handle("/organisation", Adapt(http.HandlerFunc(route.Handler.UpdateOrganisation))).Methods("PUT")
	r.Handle("/organisation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOrganisation))).Methods("PUT")
	r.Handle("/organisation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOrganisation))).Methods("PUT")
	r.Handle("/organisation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOrganisation))).Methods("DELETE")
	r.Handle("/organisation/filter", Adapt(http.HandlerFunc(route.Handler.FilterOrganisation))).Methods("POST")
}

//UserRoutes : ""
func (route *Route) UserRoutes(r *mux.Router) {
	r.Handle("/user", Adapt(http.HandlerFunc(route.Handler.SaveUser))).Methods("POST")
	r.Handle("/user", Adapt(http.HandlerFunc(route.Handler.GetSingleUser))).Methods("GET")
	r.Handle("/user", Adapt(http.HandlerFunc(route.Handler.UpdateUser))).Methods("PUT")
	r.Handle("/user/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUser))).Methods("PUT")
	r.Handle("/user/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUser))).Methods("PUT")
	r.Handle("/user/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUser))).Methods("DELETE")
	r.Handle("/user/filter", Adapt(http.HandlerFunc(route.Handler.FilterUser))).Methods("POST")
	r.Handle("/user/resetpassword", Adapt(http.HandlerFunc(route.Handler.ResetUserPassword))).Methods("PUT")
	r.Handle("/user/changepassword", Adapt(http.HandlerFunc(route.Handler.ChangePassword))).Methods("PUT")

}

//UserAuthRoutes : ""
func (route *Route) UserAuthRoutes(r *mux.Router) {
	r.Handle("/user/auth", Adapt(http.HandlerFunc(route.Handler.Login))).Methods("POST")
	r.Handle("/consumer/auth/generateotp", Adapt(http.HandlerFunc(route.Handler.ConsumerLoginSendOTP))).Methods("GET")
	r.Handle("/consumer/auth/validateotp", Adapt(http.HandlerFunc(route.Handler.ConsumerLoginValidateOTP))).Methods("POST")
}
