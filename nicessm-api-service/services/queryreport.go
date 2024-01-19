package services

import (
	"nicessm-api-service/models"
)

//FilterQueryReport :""
func (s *Service) FilterQueryReport(ctx *models.Context, Queryfilter *models.QueryReportFilter, pagination *models.Pagination) (Query []models.RefQuery, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterQueryReport(ctx, Queryfilter, pagination)

}
