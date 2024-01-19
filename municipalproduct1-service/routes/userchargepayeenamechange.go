package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) UserchargePayeeNameChangeRoutes(r *mux.Router) {
	r.Handle("/userchargepayeenamechange", Adapt(http.HandlerFunc(route.Handler.SaveUserchargePayeeNameChange))).Methods("POST")
	r.Handle("/userchargepayeenamechange", Adapt(http.HandlerFunc(route.Handler.GetSingleUserchargePayeeNameChange))).Methods("GET")
	r.Handle("/userchargepayeenamechange", Adapt(http.HandlerFunc(route.Handler.UpdateUserchargePayeeNameChange))).Methods("PUT")
	r.Handle("/userchargepayeenamechange/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUserchargePayeeNameChange))).Methods("PUT")
	r.Handle("/userchargepayeenamechange/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUserchargePayeeNameChange))).Methods("PUT")
	r.Handle("/userchargepayeenamechange/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUserchargePayeeNameChange))).Methods("DELETE")
	r.Handle("/userchargepayeenamechange/filter", Adapt(http.HandlerFunc(route.Handler.FilterUserchargePayeeNameChange))).Methods("POST")
	//Approve
	r.Handle("/userchargepayeenamechange/approve", Adapt(http.HandlerFunc(route.Handler.ApproveUserchargePayeeNameChange))).Methods("PUT")
	r.Handle("/userchargepayeenamechange/rejected", Adapt(http.HandlerFunc(route.Handler.NotApproveUserchargePayeeNameChange))).Methods("PUT")

}
