package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FarmerRoutes : ""
func (route *Route) FarmerRoutes(r *mux.Router) {
	r.Handle("/farmer", Adapt(http.HandlerFunc(route.Handler.SaveFarmer))).Methods("POST")
	r.Handle("/farmer/registration/generateotp", Adapt(http.HandlerFunc(route.Handler.GenerateotpFarmerRegistration))).Methods("POST")
	r.Handle("/farmer/registration/validateotp", Adapt(http.HandlerFunc(route.Handler.RegistrationValidateOTPFarmer))).Methods("POST")
	r.Handle("/farmer", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmer))).Methods("GET")
	r.Handle("/farmer", Adapt(http.HandlerFunc(route.Handler.UpdateFarmer))).Methods("PUT")
	r.Handle("/farmer/profileImgupload", Adapt(http.HandlerFunc(route.Handler.UpdateFarmerProfileImage))).Methods("PUT")
	r.Handle("/farmer/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFarmer))).Methods("PUT")
	r.Handle("/farmer/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFarmer))).Methods("PUT")
	r.Handle("/farmer/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFarmer))).Methods("DELETE")
	r.Handle("/farmer/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmer))).Methods("POST")
	r.Handle("/farmer/location", Adapt(http.HandlerFunc(route.Handler.FilterFarmerWithLocation))).Methods("POST")
	r.Handle("/farmer/basic/filter", Adapt(http.HandlerFunc(route.Handler.FilterFarmerBasic))).Methods("POST")
	r.Handle("/farmer/nearby", Adapt(http.HandlerFunc(route.Handler.FarmerNearBy))).Methods("POST")
	r.Handle("/farmer/registration/uniquecheck", Adapt(http.HandlerFunc(route.Handler.FarmerUniquenessCheckRegistration))).Methods("GET")
	r.Handle("/farmer/unicheckmobileorg", Adapt(http.HandlerFunc(route.Handler.GetSingleFarmerWithMobilenoAndOrg))).Methods("GET")
	r.Handle("/farmer/addprojectfarmer", Adapt(http.HandlerFunc(route.Handler.AddProjectFarmer))).Methods("POST")

}

//UserAuthRoutes : ""
func (route *Route) FarmerAuthRoutes(r *mux.Router) {
	r.Handle("/farmer/auth/generateotp", Adapt(http.HandlerFunc(route.Handler.LoginGenerateotpFarmer))).Methods("POST")
	r.Handle("/farmer/auth/otplogin/validateotp", Adapt(http.HandlerFunc(route.Handler.LoginValidateOTPFarmer))).Methods("POST")
}
