package services

import (
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) InitiateOrder(ctx *models.Context, co *models.CreateOrder) (*models.Order, error) {
	order := new(models.Order)
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}

	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		order.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORDER)
		t := time.Now().UTC()

		order.Date = &t
		switch co.Customer.Type {
		case "FPO":
			refFPO, err := s.Daos.GetSingleFPO(ctx, co.Customer.ID)
			if refFPO == nil {
				return errors.New("FPO Not Available" + err.Error())
			}
			if err != nil {
				return errors.New("Error in finding FPO customer - " + err.Error())
			}
			if refFPO != nil {
				order.Customer.ID = refFPO.UniqueID
				order.Customer.Type = "FPO"
				order.Customer.Name = refFPO.Name
				order.Customer.Address = refFPO.Address
				order.Customer.Contact = refFPO.Mobile
				order.Customer.Email = refFPO.Email
				order.Company.User = refFPO.ChairMan

				state, stateErr := s.Daos.GetSingleState(ctx, refFPO.Address.StateCode)
				if stateErr != nil {
					log.Println("Error in geting State - " + stateErr.Error())
				}
				if state != nil {
					order.Customer.State.ID = state.Code
					order.Customer.State.Label = state.Name
				}
				district, districtErr := s.Daos.GetSingleDistrict(ctx, refFPO.Address.DistrictCode)
				if districtErr != nil {
					log.Println("Error in geting district - " + stateErr.Error())
				}
				if state != nil && district != nil {
					order.Address.Billing = refFPO.Address.AL1 + " " + district.Name + " " + state.Name + " " + refFPO.Address.PostalCode
					order.Address.Shipping = refFPO.Address.AL1 + " " + district.Name + " " + state.Name + " " + refFPO.Address.PostalCode
				}
			}

		case "ULB":
			refULB, err := s.Daos.GetSingleULB(ctx, co.Customer.ID)
			if refULB == nil {
				return errors.New(" ULB Not Available" + err.Error())
			}
			if err != nil {
				return errors.New("Error in finding ULB customer - " + err.Error())
			}
			if refULB != nil {
				order.Customer.ID = refULB.UniqueID
				order.Customer.Type = "ULB"
				order.Customer.Name = refULB.Name
				order.Customer.Address = refULB.Address
				order.Customer.Contact = refULB.NodalOfficer.MobileNo
				order.Customer.Email = refULB.NodalOfficer.Email
				order.Company.User = refULB.NodalOfficer.UserName

				state, stateErr := s.Daos.GetSingleState(ctx, refULB.Address.StateCode)
				if stateErr != nil {
					log.Println("Error in geting State - " + stateErr.Error())
				}
				if state != nil {
					order.Customer.State.ID = state.Code
					order.Customer.State.Label = state.Name
				}
				district, districtErr := s.Daos.GetSingleDistrict(ctx, refULB.Address.DistrictCode)
				if districtErr != nil {
					log.Println("Error in geting district - " + stateErr.Error())
				}
				if state != nil && district != nil {
					order.Address.Billing = refULB.Address.AL1 + " " + district.Name + " " + state.Name + " " + refULB.Address.PostalCode
					order.Address.Shipping = refULB.Address.AL1 + " " + district.Name + " " + state.Name + " " + refULB.Address.PostalCode
				}
			}
		case "Customer":
			refCustomer, err := s.Daos.GetSingleCustomerwithmobileno(ctx, co.Customer.Mobile)
			if err != nil {
				return errors.New("Error in finding  customer - " + err.Error())
			}
			if refCustomer == nil {
				refCustomer = new(models.RefCustomer)
				refCustomer.Name = co.Customer.Name
				//refCustomer.PrimaryContact.Mobile = co.Customer.Mobile
				refCustomer.PrimaryContact.Ph = co.Customer.Mobile
				refCustomer.Gender = co.Customer.Gender
				refCustomer.Address.PinCode = co.Customer.PinCode
				refCustomer.Status = constants.CUSTOMERSSTATUSACTIVE
				t := time.Now()
				created := models.CreatedV2{}
				created.On = &t
				created.By = constants.SYSTEM
				refCustomer.Created = created
				refCustomer.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCUSTOMER)
				err := s.Daos.SaveCustomer(ctx, &refCustomer.Customer)
				if err != nil {
					return err
				}
			}

			order.Customer.ID = refCustomer.UniqueID
			order.Customer.Type = "Customer"
			order.Customer.Name = refCustomer.Name
			order.Customer.Gender = refCustomer.Gender
			order.Customer.Address.PinCode = refCustomer.Address.PinCode
			order.Customer.Contact = refCustomer.PrimaryContact.Ph
			order.Customer.Email = refCustomer.Email

		// order.Customer.ID = refCustomer.UniqueID
		// order.Customer.Type = "Customer"
		// order.Customer.Name = refCustomer.Name
		// order.Customer.Address.PinCode = refCustomer.PinCode
		// order.Customer.Contact = refCustomer.Mobile
		// order.Customer.Email = refCustomer.Email
		//order.Company.User = refCustomer

		// state, stateErr := s.Daos.GetSingleState(ctx, refCustomer.Ref.Address.State.Code)
		// if stateErr != nil {
		// 	log.Println("Error in geting State - " + stateErr.Error())
		// }
		// if state != nil {
		// 	order.Customer.State.ID = state.Code
		// 	order.Customer.State.Label = state.Name
		// }
		// // district, districtErr := s.Daos.GetSingleDistrict(ctx, refCustomer.Ref.Address.District.Code)
		// if districtErr != nil {
		// 	log.Println("Error in geting district - " + stateErr.Error())
		// }
		// // order.Address.Billing = refCustomer.Ref.Address.District.Name + " " + district.Name + " " + state.Name + " " + refCustomer.Ref.Address.District.PostalCode
		// order.Address.Shipping = refCustomer.Ref.Address.District.Name + " " + district.Name + " " + state.Name + " " + refCustomer.Address.PostalCode
		case "bulkcustomer":
			order.Customer.Type = "bulkcustomer"
			order.Customer.MaleCount = co.Customer.MaleCount
			order.Customer.FemaleCount = co.Customer.FemaleCount
		default:
			return nil
		}

		switch co.Company.Type {
		case "ULB":
			refULB, err := s.Daos.GetSingleULB(ctx, co.Company.ID)
			if err != nil {
				return errors.New("Error in finding ULB company - " + err.Error())
			}
			if refULB == nil {
				return errors.New("ULB is not Available - ")

			}
			order.CompanyID = refULB.UniqueID
			order.Company.ID = refULB.UniqueID
			order.Company.Type = "ULB"
			order.Company.Name = refULB.Name
			order.Company.Logo = refULB.Logo
			order.Company.Address = refULB.Address.AL1
			order.Company.User = refULB.NodalOfficer.Name
			order.Company.Contact = refULB.NodalOfficer.MobileNo
			order.Company.Email = refULB.NodalOfficer.Email

			state, stateErr := s.Daos.GetSingleState(ctx, refULB.Address.StateCode)
			if stateErr != nil {
				log.Println("Error in geting State - " + stateErr.Error())
			}
			if state != nil {
				order.Company.State.ID = state.Code
				order.Company.State.Label = state.Name
			}
		case "FPO":
			refFPO, err := s.Daos.GetSingleFPO(ctx, co.Company.ID)
			if err != nil {
				return errors.New("Error in finding ULB company - " + err.Error())
			}
			if refFPO == nil {
				return errors.New("ULB is not Available - ")

			}
			order.CompanyID = refFPO.UniqueID
			order.Company.ID = refFPO.UniqueID
			order.Company.Type = "FPO"
			order.Company.Name = refFPO.Name
			order.Company.Logo = refFPO.Logo
			order.Company.Address = refFPO.Address.AL1
			order.Company.User = refFPO.ChairMan
			order.Company.Contact = refFPO.Mobile
			order.Company.Email = refFPO.Email

			state, stateErr := s.Daos.GetSingleState(ctx, refFPO.Address.StateCode)
			if stateErr != nil {
				log.Println("Error in geting State - " + stateErr.Error())
			}
			if state != nil {
				order.Company.State.ID = state.Code
				order.Company.State.Label = state.Name
			}

		}

		var saleitem models.SaleItems
		saleitem.Quantity = co.Product.Quantity
		// Get Selling Price

		switch co.Company.Type {
		case "FPO":
			RefFPOInventory, err := s.Daos.GetSingleFPOInventoryWithCompalyID(ctx, co.Company.ID)
			if err != nil {
				return errors.New("Error in finding FPO Inventory customer - " + err.Error())
			}
			if RefFPOInventory != nil {
				saleitem.Price = RefFPOInventory.Sellingprice
			}
		case "ULB":
			ULBInventory, err := s.Daos.GetSingleULBInventoryWithCompalyID(ctx, co.Company.ID)
			if err != nil {
				return errors.New("Error in finding ULB Inventory customer - " + err.Error())
			}
			if ULBInventory != nil {
				saleitem.Price = ULBInventory.Price
			}
		}

		//Get default Product
		refProduct, err := s.Daos.GetDefaultProduct(ctx)
		if err != nil {
			log.Println("Err in geting default product - " + err.Error())
		}
		fmt.Println("product is", refProduct.Product)
		if refProduct != nil {
			saleitem.Product.ID = refProduct.UniqueID
			saleitem.Product.Name = refProduct.Name
			saleitem.Product.Unit = refProduct.Unit
			saleitem.Product.HSN = refProduct.HSN
			saleitem.Product.IsTaxable = refProduct.IsTaxable
			saleitem.Product.IntraTaxRate = refProduct.IntraTaxRate
			saleitem.Product.InterTaxate = refProduct.InterTaxRate
		}
		//get Default pakage Type

		//Get Default GST
		refGST, err := s.Daos.GetDefaultGST(ctx)
		if err != nil {
			log.Println("Err in geting default product - " + err.Error())
		}
		fmt.Println("gst is", refGST.GST)
		if refGST != nil {
			saleitem.Tax.GST.ID = refGST.UniqueID
			saleitem.Tax.GST.Percentage = refGST.Percentage
			saleitem.Tax.GST.Label = refGST.Label
			saleitem.Tax.GST.CGST = 0
			saleitem.Tax.GST.SGST = 0
			saleitem.Tax.GST.IGST = 0
			saleitem.Tax.GST.Total = 0
		}
		saleitem.Amount = saleitem.Quantity * saleitem.Price
		saleitem.Tax.Total = 0
		saleitem.TotalAmount = saleitem.Amount + saleitem.Tax.Total
		order.Items = append(order.Items, saleitem)

		for _, v := range order.Items {
			order.SubTotal = order.SubTotal + v.TotalAmount
			order.TotalTax = order.TotalTax + v.Tax.Total
		}
		order.TotalAmount = math.Round(order.SubTotal)
		var transport models.SaleTransport
		transport.Status = constants.SALETRANSPORTSTATUSPENDING
		transport.Type = constants.SALETRANSPORTATIONTYPETAKEWAY
		order.Transport = transport
		order.Created.On = &t
		order.OrderPayment = co.OrderPayment
		if err := s.Daos.SaveOrder(ctx, order); err != nil {
			return errors.New("Error in saving order - " + err.Error())
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}

		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return nil, errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return nil, err
	}
	return order, nil
}

// GetSingleOrder : ""
func (s *Service) GetSingleOrder(ctx *models.Context, UniqueID string) (*models.RefOrder, error) {
	order, err := s.Daos.GetSingleOrder(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

//FilterProduct :""
func (s *Service) FilterOrder(ctx *models.Context, filter *models.OrderFilter, pagination *models.Pagination) ([]models.RefOrder, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterOrder(ctx, filter, pagination)

}

// PlaceOrder : ""
func (s *Service) PlaceOrder(ctx *models.Context, orderID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.PlaceOrderWithoutTransaction(ctx, orderID)
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

// InitiateNotification : ""
func (s *Service) InitiateNotification(ctx *models.Context, orderID string) (*models.OrderNotification, error) {
	resOrder, err := s.Daos.GetSingleOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if resOrder == nil {
		return nil, errors.New("order is nil")
	}
	const (
		ULB = "ULB"
		FPO = "FPO"
	)
	ordernotify := new(models.OrderNotification)

	switch resOrder.Company.Type {

	case ULB:
		resULB, err := s.Daos.GetSingleULB(ctx, resOrder.Company.ID)
		if err != nil {
			return nil, err
		}

		ordernotify.CompanyMobileNo = resULB.NodalOfficer.MobileNo
		ordernotify.CompanyEmailID = resULB.NodalOfficer.Email
		ordernotify.CompanyName = resULB.Name
		ordernotify.CompanySpocName = resULB.NodalOfficer.Name
		resAppToken, err := s.Daos.GetRegTokenWithParticulars(ctx, ULB, resOrder.Company.ID)
		if err != nil {
			return nil, err
		}
		if resAppToken != nil {
			ordernotify.CompanyAppToken = resAppToken.RegistrationToken

		} else {
			log.Println("resAppToken.RegistrationToken - s nil)")

		}
	case FPO:
		resFPO, err := s.Daos.GetSingleFPO(ctx, resOrder.Company.ID)
		if err != nil {
			return nil, err
		}
		ordernotify.CompanyMobileNo = resFPO.Mobile
		ordernotify.CompanyEmailID = resFPO.Email
		ordernotify.CompanyName = resFPO.Name
		ordernotify.CompanySpocName = resFPO.ChairMan
		resAppToken, err := s.Daos.GetRegTokenWithParticulars(ctx, FPO, resOrder.Company.ID)
		if err != nil {
			return nil, err
		}
		if resAppToken != nil {
			ordernotify.CompanyAppToken = resAppToken.RegistrationToken
		} else {
			log.Println("resAppToken.RegistrationToken - s nil)")

		}
	}

	switch resOrder.Customer.Type {
	case ULB:
		resULB, err := s.Daos.GetSingleULB(ctx, resOrder.Customer.ID)
		if err != nil {
			return nil, err
		}
		ordernotify.CustomerMobileNo = resULB.NodalOfficer.MobileNo
		ordernotify.CustomerEmailID = resULB.NodalOfficer.Email
		ordernotify.CustomerName = resULB.NodalOfficer.Name
		ordernotify.CustomerFirmName = resULB.Name
		resAppToken, err := s.Daos.GetRegTokenWithParticulars(ctx, ULB, resOrder.Customer.ID)
		if err != nil {
			return nil, err
		}
		if resAppToken != nil {
			ordernotify.CustomerAppToken = resAppToken.RegistrationToken
		} else {
			log.Println("resAppToken.RegistrationToken - s nil)")

		}
	case FPO:
		resFPO, err := s.Daos.GetSingleFPO(ctx, resOrder.Customer.ID)
		if err != nil {
			return nil, err
		}
		ordernotify.CustomerMobileNo = resFPO.Mobile
		ordernotify.CustomerEmailID = resFPO.Email
		ordernotify.CustomerName = resFPO.ChairMan
		ordernotify.CustomerFirmName = resFPO.Name
		resAppToken, err := s.Daos.GetRegTokenWithParticulars(ctx, FPO, resOrder.Customer.ID)
		if err != nil {
			return nil, err
		}
		if resAppToken != nil {
			ordernotify.CustomerAppToken = resAppToken.RegistrationToken

		} else {
			log.Println("resAppToken.RegistrationToken - s nil)")
		}
	}
	for _, v := range resOrder.Items {
		ordernotify.Quantity = ordernotify.Quantity + v.Quantity
	}

	return ordernotify, nil

}
func (s *Service) OrderCancel(ctx *models.Context, order *models.OrderCancelFilter) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.OrderCancel(ctx, order)
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

func (s *Service) PlaceOrderWithoutTransaction(ctx *models.Context, orderID string) error {
	log.Println("transaction start")
	err := s.Daos.PlaceOrder(ctx, orderID)
	if err != nil {
		return errors.New("Error in saving order - " + err.Error())
	}

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
	// resInitiateOrder, err := s.InitiateNotification(ctx, orderID)
	// if err != nil {
	// 	return errors.New("Error in initiating order - " + err.Error())
	// }

	// //SMS to Company
	// msg := fmt.Sprintf("You have received an order Req from %v Quantity - %v Amount - %v", order.Customer.Name, quantity, order.TotalAmount)
	// text := fmt.Sprintf(constants.COMMONTEMPLATE, order.Company.User, "URNCCH", "Order Request ("+order.UniqueID+")", msg, "http://urncc.org")
	// if order.Company.Contact != "" {
	// 	err = s.SendSMS(order.Company.Contact, text)
	// 	if err != nil {
	// 		return errors.New("Sms Sending Error - " + err.Error())
	// 	}
	// }
	// //Email to Company
	// if order.Company.Email != "" {
	// 	err := s.SendEmail("Harit - Order Request ("+order.UniqueID+")", []string{order.Company.Email}, msg)
	// 	if err != nil {
	// 		fmt.Println("Sms Sending Error - " + err.Error())
	// 	}

	// }
	// // Notiication to company
	// if resInitiateOrder.CompanyAppToken != "" {
	// 	err = s.SendNotification(
	// 		"Harit - Order Request ("+order.UniqueID+")",
	// 		msg,
	// 		"",
	// 		[]string{resInitiateOrder.CompanyAppToken})
	// 	if err != nil {
	// 		fmt.Println("error in sending notification to company " + err.Error())
	// 	}
	// }

	// //SMS to Customer
	// msg = fmt.Sprintf("You have successfully placed order to %v -  Quantity - %v Amount - %v", order.Company.Name, quantity, order.TotalAmount)
	// text = fmt.Sprintf(constants.COMMONTEMPLATE, order.Customer.User, "URNCCH", "Order Request ("+order.UniqueID+")", msg, "http://urncc.org")
	// if order.Company.Contact != "" {
	// 	err = s.SendSMS(order.Company.Contact, text)
	// 	if err != nil {
	// 		return errors.New("Sms Sending Error - " + err.Error())
	// 	}
	// }

	// //Email to Customer
	// if order.Customer.Email != "" {
	// 	err := s.SendEmail("Harit - Order Request ("+order.UniqueID+")", []string{order.Customer.Email}, msg)
	// 	if err != nil {
	// 		fmt.Println("Sms Sending Error - " + err.Error())
	// 	}

	// }

	// // Notiication to customer
	// if resInitiateOrder.CustomerAppToken != "" {
	// 	err = s.SendNotification(
	// 		"Harit - Order Request ("+order.UniqueID+")",
	// 		msg,
	// 		"",
	// 		[]string{resInitiateOrder.CustomerAppToken})
	// 	if err != nil {
	// 		fmt.Println("error in sending notification to company " + err.Error())
	// 	}
	// }

	// // Send SMS to company
	// err = s.SendSMS(resInitiateOrder.CompanyMobileNo, fmt.Sprintf(constants.PLACEOREDERSMSTOCOMPANY, resInitiateOrder.CompanySpocName, resInitiateOrder.CompanyName, resInitiateOrder.CustomerFirmName, resInitiateOrder.Quantity, orderID))
	// if err != nil {
	// 	fmt.Println("error in sending message to company " + err.Error())
	// }
	// // Send SMS to customer
	// err = s.SendSMS(resInitiateOrder.CustomerMobileNo, fmt.Sprintf(constants.PLACEOREDERSMSTOCUSTOMER, resInitiateOrder.CustomerName, resInitiateOrder.CustomerFirmName, resInitiateOrder.CompanyName, orderID))
	// if err != nil {
	// 	fmt.Println("error in sending message to customer " + err.Error())
	// }
	// // Send Email to company
	// err = s.SendEmail(fmt.Sprintf(constants.PLACEORDEREMAILSUBJECTTOCOMPANY, resInitiateOrder.CompanySpocName, orderID), []string{resInitiateOrder.CompanyEmailID}, fmt.Sprintf(constants.PLACEORDEREMAILBODYTOCOMPANY, resInitiateOrder.CompanySpocName, resInitiateOrder.CompanyName, resInitiateOrder.CustomerFirmName, resInitiateOrder.Quantity, orderID))
	// if err != nil {
	// 	fmt.Println("error in sending email to company " + err.Error())
	// }
	// // Send Email to customer
	// err = s.SendEmail(fmt.Sprintf(constants.PLACEORDEREMAILSUBJECTTOCUSTOMER, resInitiateOrder.CustomerFirmName, orderID), []string{resInitiateOrder.CustomerEmailID}, fmt.Sprintf(constants.PLACEORDEREMAILBODYTOCUSTOMER, resInitiateOrder.CustomerName, resInitiateOrder.CustomerFirmName, resInitiateOrder.CompanyName, orderID))
	// if err != nil {
	// 	fmt.Println("error in sending email to company " + err.Error())
	// }
	// // Send Email to company
	// err = s.SendNotification(fmt.Sprintf(constants.PLACEORDERNOTIFICATIONTITILETOCOMPANY, resInitiateOrder.CompanySpocName, resInitiateOrder.CompanyName), fmt.Sprintf(constants.PLACEORDERNOTIFICATIONBODYTOCOMPANY, resInitiateOrder.CompanySpocName, resInitiateOrder.CompanyName, resInitiateOrder.CustomerFirmName, resInitiateOrder.Quantity, orderID), "", []string{resInitiateOrder.CompanyAppToken})
	// if err != nil {
	// 	fmt.Println("error in sending notification to company " + err.Error())
	// }
	// // Send Email to customer
	// err = s.SendNotification(fmt.Sprintf(constants.PLACEORDERNOTIFICATIONTITILETOCUSTOMER, resInitiateOrder.CustomerName, resInitiateOrder.CustomerFirmName), fmt.Sprintf(constants.PLACEORDERNOTIFICATIONBODYTOCUSTOMER, resInitiateOrder.CustomerName, resInitiateOrder.CustomerFirmName, resInitiateOrder.CompanyName, orderID), "", []string{resInitiateOrder.CustomerAppToken})
	// if err != nil {
	// 	fmt.Println("error in sending notification to company " + err.Error())
	// }
	return nil
}

// UpdateSelfConsumption : ""
func (s *Service) RejectedOrder(ctx *models.Context, selfconsumption *models.Order) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectedOrder(ctx, selfconsumption)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}
