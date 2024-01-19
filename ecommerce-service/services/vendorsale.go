package services

import (
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
)

func (s *Service) VendorInitiateSale(ctx *models.Context, sale *models.Sale) error {
	if sale == nil {
		return errors.New("Sale can't be nil")
	}
	sale.From.Type = constants.USERTYPEVENDOR
	if sale.From.ID == "" {
		return errors.New("Vendor id can't be nil")
	}
	vendor, err := s.Daos.GetSingleVendor(ctx, sale.From.ID)
	if err != nil {
		return errors.New("error getting vendor - " + err.Error())
	}
	if vendor != nil {
		return errors.New("vendor can't be nil - ")
	}
	return nil

}
