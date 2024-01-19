package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// UserChargeUpdateLogRoutes : ""
func (route *Route) UserChargeUpdateLogRoutes(r *mux.Router) {
	r.Handle("/userchargeupdatelog", Adapt(http.HandlerFunc(route.Handler.SaveUserChargeUpdateLog))).Methods("POST")
	r.Handle("/userchargeupdatelog", Adapt(http.HandlerFunc(route.Handler.GetSingleUserChargeUpdateLog))).Methods("GET")
	r.Handle("/userchargeupdatelog", Adapt(http.HandlerFunc(route.Handler.UpdateUserChargeUpdateLog))).Methods("PUT")
	r.Handle("/userchargeupdatelog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserChargeUpdateLog))).Methods("PUT")
	r.Handle("/userchargeupdatelog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserChargeUpdateLog))).Methods("PUT")
	r.Handle("/userchargeupdatelog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserChargeUpdateLog))).Methods("DELETE")
	r.Handle("/userchargeupdatelog/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserChargeUpdateLog))).Methods("POST")

}
