package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ExpenseCategoryRoutes : ""
func (route *Route) ExpenseCategoryRoutes(r *mux.Router) {
	r.Handle("/expensecategory", Adapt(http.HandlerFunc(route.Handler.SaveExpenseCategory))).Methods("POST")
	r.Handle("/expensecategory", Adapt(http.HandlerFunc(route.Handler.GetSingleExpenseCategory))).Methods("GET")
	r.Handle("/expensecategory", Adapt(http.HandlerFunc(route.Handler.UpdateExpenseCategory))).Methods("PUT")
	r.Handle("/expensecategory/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableExpenseCategory))).Methods("PUT")
	r.Handle("/expensecategory/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableExpenseCategory))).Methods("PUT")
	r.Handle("/expensecategory/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteExpenseCategory))).Methods("DELETE")
	r.Handle("/expensecategory/filter", Adapt(http.HandlerFunc(route.Handler.FilterExpenseCategory))).Methods("POST")

}
