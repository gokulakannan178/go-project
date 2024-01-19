package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AssetMasterRoutes
func (route *Route) AssetMasterRoutes(r *mux.Router) {
	r.Handle("/assetmaster", Adapt(http.HandlerFunc(route.Handler.SaveAssetMaster))).Methods("POST")
	r.Handle("/assetmaster", Adapt(http.HandlerFunc(route.Handler.GetSingleAssetMaster))).Methods("GET")
	r.Handle("/assetmaster", Adapt(http.HandlerFunc(route.Handler.UpdateAssetMaster))).Methods("PUT")
	r.Handle("/assetmaster/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAssetMaster))).Methods("PUT")
	r.Handle("/assetmaster/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAssetMaster))).Methods("PUT")
	r.Handle("/assetmaster/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAssetMaster))).Methods("DELETE")
	r.Handle("/assetmaster/filter", Adapt(http.HandlerFunc(route.Handler.FilterAssetMaster))).Methods("POST")

}
