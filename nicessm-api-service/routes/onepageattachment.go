package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) OnePageAttachmentRoutes(r *mux.Router) {
	r.Handle("/onepageattachment/image", Adapt(http.HandlerFunc(route.Handler.GetSingleOnePageAttachment))).Methods("GET")
}
