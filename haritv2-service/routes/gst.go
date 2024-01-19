package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) GSTRoutes(r *mux.Router) {
	// GST
	r.Handle("/gst", Adapt(http.HandlerFunc(route.Handler.SaveGST))).Methods("POST")
	r.Handle("/gst", Adapt(http.HandlerFunc(route.Handler.GetSingleGST))).Methods("GET")
	r.Handle("/gst", Adapt(http.HandlerFunc(route.Handler.UpdateGST))).Methods("PUT")
	r.Handle("/gst/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableGST))).Methods("PUT")
	r.Handle("/gst/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableGST))).Methods("PUT")
	r.Handle("/gst/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteGST))).Methods("DELETE")
	r.Handle("/gst/filter", Adapt(http.HandlerFunc(route.Handler.FilterGST))).Methods("POST")
	r.Handle("/gst/getdefaultgst", Adapt(http.HandlerFunc(route.Handler.GetDefaultGST))).Methods("GET")
}
