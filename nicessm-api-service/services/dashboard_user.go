package services

import (
	"fmt"
	"nicessm-api-service/models"
	"time"
)

func (s *Service) DashboardUserCount(ctx *models.Context, userfilter *models.DashboardUserCountFilter) (user []models.DashboardUserCountReport, err error) {
	err = s.UserDataAccess(ctx, &userfilter.UserFilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.DashboardUserCount(ctx, userfilter)
}
func (s *Service) DayWiseUserDemandChart(ctx *models.Context, userfilter *models.DashboardUserCountFilter) (user *models.DayWiseUserDemandChartReport, err error) {
	if userfilter == nil {
		userfilter = new(models.DashboardUserCountFilter)
	}
	if userfilter.CreatedFrom.StartDate == nil {
		t := time.Now()
		userfilter.CreatedFrom.StartDate = &t
	}
	sd := s.Shared.BeginningOfMonth(*userfilter.CreatedFrom.StartDate)
	ed := s.Shared.EndOfMonth(sd)
	fmt.Println(sd.Month(), " & ", ed.Month())
	userfilter.CreatedFrom.StartDate = &sd
	userfilter.CreatedFrom.EndDate = &ed
	err = s.UserDataAccess(ctx, &userfilter.UserFilter)
	if err != nil {
		return nil, err
	}
	//userfilter.DataAccess = *dataAceess
	return s.Daos.DayWiseUserDemandChart(ctx, userfilter)
}
