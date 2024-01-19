package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) ContentManagerRoutes(r *mux.Router) {
	r.Handle("/contentmanager/count", Adapt(http.HandlerFunc(route.Handler.ContentManagerCount))).Methods("POST")

}
