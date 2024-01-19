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

// InitiateShopRentPayment : ""
func (s *Service) InitiateShopRentMonthlyLegacyPayment(ctx *models.Context, ipmtr *models.InitiateShopRentMonthlyPaymentReq) (string, error) {
	// Start Transaction
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	tnxId := ""
	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("ipmtr", ipmtr)
		filter := new(models.ShopRentMonthlyCalcQueryFilter)
		filter.ShopRentID = ipmtr.ShopRentID
		filter.AddFy = ipmtr.Months
		demand, err := s.CalcShopRentMonthlyDemandForParticulars(ctx, filter)
		if err != nil {
			return errors.New("Error in calculating demand - " + err.Error())
		}
		if demand == nil {
			return errors.New("Demand is nil ")
		}

		pmt := new(models.ShopRentMonthlyPayments)
		pmt.TnxID = s.Shared.GetTransactionID(demand.ShopRent.UniqueID, 32)
		tnxId = pmt.TnxID
		pmt.ShopRentID = demand.ShopRent.UniqueID

		fy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return errors.New("Error in geting current financial year - " + err.Error())
		}
		pmt.FinancialYear = fy.FinancialYear

		pmt.Demand = models.ShopRentPaymentDemand{
			Current: models.ShopRentPaymentDemandSplitage{
				Tax:     demand.Demand.Current.Tax,
				Penalty: demand.Demand.Current.Penalty,
				Other:   demand.Demand.Current.Other,
				Total:   demand.Demand.Current.Total,
			},
			Arrear: models.ShopRentPaymentDemandSplitage{
				Tax:     demand.Demand.Arrear.Tax,
				Penalty: demand.Demand.Arrear.Penalty,
				Other:   demand.Demand.Arrear.Other,
				Total:   demand.Demand.Arrear.Total,
			},
			Total: models.ShopRentPaymentDemandSplitage{
				Tax:     demand.Demand.Total.Tax,
				Penalty: demand.Demand.Total.Penalty,
				Other:   demand.Demand.Total.Other,
				Total:   demand.Demand.Total.Total,
			},
		}
		pmt.Status = constants.SHOPRENTPAYMENTSTATUSINIT
		pmt.Scenario = constants.SHOPRENTPAYMENTSCENARIOLEGACY

		pmt.Created = models.CreatedV2{
			By:     ipmtr.By,
			ByType: ipmtr.ByType,
			On:     &t,
		}
		var pmtFys []models.ShopRenttMonthlyPaymentsfY
		for _, v := range demand.FY {
			var pmtFy models.ShopRenttMonthlyPaymentsfY
			pmtFy.TnxID = pmt.TnxID
			pmtFy.ShopRentID = demand.ShopRent.UniqueID
			pmtFy.FY = v
			pmtFy.Status = constants.SHOPRENTPAYMENTSTATUSINIT
			pmtFys = append(pmtFys, pmtFy)
		}

		var pmtBasic = new(models.ShopRentPaymentsBasics)
		pmtBasic.TnxID = pmt.TnxID
		pmtBasic.ShopRentID = demand.ShopRent.UniqueID
		sr, err := s.Daos.GetSingleShopRent(ctx, ipmtr.ShopRentID)
		if err != nil {
			return err
		}
		pmt.Address = demand.ShopRent.Address
		if sr != nil {
			pmtBasic.ShopRent = *sr
		}
		err = s.Daos.SaveShopRentMonthlyPayment(ctx, pmt)
		if err != nil {
			return errors.New("Error in saving mobile tower payment - " + err.Error())
		}
		err = s.Daos.SaveShopMonthlyRentPaymentFYs(ctx, pmtFys)
		if err != nil {
			return errors.New("Error in saving mobile tower payment fys- " + err.Error())
		}
		err = s.Daos.SaveShopRentPaymentBasic(ctx, pmtBasic)
		if err != nil {
			return errors.New("Error in saving mobile tower payment basics- " + err.Error())
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
	return tnxId, nil
}

func (s *Service) MakeShopRentMonthlyLegacyPayment(ctx *models.Context, mmtpr *models.MakeShopRentPaymentReq) (string, error) {
	shopRentID := ""
	// Start Transaction
	log.Println("transaction start")
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	t := time.Now()
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		mmtpr.CompletionDate = &t
		// status, dbErr := s.ShopRentPaymentStatusSelector(ctx, mmtpr)
		// if dbErr != nil {
		// 	return dbErr
		// }
		mmtpr.Status = constants.SHOPRENTPAYMENTSTATUSCOMPLETED
		mmtpr.Details.Collector.On = &t
		dbErr := s.Daos.MakeShopRentPayment(ctx, mmtpr)
		if dbErr != nil {
			return dbErr
		}

		payment, dbErr := s.Daos.GetSingleShopRentPayment(ctx, mmtpr.TnxID)
		if dbErr != nil {
			return dbErr
		}
		shopRentID = payment.ShopRentID
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
	return shopRentID, nil

}
