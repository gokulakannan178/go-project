package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// LetterGenerate : ""
func (route *Route) LetterGenerate(r *mux.Router) {

	// LetterGenerate
	r.Handle("/lettergenerate", Adapt(http.HandlerFunc(route.Handler.SaveLetterGenerate))).Methods("POST")
	r.Handle("/lettergenerate", Adapt(http.HandlerFunc(route.Handler.GetSingleLetterGenerate))).Methods("GET")
	r.Handle("/lettergenerate", Adapt(http.HandlerFunc(route.Handler.UpdateLetterGenerate))).Methods("PUT")
	r.Handle("/lettergenerate/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLetterGenerate))).Methods("PUT")
	r.Handle("/lettergenerate/status/approved", Adapt(http.HandlerFunc(route.Handler.ApprovedLetterGenerate))).Methods("PUT")
	r.Handle("/lettergenerate/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLetterGenerate))).Methods("PUT")
	r.Handle("/lettergenerate/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLetterGenerate))).Methods("DELETE")
	r.Handle("/lettergenerate/status/blocked", Adapt(http.HandlerFunc(route.Handler.BlockedLetterGenerate))).Methods("PUT")
	r.Handle("/lettergenerate/status/submitted", Adapt(http.HandlerFunc(route.Handler.SubmittedLetterGenerate))).Methods("PUT")
	r.Handle("/lettergenerate/upload", Adapt(http.HandlerFunc(route.Handler.UploadLetterGenerate))).Methods("PUT")
	r.Handle("/lettergenerate/filter", Adapt(http.HandlerFunc(route.Handler.FilterLetterGenerate))).Methods("POST")
	r.Handle("/lettergenerate/execute/pdf", Adapt(http.HandlerFunc(route.Handler.LetterGenerateExecute))).Methods("GET")
}
