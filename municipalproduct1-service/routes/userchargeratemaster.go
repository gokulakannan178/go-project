package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// UserChargeRateMasterRoutes : ""
func (route *Route) UserChargeRateMasterRoutes(r *mux.Router) {
	r.Handle("/userchargeratemaster", Adapt(http.HandlerFunc(route.Handler.SaveUserChargeRateMaster))).Methods("POST")
	r.Handle("/userchargeratemaster", Adapt(http.HandlerFunc(route.Handler.GetSingleUserChargeRateMaster))).Methods("GET")
	r.Handle("/userchargeratemaster", Adapt(http.HandlerFunc(route.Handler.UpdateUserChargeRateMaster))).Methods("PUT")
	r.Handle("/userchargeratemaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserChargeRateMaster))).Methods("PUT")
	r.Handle("/userchargeratemaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserChargeRateMaster))).Methods("PUT")
	r.Handle("/userchargeratemaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserChargeRateMaster))).Methods("DELETE")
	r.Handle("/userchargeratemaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserChargeRateMaster))).Methods("POST")
	r.Handle("/usercharge/property/demand", Adapt(http.HandlerFunc(route.Handler.GetUserChargeDemand))).Methods("GET")
	r.Handle("/userchargeratemaster/categoryid", Adapt(http.HandlerFunc(route.Handler.GetSingleUserChargeRateMasterWithCategoryId))).Methods("GET")

}
