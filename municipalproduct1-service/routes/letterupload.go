package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// LetterUploadRoutes : ""
func (route *Route) LetterUpload(r *mux.Router) {

	// LetterUpload
	r.Handle("/letterupload", Adapt(http.HandlerFunc(route.Handler.SaveLetterUpload))).Methods("POST")
	r.Handle("/letterupload", Adapt(http.HandlerFunc(route.Handler.GetSingleLetterUpload))).Methods("GET")
	r.Handle("/letterupload", Adapt(http.HandlerFunc(route.Handler.UpdateLetterUpload))).Methods("PUT")
	r.Handle("/letterupload/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLetterUpload))).Methods("PUT")
	r.Handle("/letterupload/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLetterUpload))).Methods("PUT")
	r.Handle("/letterupload/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLetterUpload))).Methods("DELETE")
	r.Handle("/letterupload/filter", Adapt(http.HandlerFunc(route.Handler.FilterLetterUpload))).Methods("POST")
}
