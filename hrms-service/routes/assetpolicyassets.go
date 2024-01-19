package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AssetPolicyAssetsRoutes
func (route *Route) AssetPolicyAssetsRoutes(r *mux.Router) {
	r.Handle("/assetpolicyassets", Adapt(http.HandlerFunc(route.Handler.SaveAssetPolicyAssets))).Methods("POST")
	r.Handle("/assetpolicyassets", Adapt(http.HandlerFunc(route.Handler.GetSingleAssetPolicyAssets))).Methods("GET")
	r.Handle("/assetpolicyassets", Adapt(http.HandlerFunc(route.Handler.UpdateAssetPolicyAssets))).Methods("PUT")
	r.Handle("/assetpolicyassets/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAssetPolicyAssets))).Methods("PUT")
	r.Handle("/assetpolicyassets/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAssetPolicyAssets))).Methods("PUT")
	r.Handle("/assetpolicyassets/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAssetPolicyAssets))).Methods("DELETE")
	r.Handle("/assetpolicyassets/filter", Adapt(http.HandlerFunc(route.Handler.FilterAssetPolicyAssets))).Methods("POST")

}
