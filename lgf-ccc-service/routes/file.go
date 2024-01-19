package routes

import (
	"fmt"
	"lgf-ccc-service/constants"
	"net/http"

	"github.com/gorilla/mux"
)

//FileRoutes : ""
func (route *Route) FileRoutes(r *mux.Router) {
	//	docStart := route.ConfigReader.GetString(route.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOCD)
	docStart2 := route.ConfigReader.GetString(route.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOC)

	r.Handle("/common/docupload/{scenario}", Adapt(http.HandlerFunc(route.Handler.DocumentUpload))).Methods("POST")
	r.Handle("/common/docsupload/{scenario}", Adapt(http.HandlerFunc(route.Handler.DocumentsUpload))).Methods("POST")
	// r.PathPrefix(docStart2).Handler(http.StripPrefix(docStart2, http.FileServer(http.Dir(docStart))))
	// r.PathPrefix("/documents/{folder1}/{file}").Handler(http.StripPrefix("/documents/", http.FileServer(http.Dir("documents/"))))
	// //
	stripDocURL := route.ConfigReader.GetString(route.Shared.GetCmdArg(constants.ENV) + "." + constants.STRIPDOCLOC)
	DocLOCURL := route.ConfigReader.GetString(route.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOCURL)
	fmt.Println(stripDocURL, "stripDocURL", DocLOCURL, "DocLOCURL")
	r.PathPrefix(DocLOCURL).Handler(http.StripPrefix(stripDocURL, http.FileServer(http.Dir(docStart2))))
	// r.PathPrefix(docStart + "{folder1}/{file}").Handler(http.StripPrefix(docStart, http.FileServer(http.Dir(docStart2))))
	// r.PathPrefix("/documents/}").Handler(http.StripPrefix(docStart, http.FileServer(http.Dir(docStart))))
	r.Handle("/common/docsupload/base64/{scenario}", Adapt(http.HandlerFunc(route.Handler.DocumentUploadBase64))).Methods("POST")
}
