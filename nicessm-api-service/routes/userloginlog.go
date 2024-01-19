package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserLoginLogRoutes(r *mux.Router) {
	r.Handle("/userloginlog", Adapt(http.HandlerFunc(route.Handler.SaveUserLoginLog))).Methods("POST")
	r.Handle("/userloginlog", Adapt(http.HandlerFunc(route.Handler.GetSingleUserLoginLog))).Methods("GET")
	r.Handle("/userloginlog", Adapt(http.HandlerFunc(route.Handler.UpdateUserLoginLog))).Methods("PUT")
	r.Handle("/userloginlog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserLoginLog))).Methods("PUT")
	r.Handle("/userloginlog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserLoginLog))).Methods("PUT")
	r.Handle("/userloginlog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserLoginLog))).Methods("DELETE")
	r.Handle("/userloginlog/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserLoginLog))).Methods("POST")
	r.Handle("/user/login", Adapt(http.HandlerFunc(route.Handler.UserLogin))).Methods("PUT")
	r.Handle("/user/logout", Adapt(http.HandlerFunc(route.Handler.UpdateUserLogout))).Methods("PUT")

}
