package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ShoprentPayeeNameChangeRoutes(r *mux.Router) {
	r.Handle("/shoprentpayeenamechange", Adapt(http.HandlerFunc(route.Handler.SaveShoprentPayeeNameChange))).Methods("POST")
	r.Handle("/shoprentpayeenamechange", Adapt(http.HandlerFunc(route.Handler.GetSingleShoprentPayeeNameChange))).Methods("GET")
	r.Handle("/shoprentpayeenamechange", Adapt(http.HandlerFunc(route.Handler.UpdateShoprentPayeeNameChange))).Methods("PUT")
	r.Handle("/shoprentpayeenamechange/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableShoprentPayeeNameChange))).Methods("PUT")
	r.Handle("/shoprentpayeenamechange/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableShoprentPayeeNameChange))).Methods("PUT")
	r.Handle("/shoprentpayeenamechange/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteShoprentPayeeNameChange))).Methods("DELETE")
	r.Handle("/shoprentpayeenamechange/filter", Adapt(http.HandlerFunc(route.Handler.FilterShoprentPayeeNameChange))).Methods("POST")
	//Approve
	r.Handle("/shoprentpayeenamechange/approve", Adapt(http.HandlerFunc(route.Handler.ApproveShoprentPayeeNameChange))).Methods("PUT")
	r.Handle("/shoprentpayeenamechange/rejected", Adapt(http.HandlerFunc(route.Handler.NotApproveShoprentPayeeNameChange))).Methods("PUT")

}
