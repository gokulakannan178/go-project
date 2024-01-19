package services

import (
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) CreateFPOPurchaseULBSale(ctx *models.Context, orderId string) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.CreateFPOPurchaseULBSaleWithoutTransaction(ctx, orderId)
		if err != nil {
			return err
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

func (s *Service) CreateFPOPurchaseULBSaleWithoutTransaction(ctx *models.Context, orderId string) error {
	log.Println("transaction start")
	t := time.Now()
	refOrder, err := s.Daos.GetSingleOrder(ctx, orderId)
	if err != nil {
		return errors.New("Error in geting single order - " + err.Error())
	}
	if refOrder == nil {
		return errors.New("Error in geting single order in nil ")
	}
	refOrder.Status = constants.ORDERSTATUSACTIVE
	refOrder.PaymentStatus = constants.SALEPAYMENTSTATUSCOMPLETED
	refOrder.Transport.Status = constants.SALETRANSPORTSTATUSDELIVERED
	payment := new(models.Payment)
	payment.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPAYMENT)
	payment.SaleID = refOrder.UniqueID
	switch refOrder.OrderPayment.Type {
	case "Cash":
		payment.Type = constants.PAYMENTTYPECASH
	case "NetBanking":
		payment.Type = constants.PAYMENTTYPENETBANKING
		payment.ReferenceID = refOrder.OrderPayment.TnxID
	case "Cheque":
		payment.Type = constants.PAYMENTTYPECHEQUE
		payment.ReferenceID = refOrder.OrderPayment.TnxID
	case "Credit":
		payment.Type = constants.PAYMENTTYPECREDIT
		payment.ReferenceID = refOrder.OrderPayment.TnxID
	case "DD":
		payment.Type = constants.PAYMENTTYPEDD
		payment.ReferenceID = refOrder.OrderPayment.TnxID
	default:
		payment.Type = constants.PAYMENTTYPECASH

	}

	payment.Amount = refOrder.TotalAmount
	payment.Status = constants.SALEPAYMENTSTATUSCOMPLETED
	refFPO, err := s.Daos.GetSingleFPO(ctx, refOrder.Customer.ID)
	if err != nil {
		return errors.New("Error in finding FPO customer - " + err.Error())
	}
	if refFPO != nil {
		fmt.Println("updating payment.name as chairman name")
		payment.Name = refFPO.ChairMan

	}
	payment.Created.On = &t
	payment.Date = &t
	if err := s.Daos.SavePayment(ctx, payment); err != nil {
		return errors.New("Error in Saving Payment - " + err.Error())
	}
	fmt.Println("payment saved")

	// Updating Order Collection
	err = s.Daos.UpdateOrderStatus(ctx, refOrder.UniqueID)
	if err != nil {
		return errors.New("Error in updateing order - " + err.Error())
	}
	fmt.Println("updated order status")

	// Decreasing ULB Inventory Collection
	switch refOrder.Company.Type {
	case "ULB":
		ULBInventory, err := s.Daos.GetSingleULBInventoryWithCompalyID(ctx, refOrder.CompanyID)
		if ULBInventory == nil {
			return errors.New(" ULBInventory not Available" + err.Error())
		}
		if err != nil {
			return errors.New("Error in finding ULB Inventory customer - " + err.Error())
		}
		fmt.Println("ulb inventory found")

		if ULBInventory != nil {
			ULBInventory.Quantity = ULBInventory.Quantity - refOrder.Items[0].Quantity
			err := s.Daos.UpdateULBInventoryDeliverSale(ctx, ULBInventory)
			if err != nil {
				return errors.New("Unable to Update quantity in ULB" + err.Error())
			}
		}
		fmt.Println("updated ulb inventory")
		sale := refOrder.Sale
		sale.ID = primitive.NewObjectID()
		sale.Status = constants.SALESTATUSACTIVE
		sale.Transport.Status = constants.SALETRANSPORTSTATUSDELIVERED
		sale.PaymentStatus = constants.SALEPAYMENTSTATUSCOMPLETED
		err = s.Daos.SaveSale(ctx, &sale)
		if err != nil {
			return errors.New("Error in saving sale" + err.Error())
		}
		fmt.Println("sale saved")

	case "FPO":
		FPOInventory, err := s.Daos.GetSingleFPOInventoryWithCompalyID(ctx, refOrder.CompanyID)
		if err != nil {
			return errors.New("Error in finding FPO Inventory customer - " + err.Error())
		}
		fmt.Println("ulb inventory found")

		if FPOInventory != nil {
			FPOInventory.Quantity = FPOInventory.Quantity - refOrder.Items[0].Quantity
			err := s.Daos.UpdateFPOInventoryDeliverSale(ctx, FPOInventory)
			if err != nil {
				return errors.New("Unable to Update quantity in FPO" + err.Error())
			}
		}
		fmt.Println("updated FPO inventory")

		sale := refOrder.Sale
		sale.ID = primitive.NewObjectID()
		sale.Status = constants.SALESTATUSACTIVE
		sale.Transport.Status = constants.SALETRANSPORTSTATUSDELIVERED
		sale.PaymentStatus = constants.SALEPAYMENTSTATUSCOMPLETED
		err = s.Daos.SaveSale(ctx, &sale)
		if err != nil {
			return errors.New("Error in saving sale" + err.Error())
		}
		fmt.Println("sale saved")
	}
	// Increasing FPO Inventory Collection
	switch refOrder.Customer.Type {
	case "FPO":
		FPOInventory, err := s.Daos.GetSingleFPOInventoryWithCompalyID(ctx, refOrder.Customer.ID)
		if err != nil {
			return errors.New("Error in finding FPO Inventory customer - " + err.Error())
		}
		fmt.Println("fpo inventory found")

		if FPOInventory != nil {
			FPOInventory.Quantity = FPOInventory.Quantity + refOrder.Items[0].Quantity
			err := s.Daos.UpdateFPOInventoryDeliverSale(ctx, FPOInventory)
			if err != nil {
				return errors.New("Unable to Update quantity in FPO" + err.Error())
			}
		}
		fmt.Println("fpo inventory saved")
	case "ULB":
		ULBInventory, err := s.Daos.GetSingleULBInventoryWithCompalyID(ctx, refOrder.Customer.ID)
		if err != nil {
			return errors.New("Error in finding ULB Inventory customer - " + err.Error())
		}
		fmt.Println("ulb inventory found")

		if ULBInventory != nil {
			ULBInventory.Quantity = ULBInventory.Quantity + refOrder.Items[0].Quantity
			err := s.Daos.UpdateULBInventoryDeliverSale(ctx, ULBInventory)
			if err != nil {
				return errors.New("Unable to Update quantity in FPO" + err.Error())
			}
		}
		fmt.Println("ulb inventory saved")
	}

	//Calling InitiateNotification
	resInitiateOrder, err := s.InitiateNotification(ctx, orderId)
	if err != nil {
		return errors.New("Error in initiating order - " + err.Error())
	}

	fmt.Println("resInitiateOrder.CompanyMobileNo = ", resInitiateOrder.CompanyMobileNo)
	fmt.Println("resInitiateOrder.CustomerMobileNo = ", resInitiateOrder.CustomerMobileNo)
	fmt.Println("resInitiateOrder.CompanyEmailID = ", resInitiateOrder.CompanyEmailID)
	fmt.Println("resInitiateOrder.CustomerEmailID = ", resInitiateOrder.CustomerEmailID)
	fmt.Println("resInitiateOrder.CompanyAppToken = ", resInitiateOrder.CompanyAppToken)
	fmt.Println("resInitiateOrder.CustomerAppToken = ", resInitiateOrder.CustomerAppToken)
	// Send SMS to company
	msg := fmt.Sprintf("Order %v, is delivered to %v successfuly", orderId, resInitiateOrder.CustomerFirmName)
	text := fmt.Sprintf(constants.COMMONTEMPLATE, resInitiateOrder.CompanySpocName, "URNCCH", "Order ID ("+orderId+")", msg, "http://urncc.org")
	if resInitiateOrder.CompanyMobileNo != "" {
		err = s.SendSMS(resInitiateOrder.CompanyMobileNo, text)
		if err != nil {
			return errors.New("Sms Sending Error - " + err.Error())
		}
	}
	//Email to Company
	if resInitiateOrder.CompanyEmailID != "" {
		err := s.SendEmail("Harit - Order ("+orderId+") - Delivered", []string{resInitiateOrder.CompanyEmailID}, msg)
		if err != nil {
			fmt.Println("Sms Sending Error - " + err.Error())
		}

	}
	// Send Notification to company
	if resInitiateOrder.CompanyAppToken != "" {
		err = s.SendNotification(fmt.Sprintf(constants.DELIVERORDERNOTIFICATIONTITILETOCOMPANY, orderId, resInitiateOrder.CustomerFirmName), msg, "", []string{resInitiateOrder.CompanyAppToken})
		if err != nil {
			fmt.Println("error in sending notification to company " + err.Error())
		}
	}
	//SMS to customer

	msg = fmt.Sprintf("Order %v, is delivered successfuly", orderId)
	text = fmt.Sprintf(constants.COMMONTEMPLATE, resInitiateOrder.CustomerName, "URNCCH", "Order ID ("+orderId+")", msg, "http://urncc.org")
	if resInitiateOrder.CustomerMobileNo != "" {
		err = s.SendSMS(resInitiateOrder.CustomerMobileNo, text)
		if err != nil {
			return errors.New("Sms Sending Error - " + err.Error())
		}
	}

	//Email to Customer
	if resInitiateOrder.CustomerEmailID != "" {
		err := s.SendEmail("Harit - Order ("+orderId+") - Delivered", []string{resInitiateOrder.CustomerEmailID}, msg)
		if err != nil {
			fmt.Println("Sms Sending Error - " + err.Error())
		}

	}

	// Send Notification to customer
	if resInitiateOrder.CustomerAppToken != "" {
		err = s.SendNotification(fmt.Sprintf(constants.DELIVERORDERNOTIFICATIONTITILETOCUSTOMER, orderId, resInitiateOrder.CompanyName), msg, "", []string{resInitiateOrder.CustomerAppToken})
		if err != nil {
			fmt.Println("error in sending notification to customer " + err.Error())
		}
	}
	return nil
}

func (s *Service) PlaceAndDeliverOrder(ctx *models.Context, orderId string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.PlaceOrderWithoutTransaction(ctx, orderId)
		if err != nil {
			return err
		}
		err = s.BlockChainWithoutTransaction(ctx, orderId)
		if err != nil {
			return err
		}
		err = s.CreateFPOPurchaseULBSaleWithoutTransaction(ctx, orderId)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}
