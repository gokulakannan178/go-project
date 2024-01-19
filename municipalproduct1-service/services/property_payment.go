package services

import (
	"errors"
	"fmt"
	"log"
	"math"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitiatePropertyPayment : ""
func (s *Service) InitiatePropertyPayment(ctx *models.Context, filter *models.PropertyDemandFilter) (string, error) {
	var transactionID string
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		demand, err := s.GetPropertyDemandCalc(ctx, filter, "")
		if err != nil {
			return err
		}
		if demand == nil {
			return errors.New("Nil Demand")
		}

		//Prepare Payment Details
		propertyDemandPayment := new(models.PropertyPaymentDemand)
		propertyDemandPayment.Property = demand.Property
		propertyDemandPayment.PercentAreaBuildup = demand.PercentAreaBuildup
		propertyDemandPayment.TaxableVacantLand = demand.TaxableVacantLand
		propertyDemandPayment.FYs = demand.FYs
		propertyDemandPayment.ServiceCharge = demand.ServiceCharge
		propertyDemandPayment.IsServiceChargeApplicable = demand.IsServiceChargeApplicable
		propertyDemandPayment.PropertyConfig = demand.PropertyConfig
		propertyDemandPayment.FYTax = demand.FYTax
		propertyDemandPayment.Tax = demand.Tax
		propertyDemandPayment.FlTax = demand.FlTax
		propertyDemandPayment.VlTax = demand.VlTax
		propertyDemandPayment.TotalTax = demand.TotalTax
		propertyDemandPayment.FormFee = demand.FormFee
		propertyDemandPayment.BoreCharge = demand.BoreCharge
		propertyDemandPayment.CompositeTax = demand.OverallPropertyDemand.Total.CompositeTax
		propertyDemandPayment.EducationChess = demand.OverallPropertyDemand.Total.Ecess
		propertyDemandPayment.PanelCh = demand.OverallPropertyDemand.Total.PanelCh
		propertyDemandPayment.Rebate = demand.OverallPropertyDemand.Total.Rebate
		propertyDemandPayment.OtherDemand = demand.OtherDemand

		propertyDemandPayment.PreviousCollection = demand.Property.PreviousCollection
		if len(demand.FYs) > 0 {
			for _, v := range demand.FYs {
				propertyDemandPayment.PenalCharge = propertyDemandPayment.PenalCharge + v.Penalty

			}
		}
		//Prepare Property Payment
		tnxID := s.Shared.GetTransactionID(filter.PropertyID, 32)
		payment := new(models.PropertyPayment)
		//27022023*01*0004
		ward, err := s.Daos.GetSingleWard(ctx, demand.Property.Address.WardCode)
		if err != nil {
			fmt.Println(err)
			ward = new(models.RefWard)
		}

		if ctx.ProductConfig.LocationID == "Bhagalpur" {
			str := fmt.Sprintf("%v%v%v_%v_%v", t.Day(), int(t.Month()), t.Year(), ward.Name, s.Daos.GetUniqueID(ctx, "recipt"))
			payment.ReciptNo = str
		} else {
			payment.ReciptNo = s.Daos.GetUniqueID(ctx, "recipt")

		}

		payment.TnxID = tnxID
		payment.PropertyID = filter.PropertyID
		payment.Demand = propertyDemandPayment
		payment.Status = constants.PROPERTYPAYMENTINITIATED
		payment.Address = demand.Property.Address
		isCurrentFYAvailable := false
		cfy, _ := s.GetCurrentFinancialYear(ctx)
		if cfy != nil {
			payment.FinancialYear = cfy.FinancialYear
			for _, v := range propertyDemandPayment.FYs {
				if cfy.FinancialYear.UniqueID == v.FinancialYear.UniqueID {
					isCurrentFYAvailable = true
					propertyDemandPayment.Current = propertyDemandPayment.Current + v.TotalTax
				} else {
					propertyDemandPayment.Arrear = propertyDemandPayment.Arrear + v.TotalTax
				}
			}

		}

		if isCurrentFYAvailable {
			propertyDemandPayment.Current = propertyDemandPayment.Current + propertyDemandPayment.FormFee + propertyDemandPayment.BoreCharge
		} else {
			propertyDemandPayment.Arrear = propertyDemandPayment.Arrear + propertyDemandPayment.FormFee + propertyDemandPayment.BoreCharge

		}
		payment.Demand.TotalTax = math.Ceil(payment.Demand.TotalTax)
		if filter.PartPayment.IS {
			payment.Type = "PartPayment"
			payment.RemainingAmount = payment.Demand.TotalTax - filter.PartPayment.Amount
		}

		//Prepare Property Basic details
		propertyPaymentDemandBasic := new(models.PropertyPaymentDemandBasic)
		propertyPaymentDemandBasic.TnxID = tnxID
		propertyPaymentDemandBasic.Property = demand.Property
		propertyPaymentDemandBasic.Status = constants.PROPERTYPAYMENTINITIATED

		floors, floorErr := s.Daos.GetFloorsOfProperty(ctx, filter.PropertyID)
		if floorErr != nil {
			log.Println("Error in geting Floors - " + floorErr.Error())
		}
		propertyPaymentDemandBasic.Floors = floors

		owners, ownerErr := s.Daos.GetOwnersOfProperty(ctx, filter.PropertyID)
		if ownerErr != nil {
			log.Println("Error in geting Floors - " + floorErr.Error())
		}
		propertyPaymentDemandBasic.Owners = owners

		if err := s.Daos.SavePropertyPaymentDemandBasic(ctx, propertyPaymentDemandBasic); err != nil {
			return errors.New("Errror in saving payment demand basics - " + err.Error())
		}
		//Prepare Fy Details
		ppdfys := []models.PropertyPaymentDemandFy{}
		for _, v := range propertyDemandPayment.FYs {
			propertyPaymentDemandFy := models.PropertyPaymentDemandFy{}
			propertyPaymentDemandFy.TnxID = tnxID
			propertyPaymentDemandFy.PropertyID = filter.PropertyID
			propertyPaymentDemandFy.FY = v
			propertyPaymentDemandFy.Status = constants.PROPERTYPAYMENTINITIATED
			propertyPaymentDemandFy.FY.OtherDemand = v.OtherDemand
			propertyPaymentDemandFy.FY.OtherDemandAdditionalPenalty = v.OtherDemandAdditionalPenalty
			ppdfys = append(ppdfys, propertyPaymentDemandFy)
		}
		if filter.PartPayment.IS {
			if err := s.AddPartPaymentToFys(ctx, filter.PartPayment.Amount, ppdfys); err != nil {
				return errors.New("Error in part paying - " + err.Error())
			}
			payment.Demand.TotalTax = 0
			payment.Demand.Current = 0
			payment.Demand.Arrear = 0
			for _, v := range ppdfys {
				if v.FY.IsCurrent {
					payment.Demand.Current = payment.Demand.Current + v.FY.TotalTax
					payment.Summary.CurrentTax = payment.Summary.CurrentTax + v.FY.VacantLandTax + v.FY.Tax
					payment.Summary.CurrentPenalty = payment.Summary.CurrentPenalty + v.FY.Penalty
					payment.Summary.CurrentRebate = payment.Summary.CurrentRebate + v.FY.Rebate
					payment.Summary.ArrearOtherDemand = payment.Summary.ArrearOtherDemand + v.FY.OtherDemand
				} else {
					payment.Demand.Arrear = payment.Demand.Arrear + v.FY.TotalTax
					payment.Summary.ArrearTax = payment.Summary.ArrearTax + v.FY.VacantLandTax + v.FY.Tax
					payment.Summary.ArrearPenalty = payment.Summary.ArrearPenalty + v.FY.Penalty
					payment.Summary.ArrearRebate = payment.Summary.ArrearRebate + v.FY.Rebate
					payment.Summary.CurrentOtherDemand = payment.Summary.CurrentOtherDemand + v.FY.OtherDemand
				}
				payment.Demand.TotalTax = payment.Demand.TotalTax + v.FY.TotalTax
			}
			payment.Demand.TotalTax = payment.Demand.TotalTax + payment.Demand.FormFee
		}
		for _, v := range ppdfys {
			if v.FY.IsCurrent {
				payment.Summary.CurrentTax = payment.Summary.CurrentTax + v.FY.VacantLandTax + v.FY.Tax
				payment.Summary.CurrentPenalty = payment.Summary.CurrentPenalty + v.FY.Penalty
				payment.Summary.CurrentRebate = payment.Summary.CurrentRebate + v.FY.Rebate
				payment.Summary.ArrearOtherDemand = payment.Summary.ArrearOtherDemand + v.FY.OtherDemand
				payment.Summary.TotalCurrent = payment.Summary.TotalCurrent + v.FY.VacantLandTax + v.FY.Tax + v.FY.Penalty + v.FY.OtherDemand - v.FY.Rebate

			} else {
				payment.Summary.ArrearTax = payment.Summary.ArrearTax + v.FY.VacantLandTax + v.FY.Tax
				payment.Summary.ArrearPenalty = payment.Summary.ArrearPenalty + v.FY.Penalty
				payment.Summary.ArrearRebate = payment.Summary.ArrearRebate + v.FY.Rebate
				payment.Summary.CurrentOtherDemand = payment.Summary.CurrentOtherDemand + v.FY.OtherDemand
				payment.Summary.TotalArrear = payment.Summary.TotalArrear + v.FY.VacantLandTax + v.FY.Tax + v.FY.Penalty + v.FY.OtherDemand - v.FY.Rebate

			}
			//payment.Demand.TotalTax = payment.Demand.TotalTax + v.FY.TotalTax
		}
		if err := s.Daos.SaveManyPropertyPaymentDemandFy(ctx, ppdfys); err != nil {
			return errors.New("Errror in saving payment demand fys - " + err.Error())
		}
		payment.Summary.FormFee = payment.Demand.FormFee
		payment.Summary.BoreCharge = payment.Demand.BoreCharge
		if len(ppdfys) > 0 {
			payment.Summary.FromFy = ppdfys[0].FY.Name
			payment.Summary.ToFy = ppdfys[len(ppdfys)-1].FY.Name
		}

		if err := s.Daos.SavePropertyPayment(ctx, payment); err != nil {
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

// AddPartPaymentToFys : ""
func (s *Service) AddPartPaymentToFys(ctx *models.Context, paidAmount float64, ppydf []models.PropertyPaymentDemandFy) error {
	fmt.Println("Calculating AddPartPaymentToFys")
	for k := range ppydf {
		if paidAmount <= 0 {
			return errors.New("remove unwanted financial years")
		}
		if paidAmount > ppydf[k].FY.TotalTax {
			ppydf[k].FY.PaidTotalTax = ppydf[k].FY.TotalTax - ppydf[k].FY.OtherDemand
		} else {
			ppydf[k].FY.PaidTotalTax = paidAmount - ppydf[k].FY.OtherDemand
		}

		if ppydf[k].FY.IsCurrent {
			ppydf[k].FY.PaidTax = ((ppydf[k].FY.PaidTotalTax / ((ppydf[k].FY.PenaltyRate * ppydf[k].FY.PenaltyMonths) + 100)) * 100)
			ppydf[k].FY.PaidPenalty = ppydf[k].FY.PaidTotalTax - ppydf[k].FY.PaidTax
		} else {
			ppydf[k].FY.PaidTax = ((ppydf[k].FY.PaidTotalTax / (ppydf[k].FY.PenaltyRate + 100)) * 100)
			ppydf[k].FY.PaidPenalty = ppydf[k].FY.PaidTotalTax - ppydf[k].FY.PaidTax
		}
		fmt.Println("ppydf[k].FY.PaidTotalTax", ppydf[k].FY.PaidTotalTax)
		fmt.Println("ppydf[k].FY.PaidTax", ppydf[k].FY.PaidTax)
		fmt.Println("ppydf[k].FY.PaidPenalty", ppydf[k].FY.PaidPenalty)

		ppydf[k].FY.ActualTotalTax = ppydf[k].FY.TotalTax
		ppydf[k].FY.ActualFlTax = ppydf[k].FY.Tax
		ppydf[k].FY.ActualVlTax = ppydf[k].FY.VacantLandTax

		ppydf[k].FY.ActualTax = ppydf[k].FY.Tax
		ppydf[k].FY.ActualPenalty = ppydf[k].FY.Penalty

		//need to calculate vl tax and fl tax
		//PaidTax/(Tax+Vl /100) = %?
		//(PaidTax/Tax+Vl)*100 = %?
		ppydf[k].FY.PaidPartPaymentPercentage = (ppydf[k].FY.PaidTax / (ppydf[k].FY.Tax + ppydf[k].FY.VacantLandTax)) * 100
		ppydf[k].FY.TotalTax = ppydf[k].FY.PaidTotalTax + ppydf[k].FY.OtherDemand
		ppydf[k].FY.Tax = (ppydf[k].FY.Tax * ppydf[k].FY.PaidPartPaymentPercentage) / 100
		ppydf[k].FY.VacantLandTax = (ppydf[k].FY.VacantLandTax * ppydf[k].FY.PaidPartPaymentPercentage) / 100
		ppydf[k].FY.Penalty = ppydf[k].FY.PaidPenalty
		paidAmount = paidAmount - ppydf[k].FY.PaidTotalTax
	}
	return nil
}

// GetSinglePropertyPaymentTxtID : ""
func (s *Service) GetSinglePropertyPaymentTxtID(ctx *models.Context, id string) (*models.RefPropertyPayment, error) {
	refPropertyPayment := new(models.RefPropertyPayment)
	payment, err := s.Daos.GetSinglePropertyPaymentWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property payment - " + err.Error())
	}
	propertyDemandBasic, err := s.Daos.GetSinglePropertyPaymentDemandBasicWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property payment damand basic")
	}
	propertyDemandFys, err := s.Daos.GetPropertyPaymentDemandFycWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property payment damand fys")
	}
	ppFilter := new(models.PropertyPartPaymentFilter)
	ppFilter.TnxID = []string{id}
	ppFilter.Status = []string{constants.PROPERTYPAYMENTCOMPLETED}
	//FilterPropertyPartPayment(ctx *models.Context, filter *models.PropertyPartPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyPartPayment, error) {
	refPartPayments, _ := s.Daos.FilterPropertyPartPayment(ctx, ppFilter, nil)
	if refPartPayments != nil {
		refPropertyPayment.Ref.PartPayments = refPartPayments
		for _, v := range refPartPayments {
			refPropertyPayment.Ref.PartAmountCollected = refPropertyPayment.Ref.PartAmountCollected + v.Details.Amount
		}
	}
	refPropertyPayment.PropertyPayment = *payment
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

// GetSinglePropertyPaymentTxtID : ""
func (s *Service) GetSinglePropertyPaymentReceiptNo(ctx *models.Context, id string) (*models.RefPropertyPayment, error) {
	refPropertyPayment := new(models.RefPropertyPayment)
	payment, err := s.Daos.GetSinglePropertyPaymentReceiptNo(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property payment - " + err.Error())
	}
	propertyDemandBasic, err := s.Daos.GetSinglePropertyPaymentDemandBasicWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property payment damand basic")
	}
	propertyDemandFys, err := s.Daos.GetPropertyPaymentDemandFycWithTxtID(ctx, id)
	if err != nil {
		return nil, errors.New("Error in geting property payment damand fys")
	}
	ppFilter := new(models.PropertyPartPaymentFilter)
	ppFilter.TnxID = []string{id}
	ppFilter.Status = []string{constants.PROPERTYPAYMENTCOMPLETED}
	//FilterPropertyPartPayment(ctx *models.Context, filter *models.PropertyPartPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyPartPayment, error) {
	refPartPayments, _ := s.Daos.FilterPropertyPartPayment(ctx, ppFilter, nil)
	if refPartPayments != nil {
		refPropertyPayment.Ref.PartPayments = refPartPayments
		for _, v := range refPartPayments {
			refPropertyPayment.Ref.PartAmountCollected = refPropertyPayment.Ref.PartAmountCollected + v.Details.Amount
		}
	}
	refPropertyPayment.PropertyPayment = *payment
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

// PropertyMakePayment : ""
func (s *Service) PropertyMakePayment(ctx *models.Context, payment *models.PropertyMakePayment) (string, error) {
	var propertyId string
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		/*
				r := NewRequestPdf("")
				docStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOCD)
				templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
				//html template path
				templatePath := templatePathStart + "Receipt_page.html"
			t := time.Now()
			filename := fmt.Sprintf("Recipt_%v%v%v_%v%v%v", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
			//path for download pdf
			outputPath := docStart + "recipts/" + filename + ".pdf"
			payment.ReciptURL = "/" + outputPath
		*/
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
		py, err := s.GetSinglePropertyPaymentTxtID(ctx, payment.TnxID)
		if err != nil {
			fmt.Println("problem in geting property payment for transaction id - " + err.Error())
		}
		// if py.Details != nil {
		// 	py.Details.AmountPaid = py.Details.AmountPaid + payment.Details.IncomingPayment
		// }

		if err := s.Daos.CompletePropertyPaymentWithTxtID(ctx, payment, false); err != nil {
			return errors.New("Error in updating in - Payment" + err.Error())
		}
		if err := s.Daos.CompleteSinglePropertyPaymentDemandBasicWithTxtID(ctx, payment); err != nil {
			return errors.New("Error in updating in - basic" + err.Error())
		}
		if err := s.Daos.CompletePropertyPaymentDemandFycWithTxtID(ctx, payment); err != nil {
			return errors.New("Error in updating in - Fys" + err.Error())
		}

		propertyId = py.PropertyID
		if py != nil {
			if err := s.Daos.UpdateBoringStatusToProperty(ctx, py.PropertyID, true); err != nil {
				return errors.New("Error in updating in - Boring Charge" + err.Error())
			}
			fmt.Println("Boring Charge Updated")
			if err := s.Daos.UpdateFormFeeStatusToProperty(ctx, py.PropertyID, true); err != nil {
				return errors.New("Error in updating in - Form Fee" + err.Error())
			}
			fmt.Println("Form Fee Update")
			//
			if !py.Demand.Property.PreviousCollection.IsCalculated {
				if err := s.Daos.UpdatePayedPropertyPreviousYrCollection(ctx, py.PropertyID); err != nil {
					return errors.New("Error in previous year collection" + err.Error())
				}
				fmt.Println("previous years collection updated ")
			}

		} else {
			fmt.Println("property not found")
		}
		/*
			if err := r.ParseTemplate(templatePath, py); err == nil {
				ok, _ := r.GeneratePDF(outputPath)
				fmt.Println(ok, "pdf generated successfully")
			} else {
				fmt.Println(err)
			}*/

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
				fmt.Println("Error in sending SMS - " + err.Error())
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

// PropertyMakePaymentV2 : ""
func (s *Service) PropertyMakePaymentV2(ctx *models.Context, payment *models.PropertyMakePayment) (string, error) {
	var propertyId string
	if err := ctx.Session.StartTransaction(); err != nil {
		return "", err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		py, err := s.GetSinglePropertyPaymentTxtID(ctx, payment.TnxID)
		if err != nil {
			fmt.Println("problem in geting property payment for transaction id - " + err.Error())
		}

		if err := s.Daos.RecordPartPayment(ctx, payment); err != nil {
			return errors.New("Error in inserting in - Part Payment" + err.Error())
		}
		if err := s.Daos.CompletePropertyPaymentWithTxtID(ctx, payment, false); err != nil {
			return errors.New("Error in updating in - Payment" + err.Error())
		}
		if err := s.Daos.CompleteSinglePropertyPaymentDemandBasicWithTxtID(ctx, payment); err != nil {
			return errors.New("Error in updating in - basic" + err.Error())
		}
		if err := s.Daos.CompletePropertyPaymentDemandFycWithTxtID(ctx, payment); err != nil {
			return errors.New("Error in updating in - Fys" + err.Error())
		}

		propertyId = py.PropertyID
		if py != nil {
			if err := s.Daos.UpdateBoringStatusToProperty(ctx, py.PropertyID, true); err != nil {
				return errors.New("Error in updating in - Boring Charge" + err.Error())
			}
			fmt.Println("Boring Charge Updated")
			if err := s.Daos.UpdateFormFeeStatusToProperty(ctx, py.PropertyID, true); err != nil {
				return errors.New("Error in updating in - Form Fee" + err.Error())
			}
			fmt.Println("Form Fee Update")
			//
			if !py.Demand.Property.PreviousCollection.IsCalculated {
				if err := s.Daos.UpdatePayedPropertyPreviousYrCollection(ctx, py.PropertyID); err != nil {
					return errors.New("Error in previous year collection" + err.Error())
				}
				fmt.Println("previous years collection updated ")
			}

		} else {
			fmt.Println("property not found")
		}
		/*
			if err := r.ParseTemplate(templatePath, py); err == nil {
				ok, _ := r.GeneratePDF(outputPath)
				fmt.Println(ok, "pdf generated successfully")
			} else {
				fmt.Println(err)
			}*/

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

func (s *Service) PropertyUpdateCollection(ctx *models.Context, propertyID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if err := s.Daos.PropertyUpdateCollection(ctx, propertyID); err != nil {
			return errors.New("Not update collection - " + err.Error())
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

// GetAllPaymentsForProperty : ""
func (s *Service) GetAllPaymentsForProperty(ctx *models.Context, id string) ([]models.RefPropertyPayment, error) {
	data, err := s.Daos.GetAllPaymentsForProperty(ctx, id)
	return data, err
}

// DashboardTotalCollection : ""
func (s *Service) DashboardTotalCollection(ctx *models.Context, filter *models.DashboardTotalCollectionFilter) (*models.DashboardTotalCollection, error) {
	refData, err := s.Daos.DashboardTotalCollection(ctx, filter)
	if err != nil {
		return nil, errors.New("Error in geting queried data - " + err.Error())
	}
	data := new(models.DashboardTotalCollection)
	if len(refData) > 0 {
		for _, v := range refData {
			data.Current = data.Current + v.CurrentYear.FY.TotalTax + v.CurrentYear.FY.ServiceCharge
			for _, v1 := range v.ArriearYears {
				data.Arriear = data.Arriear + +v1.FY.TotalTax + v1.FY.ServiceCharge
			}
		}
		data.Total = data.Current + data.Arriear
	}
	return data, nil
}

// FilterPropertyPayment : ""
func (s *Service) FilterPropertyPayment(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyPayment, error) {

	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterPropertyPayment(ctx, filter, pagination)
}

// FilterPropertyPaymentExcel : ""
func (s *Service) FilterPropertyPaymentExcel(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Payments"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "K3")
	excel.MergeCell(sheet1, "C4", "K5")
	excel.MergeCell(sheet1, "A6", "K6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Payment List")
	rowNo++
	rowNo++

	//
	reportFromMsg := "Report"
	t := time.Now()
	toDate := t.Format("02-January-2006")
	if filter != nil {
		if filter.DateRange != nil {
			fmt.Println(filter.DateRange.From, filter.DateRange.To)
			if filter.DateRange.From != nil && filter.DateRange.To == nil {
				reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + toDate
			}
			if filter.DateRange.From != nil && filter.DateRange.To != nil {
				reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
			}
			if filter.DateRange.From == nil && filter.DateRange.To == nil {
				fmt.Println("everything is nil")
			}

		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++
	//
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "District")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Payee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Mode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Payment Made At")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Collected By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Rejected By")

	fmt.Println("'res length==>'", len(res))
	var totalAmount float64
	for i, v := range res {
		totalAmount = totalAmount + func() float64 {
			if v.Details != nil {
				return v.Details.Amount
			}
			return 0
		}()

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Basic.Property.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.Ref.Address.District != nil {
				return v.Ref.Address.District.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Details != nil {
				return v.Details.PayeeName
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() interface{} {
			if v.Details != nil {
				return v.Details.Amount
			}
			return "NA"
		}())
		if v.CompletionDate != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.CompletionDate.Format("2006-01-02"))
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Details != nil {
				return v.Details.MOP.Mode
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MadeAt != nil {
					return v.Details.MadeAt.At
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Ref.Collector.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), func() string {
			if v.Ref.RejectedBy.Name != "" {
				return v.Ref.RejectedBy.Name
			}
			return "NA"
		}())
	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%.0f", totalAmount))

	return excel, nil
}

// FilterPropertyPaymentPDF : ""
func (s *Service) FilterPropertyPaymentPDF(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]byte, error) {

	data, err := s.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	r := NewRequestPdf("")
	//r.Orentation = "Landscape"
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "PaymentFilter.html"
	err = r.ParseTemplate(templatePath, data)
	if err != nil {
		return nil, err
	}
	ok, file, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	fmt.Println(ok, "pdf generated successfully")

	return file, nil
}

// FilterPropertyPaymentPDFV2 : ""
func (s *Service) FilterPropertyPaymentPDFV2(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["Payment"] = data
	m2["currentDate"] = time.Now()
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in geting product config" + err.Error())
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	pdfdata.Config = productConfig.ProductConfiguration

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "allpropertypaymentfilter.html"
	err = r.ParseTemplate(templatePath, pdfdata)
	if err != nil {
		return nil, err
	}
	ok, file, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	fmt.Println(ok, "pdf generated successfully")

	return file, nil
}

// FilterPropertyPaymentPDFV3 : ""
func (s *Service) FilterPropertyPaymentPDFV3(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	// fmt.Println("data =======>", data)

	var fyNameFrom, fyNameTo string
	if len(data) > 0 {
		for _, v := range data {
			fyNameFrom = v.Fys[0].FY.Name
			fyNameTo = v.Fys[len(v.Fys)-1].FY.Name
		}
	}

	fmt.Println("fyNameFrom=====>", fyNameFrom, "to", fyNameTo)
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})

	m["Payment"] = data
	m2["currentDate"] = time.Now()
	m3["financialYears"] = fyNameFrom + "to" + fyNameTo
	fmt.Println("from - to =======>", fyNameFrom+"to"+fyNameTo)

	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in geting product config" + err.Error())
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	pdfdata.Config = productConfig.ProductConfiguration

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "allpropertypaymentv2.html"
	err = r.ParseTemplate(templatePath, pdfdata)
	if err != nil {
		return nil, err
	}
	ok, file, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	fmt.Println(ok, "pdf generated successfully")

	return file, nil
}

// CheckBounceReport : ""
func (s *Service) ChequeBounceReport(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Cheque Bounce"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "O3")
		excel.MergeCell(sheet1, "A4", "O5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "O3")
		excel.MergeCell(sheet1, "C4", "O5")
	}

	excel.MergeCell(sheet1, "A6", "O6")
	excel.MergeCell(sheet1, "A7", "O7")

	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style2)

	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Cheque Bounce Report")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Cheque Bounce Report")
	}
	rowNo++
	rowNo++

	reportFromMsg := "Report"

	if filter != nil {
		if filter.DateRange != nil {
			fmt.Println(filter.DateRange.From, filter.DateRange.To)
			if filter.DateRange.From != nil && filter.DateRange.To == nil {
				reportFromMsg = reportFromMsg + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
			}
			if filter.DateRange.From != nil && filter.DateRange.To != nil {
				reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
			}
			if filter.DateRange.From == nil && filter.DateRange.To == nil {
				fmt.Println("everything is nil")
			}

		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Cheque bounce report from_____to________")

	rowNo++

	t2 := time.Now()
	toDate := t2.Format("02-January-2006")
	reportFromMsg2 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style1)

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Property ID")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Transaction Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Mode")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Cheque No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Cheque Details")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Bank Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Branch Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "From")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "To")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "User Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Remarks")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Bounce Date")

	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "District")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Payee")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Payment Made At")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Collected By")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Cheque No")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Cheque Details")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Cheque Date")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Remarks")

	var totalAmount float64

	for i, v := range res {
		totalAmount = totalAmount + v.Demand.TotalTax

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.PropertyPayment.PropertyID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if len(v.Basic.Owners) > 0 {
				return v.Basic.Owners[0].Name
			} else {
				return "NA"
			}
		}())
		if v.CompletionDate != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.CompletionDate.Format("2006-01-02"))

		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "NA")

		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Cheque")
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.No
				}
				return "NA"
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.Bank
				}
				return "NA"

			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.Branch
				}
				return "NA"
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Summary.FromFy)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Summary.ToFy)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Demand.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Ref.Collector.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.BouncedInfo.Remark)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), func() string {
			if v.BouncedInfo.BouncedDate != nil {
				return v.BouncedInfo.BouncedDate.Format("2006-01-02")
			}
			return "NA"

		}())

		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
		// 	if v.Details != nil {
		// 		return v.Details.PayeeName
		// 	}
		// 	return "NA"
		// }())

		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
		// 	if v.Details != nil {
		// 		return v.Details.MadeAt.At
		// 	}
		// 	return "NA"
		// }())
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), func() string {
		// 	if v.Details != nil {
		// 		return v.Details.Collector.ID
		// 	}
		// 	return "NA"
		// }())
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), func() string {
		// 	if v.Details != nil {
		// 		if v.Details.MOP.Cheque != nil {
		// 			return v.Details.MOP.Cheque.No
		// 		}
		// 		return "NA"
		// 	}
		// 	return "NA"
		// }())
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), func() string {
		// 	if v.Details != nil {
		// 		if v.Details.MOP.Cheque != nil {

		// 			if v.Details.MOP.Cheque.Date != nil {
		// 				return v.Details.MOP.Cheque.Date.Format("2006-01-02")
		// 			}
		// 			return "NA"

		// 		}
		// 		return "NA"
		// 	}
		// 	return "NA"
		// }())

		// if v.Ref.Address.District != nil {
		// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Address.DistrictCode)
		// } else {
		// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "NA")
		// }
	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "N", rowNo), fmt.Sprintf("%v%v", "N", rowNo), style3)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "O", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style3)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%.2f", totalAmount))

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v", totalConsumer))

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%.2f", totalAmount))
	return excel, nil

}

// ChequeBounceReport:""
func (s *Service) ChequeBounceReportPdf(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]byte, error) {
	res, _ := s.FilterPropertyPayment(ctx, filter, pagination)
	r := NewRequestPdf("")
	//html template path
	templatePath := "templates/bouncereport.html"
	//path for download pdf
	//outputPath := "storage/recipt4.pdf"
	err := r.ParseTemplate(templatePath, res)
	if err != nil {
		return nil, err
	}
	_, data, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// PendingChequeReport : ""
func (s *Service) PendingChequeReportPdf(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]byte, error) {
	res, _ := s.FilterPropertyPayment(ctx, filter, pagination)
	r := NewRequestPdf("")
	//html template path
	templatePath := "templates/bouncereport.html"
	//path for download pdf
	//outputPath := "storage/recipt4.pdf"
	err := r.ParseTemplate(templatePath, res)
	if err != nil {
		return nil, err
	}
	_, data, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// VerifyChequeReport : ""
func (s *Service) VerifyChequeReportPdf(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]byte, error) {
	res, _ := s.FilterPropertyPayment(ctx, filter, pagination)
	r := NewRequestPdf("")
	//html template path
	templatePath := "templates/bouncereport.html"
	//path for download pdf
	//outputPath := "storage/recipt4.pdf"
	err := r.ParseTemplate(templatePath, res)
	if err != nil {
		return nil, err
	}
	_, data, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// PendingChequeReport : ""
func (s *Service) PendingChequeReport(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Cheque Bounce"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)

	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "M3")
		excel.MergeCell(sheet1, "A4", "M5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "M3")
		excel.MergeCell(sheet1, "C4", "M5")
	}
	excel.MergeCell(sheet1, "A6", "M6")

	style, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Pending Cheque Report")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Pending Cheque Report")
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "District")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Payee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Payment Made At")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Collected By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Cheque No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Cheque Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Bank Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Branch Name")

	for i, v := range res {
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.PropertyPayment.PropertyID)
		if v.Ref.Address.District != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Address.DistrictCode)

		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "NA")

		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Details != nil {
				return v.Details.PayeeName
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Demand.TotalTax)
		if v.CompletionDate != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.CompletionDate.Format("2006-01-02"))

		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "NA")

		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Details != nil {
				return v.Details.MadeAt.At
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if v.Details != nil {
				return v.Details.Collector.ID
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.No
				}
				return "NA"
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {

					if v.Details.MOP.Cheque.Date != nil {
						return v.Details.MOP.Cheque.Date.Format("2006-01-02")
					}
					return "NA"

				}
				return "NA"
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.Bank
				}
				return "NA"

			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.Branch
				}
				return "NA"
			}
			return "NA"
		}())
	}
	return excel, nil

}

// VerifyChequeReport : ""
func (s *Service) VerifyChequeReport(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))
	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Cheque Bounce"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "M3")
		excel.MergeCell(sheet1, "A4", "M5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "M3")
		excel.MergeCell(sheet1, "C4", "M5")
	}
	excel.MergeCell(sheet1, "A6", "M6")
	style, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Verified Cheque Report")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Verified Cheque Report")
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "District")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Payee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Payment Made At")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Collected By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Cheque No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Cheque Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Bank Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Branch Name")

	for i, v := range res {
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.PropertyPayment.PropertyID)
		if v.Ref.Address.District != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Address.DistrictCode)

		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "NA")

		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Details != nil {
				return v.Details.PayeeName
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Demand.TotalTax)
		if v.CompletionDate != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.CompletionDate.Format("2006-01-02"))

		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "NA")

		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Details != nil {
				return v.Details.MadeAt.At
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if v.Details != nil {
				return v.Details.Collector.ID
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.No
				}
				return "NA"
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {

					if v.Details.MOP.Cheque.Date != nil {
						return v.Details.MOP.Cheque.Date.Format("2006-01-02")
					}
					return "NA"

				}
				return "NA"
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.Bank
				}
				return "NA"

			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.Branch
				}
				return "NA"
			}
			return "NA"
		}())
	}
	return excel, nil

}

// CounterReport : ""
func (s *Service) CounterReport(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Counter Report"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "Q3")
		excel.MergeCell(sheet1, "A4", "Q5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "Q3")
		excel.MergeCell(sheet1, "C4", "Q5")
	}
	excel.MergeCell(sheet1, "A6", "Q6")
	excel.MergeCell(sheet1, "A7", "Q7")
	excel.MergeCell(sheet1, "A8", "Q8")
	excel.MergeCell(sheet1, "A9", "Q9")
	excel.MergeCell(sheet1, "A10", "Q10")

	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style, err := excel.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++
	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Counter Report")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Counter Report")
	}
	rowNo++
	rowNo++

	reportFromMsg := "Report"

	if filter != nil {
		if filter.DateRange != nil {
			fmt.Println(filter.DateRange.From, filter.DateRange.To)
			if filter.DateRange.From != nil && filter.DateRange.To == nil {
				reportFromMsg = reportFromMsg + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
			}
			if filter.DateRange.From != nil && filter.DateRange.To != nil {
				reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
			}
			if filter.DateRange.From == nil && filter.DateRange.To == nil {
				fmt.Println("everything is nil")
			}

		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg2 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	var totalAmountRow, totalConsumerRow, totalPaymentsRow int
	fmt.Println(totalPaymentsRow)
	// Total Consumer

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalConsumerRow = rowNo
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalPaymentsRow = rowNo
	rowNo++

	// Total Collection
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalAmountRow = rowNo
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "Q", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Application No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Transaction Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Receipt No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Payment Mode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "From Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Upto Period")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Cheque Code")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Bank Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Branch Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), "Total Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Tax Collector Name")

	var totalAmount float64
	var totalNoOfPayments int
	var totalPayments int
	totalPayments = len(res)
	totalConsumerCalc := make(map[string]int)
	for i, v := range res {
		totalPayments++
		totalConsumerCalc[v.Basic.Property.UniqueID] = totalConsumerCalc[v.Basic.Property.UniqueID] + 1
		totalAmount = totalAmount + func() float64 {
			if v.Details != nil {
				return v.Details.Amount
			}
			return 0
		}()
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.PropertyID)

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Basic.Property.ApplicationNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if len(v.Basic.Owners) > 0 {
				return v.Basic.Owners[0].Name
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if len(v.Basic.Owners) > 0 {
				return v.Basic.Owners[0].Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Address.AL1+" "+v.Address.Al2)

		if v.CompletionDate != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.CompletionDate.Format("2006-01-02"))

		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "NA")

		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.ReciptNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if v.Details != nil {
				switch v.Details.MOP.Mode {
				case constants.MOPCASH:
					return "Cash"
				case constants.MOPCHEQUE:
					return "Cheque"
				case constants.MOPDD:
					return "DD"
				case constants.MOPNETBANKING:
					return "Online"
				default:
					return "Invalid"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), func() string {

			fmt.Println(len(v.Fys))
			if len(v.Fys) > 0 {
				return v.Fys[0].FY.Name
			}

			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), func() string {

			if len(v.Fys) > 0 {
				if len(v.Fys) == 1 {
					return v.Fys[0].FY.Name

				}
				if len(v.Fys) > 1 {
					return v.Fys[len(v.Fys)-1].FY.Name

				}
			}

			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Cheque != nil {
					return v.Details.MOP.Cheque.No
				}
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), func() string {
			if v.Details != nil {
				switch v.Details.MOP.Mode {
				case constants.MOPCASH:
					return "NA"
				case constants.MOPCHEQUE:
					return func() string {
						if v.Details.MOP.Cheque != nil {
							return v.Details.MOP.Cheque.Bank
						}
						return "NA"
					}()
				case constants.MOPDD:
					return func() string {
						if v.Details.MOP.DD != nil {
							return v.Details.MOP.DD.Bank
						}
						return "NA"
					}()
				case constants.MOPNETBANKING:
					return func() string {
						if v.Details.MOP.CardRNet != nil {
							return v.Details.MOP.CardRNet.Bank
						}
						return "NA"
					}()
				default:
					return "Invalid"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), func() string {
			if v.Details != nil {
				switch v.Details.MOP.Mode {
				case constants.MOPCASH:
					return "NA"
				case constants.MOPCHEQUE:
					return func() string {
						if v.Details.MOP.Cheque != nil {
							return v.Details.MOP.Cheque.Branch
						}
						return "NA"
					}()
				case constants.MOPDD:
					return func() string {
						if v.Details.MOP.DD != nil {
							return v.Details.MOP.DD.Branch
						}
						return "NA"
					}()
				case constants.MOPNETBANKING:
					return func() string {
						if v.Details.MOP.CardRNet != nil {
							return v.Details.MOP.CardRNet.Branch
						}
						return "NA"
					}()
				default:
					return "Invalid"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), func() string {
			if v.Details != nil {
				return fmt.Sprintf("%v", v.Details.Amount)
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), v.Ref.Collector.Name)

	}
	fmt.Println("total consumner length ======>", len(totalConsumerCalc))

	for _, v := range totalConsumerCalc {
		totalNoOfPayments = totalNoOfPayments + v
	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "Q", rowNo), style4)

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), fmt.Sprintf("%.2f", totalAmount))

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalConsumerRow), fmt.Sprintf("Total Holdings - %v", len(totalConsumerCalc)))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalPaymentsRow), fmt.Sprintf("Total No of Payments - %v", totalNoOfPayments))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalAmountRow), fmt.Sprintf("Total Collection - %.2f", totalAmount))

	return excel, nil

}

// GenerateReciptForAPayment : ""
func (s *Service) GenerateReciptForAPayment(ctx *models.Context, id string) (*models.RefPropertyPayment, error) {
	payment, _ := s.GetSinglePropertyPaymentTxtID(ctx, id)
	r := NewRequestPdf("")
	//html template path
	templatePath := "templates/Receipt_page.html"
	//path for download pdf
	outputPath := "storage/recipt4.pdf"
	if err := r.ParseTemplate(templatePath, payment); err == nil {
		ok, _ := r.GeneratePDF(outputPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
	return payment, nil
}

// VerifyPayment : ""
func (s *Service) VerifyPayment(ctx *models.Context, vp *models.VerifyPayment) (string, error) {
	t := time.Now()
	vp.ActionDate = &t
	if vp.Date == nil {
		vp.Date = &t
	}
	err := s.Daos.VerifyPayment(ctx, vp)
	if err != nil {
		return "", err
	}
	propertypayment, err := s.Daos.GetSinglePropertyPaymentWithTxtID(ctx, vp.TnxID)
	if err != nil {
		return "", err
	}
	return propertypayment.PropertyID, err
}

// NotVerifiedPayment : ""
func (s *Service) NotVerifiedPayment(ctx *models.Context, vp *models.NotVerifiedPayment) error {
	t := time.Now()
	vp.ActionDate = &t
	if vp.Date == nil {
		vp.Date = &t
	}
	err := s.Daos.NotVerifiedPayment(ctx, vp)
	return err
}

// BouncePayment : ""
func (s *Service) BouncePayment(ctx *models.Context, bp *models.BouncePayment) (string, error) {
	t := time.Now()
	bp.ActionDate = &t
	if bp.Date == nil {
		bp.Date = &t
	}
	err := s.Daos.BouncePayment(ctx, bp)
	if err != nil {
		return "", err
	}
	propertypayment, err := s.Daos.GetSinglePropertyPaymentWithTxtID(ctx, bp.TnxID)
	if err != nil {
		return "", err
	}
	return propertypayment.PropertyID, err
}

// RejectPayment : ""
func (s *Service) RejectPayment(ctx *models.Context, rp *models.RejectPayment) (string, error) {
	t := time.Now()
	rp.ActionDate = &t
	if rp.Date == nil {
		rp.Date = &t
	}
	err := s.Daos.RejectPayment(ctx, rp)
	if err != nil {
		return "", err
	}
	propertypayment, err := s.Daos.GetSinglePropertyPaymentWithTxtID(ctx, rp.TnxID)
	if err != nil {
		return "", err
	}
	return propertypayment.PropertyID, err
}

// RejectPaymentByReceiptNo
func (s *Service) RejectPaymentByReceiptNo(ctx *models.Context, rp *models.RejectPayment) (string, error) {
	t := time.Now()
	rp.ActionDate = &t
	if rp.Date == nil {
		rp.Date = &t
	}
	propertypayment, err := s.Daos.GetSinglePropertyPaymentReceiptNo(ctx, rp.ReceiptNo)
	if err != nil {
		return "", err
	}
	if propertypayment == nil {
		return "", errors.New("Property Payment nil")
	}
	rp.TnxID = propertypayment.TnxID
	err = s.Daos.RejectPaymentByReceiptNo(ctx, rp)
	if err != nil {
		return "", err
	}

	return propertypayment.PropertyID, err
}

// DateRangeWisePropertyPaymentReport : ""
func (s *Service) DateRangeWisePropertyPaymentReport(ctx *models.Context, filter *models.DateWisePropertyPaymentReportFilter) (*models.RefDateWisePropertyPaymentReport, error) {
	res, err := s.Daos.DateRangeWisePropertyPaymentReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	if ctx.ProductConfig.LocationID == "Munger" {
		if res != nil {
			// res.Report.Year.TotalCollection = res.Report.Year.TotalCollection - 3817620
			res.Report.Year.TotalCollection = 28794492 + res.Report.Month.TotalCollection
			res.Report.Year.PropertyCount = 6582 + res.Report.Month.PropertyCount
			res.Report.Overall.CurrentCollection = res.Report.Overall.CurrentCollection - 3817620
			res.Report.Overall.TotalCollection = res.Report.Overall.TotalCollection - 3817620
		}
	}
	if ctx.ProductConfig.LocationID == "Munger" {
		if res != nil {
			// res.Report.Year.TotalCollection = 28953524
			// res.Report.Year.PropertyCount = 6464
			// res.Report.Overall.CurrentCollection = 20367635.81085408
			// res.Report.Overall.ArrearCollection = 28326726
			// res.Report.Overall.TotalCollection = 48694361
		}
	}
	if ctx.ProductConfig.LocationID != "Munger" {
		err = s.Daos.UpdateOverallDashBoard(ctx, res.Report.Overall.ArrearCollection, res.Report.Overall.CurrentCollection, res.Report.Overall.TotalCollection, res.Report.Overall.PropertyCount, res)
		if err != nil {
			return nil, errors.New("error in updating the overall dashboard collection" + err.Error())
		}
	}
	return res, nil

}

func (s *Service) DateWisePropertyPaymentReport(ctx *models.Context, filter *models.DateWisePropertyPaymentReportFilter) ([]models.DateWisePropertyPaymentReport, error) {
	day := filter.Date.Format("2006-01-02")
	fmt.Println("date===========>", day)
	res, err := s.Daos.DateWisePropertyPaymentReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	if res != nil {
		for _, v := range res {
			v.Date = day

		}

	}
	return res, nil

}
func (s *Service) DayWisePropertyPaymentExcel(ctx *models.Context, filter *models.DateWisePropertyPaymentReportFilter) (*excelize.File, error) {
	data, err := s.DateWisePropertyPaymentReport(ctx, filter)
	if err != nil {
		return nil, err
	}

	excel := excelize.NewFile()
	sheet1 := " Day Wise Collection Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "j1")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "No.Of.Holdings")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Arrer Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "ArrearPenalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "CurrentAmount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "CurrentPenalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "RebateAmount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "AdvanceAmount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "TotalAmount")
	rowNo++
	var totalAmount float64

	for i, v := range data {

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Date)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.PropertyCount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.ArrearCollection)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.ArrearPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.CurrentCollection)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.CurrentPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.RebateAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.AdvanceAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.TotalCollection)
		totalAmount = totalAmount + v.TotalCollection
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

func (s *Service) CollectedPropertyPayment(ctx *models.Context, crr *models.CollectionReceivedRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		for _, v := range crr.TnxIDs {
			CollectionReceived := new(models.CollectionReceived)
			CollectionReceived.By = crr.By
			CollectionReceived.ByType = crr.ByType
			CollectionReceived.TnxID = v
			CollectionReceived.Status = constants.PROPERTYPAYMENTCOLLECTED
			CollectionReceived.Date = &t
			err := s.Daos.CollectedPropertyPayment(ctx, CollectionReceived)

			if err != nil {
				return err
			}

		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if err := ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Error while aborting transaction" + err.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

func (s *Service) RejectedPropertyPayment(ctx *models.Context, crr *models.CollectionReceivedRequest) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	t := time.Now()
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		CollectionReceived := new(models.CollectionReceived)
		for _, v := range crr.TnxIDs {
			CollectionReceived.By = crr.By
			CollectionReceived.ByType = crr.ByType
			CollectionReceived.TnxID = v
			CollectionReceived.Status = constants.PROPERTYPAYMENREJECTED
			CollectionReceived.Date = &t
			err := s.Daos.RejectedPropertyPayment(ctx, CollectionReceived)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if err := ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Error while aborting transaction" + err.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

// PropertyPaymentArrerAndCurrentCollection : ""
func (s *Service) PropertyPaymentArrerAndCurrentCollection(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.ArrerAndCurrentReport, error) {

	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.PropertyPaymentArrerAndCurrentCollection(ctx, filter, pagination)
}

// PropertyPaymentArrerAndCurrentCollectionExcel : ""
func (s *Service) PropertyPaymentArrerAndCurrentCollectionExcel(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.PropertyPaymentArrerAndCurrentCollection(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	// fmt.Println("'res length==>'", len(res))

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Payments"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "G1")
	excel.MergeCell(sheet1, "A2", "G2")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", sheet1))
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", "ARREAR AND CURRENT COLLECTION REPORT"))
	rowNo++
	var date string
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.MergeCell(sheet1, "A3", "G3")
	if filter.StartDate != nil {
		sd := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 1, 0, 0, 0, 0, filter.StartDate.Location())
		ed := time.Date(filter.StartDate.Year(), filter.StartDate.Month()+1, 0, 23, 59, 59, 999999999, filter.StartDate.Location())
		date = fmt.Sprintf("From %v-%v-%v To %v-%v-%v", sd.Day(), int(sd.Month()), sd.Year(), ed.Day(), int(ed.Month()), ed.Year())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", date))
		// excel.MergeCell(sheet1, "A3", "G3")

	}

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", date))
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.MergeCell(sheet1, "A4", "A5")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v", "DATE"))
	excel.MergeCell(sheet1, "B4", "C4")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v", "PROPERTY TAX"))
	excel.MergeCell(sheet1, "D4", "E4")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v", "PENALTY"))
	excel.MergeCell(sheet1, "F4", "F5")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v", "FORM FEES"))
	excel.MergeCell(sheet1, "G4", "G5")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v", "TOTAL"))
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v", "ARREAR"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", "CURRENT"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v", "ARREAR"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v", "CURRENT"))
	rowNo++
	for k, v := range res {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v", k+1))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v", v.ArrearTax))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", v.CurrentTax))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v", v.ArrearPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v", v.CurrentPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v", v.CurrentPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v", v.Formfee))
		//	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v", v.T))

	}
	return excel, nil
}

// UpdatePropertyPaymentBasicPropertyID : ""
func (s *Service) UpdatePropertyPaymentBasicPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range uniqueIds.UniqueIDs {
			resProperty, err := s.GetSingleProperty(ctx, v)
			if err != nil {
				return errors.New("Not able to get property - " + err.Error())
			}

			uniqueIds.UniqueID = resProperty.OldUniqueID
			uniqueIds.OldUniqueID = resProperty.OldUniqueID
			uniqueIds.NewUniqueID = resProperty.NewUniqueID
			err = s.Daos.UpdatePropertyPaymentBasicPropertyID(ctx, uniqueIds)
			if err != nil {
				return err
			}
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

// UpdatePropertyPaymentFYPropertyID : ""
func (s *Service) UpdatePropertyPaymentFYPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range uniqueIds.UniqueIDs {
			resProperty, err := s.GetSingleProperty(ctx, v)
			if err != nil {
				return errors.New("Not able to get property - " + err.Error())
			}

			uniqueIds.UniqueID = resProperty.OldUniqueID
			uniqueIds.OldUniqueID = resProperty.OldUniqueID
			uniqueIds.NewUniqueID = resProperty.NewUniqueID
			err = s.Daos.UpdatePropertyPaymentFYPropertyID(ctx, uniqueIds)
			if err != nil {
				return err
			}
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

// UpdatePropertyPaymentPropertyID : ""
func (s *Service) UpdatePropertyPaymentPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range uniqueIds.UniqueIDs {
			resProperty, err := s.GetSingleProperty(ctx, v)
			if err != nil {
				return errors.New("Not able to get property - " + err.Error())
			}

			uniqueIds.UniqueID = resProperty.OldUniqueID
			uniqueIds.OldUniqueID = resProperty.OldUniqueID
			uniqueIds.NewUniqueID = resProperty.NewUniqueID
			err = s.Daos.UpdatePropertyPaymentPropertyID(ctx, uniqueIds)
			if err != nil {
				return err
			}
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

func (s *Service) PropertyPaymentSummaryUpdate(ctx *models.Context, TxnId string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		propertypayment, err := s.Daos.GetSinglePropertyPaymentWithTxtID(ctx, TxnId)
		if err != nil {
			return err
		}
		propertypaymentfys, err := s.Daos.GetPropertyPaymentDemandFycWithTxtID(ctx, TxnId)
		if err != nil {
			return err
		}
		payment := new(models.Summary)
		payment.TnxID = TxnId
		for _, v := range propertypaymentfys {
			if v.FY.IsCurrent {
				payment.CurrentTax = payment.CurrentTax + v.FY.VacantLandTax + v.FY.Tax
				payment.CurrentPenalty = payment.CurrentPenalty + v.FY.Penalty
				payment.CurrentRebate = payment.CurrentRebate + v.FY.Rebate
				payment.CurrentOtherDemand = payment.CurrentOtherDemand + v.FY.OtherDemand
				payment.TotalCurrent = payment.TotalCurrent + v.FY.VacantLandTax + v.FY.Tax + v.FY.Penalty + v.FY.OtherDemand - v.FY.Rebate
			} else {
				payment.ArrearTax = payment.ArrearTax + v.FY.VacantLandTax + v.FY.Tax
				payment.ArrearPenalty = payment.ArrearPenalty + v.FY.Penalty
				payment.ArrearRebate = payment.ArrearRebate + v.FY.Rebate
				payment.ArrearOtherDemand = payment.ArrearOtherDemand + v.FY.OtherDemand
				payment.TotalArrear = payment.TotalArrear + v.FY.VacantLandTax + v.FY.Tax + v.FY.Penalty + v.FY.OtherDemand - v.FY.Rebate
			}

			payment.TotalTax = payment.TotalTax + v.FY.VacantLandTax + v.FY.Tax + v.FY.Penalty + v.FY.OtherDemand - v.FY.Rebate

		}
		payment.FormFee = payment.FormFee + propertypayment.Demand.FormFee
		payment.BoreCharge = payment.BoreCharge + propertypayment.Demand.BoreCharge
		if len(propertypaymentfys) > 0 {
			payment.ToFy = propertypaymentfys[0].FY.Name
			payment.FromFy = propertypaymentfys[len(propertypaymentfys)-1].FY.Name
		}
		if err := s.Daos.UpdatePropertyPaymentSummary(ctx, payment); err != nil {
			return errors.New("Not update collection - " + err.Error())
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
