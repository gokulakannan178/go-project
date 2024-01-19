package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (route *Route) BankInformationRoutes(r *mux.Router) {
	// BankInformation
	r.Handle("/bankInformation", Adapt(http.HandlerFunc(route.Handler.SaveBankInformation))).Methods("POST")
	r.Handle("/bankInformation", Adapt(http.HandlerFunc(route.Handler.GetSingleBankInformation))).Methods("GET")
	r.Handle("/bankInformation", Adapt(http.HandlerFunc(route.Handler.UpdateBankInformation))).Methods("PUT")
	r.Handle("/bankInformation/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableBankInformation))).Methods("PUT")
	r.Handle("/bankInformation/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableBankInformation))).Methods("PUT")
	r.Handle("/bankInformation/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteBankInformation))).Methods("DELETE")
	r.Handle("/bankInformation/filter", Adapt(http.HandlerFunc(route.Handler.FilterBankInformation))).Methods("POST")
	r.Handle("/bankInformation/updateEmployeBankInfromation", Adapt(http.HandlerFunc(route.Handler.UpdateEmployeeBankInformation))).Methods("PUT")
	r.Handle("/bankInformation/employee", Adapt(http.HandlerFunc(route.Handler.GetSingleBankInformationWithEmployee))).Methods("GET")

}
