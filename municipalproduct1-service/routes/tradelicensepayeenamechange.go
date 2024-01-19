package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TradelicensePayeeNameChangeRoutes(r *mux.Router) {
	r.Handle("/tradelicensepayeenamechange", Adapt(http.HandlerFunc(route.Handler.SaveTradelicensePayeeNameChange))).Methods("POST")
	r.Handle("/tradelicensepayeenamechange", Adapt(http.HandlerFunc(route.Handler.GetSingleTradelicensePayeeNameChange))).Methods("GET")
	r.Handle("/tradelicensepayeenamechange", Adapt(http.HandlerFunc(route.Handler.UpdateTradelicensePayeeNameChange))).Methods("PUT")
	r.Handle("/tradelicensepayeenamechange/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTradelicensePayeeNameChange))).Methods("PUT")
	r.Handle("/tradelicensepayeenamechange/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTradelicensePayeeNameChange))).Methods("PUT")
	r.Handle("/tradelicensepayeenamechange/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTradelicensePayeeNameChange))).Methods("DELETE")
	r.Handle("/tradelicensepayeenamechange/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradelicensePayeeNameChange))).Methods("POST")
	//Approvt
	r.Handle("/tradelicensepayeenamechange/approve", Adapt(http.HandlerFunc(route.Handler.ApproveTradelicensePayeeNameChange))).Methods("PUT")
	r.Handle("/tradelicensepayeenamechange/rejected", Adapt(http.HandlerFunc(route.Handler.NotApproveTradelicensePayeeNameChange))).Methods("PUT")

}
