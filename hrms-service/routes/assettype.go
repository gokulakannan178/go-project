package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AssetTypeRoutes
func (route *Route) AssetTypeRoutes(r *mux.Router) {
	r.Handle("/assettype", Adapt(http.HandlerFunc(route.Handler.SaveAssetType))).Methods("POST")
	r.Handle("/assettype/withproperty", Adapt(http.HandlerFunc(route.Handler.SaveAssetTypeWithPropertys))).Methods("POST")
	r.Handle("/assettype", Adapt(http.HandlerFunc(route.Handler.GetSingleAssetType))).Methods("GET")
	r.Handle("/assettype", Adapt(http.HandlerFunc(route.Handler.UpdateAssetType))).Methods("PUT")
	r.Handle("/assettype/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAssetType))).Methods("PUT")
	r.Handle("/assettype/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAssetType))).Methods("PUT")
	r.Handle("/assettype/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAssetType))).Methods("DELETE")
	r.Handle("/assettype/filter", Adapt(http.HandlerFunc(route.Handler.FilterAssetType))).Methods("POST")
	r.Handle("/assettype/updateassettypeproperty", Adapt(http.HandlerFunc(route.Handler.UpdateAssetTypeWithProperty))).Methods("PUT")
	//r.Handle("/assettype/getproperty", Adapt(http.HandlerFunc(route.Handler.GetAssetTypePropertys))).Methods("GET")

}
