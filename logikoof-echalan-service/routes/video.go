package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//LiveVideoRoutes : ""
func (route *Route) LiveVideoRoutes(r *mux.Router) {
	r.Handle("/livevideo", Adapt(http.HandlerFunc(route.Handler.SaveLiveVideo))).Methods("POST")
	r.Handle("/livevideo", Adapt(http.HandlerFunc(route.Handler.GetSingleLiveVideo))).Methods("GET")
	r.Handle("/livevideo", Adapt(http.HandlerFunc(route.Handler.UpdateLiveVideo))).Methods("PUT")
	r.Handle("/livevideo/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableLiveVideo))).Methods("PUT")
	r.Handle("/livevideo/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableLiveVideo))).Methods("PUT")
	r.Handle("/livevideo/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteLiveVideo))).Methods("DELETE")
	r.Handle("/livevideo/filter", Adapt(http.HandlerFunc(route.Handler.FilterLiveVideo))).Methods("POST")
}

//OffenceVideoRoutes : ""
func (route *Route) OffenceVideoRoutes(r *mux.Router) {
	r.Handle("/offencevideo", Adapt(http.HandlerFunc(route.Handler.SaveOffenceVideo))).Methods("POST")
	r.Handle("/offencevideo", Adapt(http.HandlerFunc(route.Handler.GetSingleOffenceVideo))).Methods("GET")
	r.Handle("/offencevideo", Adapt(http.HandlerFunc(route.Handler.UpdateOffenceVideo))).Methods("PUT")
	r.Handle("/offencevideo/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOffenceVideo))).Methods("PUT")
	r.Handle("/offencevideo/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOffenceVideo))).Methods("PUT")
	r.Handle("/offencevideo/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOffenceVideo))).Methods("DELETE")
	r.Handle("/offencevideo/filter", Adapt(http.HandlerFunc(route.Handler.FilterOffenceVideo))).Methods("POST")
}
