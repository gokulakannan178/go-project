package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// InitiatePropertyOtherDemandPayment : ""
func (s *Service) InitiatePropertyOtherDemandPayment(ctx *models.Context, filter *models.InitiatePropertyOtherDemandFilter) (string, error) {
	var transactionID string
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		podFilter := new(models.PropertyOtherDemandFilter)

		podFilter.PropertyID = append(podFilter.PropertyID, filter.PropertyID)
		podFilter.UniqueIDs = filter.RecordID
		podFilter.PaymentStatus = append(podFilter.Status, constants.PROPERTYOTHERDEMANDPAYMENTSTATUSNOTPAID)
		fmt.Println("podFilter.PropertyID ===>", podFilter.PropertyID)
		fmt.Println("podFilter.UniqueIDs ===>", podFilter.UniqueIDs)
		fmt.Println("podFilter.PaymentStatus ===>", podFilter.PaymentStatus)

		res, err := s.FilterPropertyOtherDemand(ctx, podFilter, nil)
		if err != nil {
			return errors.New("error in getting the property other demand" + err.Error())
		}
		//Prepare Property Other Demand Payment
		tnxID := s.Shared.GetTransactionID(filter.PropertyID, 32)
		payment := new(models.PropertyOtherDemandPayment)
		payment.ReciptNo = s.Daos.GetUniqueID(ctx, "recipt")
		for _, v := range res {
			payment.PropertyID = v.PropertyID
			payment.Status = constants.PROPERTYOTHERDEMANDPAYMENTSTATUSINIT
			payment.RecordID = v.UniqueID
			payment.TotalTax = v.Amount

		}
		payment.TnxID = tnxID
		payment.PropertyID = filter.PropertyID
		payment.Status = constants.PROPERTYOTHERDEMANDPAYMENTSTATUSINIT

		//Prepare Property Basic details
		propertyPaymentDemandBasic := new(models.PropertyOtherDemandPaymentDemandBasic)
		propertyPaymentDemandBasic.TnxID = tnxID
		// propertyPaymentDemandBasic.Property = demand.Property
		propertyPaymentDemandBasic.Status = constants.PROPERTYOTHERDEMANDPAYMENTSTATUSINIT

		// floors, floorErr := s.Daos.GetFloorsOfProperty(ctx, filter.PropertyID)
		// if floorErr != nil {
		// 	log.Println("Error in geting Floors - " + floorErr.Error())
		// }
		// propertyPaymentDemandBasic.Floors = floors

		owners, ownerErr := s.Daos.GetOwnersOfProperty(ctx, filter.PropertyID)
		if ownerErr != nil {
			log.Println("Error in geting Owner - " + ownerErr.Error())
		}
		propertyPaymentDemandBasic.Owners = owners

		if err := s.Daos.SavePropertyOtherDemandPaymentDemandBasic(ctx, propertyPaymentDemandBasic); err != nil {
			return errors.New("Errror in saving other demand payment basics - " + err.Error())
		}
		//Prepare Fy Details
		ppdfys := []models.PropertyOtherDemandPaymentDemandFy{}
		for _, v := range res {
			propertyPaymentDemandFy := models.PropertyOtherDemandPaymentDemandFy{}
			propertyPaymentDemandFy.TnxID = tnxID
			propertyPaymentDemandFy.PropertyID = filter.PropertyID
			propertyPaymentDemandFy.FY.FinancialYear = v.Ref.FY
			propertyPaymentDemandFy.Status = constants.PROPERTYOTHERDEMANDPAYMENTSTATUSINIT
			ppdfys = append(ppdfys, propertyPaymentDemandFy)
		}

		if err := s.Daos.SaveManyPropertyOtherDemandPaymentDemandFy(ctx, ppdfys); err != nil {
			return errors.New("Errror in saving payment demand fys - " + err.Error())
		}
		if err := s.Daos.SavePropertyOtherDemandPayment(ctx, payment); err != nil {
			return errors.New("Errror in saving payment - " + err.Error())
		}

		transactionID = tnxID
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return "", errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return "", err
	}
	return transactionID, nil

}

// GetSinglePropertyOtherDemandPaymentTxtID : ""
func (s *Service) GetSinglePropertyOtherDemandPaymentTxtID(ctx *models.Context, id string) (*models.RefPropertyOtherDemandPayment, error) {
	refPropertyPayment := new(models.RefPropertyOtherDemandPayment)
	payment, err := s.Daos.GetSinglePropertyOtherDemandPaymentWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property other demand payment - " + err.Error())
	}
	propertyDemandBasic, err := s.Daos.GetSinglePropertyOtherDemandPaymentDemandBasicWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property other demand payment damand basic")
	}
	propertyDemandFys, err := s.Daos.GetPropertyOtherDemandPaymentDemandFycWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property other demand payment damand fys")
	}
	// ppFilter := new(models.PropertyOtherDemandPartPaymentFilter)
	// ppFilter.TnxID = []string{id}
	// ppFilter.Status = []string{constants.PROPERTYOTHERDEMANDPAYMENTCOMPLETED}
	// refPartPayments, _ := s.Daos.FilterPropertyOtherDemandPartPayment(ctx, ppFilter, nil)
	// if refPartPayments != nil {
	// 	refPropertyPayment.Ref.PartPayments = refPartPayments
	// 	for _, v := range refPartPayments {
	// 		refPropertyPayment.Ref.PartAmountCollected = refPropertyPayment.Ref.PartAmountCollected + v.Details.Amount
	// 	}
	// }
	refPropertyPayment.PropertyOtherDemandPayment = *payment
	refPropertyPayment.Basic = *propertyDemandBasic
	refPropertyPayment.Fys = propertyDemandFys
	state, err := s.Daos.GetSingleState(ctx, payment.Address.StateCode)
	if state != nil {
		refPropertyPayment.Ref.Address.State = &state.State
	}
	fmt.Println(err)
	district, err := s.Daos.GetSingleDistrict(ctx, payment.Address.DistrictCode)
	if district != nil {
		refPropertyPayment.Ref.Address.District = &district.District
	}
	fmt.Println(err)
	village, err := s.Daos.GetSingleVillage(ctx, payment.Address.VillageCode)
	if village != nil {
		refPropertyPayment.Ref.Address.Village = &village.Village
	}
	fmt.Println(err)
	zone, err := s.Daos.GetSingleZone(ctx, payment.Address.ZoneCode)
	if zone != nil {
		refPropertyPayment.Ref.Address.Zone = &zone.Zone
	}
	fmt.Println(err)
	ward, err := s.Daos.GetSingleWard(ctx, payment.Address.WardCode)
	if ward != nil {
		refPropertyPayment.Ref.Address.Ward = &ward.Ward
	}
	fmt.Println(err)
	if payment.Details != nil {
		collector, err := s.Daos.GetSingleUser(ctx, payment.Details.Collector.ID)
		if collector != nil {
			refPropertyPayment.Ref.Collector = collector.User
		}
		fmt.Println(err)
	}

	return refPropertyPayment, nil
}

// PropertyOtherDemandMakePayment : ""
func (s *Service) PropertyOtherDemandMakePayment(ctx *models.Context, payment *models.PropertyOtherDemandMakePayment) (string, error) {
	var propertyId string
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		user, err := s.GetSingleUser(ctx, payment.Details.Collector.ID)
		if err != nil {
			fmt.Println("problem in geting user data " + err.Error())
		}
		if payment.Details.Collector.ID != "system" {
			fmt.Println("user ======> ", user)
			if user.Status != constants.USERSTATUSACTIVE {
				return errors.New("non active user")
			}
		}
		py, err := s.GetSinglePropertyOtherDemandPaymentTxtID(ctx, payment.TnxID)
		if err != nil {
			fmt.Println("problem in geting property payment for transaction id - " + err.Error())
		}

		if err := s.Daos.CompletePropertyOtherDemandPaymentWithTxtID(ctx, payment, false); err != nil {
			return errors.New("Error in updating in - Payment" + err.Error())
		}
		if err := s.Daos.CompleteSinglePropertyOtherDemandPaymentDemandBasicWithTxtID(ctx, payment); err != nil {
			return errors.New("Error in updating in - basic" + err.Error())
		}
		if err := s.Daos.CompletePropertyOtherDemandPaymentDemandFycWithTxtID(ctx, payment); err != nil {
			return errors.New("Error in updating in - Fys" + err.Error())
		}

		if err := s.Daos.UpdatePropertyOtherDemandStatus(ctx, py.RecordID); err != nil {
			return errors.New("Error in updating propertydemand status - Fys" + err.Error())
		}
		propertyId = py.PropertyID
		// if py != nil {
		// 	if err := s.Daos.UpdateBoringStatusToProperty(ctx, py.PropertyID, true); err != nil {
		// 		return errors.New("Error in updating in - Boring Charge" + err.Error())
		// 	}
		// 	fmt.Println("Boring Charge Updated")
		// 	if err := s.Daos.UpdateFormFeeStatusToProperty(ctx, py.PropertyID, true); err != nil {
		// 		return errors.New("Error in updating in - Form Fee" + err.Error())
		// 	}
		// 	fmt.Println("Form Fee Update")
		// 	//
		// 	if !py.Demand.Property.PreviousCollection.IsCalculated {
		// 		if err := s.Daos.UpdatePayedPropertyPreviousYrCollection(ctx, py.PropertyID); err != nil {
		// 			return errors.New("Error in previous year collection" + err.Error())
		// 		}
		// 		fmt.Println("previous years collection updated ")
		// 	}

		// } else {
		// 	fmt.Println("property not found")
		// }
		// s.Daos.Update
		// Generate Message
		msg := fmt.Sprintf("Payment of amount %v is received against property Tax for Service No. %v.", payment.Details.Amount, py.ReciptNo)
		if payment.Details.MOP.Mode == constants.MOPCASH {
			msg = fmt.Sprintf("Payment of amount %v is received against property Tax for Service No. %v.", payment.Details.Amount, py.ReciptNo)
		}
		if payment.Details.MOP.Mode == constants.MOPCHEQUE {
			msg = fmt.Sprintf("Cheque for payment amount %v is received against property Tax for Service No. %v.", payment.Details.Amount, py.ReciptNo)
		}
		if payment.Details.MOP.Mode == constants.MOPNETBANKING {
			msg = fmt.Sprintf("Payment of amount %v is initiated through net banking against property Tax for Service No. %v.", payment.Details.Amount, py.ReciptNo)
		}
		if payment.Details.MOP.Mode == constants.MOPDD {
			msg = fmt.Sprintf("DD for payment amount %v is received against property Tax for Service No. %v.", payment.Details.Amount, py.ReciptNo)
		}
		if len(py.Basic.Owners) > 0 {
			mobileNo := py.Basic.Owners[0].Mobile
			err = s.SendSMS(mobileNo, msg)
			if err != nil {
				return err
			}
		}

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return "", errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return "", err
	}
	return propertyId, nil
}

// PropertyOtherDemandVerifyPayment : ""
func (s *Service) PropertyOtherDemandVerifyPayment(ctx *models.Context, vp *models.PropertyOtherDemandVerifyPayment) (string, error) {
	t := time.Now()
	vp.ActionDate = &t
	if vp.Date == nil {
		vp.Date = &t
	}
	err := s.Daos.PropertyOtherDemandVerifyPayment(ctx, vp)
	if err != nil {
		return "", err
	}
	py, err := s.GetSinglePropertyOtherDemandPaymentTxtID(ctx, vp.TnxID)
	if err != nil {
		fmt.Println("problem in geting property payment for transaction id - " + err.Error())
	}
	if err := s.Daos.UpdatePropertyOtherDemandStatus(ctx, py.RecordID); err != nil {
		return "", err
	}
	propertypayment, err := s.Daos.GetSinglePropertyOtherDemandPaymentWithTxtID(ctx, vp.TnxID)
	if err != nil {
		return "", err
	}
	return propertypayment.PropertyID, err
}

// PropertyOtherDemandNotVerifiedPayment : ""
func (s *Service) PropertyOtherDemandNotVerifiedPayment(ctx *models.Context, vp *models.PropertyOtherDemandNotVerifiedPayment) error {
	t := time.Now()
	vp.ActionDate = &t
	if vp.Date == nil {
		vp.Date = &t
	}
	err := s.Daos.PropertyOtherDemandNotVerifiedPayment(ctx, vp)
	return err
}

// PropertyOtherDemandRejectPayment : ""
func (s *Service) PropertyOtherDemandRejectPayment(ctx *models.Context, rp *models.PropertyOtherDemandRejectPayment) (string, error) {
	t := time.Now()
	rp.ActionDate = &t
	if rp.Date == nil {
		rp.Date = &t
	}
	err := s.Daos.PropertyOtherDemandRejectPayment(ctx, rp)
	if err != nil {
		return "", err
	}
	propertypayment, err := s.Daos.GetSinglePropertyOtherDemandPaymentWithTxtID(ctx, rp.TnxID)
	if err != nil {
		return "", err
	}
	return propertypayment.PropertyID, err
}
