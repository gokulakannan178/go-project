package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AssetPolicyRoutes
func (route *Route) AssetPolicyRoutes(r *mux.Router) {
	r.Handle("/assetpolicy", Adapt(http.HandlerFunc(route.Handler.SaveAssetPolicy))).Methods("POST")
	r.Handle("/assetpolicy", Adapt(http.HandlerFunc(route.Handler.GetSingleAssetPolicy))).Methods("GET")
	r.Handle("/assetpolicy", Adapt(http.HandlerFunc(route.Handler.UpdateAssetPolicy))).Methods("PUT")
	r.Handle("/assetpolicy/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAssetPolicy))).Methods("PUT")
	r.Handle("/assetpolicy/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAssetPolicy))).Methods("PUT")
	r.Handle("/assetpolicy/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAssetPolicy))).Methods("DELETE")
	r.Handle("/assetpolicy/filter", Adapt(http.HandlerFunc(route.Handler.FilterAssetPolicy))).Methods("POST")

}
