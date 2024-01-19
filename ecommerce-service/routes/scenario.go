package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ScenarioRoutes(r *mux.Router) {
	// Scenario
	r.Handle("/scenario", Adapt(http.HandlerFunc(route.Handler.SaveScenario))).Methods("POST")
	r.Handle("/scenario", Adapt(http.HandlerFunc(route.Handler.GetSingleScenario))).Methods("GET")
	r.Handle("/scenario", Adapt(http.HandlerFunc(route.Handler.UpdateScenario))).Methods("PUT")
	r.Handle("/scenario/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableScenario))).Methods("PUT")
	r.Handle("/scenario/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableScenario))).Methods("PUT")
	r.Handle("/scenario/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteScenario))).Methods("DELETE")
	r.Handle("/scenario/filter", Adapt(http.HandlerFunc(route.Handler.FilterScenario))).Methods("POST")
}
