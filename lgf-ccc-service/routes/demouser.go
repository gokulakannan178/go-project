package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) DemoUserRoutes(r *mux.Router) {
	// DemoUser
	r.Handle("/demoUser", Adapt(http.HandlerFunc(route.Handler.SaveDemoUser))).Methods("POST")
	r.Handle("/demoUser", Adapt(http.HandlerFunc(route.Handler.GetSingleDemoUser))).Methods("GET")
	r.Handle("/demoUser", Adapt(http.HandlerFunc(route.Handler.UpdateDemoUser))).Methods("PUT")
	r.Handle("/demoUser/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableDemoUser))).Methods("PUT")
	r.Handle("/demoUser/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableDemoUser))).Methods("PUT")
	r.Handle("/demoUser/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteDemoUser))).Methods("DELETE")
	r.Handle("/demoUser/filter", Adapt(http.HandlerFunc(route.Handler.FilterDemoUser))).Methods("POST")

}
