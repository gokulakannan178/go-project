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
	r.Handle("/user/forgetpassword/generateotp", Adapt(http.HandlerFunc(route.Handler.ForgetPasswordGenerateOTP))).Methods("GET")
	r.Handle("/user/forgetpassword/validateotp", Adapt(http.HandlerFunc(route.Handler.ForgetPasswordValidateOTP))).Methods("GET")
	r.Handle("/user/passwordupdate", Adapt(http.HandlerFunc(route.Handler.PasswordUpdate))).Methods("PUT")
	//UserCollectionLimit updates
	r.Handle("/user/collectionlimitupdate", Adapt(http.HandlerFunc(route.Handler.UserCollectionLimit))).Methods("PUT")
	r.Handle("/user/idcards", Adapt(http.HandlerFunc(route.Handler.IDCaredPDF))).Methods("POST")
	// AccessPrivileage
	r.Handle("/user/accessprivilege", Adapt(http.HandlerFunc(route.Handler.UpdateAccessPrivilege))).Methods("PUT")
	//AppVersion
	r.Handle("/user/update/appversion", Adapt(http.HandlerFunc(route.Handler.UpdateAppVersionUser))).Methods("PUT")
	//mpin
	r.Handle("/user/mpin/register", Adapt(http.HandlerFunc(route.Handler.UserMpinRegistration))).Methods("PUT")
	r.Handle("/user/mpin/verify", Adapt(http.HandlerFunc(route.Handler.VerifyUserMpinRegistration))).Methods("PUT")
	r.Handle("/user/mpin/login", Adapt(http.HandlerFunc(route.Handler.UserMpinLogin))).Methods("POST")
	r.Handle("/user/removedtoken", Adapt(http.HandlerFunc(route.Handler.RemovedUserToken))).Methods("PUT")

}

//UserAuthRoutes : ""
func (route *Route) UserAuthRoutes(r *mux.Router) {
	r.Handle("/user/auth", Adapt(http.HandlerFunc(route.Handler.Login))).Methods("POST")
	r.Handle("/user/auth/generateotp", Adapt(http.HandlerFunc(route.Handler.LoginGenerateOTPV2))).Methods("GET")
	r.Handle("/user/auth/validateotp", Adapt(http.HandlerFunc(route.Handler.LoginValidateOTPV2))).Methods("GET")

}

//UserTypeRoutes : ""
func (route *Route) UserTypeRoutes(r *mux.Router) {
	r.Handle("/usertype", Adapt(http.HandlerFunc(route.Handler.SaveUserType))).Methods("POST")
	r.Handle("/usertype", Adapt(http.HandlerFunc(route.Handler.GetSingleUserType))).Methods("GET")
	r.Handle("/usertype", Adapt(http.HandlerFunc(route.Handler.UpdateUserType))).Methods("PUT")
	r.Handle("/usertype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserType))).Methods("PUT")
	r.Handle("/usertype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserType))).Methods("PUT")
	r.Handle("/usertype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserType))).Methods("DELETE")
	r.Handle("/usertype/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserType))).Methods("POST")
}

//UserLocationRoutes : ""
func (route *Route) UserLocationRoutes(r *mux.Router) {
	r.Handle("/userlocation/general", Adapt(http.HandlerFunc(route.Handler.SaveUserLocation))).Methods("POST")

}

//ConsumerRoutes : ""
func (route *Route) ConsumerRoutes(r *mux.Router) {
	r.Handle("/consumer/login/getotp", Adapt(http.HandlerFunc(route.Handler.SendOTPConsumerLogin))).Methods("GET")
	r.Handle("/consumer/login/validateotp", Adapt(http.HandlerFunc(route.Handler.ConsumerLoginValidateOTP))).Methods("POSt")
}
