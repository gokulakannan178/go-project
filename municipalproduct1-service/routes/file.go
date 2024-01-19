package routes

import (
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//FileRoutes : ""
func (route *Route) FileRoutes(r *mux.Router) {
	// docStart := route.ConfigReader.GetString(route.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOCD)
	docStart2 := route.ConfigReader.GetString(route.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOC)

	r.Handle("/common/docupload/{scenario}", Adapt(http.HandlerFunc(route.Handler.DocumentUpload))).Methods("POST")
	r.Handle("/common/docsupload/{scenario}", Adapt(http.HandlerFunc(route.Handler.DocumentsUpload))).Methods("POST")
	// r.PathPrefix(docStart2).Handler(http.StripPrefix(docStart2, http.FileServer(http.Dir(docStart))))
	stripDocURL := route.ConfigReader.GetString(route.Shared.GetCmdArg(constants.ENV) + "." + constants.STRIPDOCLOC)
	DocLOCURL := route.ConfigReader.GetString(route.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOCURL)
	fmt.Println(stripDocURL, "stripDocURL", DocLOCURL, "DocLOCURL")
	// r.PathPrefix(DocLOCURL).Handler(http.StripPrefix(stripDocURL, http.FileServer(http.Dir(docStart2))))
	r.Handle(DocLOCURL, http.StripPrefix(stripDocURL, GetFile(http.FileServer(http.Dir(docStart2))))).Methods("GET")

	// r.PathPrefix("/documents/{folder1}/{file}").Handler(http.StripPrefix("/documents/", http.FileServer(http.Dir("documents/"))))
	//
	// r.PathPrefix(docStart + "{folder1}/{file}").Handler(http.StripPrefix(docStart, http.FileServer(http.Dir(docStart2))))
	// r.PathPrefix("/documents/}").Handler(http.StripPrefix(docStart, http.FileServer(http.Dir(docStart))))
	r.Handle("/common/docsupload/base64/{scenario}", Adapt(http.HandlerFunc(route.Handler.DocumentUploadBase64))).Methods("POST")
}

//Log : ""
func GetFile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		log.Println(r.URL)
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
		duration := time.Since(t)
		log.Println("API ==>", r.RequestURI, " Time taken ===> ", duration.Minutes(), "m")
		fmt.Println()
		fmt.Println()
		fmt.Println()
	})
}
