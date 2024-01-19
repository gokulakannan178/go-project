package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DisseminationRoutes(r *mux.Router) {
	r.Handle("/dissemination", Adapt(http.HandlerFunc(route.Handler.SaveDissemination))).Methods("POST")
	r.Handle("/dissemination", Adapt(http.HandlerFunc(route.Handler.GetSingleDissemination))).Methods("GET")
	r.Handle("/dissemination", Adapt(http.HandlerFunc(route.Handler.UpdateDissemination))).Methods("PUT")
	r.Handle("/dissemination/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDissemination))).Methods("PUT")
	r.Handle("/dissemination/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDissemination))).Methods("PUT")
	r.Handle("/dissemination/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDissemination))).Methods("DELETE")
	r.Handle("/dissemination/filter", Adapt(http.HandlerFunc(route.Handler.FilterDissemination))).Methods("POST")
	r.Handle("/content/dissemination/sendnow", Adapt(http.HandlerFunc(route.Handler.SaveSendNow))).Methods("POST")
	r.Handle("/content/dissemination/sendlater", Adapt(http.HandlerFunc(route.Handler.SaveSendLater))).Methods("POST")
	r.Handle("/dissemination/pdf", Adapt(http.HandlerFunc(route.Handler.DisseminationPDF))).Methods("POST")
	r.Handle("/dissemination/report/filter", Adapt(http.HandlerFunc(route.Handler.DisseminationReport))).Methods("POST")

}
