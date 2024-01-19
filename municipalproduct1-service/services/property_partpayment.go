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

func (s *Service) MakePaymentV2(ctx *models.Context, propertyPayment *models.PropertyPayment) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		partPayment := new(models.PropertyPartPayment)
		partPayment.TnxID = propertyPayment.TnxID
		partPayment.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYPARTPAYMENT)
		if propertyPayment.Details != nil {
			partPayment.Details = *propertyPayment.Details
		}
		py, err := s.GetSinglePropertyPaymentTxtID(ctx, propertyPayment.TnxID)
		if err != nil {
			fmt.Println("problem in geting property payment for transaction id - " + err.Error())
		}
		partPayment.ReciptNo = py.ReciptNo
		partPayment.PaymentDate = &t
		partPayment.PropertyID = py.PropertyID
		partPayment.Address = py.Address
		dberr := s.Daos.SavePropertyPartPayment(ctx, partPayment)
		if dberr != nil {
			return errors.New("Error in saving Part Payment  - " + dberr.Error())
		}
		dberr = s.Daos.PartiallyCompletedPropertyPaymentWithTxtIDForceFulForPartPayment(ctx, propertyPayment.TnxID)
		if dberr != nil {
			return errors.New("Error in Partially Completing main Payment - " + dberr.Error())
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

func (s *Service) MakePaymentV2AdditionalPayment(ctx *models.Context, propertyPayment *models.PropertyPayment) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		partPayment := new(models.PropertyPartPayment)
		partPayment.TnxID = propertyPayment.TnxID
		partPayment.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYPARTPAYMENT)
		if propertyPayment.Details != nil {
			partPayment.Details = *propertyPayment.Details
		}
		py, err := s.GetSinglePropertyPaymentTxtID(ctx, propertyPayment.TnxID)
		if err != nil {
			fmt.Println("problem in geting property payment for transaction id - " + err.Error())
		}
		partPayment.ReciptNo = py.ReciptNo
		partPayment.PaymentDate = &t
		partPayment.PropertyID = py.PropertyID
		partPayment.Address = py.Address
		dberr := s.Daos.SavePropertyPartPayment(ctx, partPayment)
		if dberr != nil {
			return errors.New("Error in saving Part Payment  - " + dberr.Error())
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

func (s *Service) ValidatePartPayments(ctx *models.Context, TnxID string) (string, error) {
	propertyId := ""
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return propertyId, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		propertyPayment, err := s.Daos.GetSinglePropertyPaymentWithTxtID(ctx, TnxID)
		if err != nil {
			return err
		}
		fmt.Println(propertyPayment)
		propertyId = propertyPayment.PropertyID
		filter := new(models.PropertyPartPaymentFilter)
		filter.TnxID = []string{TnxID}
		propertyPartPayment, err := s.Daos.FilterPropertyPartPayment(ctx, filter, nil)
		if err != nil {
			return err
		}
		fmt.Println(propertyPartPayment)

		//Calculating paidAmount
		var paidAmount float64
		for _, v := range propertyPartPayment {
			paidAmount = paidAmount + v.Details.Amount
		}
		if paidAmount >= propertyPayment.Demand.TotalTax {
			if err := s.Daos.CompletePropertyPaymentWithTxtIDForceFulForPartPayment(ctx, TnxID); err != nil {
				return errors.New("Er in completing payment - " + err.Error())
			}
			if err := s.Daos.UpdatePendingAmount(ctx, TnxID, 0); err != nil {
				return errors.New("Er in completing payment - " + err.Error())
			}
			if err := s.Daos.UpdateDetailsAmountForCompletedPartPayment(ctx, TnxID, paidAmount); err != nil {
				return errors.New("Er in UpdateDetailsAmountForCompletedPartPayment - " + err.Error())
			}
		} else {
			pendingAmount := propertyPayment.Demand.TotalTax - paidAmount
			if err := s.Daos.UpdatePendingAmount(ctx, TnxID, pendingAmount); err != nil {
				return errors.New("Er in completing payment - " + err.Error())
			}
			if err := s.Daos.PartiallyCompletedPropertyPaymentWithTxtIDForceFulForPartPayment(ctx, TnxID); err != nil {
				return errors.New("Er in PartiallyCompletedPropertyPaymentWithTxtIDForceFulForPartPayment - " + err.Error())
			}

		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return propertyId, errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return propertyId, err
	}
	return propertyId, nil
}
func (s *Service) ValidateMainPayment(ctx *models.Context, propertyID string) error {
	// refProperty, err := s.Daos.GetSingleProperty(ctx, propertyID)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("payment is completed so updating demand and collection")
	err := s.SavePropertyDemand(ctx, propertyID)
	if err != nil {
		return err
	}
	err = s.PropertyUpdateCollection(ctx, propertyID)
	if err != nil {
		return err
	}

	return nil
}

// VerifyPropertyPartPayment : ""
func (s *Service) VerifyPropertyPartPayment(ctx *models.Context, makeAction *models.MakePropertyPartPaymentsAction) (string, error) {
	log.Println("transaction start")
	propertyPartPaymentID := ""
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		var dberr error

		propertyPartPaymentID, dberr = s.Daos.VerifyPropertyPartPayment(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
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
	return propertyPartPaymentID, nil
}

// NotVerifyPropertyPartPayment : ""
func (s *Service) NotVerifyPropertyPartPayment(ctx *models.Context, makeAction *models.MakePropertyPartPaymentsAction) (string, error) {
	propertyPartPaymentID := ""
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		var dberr error

		propertyPartPaymentID, dberr = s.Daos.NotVerifyPropertyPartPayment(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
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
	return propertyPartPaymentID, nil
}

// RejectPropertyPartPayment : ""
func (s *Service) RejectPropertyPartPayment(ctx *models.Context, makeAction *models.MakePropertyPartPaymentsAction) (string, error) {
	log.Println("transaction start")
	propertyPartPaymentID := ""
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		makeAction.ActionDate = &t
		var dberr error
		propertyPartPaymentID, dberr = s.Daos.RejectPropertyPartPayment(ctx, makeAction)
		if dberr != nil {
			return dberr
		}
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
	return propertyPartPaymentID, nil
}

//FilterPropertyWallet :""
func (s *Service) FilterPropertyPartPayment(ctx *models.Context, filter *models.PropertyPartPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyPartPayment, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterPropertyPartPayment(ctx, filter, pagination)
}

// //CalcPropertyPartPaymentDemandForParticulars : ""
// func (s *Service) CalcPropertyPartPaymentDemandForParticulars(ctx *models.Context, filter *models.PropertyPartPaymentCalcQueryFilter) (*models.ShopRentDemand, error) {
// 	mtd := new(models.PropertyPartPaymentDemand)
// 	mtd.RefPropertyPartPayment.UniqueID = filter.PropertyPartPaymentID
// 	var dberr error
// 	mainpipeline, err := mtd.CalcQuery(filter)
// 	if err != nil {
// 		return nil, errors.New("Error in generating Query - " + err.Error())
// 	}
// 	mtd, dberr = s.Daos.CalcPropertyPartPaymentDemand(ctx, mainpipeline)
// 	if dberr != nil {
// 		return nil, dberr
// 	}
// 	dberr = mtd.CalcDemand()
// 	if dberr != nil {
// 		return nil, errors.New("Error in calculating demand " + dberr.Error())
// 	}
// 	return mtd, nil
// }

// func (s *Service) PropertyPartPaymentGetOutstandingDemand(ctx *models.Context, uniqueID string) (*models.PropertyPartPaymentDemand, error) {
// 	filter := new(models.PropertyPartPaymentCalcQueryFilter)

// 	filter.PropertyPartPaymentID = uniqueID
// 	omitFys, err := s.Daos.GetPayedFinancialYearsOfPropertyPartPayment(ctx, uniqueID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(omitFys) > 0 {
// 		filter.OmitFy = append(filter.OmitFy, omitFys...)
// 	}
// 	fmt.Println("omity = >", filter.OmitFy)
// 	return s.CalcShopRentDemandForParticulars(ctx, filter)
// }

// func (s *Service) UpdatePropertyPartpaymentDemandAndCollections(ctx *models.Context, PartPaymentID string) error {

// 	demand, err := s.ShopRentGetOutstandingDemand(ctx, shopID)
// 	if err != nil {
// 		return errors.New("Eror in calc demand - " + err.Error())
// 	}
// 	demand.ShopRent.Collections = models.ShopRentTotalCollection{}
// 	demand.ShopRent.PendingCollections = models.ShopRentTotalCollection{}
// 	demand.ShopRent.OutStanding = models.ShopRentTotalOutStanding{}
// 	mainpipeline, err := demand.CalcCollectionQuery()
// 	if err != nil {
// 		return errors.New("Error in Calc generating Query - " + err.Error())
// 	}
// 	payments, err := s.Daos.CalcShopRentPaymens(ctx, mainpipeline)
// 	if err != nil {
// 		return errors.New("Error in Payments - " + err.Error())
// 	}
// 	err = demand.CalcCollection(payments)
// 	if err != nil {
// 		return errors.New("Error in Calc Collection - " + err.Error())
// 	}
// 	mainpipeline2, err := demand.CalcPendingCollectionQuery()
// 	pendingpayments, err := s.Daos.CalcShopRentPendingPaymens(ctx, mainpipeline2)
// 	if err != nil {
// 		return errors.New("Error in Pending Payments - " + err.Error())
// 	}
// 	err = demand.CalcPendingCollection(pendingpayments)
// 	if err != nil {
// 		return errors.New("Error in Calc pending Collection - " + err.Error())
// 	}
// 	err = demand.CalcOutStanding()
// 	if err != nil {
// 		return errors.New("Error in Calc outstanding demand - " + err.Error())
// 	}
// 	err = s.Daos.UpdateShopRentCalc(ctx, demand)
// 	if err != nil {
// 		return errors.New("Error in updating demand - " + err.Error())
// 	}
// 	return nil

// }

//GetSingleModuleUserType : ""
func (s *Service) GetPropertyPaymentsWithPartPayments(ctx *models.Context, tnxId string) (*models.RefMOPPartPayment, error) {
	data, err := s.Daos.GetPropertyPaymentsWithPartPayments(ctx, tnxId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
