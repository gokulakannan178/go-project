package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CompendiumRoutes(r *mux.Router) {
	r.Handle("/compendium", Adapt(http.HandlerFunc(route.Handler.SaveCompendium))).Methods("POST")
	r.Handle("/compendium", Adapt(http.HandlerFunc(route.Handler.GetSingleCompendium))).Methods("GET")
	r.Handle("/compendium", Adapt(http.HandlerFunc(route.Handler.UpdateCompendium))).Methods("PUT")
	r.Handle("/compendium/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCompendium))).Methods("PUT")
	r.Handle("/compendium/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCompendium))).Methods("PUT")
	r.Handle("/compendium/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCompendium))).Methods("DELETE")
	r.Handle("/compendium/filter", Adapt(http.HandlerFunc(route.Handler.FilterCompendium))).Methods("POST")
	r.Handle("/compendium/fileupload", Adapt(http.HandlerFunc(route.Handler.CompendiumUploadWord))).Methods("POST")
	r.Handle("/compendium/fileupload/v2", Adapt(http.HandlerFunc(route.Handler.CompendiumUploadWordV2))).Methods("POST")
	r.Handle("/compendium/fileupload/v3", Adapt(http.HandlerFunc(route.Handler.CompendiumUploadWordV3))).Methods("POST")
}
