package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) CommonLanguageTranslationsRoutes(r *mux.Router) {
	r.Handle("/commonlanguagetranslations", Adapt(http.HandlerFunc(route.Handler.SaveCommonLanguageTranslations))).Methods("POST")
	r.Handle("/commonlanguagetranslations", Adapt(http.HandlerFunc(route.Handler.GetSingleCommonLanguageTranslations))).Methods("GET")
	r.Handle("/commonlanguagetranslations", Adapt(http.HandlerFunc(route.Handler.UpdateCommonLanguageTranslations))).Methods("PUT")
	r.Handle("/commonlanguagetranslations/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableCommonLanguageTranslations))).Methods("PUT")
	r.Handle("/commonlanguagetranslations/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableCommonLanguageTranslations))).Methods("PUT")
	r.Handle("/commonlanguagetranslations/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteCommonLanguageTranslations))).Methods("DELETE")
	r.Handle("/commonlanguagetranslations/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommonLanguageTranslations))).Methods("POST")
	r.Handle("/commonlanguagetranslations/filter", Adapt(http.HandlerFunc(route.Handler.FilterCommonLanguageTranslations))).Methods("POST")
	r.Handle("/commonlanguagetranslations/language", Adapt(http.HandlerFunc(route.Handler.GetSingleCommonLanguageTranslationsWithType))).Methods("GET")
}
