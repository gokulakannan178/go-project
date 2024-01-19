package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DocumentScenarioRoutes(r *mux.Router) {
	// DocumentScenario
	r.Handle("/documentScenario", Adapt(http.HandlerFunc(route.Handler.SaveDocumentScenario))).Methods("POST")
	r.Handle("/documentScenario", Adapt(http.HandlerFunc(route.Handler.GetSingleDocumentScenario))).Methods("GET")
	r.Handle("/documentScenario", Adapt(http.HandlerFunc(route.Handler.UpdateDocumentScenario))).Methods("PUT")
	r.Handle("/documentScenario/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDocumentScenario))).Methods("PUT")
	r.Handle("/documentScenario/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDocumentScenario))).Methods("PUT")
	r.Handle("/documentScenario/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDocumentScenario))).Methods("DELETE")
	r.Handle("/documentScenario/filter", Adapt(http.HandlerFunc(route.Handler.FilterDocumentScenario))).Methods("POST")

}
