package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AssetTypePropertysRoutes
func (route *Route) AssetTypePropertysRoutes(r *mux.Router) {
	r.Handle("/assettypepropertys", Adapt(http.HandlerFunc(route.Handler.SaveAssetTypePropertys))).Methods("POST")
	r.Handle("/assettypepropertys", Adapt(http.HandlerFunc(route.Handler.GetSingleAssetTypePropertys))).Methods("GET")
	r.Handle("/assettypepropertys", Adapt(http.HandlerFunc(route.Handler.UpdateAssetTypePropertys))).Methods("PUT")
	r.Handle("/assettypepropertys/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAssetTypePropertys))).Methods("PUT")
	r.Handle("/assettypepropertys/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAssetTypePropertys))).Methods("PUT")
	r.Handle("/assettypepropertys/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAssetTypePropertys))).Methods("DELETE")
	r.Handle("/assettypepropertys/filter", Adapt(http.HandlerFunc(route.Handler.FilterAssetTypePropertys))).Methods("POST")

}
