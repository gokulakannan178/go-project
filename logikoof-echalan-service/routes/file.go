package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//FileRoutes : ""
func (route *Route) FileRoutes(r *mux.Router) {
	r.Handle("/common/docupload/{scenario}", Adapt(http.HandlerFunc(route.Handler.DocumentUpload))).Methods("POST")
	r.Handle("/common/docsupload/{scenario}", Adapt(http.HandlerFunc(route.Handler.DocumentsUpload))).Methods("POST")
	r.PathPrefix("/documents/{folder1}/{file}").Handler(http.StripPrefix("/documents/", http.FileServer(http.Dir("documents/"))))
	r.Handle("/common/docsupload/base64/{scenario}", Adapt(http.HandlerFunc(route.Handler.DocumentUploadBase64))).Methods("POST")
}
