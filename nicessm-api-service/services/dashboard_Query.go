package services

import (
	"fmt"
	"nicessm-api-service/models"
	"time"
)

func (s *Service) DashboardQueryCount(ctx *models.Context, queryfilter *models.DashboardQueryCountFilter) (content []models.DashboardQueryCountReport, err error) {
	err = s.QueryDataAccess(ctx, &queryfilter.QueryFilter)
	if err != nil {
		return nil, err
	}
	return s.Daos.DashboardQueryCount(ctx, queryfilter)
}
func (s *Service) DayWiseQueryDemandChart(ctx *models.Context, queryfilter *models.DashboardQueryCountFilter) (query *models.DayWiseQueryDemandChartReport, err error) {
	if queryfilter == nil {
		queryfilter = new(models.DashboardQueryCountFilter)
	}
	if queryfilter.CreatedFrom.StartDate == nil {
		t := time.Now()
		queryfilter.CreatedFrom.StartDate = &t
	}
	sd := s.Shared.BeginningOfMonth(*queryfilter.CreatedFrom.StartDate)
	ed := s.Shared.EndOfMonth(sd)
	fmt.Println(sd.Month(), " & ", ed.Month())
	queryfilter.CreatedFrom.StartDate = &sd
	queryfilter.CreatedFrom.EndDate = &ed
	err = s.QueryDataAccess(ctx, &queryfilter.QueryFilter)
	if err != nil {
		return nil, err
	}
	//queryfilter.DataAccess = *dataAceess
	return s.Daos.DayWiseQueryDemandChart(ctx, queryfilter)
}
