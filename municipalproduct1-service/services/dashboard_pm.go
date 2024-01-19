package services

import (
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
)

// PmDashboard : ""
func (s *Service) PmDashboard(ctx *models.Context, PmFilter *models.PmDashboardFilter) (*models.PmDashboard, error) {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Pmdashboard := new(models.PmDashboard)
	Pmdashboard.TodaysCollectedAmount = 10000
	Pmdashboard.CurrentMonthTarget.AcheivedAmount = 300000
	Pmdashboard.CurrentMonthTarget.PendingAmount = 200000
	Pmdashboard.CurrentMonthTarget.TargetAmount = 500000
	Pmdashboard.MyAccess.CollectionOfHouses = 350
	Pmdashboard.MyAccess.TotalCollection = 800000
	Pmdashboard.MyAccess.TotalDemand = 18800000
	Pmdashboard.MyAccess.TotalSurvey = 2500
	Pmdashboard.Overall = new(models.DashboardDemandAndCollection)
	//Filling OverAll Demand
	demandfilter := new(models.DashboardDemandAndCollectionFilter)
	demandfilter.Status = append(demandfilter.Status, constants.PROPERTYSTATUSACTIVE, constants.PROPERTYSTATUSPENDING)

	overall, err := s.DashboardDemandAndCollection(ctx, demandfilter)
	if err != nil {
		fmt.Println(err)
		overall = new(models.DashboardDemandAndCollection)
	}

	Pmdashboard.Overall = overall

	return Pmdashboard, nil
}
