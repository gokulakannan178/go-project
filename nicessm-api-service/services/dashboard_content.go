package services

import (
	"fmt"
	"nicessm-api-service/models"
	"time"
)

func (s *Service) DashboardContentSmsCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) (content []models.DashboardContentCountReport, err error) {
	err = s.ContentDataAccess(ctx, &contentfilter.ContentFilter)
	if err != nil {
		return nil, err
	}
	//contentfilter.DataAccess = *dataAceess
	return s.Daos.DashboardContentSmsCount(ctx, contentfilter)
}
func (s *Service) DashboardContentVoiceCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) (content []models.DashboardContentCountReport, err error) {
	err = s.ContentDataAccess(ctx, &contentfilter.ContentFilter)
	if err != nil {
		return nil, err
	}
	//contentfilter.DataAccess = *dataAceess
	return s.Daos.DashboardContentVoiceCount(ctx, contentfilter)
}

func (s *Service) DashboardContentVideoCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) (content []models.DashboardContentCountReport, err error) {
	err = s.ContentDataAccess(ctx, &contentfilter.ContentFilter)
	if err != nil {
		return nil, err
	}
	//contentfilter.DataAccess = *dataAceess
	return s.Daos.DashboardContentVideoCount(ctx, contentfilter)
}

func (s *Service) DashboardContentPosterCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) (content []models.DashboardContentCountReport, err error) {
	err = s.ContentDataAccess(ctx, &contentfilter.ContentFilter)
	if err != nil {
		return nil, err
	}
	//contentfilter.DataAccess = *dataAceess
	return s.Daos.DashboardContentPosterCount(ctx, contentfilter)
}

func (s *Service) DashboardContentDocmentCount(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) (content []models.DashboardContentCountReport, err error) {
	err = s.ContentDataAccess(ctx, &contentfilter.ContentFilter)
	if err != nil {
		return nil, err
	}
	//contentfilter.DataAccess = *dataAceess
	return s.Daos.DashboardContentDocmentCount(ctx, contentfilter)
}
func (s *Service) DayWiseContentDemandChart(ctx *models.Context, contentfilter *models.DashboardContentCountFilter) (content *models.DayWiseContentDemandChartReport, err error) {
	if contentfilter == nil {
		contentfilter = new(models.DashboardContentCountFilter)
	}
	if contentfilter.ContentFilter.CreatedFrom.StartDate == nil {
		t := time.Now()
		contentfilter.ContentFilter.CreatedFrom.StartDate = &t
	}
	sd := s.Shared.BeginningOfMonth(*contentfilter.ContentFilter.CreatedFrom.StartDate)
	ed := s.Shared.EndOfMonth(sd)
	fmt.Println(sd.Month(), " & ", ed.Month())
	contentfilter.ContentFilter.CreatedFrom.StartDate = &sd
	contentfilter.ContentFilter.CreatedFrom.EndDate = &ed
	err = s.ContentDataAccess(ctx, &contentfilter.ContentFilter)
	if err != nil {
		return nil, err
	}
	//contentfilter.DataAccess = *dataAceess
	return s.Daos.DayWiseContentDemandChart(ctx, contentfilter)
}
