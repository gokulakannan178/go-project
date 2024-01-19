package services

import (
	"fmt"
	"nicessm-api-service/models"
	"time"
)

func (s *Service) DashboardFarmerCount(ctx *models.Context, farmerfilter *models.DashboardFarmerCountFilter) (content []models.DashboardFarmerCountReport, err error) {
	err = s.FarmerDataAccess(ctx, &farmerfilter.FarmerFilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.DashboardFarmerCount(ctx, farmerfilter)
}
func (s *Service) DayWiseFarmerDemandChart(ctx *models.Context, farmerfilter *models.DashboardFarmerCountFilter) (farmer *models.DayWiseFarmerDemandChartReport, err error) {
	if farmerfilter == nil {
		farmerfilter = new(models.DashboardFarmerCountFilter)
	}
	if farmerfilter.CreatedDate.From == nil {
		t := time.Now()
		farmerfilter.CreatedDate.From = &t
	}
	sd := s.Shared.BeginningOfMonth(*farmerfilter.CreatedDate.From)
	ed := s.Shared.EndOfMonth(sd)
	fmt.Println(sd.Month(), " & ", ed.Month())
	farmerfilter.CreatedDate.From = &sd
	farmerfilter.FarmerFilter.CreatedDate.To = &ed
	err = s.FarmerDataAccess(ctx, &farmerfilter.FarmerFilter)
	if err != nil {
		return nil, err
	}
	//farmerfilter.DataAccess = *dataAceess
	return s.Daos.DayWiseFarmerDemandChart(ctx, farmerfilter)
}
