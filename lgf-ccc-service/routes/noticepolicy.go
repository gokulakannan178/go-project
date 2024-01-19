package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NoticePolicyRoutes : ""
func (route *Route) NoticePolicyRoutes(r *mux.Router) {
	r.Handle("/noticepolicy", Adapt(http.HandlerFunc(route.Handler.SaveNoticePolicy))).Methods("POST")
	r.Handle("/noticepolicy", Adapt(http.HandlerFunc(route.Handler.GetSingleNoticePolicy))).Methods("GET")
	r.Handle("/noticepolicy", Adapt(http.HandlerFunc(route.Handler.UpdateNoticePolicy))).Methods("PUT")
	r.Handle("/noticepolicy/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableNoticePolicy))).Methods("PUT")
	r.Handle("/noticepolicy/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableNoticePolicy))).Methods("PUT")
	r.Handle("/noticepolicy/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteNoticePolicy))).Methods("DELETE")
	r.Handle("/noticepolicy/filter", Adapt(http.HandlerFunc(route.Handler.FilterNoticePolicy))).Methods("POST")

}
