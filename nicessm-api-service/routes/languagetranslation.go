package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) LanguageTranslationRoutes(r *mux.Router) {
	r.Handle("/langauagetranslation", Adapt(http.HandlerFunc(route.Handler.SaveLanguageTranslation))).Methods("POST")
	r.Handle("/langauagetranslation", Adapt(http.HandlerFunc(route.Handler.GetSingleLanguageTranslation))).Methods("GET")
	r.Handle("/langauagetranslation", Adapt(http.HandlerFunc(route.Handler.UpdateLanguageTranslation))).Methods("PUT")
	r.Handle("/langauagetranslation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLanguageTranslation))).Methods("PUT")
	r.Handle("/langauagetranslation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLanguageTranslation))).Methods("PUT")
	r.Handle("/langauagetranslation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLanguageTranslation))).Methods("DELETE")
	r.Handle("/langauagetranslation/filter", Adapt(http.HandlerFunc(route.Handler.FilterLanguageTranslation))).Methods("POST")
	r.Handle("/langauagetranslation/filter", Adapt(http.HandlerFunc(route.Handler.FilterLanguageTranslation))).Methods("POST")
	r.Handle("/langauagetranslation/loginpage", Adapt(http.HandlerFunc(route.Handler.GetSingleLanguageTranslationWithType))).Methods("GET")
}
