package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) PmAchievementRoutes(r *mux.Router) {
	// PmAchievement
	r.Handle("/pmachievement", Adapt(http.HandlerFunc(route.Handler.SavePmAchievement))).Methods("POST")
	r.Handle("/pmachievement", Adapt(http.HandlerFunc(route.Handler.GetSinglePmAchievement))).Methods("GET")
	r.Handle("/pmachievement", Adapt(http.HandlerFunc(route.Handler.UpdatePmAchievement))).Methods("PUT")
	r.Handle("/pmachievement/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePmAchievement))).Methods("PUT")
	r.Handle("/pmachievement/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePmAchievement))).Methods("PUT")
	r.Handle("/pmachievement/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePmAchievement))).Methods("DELETE")
	r.Handle("/pmachievement/filter", Adapt(http.HandlerFunc(route.Handler.FilterPmAchievement))).Methods("POST")
}
