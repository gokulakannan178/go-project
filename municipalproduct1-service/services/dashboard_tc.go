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

// TcDashboard : ""
func (s *Service) TcDashboard(ctx *models.Context, tcdFilter *models.TcDashboardFilter) (*models.Tcdashboard, error) {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Tcdashboard := new(models.Tcdashboard)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		sd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		//Inhand Cash
		var propertyPaymentFilter = new(models.PropertyPaymentFilter)
		propertyPaymentFilter.Status = append(propertyPaymentFilter.Status, constants.PROPERTYPAYMENTCOMPLETED)
		propertyPaymentFilter.Collector = append(propertyPaymentFilter.Collector, tcdFilter.CollectorID)
		propertyPaymentFilter.DateRange = new(models.DateRange)
		propertyPaymentFilter.DateRange.From = &sd
		propertyPaymentFilter.DateRange.To = &ed
		propertyPaymentFilter.MOP = append(propertyPaymentFilter.MOP, constants.MOPCASH)

		payments, err := s.Daos.FilterPropertyPayment(ctx, propertyPaymentFilter, nil)
		if err != nil {
			return err
			fmt.Println(err.Error())
			payments = []models.RefPropertyPayment{}
		}

		for _, v := range payments {
			Tcdashboard.InhandBalance = v.Details.Amount + Tcdashboard.InhandBalance
		}

		//Today target
		propertyvistlogFilter := new(models.PropertyVisitLogFilter)
		propertyvistlogFilter.UserId = append(propertyvistlogFilter.UserId, tcdFilter.CollectorID)
		propertyvistlogFilter.NextDateRange = new(models.DateRange)
		propertyvistlogFilter.NextDateRange.From = &sd
		propertyvistlogFilter.NextDateRange.To = &ed
		visitlogs, err := s.Daos.FilterPropertyVisitLog(ctx, propertyvistlogFilter, nil)
		if err != nil {
			return err
		}

		var visitlogpropertyIDS []string
		for _, v := range visitlogs {
			visitlogpropertyIDS = append(visitlogpropertyIDS, v.PropertyID)
		}
		propertyDemandFilter := new(models.PropertyDemandFilter)
		if len(visitlogpropertyIDS) > 0 {

			propertyDemandFilter.PropertyIDs = append(propertyDemandFilter.PropertyIDs, visitlogpropertyIDS...)
			propertyDemandFilter.Status = append(propertyDemandFilter.Status, constants.PROPERTYSTATUSACTIVE, constants.PROPERTYSTATUSPENDING)

			propertydemand, err := s.GetMultiplePropertyDemandCalc(ctx, propertyDemandFilter, nil)
			if err != nil {
				return err
			}

			for _, v := range propertydemand {
				Tcdashboard.TodaysTarget = Tcdashboard.TodaysTarget + v.TotalTax
			}
		}

		//Todays Collection Target
		todaysCollectionFilterTarget := new(models.PropertyPaymentFilter)
		if len(visitlogpropertyIDS) > 0 {
			todaysCollectionFilterTarget.PropertyIds = append(todaysCollectionFilterTarget.PropertyIds, visitlogpropertyIDS...)
		}

		todaysCollectionFilterTarget.DateRange = new(models.DateRange)
		todaysCollectionFilterTarget.DateRange.From = &sd
		todaysCollectionFilterTarget.DateRange.To = &ed
		todaysCollectionFilterTarget.Status = append(todaysCollectionFilterTarget.Status, constants.PROPERTYPAYMENTCOMPLETED)
		todaysCollectionFilterTarget.Collector = append(todaysCollectionFilterTarget.Collector, tcdFilter.CollectorID)

		propertyPaymenttarget, err := s.Daos.FilterPropertyPayment(ctx, todaysCollectionFilterTarget, nil)
		if err != nil {
			return err
		}
		for _, v := range propertyPaymenttarget {
			Tcdashboard.TodaysTargetCollection = Tcdashboard.TodaysTargetCollection + v.Details.Amount
		}

		//Todays Collection
		todaysCollectionFilter := new(models.PropertyPaymentFilter)
		todaysCollectionFilter.DateRange = new(models.DateRange)
		todaysCollectionFilter.DateRange.From = &sd
		todaysCollectionFilter.DateRange.To = &ed
		todaysCollectionFilter.Status = append(todaysCollectionFilter.Status, constants.PROPERTYPAYMENTCOMPLETED)
		todaysCollectionFilter.Collector = append(todaysCollectionFilter.Collector, tcdFilter.CollectorID)

		propertyPayment, err := s.Daos.FilterPropertyPayment(ctx, todaysCollectionFilter, nil)
		if err != nil {
			return err
		}
		for _, v := range propertyPayment {
			Tcdashboard.TodaysCollection = Tcdashboard.TodaysCollection + v.Details.Amount
		}

		// Yesterdays target
		psd := time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, t.Location())
		ped := time.Date(t.Year(), t.Month(), t.Day()-1, 23, 59, 59, 0, t.Location())

		yesterdayPropertyvistlogFilter := new(models.PropertyVisitLogFilter)
		yesterdayPropertyvistlogFilter.UserId = append(yesterdayPropertyvistlogFilter.UserId, tcdFilter.CollectorID)
		yesterdayPropertyvistlogFilter.NextDateRange = new(models.DateRange)
		yesterdayPropertyvistlogFilter.NextDateRange.From = &psd
		yesterdayPropertyvistlogFilter.NextDateRange.To = &ped
		yesterdayVisitlogs, err := s.Daos.FilterPropertyVisitLog(ctx, yesterdayPropertyvistlogFilter, nil)
		if err != nil {
			return err
		}
		var yesterdayVisitlogpropertyIDS []string

		for _, v := range yesterdayVisitlogs {
			yesterdayVisitlogpropertyIDS = append(yesterdayVisitlogpropertyIDS, v.PropertyID)
		}

		yesterdayPropertyDemandFilter := new(models.PropertyDemandFilter)
		if len(yesterdayVisitlogpropertyIDS) > 0 {
			yesterdayPropertyDemandFilter.PropertyIDs = append(yesterdayPropertyDemandFilter.PropertyIDs, yesterdayVisitlogpropertyIDS...)
			yesterdayPropertyDemandFilter.Status = append(yesterdayPropertyDemandFilter.Status, constants.PROPERTYSTATUSACTIVE, constants.PROPERTYSTATUSPENDING)

			yesterdayPropertydemand, err := s.GetMultiplePropertyDemandCalc(ctx, propertyDemandFilter, nil)
			if err != nil {
				return err
			}

			for _, v := range yesterdayPropertydemand {
				Tcdashboard.YesterdayTarget = Tcdashboard.YesterdayTarget + v.TotalTax
			}

		}

		//yesterday Collection
		yesterdaysTargetCollectionFilter := new(models.PropertyPaymentFilter)
		yesterdaysTargetCollectionFilter.DateRange = new(models.DateRange)
		yesterdaysTargetCollectionFilter.DateRange.From = &psd
		yesterdaysTargetCollectionFilter.DateRange.To = &ped
		yesterdaysTargetCollectionFilter.Status = append(yesterdaysTargetCollectionFilter.Status, constants.PROPERTYPAYMENTCOMPLETED)
		yesterdaysTargetCollectionFilter.Collector = append(yesterdaysTargetCollectionFilter.Collector, tcdFilter.CollectorID)

		yesterdayPropertyPaymentTarget, err := s.Daos.FilterPropertyPayment(ctx, yesterdaysTargetCollectionFilter, nil)
		if err != nil {
			return err
		}
		for _, v := range yesterdayPropertyPaymentTarget {
			Tcdashboard.YesterdayTargetCollection = Tcdashboard.YesterdayTargetCollection + v.Details.Amount
		}

		//yesterday Collection
		yesterdaysCollectionFilter := new(models.PropertyPaymentFilter)
		yesterdaysCollectionFilter.DateRange = new(models.DateRange)
		yesterdaysCollectionFilter.DateRange.From = &psd
		yesterdaysCollectionFilter.DateRange.To = &ped
		yesterdaysCollectionFilter.Status = append(yesterdaysCollectionFilter.Status, constants.PROPERTYPAYMENTCOMPLETED)
		yesterdaysCollectionFilter.Collector = append(yesterdaysCollectionFilter.Collector, tcdFilter.CollectorID)

		yesterdayPropertyPayment, err := s.Daos.FilterPropertyPayment(ctx, yesterdaysCollectionFilter, nil)
		if err != nil {
			return err
		}
		for _, v := range yesterdayPropertyPayment {
			Tcdashboard.YesterdayCollection = Tcdashboard.YesterdayCollection + v.Details.Amount
		}

		//Filling OverAll Demand
		demandfilter := new(models.DashboardDemandAndCollectionFilter)
		demandfilter.Status = append(demandfilter.Status, constants.PROPERTYSTATUSACTIVE, constants.PROPERTYSTATUSPENDING)

		overall, err := s.DashboardDemandAndCollection(ctx, demandfilter)
		if err != nil {
			fmt.Println(err)
			overall = new(models.DashboardDemandAndCollection)
		}

		Tcdashboard.Overall = overall
		//Commit transcation
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
	return Tcdashboard, nil
}

func (s *Service) GetSingleDefaultProductConfiguration(ctx *models.Context) (*models.RefProductConfiguration, error) {
	tower, err := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		return nil, err
	}
	return tower, nil
}
