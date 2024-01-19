package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PkgTypeRoutes(r *mux.Router) {
	// PkgType
	r.Handle("/pkgtype", Adapt(http.HandlerFunc(route.Handler.SavePkgType))).Methods("POST")
	r.Handle("/pkgtype", Adapt(http.HandlerFunc(route.Handler.GetSinglePkgType))).Methods("GET")
	r.Handle("/pkgtype", Adapt(http.HandlerFunc(route.Handler.UpdatePkgType))).Methods("PUT")
	r.Handle("/pkgtype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePkgType))).Methods("PUT")
	r.Handle("/pkgtype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePkgType))).Methods("PUT")
	r.Handle("/pkgtype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePkgType))).Methods("DELETE")
	r.Handle("/pkgtype/filter", Adapt(http.HandlerFunc(route.Handler.FilterPkgType))).Methods("POST")
	r.Handle("/pkgtype/getdefaultpkgtype", Adapt(http.HandlerFunc(route.Handler.GetDefaultPkgType))).Methods("GET")

}
