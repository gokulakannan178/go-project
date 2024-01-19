package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DumpSiteRoutes
func (route *Route) DumpSiteRoutes(r *mux.Router) {
	r.Handle("/dumpSite", Adapt(http.HandlerFunc(route.Handler.SaveDumpSite))).Methods("POST")
	r.Handle("/dumpSite", Adapt(http.HandlerFunc(route.Handler.GetSingleDumpSite))).Methods("GET")
	r.Handle("/dumpSite", Adapt(http.HandlerFunc(route.Handler.UpdateDumpSite))).Methods("PUT")
	r.Handle("/dumpSite/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDumpSite))).Methods("PUT")
	r.Handle("/dumpSite/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDumpSite))).Methods("PUT")
	r.Handle("/dumpSite/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDumpSite))).Methods("DELETE")
	r.Handle("/dumpSite/filter", Adapt(http.HandlerFunc(route.Handler.FilterDumpSite))).Methods("POST")
	//r.Handle("/DumpSite/assign", Adapt(http.HandlerFunc(route.Handler.DumpSiteAssign))).Methods("POST")
	//r.Handle("/DumpSite/revoke", Adapt(http.HandlerFunc(route.Handler.RevokeDumpSite))).Methods("PUT")

}

// DumpSiteRoutes
func (route *Route) DumpHistoryRoutes(r *mux.Router) {
	r.Handle("/dumphistory", Adapt(http.HandlerFunc(route.Handler.SaveDumpHistory))).Methods("POST")
	r.Handle("/dumphistory", Adapt(http.HandlerFunc(route.Handler.GetSingleDumpHistory))).Methods("GET")
	r.Handle("/dumphistory", Adapt(http.HandlerFunc(route.Handler.UpdateDumpHistory))).Methods("PUT")
	r.Handle("/dumphistory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDumpHistory))).Methods("PUT")
	r.Handle("/dumphistory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDumpHistory))).Methods("PUT")
	r.Handle("/dumphistory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDumpHistory))).Methods("DELETE")
	r.Handle("/dumphistory/filter", Adapt(http.HandlerFunc(route.Handler.FilterDumpHistory))).Methods("POST")
	r.Handle("/dumphistory/getquantity", Adapt(http.HandlerFunc(route.Handler.GetQuantityByManagerId))).Methods("POST")
	r.Handle("/dumphistory/datewise/quantity", Adapt(http.HandlerFunc(route.Handler.DateWiseDumpHistory))).Methods("POST")
	//r.Handle("/DumpSite/assign", Adapt(http.HandlerFunc(route.Handler.DumpSiteAssign))).Methods("POST")
	//r.Handle("/DumpSite/revoke", Adapt(http.HandlerFunc(route.Handler.RevokeDumpSite))).Methods("PUT")

}
