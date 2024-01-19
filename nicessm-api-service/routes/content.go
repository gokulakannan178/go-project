package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ContentRoutes(r *mux.Router) {
	r.Handle("/content", Adapt(http.HandlerFunc(route.Handler.SaveContent))).Methods("POST")
	r.Handle("/content", Adapt(http.HandlerFunc(route.Handler.GetSingleContent))).Methods("GET")
	r.Handle("/content", Adapt(http.HandlerFunc(route.Handler.UpdateContent))).Methods("PUT")
	r.Handle("/content/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableContent))).Methods("PUT")
	r.Handle("/content/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableContent))).Methods("PUT")
	r.Handle("/content/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteContent))).Methods("DELETE")
	r.Handle("/content/filter", Adapt(http.HandlerFunc(route.Handler.FilterContent))).Methods("POST")
	r.Handle("/content/manager", Adapt(http.HandlerFunc(route.Handler.ContentManager))).Methods("POST")
	r.Handle("/content/status/approved", Adapt(http.HandlerFunc(route.Handler.ApprovedContent))).Methods("PUT")
	r.Handle("/content/status/reject", Adapt(http.HandlerFunc(route.Handler.RejectedContent))).Methods("PUT")
	r.Handle("/content/editapproved", Adapt(http.HandlerFunc(route.Handler.EditApprovedContent))).Methods("PUT")
	r.Handle("/content/editrejected", Adapt(http.HandlerFunc(route.Handler.EditRejectedContent))).Methods("PUT")
	r.Handle("/content/inccount", Adapt(http.HandlerFunc(route.Handler.ContentViewCountIncrement))).Methods("PUT")

}
