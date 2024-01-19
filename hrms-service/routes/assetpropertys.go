package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AssetPropertysRoutes
func (route *Route) AssetPropertysRoutes(r *mux.Router) {
	r.Handle("/assetpropertys", Adapt(http.HandlerFunc(route.Handler.SaveAssetPropertys))).Methods("POST")
	r.Handle("/assetpropertys", Adapt(http.HandlerFunc(route.Handler.GetSingleAssetPropertys))).Methods("GET")
	r.Handle("/assetpropertys", Adapt(http.HandlerFunc(route.Handler.UpdateAssetPropertys))).Methods("PUT")
	r.Handle("/assetpropertys/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAssetPropertys))).Methods("PUT")
	r.Handle("/assetpropertys/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAssetPropertys))).Methods("PUT")
	r.Handle("/assetpropertys/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAssetPropertys))).Methods("DELETE")
	r.Handle("/assetpropertys/filter", Adapt(http.HandlerFunc(route.Handler.FilterAssetPropertys))).Methods("POST")

}
