package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) TradeLicenseCategoryTypeRoutes(r *mux.Router) {
	// TradeLicenseCategoryType
	r.Handle("/tradelicensecategorytype", Adapt(http.HandlerFunc(route.Handler.SaveTradeLicenseCategoryType))).Methods("POST")
	r.Handle("/tradelicensecategorytype", Adapt(http.HandlerFunc(route.Handler.GetSingleTradeLicenseCategoryType))).Methods("GET")
	r.Handle("/tradelicensecategorytype", Adapt(http.HandlerFunc(route.Handler.UpdateTradeLicenseCategoryType))).Methods("PUT")
	r.Handle("/tradelicensecategorytype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableTradeLicenseCategoryType))).Methods("PUT")
	r.Handle("/tradelicensecategorytype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableTradeLicenseCategoryType))).Methods("PUT")
	r.Handle("/tradelicensecategorytype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteTradeLicenseCategoryType))).Methods("DELETE")
	r.Handle("/tradelicensecategorytype/filter", Adapt(http.HandlerFunc(route.Handler.FilterTradeLicenseCategoryType))).Methods("POST")
}
