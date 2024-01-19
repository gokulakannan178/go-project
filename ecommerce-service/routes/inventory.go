package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) InventoryRoutes(r *mux.Router) {
	// Inventory
	r.Handle("/inventory", Adapt(http.HandlerFunc(route.Handler.SaveInventory))).Methods("POST")
	r.Handle("/inventory", Adapt(http.HandlerFunc(route.Handler.GetSingleInventory))).Methods("GET")
	r.Handle("/inventory", Adapt(http.HandlerFunc(route.Handler.UpdateInventory))).Methods("PUT")
	r.Handle("/inventory/quantitydetail", Adapt(http.HandlerFunc(route.Handler.UpdateInventoryQuantityDetails))).Methods("PUT")
	r.Handle("/inventory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableInventory))).Methods("PUT")
	r.Handle("/inventory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableInventory))).Methods("PUT")
	r.Handle("/inventory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteInventory))).Methods("DELETE")
	r.Handle("/inventory/filter", Adapt(http.HandlerFunc(route.Handler.FilterInventory))).Methods("POST")
	// r.Handle("/inventory/generatemesh", Adapt(http.HandlerFunc(route.Handler.CreateMesh))).Methods("POST")
	r.Handle("/inventory/image", Adapt(http.HandlerFunc(route.Handler.ImageInventory))).Methods("PUT")
	r.Handle("/inventory/images", Adapt(http.HandlerFunc(route.Handler.ImagesInventory))).Methods("PUT")
	r.Handle("/inventory/getbybarcodeandvendor", Adapt(http.HandlerFunc(route.Handler.GetbyBarcodeAndVendor))).Methods("GET")
	r.Handle("/inventory/chkuniqueness/barcode", Adapt(http.HandlerFunc(route.Handler.ChkUniqueness))).Methods("GET")
}
