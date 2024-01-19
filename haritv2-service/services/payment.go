package services

import (
	"haritv2-service/models"
)

//FilterPayment :""
func (s *Service) FilterPayment(ctx *models.Context, filter *models.PaymentFilter, pagination *models.Pagination) ([]models.RefPayment, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterPayment(ctx, filter, pagination)

}
