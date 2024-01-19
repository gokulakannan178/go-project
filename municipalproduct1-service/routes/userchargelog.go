package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// UserChargeLogRoutes : ""
func (route *Route) UserChargeLogRoutes(r *mux.Router) {
	r.Handle("/userchargelog", Adapt(http.HandlerFunc(route.Handler.SaveUserChargeLog))).Methods("POST")
	r.Handle("/userchargelog", Adapt(http.HandlerFunc(route.Handler.GetSingleUserChargeLog))).Methods("GET")
	r.Handle("/userchargelog", Adapt(http.HandlerFunc(route.Handler.UpdateUserChargeLog))).Methods("PUT")
	r.Handle("/userchargelog/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserChargeLog))).Methods("PUT")
	r.Handle("/userchargelog/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserChargeLog))).Methods("PUT")
	r.Handle("/userchargelog/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserChargeLog))).Methods("DELETE")
	r.Handle("/userchargelog/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserChargeLog))).Methods("POST")

}
