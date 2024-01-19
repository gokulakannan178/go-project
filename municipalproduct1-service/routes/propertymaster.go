package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//ConstructionTypeRoutes : ""
func (route *Route) ConstructionTypeRoutes(r *mux.Router) {
	r.Handle("/constructionType", Adapt(http.HandlerFunc(route.Handler.SaveConstructionType))).Methods("POST")
	r.Handle("/constructionType", Adapt(http.HandlerFunc(route.Handler.GetSingleConstructionType))).Methods("GET")
	r.Handle("/constructionType", Adapt(http.HandlerFunc(route.Handler.UpdateConstructionType))).Methods("PUT")
	r.Handle("/constructionType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableConstructionType))).Methods("PUT")
	r.Handle("/constructionType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableConstructionType))).Methods("PUT")
	r.Handle("/constructionType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteConstructionType))).Methods("DELETE")
	r.Handle("/constructionType/filter", Adapt(http.HandlerFunc(route.Handler.FilterConstructionType))).Methods("POST")
}

//FloorTypeRoutes : ""
func (route *Route) FloorTypeRoutes(r *mux.Router) {
	r.Handle("/floorType", Adapt(http.HandlerFunc(route.Handler.SaveFloorType))).Methods("POST")
	r.Handle("/floorType", Adapt(http.HandlerFunc(route.Handler.GetSingleFloorType))).Methods("GET")
	r.Handle("/floorType", Adapt(http.HandlerFunc(route.Handler.UpdateFloorType))).Methods("PUT")
	r.Handle("/floorType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFloorType))).Methods("PUT")
	r.Handle("/floorType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFloorType))).Methods("PUT")
	r.Handle("/floorType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFloorType))).Methods("DELETE")
	r.Handle("/floorType/filter", Adapt(http.HandlerFunc(route.Handler.FilterFloorType))).Methods("POST")
}

//MunicipalTypeRoutes : ""
func (route *Route) MunicipalTypeRoutes(r *mux.Router) {
	r.Handle("/municipalType", Adapt(http.HandlerFunc(route.Handler.SaveMunicipalType))).Methods("POST")
	r.Handle("/municipalType", Adapt(http.HandlerFunc(route.Handler.GetSingleMunicipalType))).Methods("GET")
	r.Handle("/municipalType", Adapt(http.HandlerFunc(route.Handler.UpdateMunicipalType))).Methods("PUT")
	r.Handle("/municipalType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableMunicipalType))).Methods("PUT")
	r.Handle("/municipalType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableMunicipalType))).Methods("PUT")
	r.Handle("/municipalType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteMunicipalType))).Methods("DELETE")
	r.Handle("/municipalType/filter", Adapt(http.HandlerFunc(route.Handler.FilterMunicipalType))).Methods("POST")
	r.Handle("/municipalType/selectable", Adapt(http.HandlerFunc(route.Handler.GetSelectableMunicipalType))).Methods("GET")

}

//OccupancyTypeRoutes : ""
func (route *Route) OccupancyTypeRoutes(r *mux.Router) {
	r.Handle("/occupancyType", Adapt(http.HandlerFunc(route.Handler.SaveOccupancyType))).Methods("POST")
	r.Handle("/occupancyType", Adapt(http.HandlerFunc(route.Handler.GetSingleOccupancyType))).Methods("GET")
	r.Handle("/occupancyType", Adapt(http.HandlerFunc(route.Handler.UpdateOccupancyType))).Methods("PUT")
	r.Handle("/occupancyType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOccupancyType))).Methods("PUT")
	r.Handle("/occupancyType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOccupancyType))).Methods("PUT")
	r.Handle("/occupancyType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOccupancyType))).Methods("DELETE")
	r.Handle("/occupancyType/filter", Adapt(http.HandlerFunc(route.Handler.FilterOccupancyType))).Methods("POST")
}

//RoadTypeRoutes : ""
func (route *Route) RoadTypeRoutes(r *mux.Router) {
	r.Handle("/roadType", Adapt(http.HandlerFunc(route.Handler.SaveRoadType))).Methods("POST")
	r.Handle("/roadType", Adapt(http.HandlerFunc(route.Handler.GetSingleRoadType))).Methods("GET")
	r.Handle("/roadType", Adapt(http.HandlerFunc(route.Handler.UpdateRoadType))).Methods("PUT")
	r.Handle("/roadType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableRoadType))).Methods("PUT")
	r.Handle("/roadType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableRoadType))).Methods("PUT")
	r.Handle("/roadType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteRoadType))).Methods("DELETE")
	r.Handle("/roadType/filter", Adapt(http.HandlerFunc(route.Handler.FilterRoadType))).Methods("POST")
}

//UsageTypeRoutes : ""
func (route *Route) UsageTypeRoutes(r *mux.Router) {
	r.Handle("/usageType", Adapt(http.HandlerFunc(route.Handler.SaveUsageType))).Methods("POST")
	r.Handle("/usageType", Adapt(http.HandlerFunc(route.Handler.GetSingleUsageType))).Methods("GET")
	r.Handle("/usageType", Adapt(http.HandlerFunc(route.Handler.UpdateUsageType))).Methods("PUT")
	r.Handle("/usageType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableUsageType))).Methods("PUT")
	r.Handle("/usageType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableUsageType))).Methods("PUT")
	r.Handle("/usageType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteUsageType))).Methods("DELETE")
	r.Handle("/usageType/filter", Adapt(http.HandlerFunc(route.Handler.FilterUsageType))).Methods("POST")
}

//PropertyTypeRoutes : ""
func (route *Route) PropertyTypeRoutes(r *mux.Router) {
	r.Handle("/propertyType", Adapt(http.HandlerFunc(route.Handler.SavePropertyType))).Methods("POST")
	r.Handle("/propertyType", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyType))).Methods("GET")
	r.Handle("/propertyType", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyType))).Methods("PUT")
	r.Handle("/propertyType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyType))).Methods("PUT")
	r.Handle("/propertyType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyType))).Methods("PUT")
	r.Handle("/propertyType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyType))).Methods("DELETE")
	r.Handle("/propertyType/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyType))).Methods("POST")
}

//FinancialYearRoutes : ""
func (route *Route) FinancialYearRoutes(r *mux.Router) {
	r.Handle("/financialYear", Adapt(http.HandlerFunc(route.Handler.SaveFinancialYear))).Methods("POST")
	r.Handle("/financialYear", Adapt(http.HandlerFunc(route.Handler.GetSingleFinancialYear))).Methods("GET")
	r.Handle("/financialYear", Adapt(http.HandlerFunc(route.Handler.UpdateFinancialYear))).Methods("PUT")
	r.Handle("/financialYear/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFinancialYear))).Methods("PUT")
	r.Handle("/financialYear/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFinancialYear))).Methods("PUT")
	r.Handle("/financialYear/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFinancialYear))).Methods("DELETE")
	r.Handle("/financialYear/filter", Adapt(http.HandlerFunc(route.Handler.FilterFinancialYear))).Methods("POST")
	r.Handle("/financialYear/currentyear", Adapt(http.HandlerFunc(route.Handler.MakeCurrentFinancialYear))).Methods("PUT")
	r.Handle("/financialYear/currentyear", Adapt(http.HandlerFunc(route.Handler.GetCurrentFinancialYear))).Methods("GET")

}

//OwnershipRoutes : ""
func (route *Route) OwnershipRoutes(r *mux.Router) {
	r.Handle("/ownership", Adapt(http.HandlerFunc(route.Handler.SaveOwnership))).Methods("POST")
	r.Handle("/ownership", Adapt(http.HandlerFunc(route.Handler.GetSingleOwnership))).Methods("GET")
	r.Handle("/ownership", Adapt(http.HandlerFunc(route.Handler.UpdateOwnership))).Methods("PUT")
	r.Handle("/ownership/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOwnership))).Methods("PUT")
	r.Handle("/ownership/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOwnership))).Methods("PUT")
	r.Handle("/ownership/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOwnership))).Methods("DELETE")
	r.Handle("/ownership/filter", Adapt(http.HandlerFunc(route.Handler.FilterOwnership))).Methods("POST")
}

//VacantLandRateRoutes : ""
func (route *Route) VacantLandRateRoutes(r *mux.Router) {
	r.Handle("/vacantLandRate", Adapt(http.HandlerFunc(route.Handler.SaveVacantLandRate))).Methods("POST")
	r.Handle("/vacantLandRate", Adapt(http.HandlerFunc(route.Handler.GetSingleVacantLandRate))).Methods("GET")
	r.Handle("/vacantLandRate", Adapt(http.HandlerFunc(route.Handler.UpdateVacantLandRate))).Methods("PUT")
	r.Handle("/vacantLandRate/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableVacantLandRate))).Methods("PUT")
	r.Handle("/vacantLandRate/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableVacantLandRate))).Methods("PUT")
	r.Handle("/vacantLandRate/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteVacantLandRate))).Methods("DELETE")
	r.Handle("/vacantLandRate/filter", Adapt(http.HandlerFunc(route.Handler.FilterVacantLandRate))).Methods("POST")
}

//FloorRatableAreaRoutes : ""
func (route *Route) FloorRatableAreaRoutes(r *mux.Router) {
	r.Handle("/floorRatableArea", Adapt(http.HandlerFunc(route.Handler.SaveFloorRatableArea))).Methods("POST")
	r.Handle("/floorRatableArea", Adapt(http.HandlerFunc(route.Handler.GetSingleFloorRatableArea))).Methods("GET")
	r.Handle("/floorRatableArea", Adapt(http.HandlerFunc(route.Handler.UpdateFloorRatableArea))).Methods("PUT")
	r.Handle("/floorRatableArea/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableFloorRatableArea))).Methods("PUT")
	r.Handle("/floorRatableArea/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableFloorRatableArea))).Methods("PUT")
	r.Handle("/floorRatableArea/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteFloorRatableArea))).Methods("DELETE")
	r.Handle("/floorRatableArea/filter", Adapt(http.HandlerFunc(route.Handler.FilterFloorRatableArea))).Methods("POST")
}

//AVRRoutes : ""
func (route *Route) AVRRoutes(r *mux.Router) {
	r.Handle("/avr", Adapt(http.HandlerFunc(route.Handler.SaveAVR))).Methods("POST")
	r.Handle("/avr", Adapt(http.HandlerFunc(route.Handler.GetSingleAVR))).Methods("GET")
	r.Handle("/avr", Adapt(http.HandlerFunc(route.Handler.UpdateAVR))).Methods("PUT")
	r.Handle("/avr/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableAVR))).Methods("PUT")
	r.Handle("/avr/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableAVR))).Methods("PUT")
	r.Handle("/avr/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteAVR))).Methods("DELETE")
	r.Handle("/avr/filter", Adapt(http.HandlerFunc(route.Handler.FilterAVR))).Methods("POST")
}

//NonResidentialUsageFactorRoutes : ""
func (route *Route) NonResidentialUsageFactorRoutes(r *mux.Router) {
	r.Handle("/nonResidentialUsageFactor", Adapt(http.HandlerFunc(route.Handler.SaveNonResidentialUsageFactor))).Methods("POST")
	r.Handle("/nonResidentialUsageFactor", Adapt(http.HandlerFunc(route.Handler.GetSingleNonResidentialUsageFactor))).Methods("GET")
	r.Handle("/nonResidentialUsageFactor", Adapt(http.HandlerFunc(route.Handler.UpdateNonResidentialUsageFactor))).Methods("PUT")
	r.Handle("/nonResidentialUsageFactor/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableNonResidentialUsageFactor))).Methods("PUT")
	r.Handle("/nonResidentialUsageFactor/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableNonResidentialUsageFactor))).Methods("PUT")
	r.Handle("/nonResidentialUsageFactor/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteNonResidentialUsageFactor))).Methods("DELETE")
	r.Handle("/nonResidentialUsageFactor/filter", Adapt(http.HandlerFunc(route.Handler.FilterNonResidentialUsageFactor))).Methods("POST")
}

//PropertyTaxRoutes : ""
func (route *Route) PropertyTaxRoutes(r *mux.Router) {
	r.Handle("/propertyTax", Adapt(http.HandlerFunc(route.Handler.SavePropertyTax))).Methods("POST")
	r.Handle("/propertyTax", Adapt(http.HandlerFunc(route.Handler.GetSinglePropertyTax))).Methods("GET")
	r.Handle("/propertyTax", Adapt(http.HandlerFunc(route.Handler.UpdatePropertyTax))).Methods("PUT")
	r.Handle("/propertyTax/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePropertyTax))).Methods("PUT")
	r.Handle("/propertyTax/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePropertyTax))).Methods("PUT")
	r.Handle("/propertyTax/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePropertyTax))).Methods("DELETE")
	r.Handle("/propertyTax/filter", Adapt(http.HandlerFunc(route.Handler.FilterPropertyTax))).Methods("POST")
}

//RebateRoutes : ""
func (route *Route) RebateRoutes(r *mux.Router) {
	r.Handle("/rebate", Adapt(http.HandlerFunc(route.Handler.SaveRebate))).Methods("POST")
	r.Handle("/rebate", Adapt(http.HandlerFunc(route.Handler.GetSingleRebate))).Methods("GET")
	r.Handle("/rebate", Adapt(http.HandlerFunc(route.Handler.UpdateRebate))).Methods("PUT")
	r.Handle("/rebate/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableRebate))).Methods("PUT")
	r.Handle("/rebate/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableRebate))).Methods("PUT")
	r.Handle("/rebate/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteRebate))).Methods("DELETE")
	r.Handle("/rebate/filter", Adapt(http.HandlerFunc(route.Handler.FilterRebate))).Methods("POST")
}

//PenaltyRoutes : ""
func (route *Route) PenaltyRoutes(r *mux.Router) {
	r.Handle("/penalty", Adapt(http.HandlerFunc(route.Handler.SavePenalty))).Methods("POST")
	r.Handle("/penalty", Adapt(http.HandlerFunc(route.Handler.GetSinglePenalty))).Methods("GET")
	r.Handle("/penalty", Adapt(http.HandlerFunc(route.Handler.UpdatePenalty))).Methods("PUT")
	r.Handle("/penalty/status/enable", Adapt(http.HandlerFunc(route.Handler.EnablePenalty))).Methods("PUT")
	r.Handle("/penalty/status/disable", Adapt(http.HandlerFunc(route.Handler.DisablePenalty))).Methods("PUT")
	r.Handle("/penalty/status/delete", Adapt(http.HandlerFunc(route.Handler.DeletePenalty))).Methods("DELETE")
	r.Handle("/penalty/filter", Adapt(http.HandlerFunc(route.Handler.FilterPenalty))).Methods("POST")
}

//ResidentialTypeRoutes : ""
func (route *Route) ResidentialTypeRoutes(r *mux.Router) {
	r.Handle("/residentialType", Adapt(http.HandlerFunc(route.Handler.SaveResidentialType))).Methods("POST")
	r.Handle("/residentialType", Adapt(http.HandlerFunc(route.Handler.GetSingleResidentialType))).Methods("GET")
	r.Handle("/residentialType", Adapt(http.HandlerFunc(route.Handler.UpdateResidentialType))).Methods("PUT")
	r.Handle("/residentialType/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableResidentialType))).Methods("PUT")
	r.Handle("/residentialType/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableResidentialType))).Methods("PUT")
	r.Handle("/residentialType/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteResidentialType))).Methods("DELETE")
	r.Handle("/residentialType/filter", Adapt(http.HandlerFunc(route.Handler.FilterResidentialType))).Methods("POST")
}

// PropertyTaxCalculation : ""
func (route *Route) PropertyTaxCalculation(r *mux.Router) {
	r.Handle("/propertyTaxCalculation", Adapt(http.HandlerFunc(route.Handler.GetPropertyTaxCalculation))).Methods("GET")
}

//OtherChargesRoutes : ""
func (route *Route) OtherChargesRoutes(r *mux.Router) {
	r.Handle("/otherCharges", Adapt(http.HandlerFunc(route.Handler.SaveOtherCharges))).Methods("POST")
	r.Handle("/otherCharges", Adapt(http.HandlerFunc(route.Handler.GetSingleOtherCharges))).Methods("GET")
	r.Handle("/otherCharges", Adapt(http.HandlerFunc(route.Handler.UpdateOtherCharges))).Methods("PUT")
	r.Handle("/otherCharges/status/enable", Adapt(http.HandlerFunc(route.Handler.EnableOtherCharges))).Methods("PUT")
	r.Handle("/otherCharges/status/disable", Adapt(http.HandlerFunc(route.Handler.DisableOtherCharges))).Methods("PUT")
	r.Handle("/otherCharges/status/delete", Adapt(http.HandlerFunc(route.Handler.DeleteOtherCharges))).Methods("DELETE")
	r.Handle("/otherCharges/filter", Adapt(http.HandlerFunc(route.Handler.FilterOtherCharges))).Methods("POST")

}

// SurveyAndTax : ""
func (route *Route) SurveyAndTax(r *mux.Router) {
	r.Handle("/surveyandtax", Adapt(http.HandlerFunc(route.Handler.SaveSurveyAndTax))).Methods("POST")
	r.Handle("/surevyandtax", Adapt(http.HandlerFunc(route.Handler.GetSingleSurveyAndTax))).Methods("GET")
	r.Handle("/surveyandtax/pushnotification", Adapt(http.HandlerFunc(route.Handler.PushNotification))).Methods("POST")
	r.Handle("/surveyandtaxfilter", Adapt(http.HandlerFunc(route.Handler.SurveyAndTaxFilter))).Methods("POST")

}
