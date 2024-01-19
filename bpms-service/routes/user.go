package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

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

//UserRoutes : ""
func (route *Route) UserRoutes(r *mux.Router) {
	r.Handle("/user", Adapt(http.HandlerFunc(route.Handler.SaveUser))).Methods("POST")
	r.Handle("/user", Adapt(http.HandlerFunc(route.Handler.GetSingleUser))).Methods("GET")
	r.Handle("/user", Adapt(http.HandlerFunc(route.Handler.UpdateUser))).Methods("PUT")
	r.Handle("/user/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUser))).Methods("PUT")
	r.Handle("/user/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUser))).Methods("PUT")
	r.Handle("/user/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUser))).Methods("DELETE")
	r.Handle("/user/filter", Adapt(http.HandlerFunc(route.Handler.FilterUser))).Methods("POST")
}

//UserAuthRoutes : ""
func (route *Route) UserAuthRoutes(r *mux.Router) {
	r.Handle("/user/auth", Adapt(http.HandlerFunc(route.Handler.Login))).Methods("POST")
}
