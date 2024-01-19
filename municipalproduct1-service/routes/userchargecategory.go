package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// UserChargeCategoryRoutes : ""
func (route *Route) UserChargeCategoryRoutes(r *mux.Router) {
	r.Handle("/userchargecategory", Adapt(http.HandlerFunc(route.Handler.SaveUserChargeCategory))).Methods("POST")
	r.Handle("/userchargecategory", Adapt(http.HandlerFunc(route.Handler.GetSingleUserChargeCategory))).Methods("GET")
	r.Handle("/userchargecategory", Adapt(http.HandlerFunc(route.Handler.UpdateUserChargeCategory))).Methods("PUT")
	r.Handle("/userchargecategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserChargeCategory))).Methods("PUT")
	r.Handle("/userchargecategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserChargeCategory))).Methods("PUT")
	r.Handle("/userchargecategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserChargeCategory))).Methods("DELETE")
	r.Handle("/userchargecategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserChargeCategory))).Methods("POST")
}
