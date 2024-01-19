package services

import (
	"errors"
	"haritv2-service/models"
	"log"
)

func (s *Service) BlockChainWithoutTransaction(ctx *models.Context, orderID string) error {
	log.Println("transaction start")
	order, err := s.Daos.GetSingleOrder(ctx, orderID)
	if err != nil {
		return errors.New("Error in geting order - " + err.Error())
	}
	if order == nil {
		return errors.New("order is nil")

	}
	var quantity float64
	for _, v := range order.Items {
		quantity = quantity + v.Quantity
	}

	inventory := new(models.Inventory)
	inventory.ID = orderID
	inventory.From.ID = order.Company.ID
	inventory.From.Name = order.Company.Name
	inventory.From.Type = order.Company.Type
	inventory.To.ID = order.Customer.ID
	inventory.To.Name = order.Customer.Name
	inventory.To.Type = order.Customer.Type
	inventory.Quantity = quantity
	inventory.Price = order.TotalAmount
	inventory.TimeStramp = *order.Date
	switch order.Company.Type {
	case "FPO":
		RefFPOInventory, err := s.Daos.GetSingleFPOInventoryWithCompalyID(ctx, order.Company.ID)
		if err != nil {
			return errors.New("Error in finding FPO Inventory customer - " + err.Error())
		}
		if RefFPOInventory != nil {
			inventory.From.BeforeInventory = RefFPOInventory.Quantity
		}
	case "ULB":
		ULBInventory, err := s.Daos.GetSingleULBInventoryWithCompalyID(ctx, order.Company.ID)
		if err != nil {
			return errors.New("Error in finding ULB Inventory customer - " + err.Error())
		}
		if ULBInventory != nil {
			inventory.From.BeforeInventory = ULBInventory.Quantity
		}
	}
	switch order.Customer.Type {
	case "FPO":
		RefFPOInventory, err := s.Daos.GetSingleFPOInventoryWithCompalyID(ctx, order.Company.ID)
		if err != nil {
			return errors.New("Error in finding FPO Inventory customer - " + err.Error())
		}
		if RefFPOInventory != nil {
			inventory.To.BeforeInventory = RefFPOInventory.Quantity
		}
	case "ULB":
		ULBInventory, err := s.Daos.GetSingleULBInventoryWithCompalyID(ctx, order.Company.ID)
		if err != nil {
			return errors.New("Error in finding ULB Inventory customer - " + err.Error())
		}
		if ULBInventory != nil {
			inventory.To.BeforeInventory = ULBInventory.Quantity
		}
	}
	err = s.Daos.SaveBlockChain(ctx, inventory)
	if err != nil {
		return errors.New("Error in saving sale" + err.Error())
	}
	return nil
}
