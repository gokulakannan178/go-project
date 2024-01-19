package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveDashBoardProperty : ""
func (s *Service) SaveDashBoardProperty(ctx *models.Context, property *models.PropertyDashBoard) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	property.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDASHBOARDPROPERTY)
	property.Status = constants.DASHBOARDPROPERTYSTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDashBoardProperty(ctx, property)
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
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

//GetSingleDashBoardProperty :""
func (s *Service) GetSingleDashBoardProperty(ctx *models.Context, UniqueID string) (*models.RefPropertyDashBoard, error) {
	tower, err := s.Daos.GetSingleDashBoardProperty(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateDashBoardProperty : ""
func (s *Service) UpdateDashBoardProperty(ctx *models.Context, property *models.PropertyDashBoard) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDashBoardProperty(ctx, property)
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

// EnableDashBoardProperty : ""
func (s *Service) EnableDashBoardProperty(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDashBoardProperty(ctx, UniqueID)
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

//DisableDashBoardProperty : ""
func (s *Service) DisableDashBoardProperty(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDashBoardProperty(ctx, UniqueID)
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

//DeleteDashBoardProperty : ""
func (s *Service) DeleteDashBoardProperty(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDashBoardProperty(ctx, UniqueID)
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

// FilterPropertyDayWise : ""
func (s *Service) FilterDashBoardProperty(ctx *models.Context, filter *models.PropertyDashBoardFilter, pagination *models.Pagination) ([]models.RefPropertyDashBoard, error) {
	return s.Daos.FilterDashBoardProperty(ctx, filter, pagination)

}

//Demand - overall
func (s *Service) PropertyOverallDemandCron() {
	fmt.Println("PropertyOverallDemandCron STARTED")
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	cronLog := new(models.CronLog)
	cronLog.Name = constants.PROPERTYOVERALLDEMANDCRON
	cronLog.Status = constants.CRONLOGSTATUSINIT
	t := time.Now()
	cronLog.DateStr = t.Format("2006-Jan-02")
	cronLog.Date = &t
	cronLog.StartTime = &t
	cronLog.IsCurrentScript = true
	isSuccess := true
	err := s.Daos.CloseOldCron(ctx, cronLog.Name)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Daos.InitCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
		isSuccess = false
		cronLog.ErrorMessage = "Init CronLog error - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	propertyDemandFilter := new(models.PropertyFilter)
	propertyDemandFilter.Status = append(propertyDemandFilter.Status, []string{constants.PROPERTYSTATUSACTIVE, constants.PROPERTYSTATUSPENDING}...)
	if err := s.Daos.PropertyOverallDemandCron(ctx, propertyDemandFilter); err != nil {
		log.Println("Error in Property Overall Demand Cron - " + err.Error())
		isSuccess = false
		cronLog.ErrorMessage = "Error in Property Overall Demand Cron - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	if isSuccess {
		cronLog.Status = constants.CRONLOGSTATUSCOMPLETED
	}
	t1 := time.Now()
	cronLog.EndTime = &t1
	err = s.Daos.EndCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
	}

}

//Demand - today

func (s *Service) PropertyTodayDemandCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	cronLog := new(models.CronLog)
	cronLog.Name = constants.PROPERTYTODAYDEMANDCRON
	cronLog.Status = constants.CRONLOGSTATUSINIT
	t1 := time.Now()
	cronLog.DateStr = t1.Format("2006-Jan-02")
	cronLog.Date = &t1
	cronLog.StartTime = &t1
	cronLog.IsCurrentScript = true
	isSuccess := true
	err := s.Daos.CloseOldCron(ctx, cronLog.Name)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Daos.InitCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
		isSuccess = false
		cronLog.ErrorMessage = "Init CronLog error - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	propertyDemandFilter := new(models.PropertyFilter)
	propertyDemandFilter.Status = append(propertyDemandFilter.Status, []string{constants.PROPERTYSTATUSACTIVE, constants.PROPERTYSTATUSPENDING}...)
	propertyDemandFilter.AppliedRange = new(models.PropertyAppliedRange)
	t := time.Now()
	// t := time.Now().AddDate(0, -4, 0)
	propertyDemandFilter.AppliedRange.From = &t
	if err := s.Daos.PropertyTodayDemandCron(ctx, propertyDemandFilter); err != nil {
		log.Println("Error in Property Today Demand Cron - " + err.Error())
		isSuccess = false
		cronLog.ErrorMessage = "Error in Property Today Demand Cron - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	if isSuccess {
		cronLog.Status = constants.CRONLOGSTATUSCOMPLETED
	}
	t2 := time.Now()
	cronLog.EndTime = &t2
	err = s.Daos.EndCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
	}

}

//Collection - overall

func (s *Service) PropertyOverallCollectionCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	cronLog := new(models.CronLog)
	cronLog.Name = constants.PROPERTYOVERALLCOLLECTIONCRON
	cronLog.Status = constants.CRONLOGSTATUSINIT
	t := time.Now()
	cronLog.DateStr = t.Format("2006-Jan-02")
	cronLog.Date = &t
	cronLog.StartTime = &t
	cronLog.IsCurrentScript = true
	isSuccess := true
	err := s.Daos.CloseOldCron(ctx, cronLog.Name)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Daos.InitCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
		isSuccess = false
		cronLog.ErrorMessage = "Init CronLog error - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	filter := new(models.PropertyPaymentFilter)
	filter.Status = append(filter.Status, []string{constants.PROPERTYPAYMENTCOMPLETED}...)
	if err := s.Daos.PropertyOverallCollectionCron(ctx, filter); err != nil {
		log.Println("Error in Property Overall Collection Cron - " + err.Error())
		isSuccess = false
		cronLog.ErrorMessage = "Error in Property Overall Collection Cron - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	if isSuccess {
		cronLog.Status = constants.CRONLOGSTATUSCOMPLETED
	}
	t1 := time.Now()
	cronLog.EndTime = &t1
	err = s.Daos.EndCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
	}

}

//Collection - today

func (s *Service) PropertyTodayCollectionCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	cronLog := new(models.CronLog)
	cronLog.Name = constants.PROPERTYTODAYCOLLECTIONCRON
	cronLog.Status = constants.CRONLOGSTATUSINIT
	t1 := time.Now()
	cronLog.DateStr = t1.Format("2006-Jan-02")
	cronLog.Date = &t1
	cronLog.StartTime = &t1
	cronLog.IsCurrentScript = true
	isSuccess := true
	err := s.Daos.CloseOldCron(ctx, cronLog.Name)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Daos.InitCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
		isSuccess = false
		cronLog.ErrorMessage = "Init CronLog error - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	filter := new(models.PropertyPaymentFilter)
	filter.Status = append(filter.Status, []string{constants.PROPERTYPAYMENTCOMPLETED}...)
	filter.DateRange = new(models.DateRange)
	// t := time.Now().AddDate(0, -4, 0)
	t := time.Now()
	filter.DateRange.From = &t
	if err := s.Daos.PropertyTodayCollectionCron(ctx, filter); err != nil {
		log.Println("Error in Property Today Collection Cron - " + err.Error())
		isSuccess = false
		cronLog.ErrorMessage = "Error in Property Today Collection Cron - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	if isSuccess {
		cronLog.Status = constants.CRONLOGSTATUSCOMPLETED
	}
	t2 := time.Now()
	cronLog.EndTime = &t2
	err = s.Daos.EndCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
	}

}

//GetOverAllPropertyDashBoard : ""
func (s *Service) GetOverAllPropertyDashBoard(ctx *models.Context) (*models.OverallDashBoard, error) {
	tower, err := s.Daos.GetOverAllPropertyDashBoard(ctx)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// Demand - Month

func (s *Service) PropertyMonthDemandCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	cronLog := new(models.CronLog)
	cronLog.Name = constants.PROPERTYMONTHDEMANDCRON
	cronLog.Status = constants.CRONLOGSTATUSINIT
	t1 := time.Now()
	cronLog.DateStr = t1.Format("2006-Jan-02")
	cronLog.Date = &t1
	cronLog.StartTime = &t1
	cronLog.IsCurrentScript = true
	isSuccess := true
	err := s.Daos.CloseOldCron(ctx, cronLog.Name)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Daos.InitCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
		isSuccess = false
		cronLog.ErrorMessage = "Init CronLog error - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	propertyDemandFilter := new(models.PropertyFilter)
	propertyDemandFilter.Status = append(propertyDemandFilter.Status, []string{constants.PROPERTYSTATUSACTIVE, constants.PROPERTYSTATUSPENDING}...)
	propertyDemandFilter.AppliedRange = new(models.PropertyAppliedRange)
	t := time.Now()
	// t := time.Now().AddDate(0, -4, 0)
	//
	var sd, ed *time.Time
	if propertyDemandFilter.AppliedRange.From == nil {
		sd := s.Shared.BeginningOfMonth(t)
		propertyDemandFilter.AppliedRange.From = &sd
		ed := s.Shared.EndOfMonth(t)
		propertyDemandFilter.AppliedRange.To = &ed
	}
	sdt := s.Shared.BeginningOfMonth(*propertyDemandFilter.AppliedRange.From)
	sd = &sdt
	edt := s.Shared.EndOfMonth(*propertyDemandFilter.AppliedRange.To)
	ed = &edt
	//

	propertyDemandFilter.AppliedRange.From = sd
	propertyDemandFilter.AppliedRange.To = ed
	if err := s.Daos.PropertyMonthDemandCron(ctx, propertyDemandFilter); err != nil {
		log.Println("Error in Property Today Demand Cron - " + err.Error())
		isSuccess = false
		cronLog.ErrorMessage = "Error in Property Today Demand Cron - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	if isSuccess {
		cronLog.Status = constants.CRONLOGSTATUSCOMPLETED
	}
	t2 := time.Now()
	cronLog.EndTime = &t2
	err = s.Daos.EndCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
	}

}

//Collection - Month

func (s *Service) PropertyMonthCollectionCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	cronLog := new(models.CronLog)
	cronLog.Name = constants.PROPERTYMONTHCOLLECTIONCRON
	cronLog.Status = constants.CRONLOGSTATUSINIT
	t1 := time.Now()
	cronLog.DateStr = t1.Format("2006-Jan-02")
	cronLog.Date = &t1
	cronLog.StartTime = &t1
	cronLog.IsCurrentScript = true
	isSuccess := true
	err := s.Daos.CloseOldCron(ctx, cronLog.Name)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Daos.InitCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
		isSuccess = false
		cronLog.ErrorMessage = "Init CronLog error - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	filter := new(models.PropertyPaymentFilter)
	filter.Status = append(filter.Status, []string{constants.PROPERTYPAYMENTCOMPLETED}...)
	filter.DateRange = new(models.DateRange)
	// t := time.Now().AddDate(0, -4, 0)
	t := time.Now()

	//
	var sd, ed *time.Time
	if filter.DateRange.From == nil {
		sd := s.Shared.BeginningOfMonth(t)
		filter.DateRange.From = &sd
		ed := s.Shared.EndOfMonth(t)
		filter.DateRange.To = &ed
	}
	sdt := s.Shared.BeginningOfMonth(*filter.DateRange.From)
	sd = &sdt
	edt := s.Shared.EndOfMonth(*filter.DateRange.To)
	ed = &edt
	//
	filter.DateRange.From = sd
	filter.DateRange.To = ed
	if err := s.Daos.PropertyMonthCollectionCron(ctx, filter); err != nil {
		log.Println("Error in Property Today Collection Cron - " + err.Error())
		isSuccess = false
		cronLog.ErrorMessage = "Error in Property Today Collection Cron - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	if isSuccess {
		cronLog.Status = constants.CRONLOGSTATUSCOMPLETED
	}
	t2 := time.Now()
	cronLog.EndTime = &t2
	err = s.Daos.EndCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
	}

}

// Demand - Year

func (s *Service) PropertyYearDemandCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	cronLog := new(models.CronLog)
	cronLog.Name = constants.PROPERTYYEARDEMANDCRON
	cronLog.Status = constants.CRONLOGSTATUSINIT
	t1 := time.Now()
	cronLog.DateStr = t1.Format("2006-Jan-02")
	cronLog.Date = &t1
	cronLog.StartTime = &t1
	cronLog.IsCurrentScript = true
	isSuccess := true
	err := s.Daos.CloseOldCron(ctx, cronLog.Name)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Daos.InitCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
		isSuccess = false
		cronLog.ErrorMessage = "Init CronLog error - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	propertyDemandFilter := new(models.PropertyFilter)
	propertyDemandFilter.Status = append(propertyDemandFilter.Status, []string{constants.PROPERTYSTATUSACTIVE, constants.PROPERTYSTATUSPENDING}...)
	propertyDemandFilter.AppliedRange = new(models.PropertyAppliedRange)
	// t := time.Now()
	// t := time.Now().AddDate(0, -4, 0)
	//
	resFY, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		fmt.Println(err)
	}
	sd := time.Date(resFY.From.Year(), resFY.From.Month(), resFY.From.Day(), 0, 0, 0, 0, resFY.From.Location())
	ed := time.Date(resFY.To.Year(), resFY.To.Month(), resFY.To.Day(), 23, 59, 59, 0, resFY.To.Location())
	//

	propertyDemandFilter.AppliedRange.From = &sd
	propertyDemandFilter.AppliedRange.To = &ed
	if err := s.Daos.PropertyYearDemandCron(ctx, propertyDemandFilter); err != nil {
		log.Println("Error in Property Today Demand Cron - " + err.Error())
		isSuccess = false
		cronLog.ErrorMessage = "Error in Property Today Demand Cron - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	if isSuccess {
		cronLog.Status = constants.CRONLOGSTATUSCOMPLETED
	}
	t2 := time.Now()
	cronLog.EndTime = &t2
	err = s.Daos.EndCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
	}

}

//Collection - Year

func (s *Service) PropertyYearCollectionCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	cronLog := new(models.CronLog)
	cronLog.Name = constants.PROPERTYYEARCOLLECTIONCRON
	cronLog.Status = constants.CRONLOGSTATUSINIT
	t1 := time.Now()
	cronLog.DateStr = t1.Format("2006-Jan-02")
	cronLog.Date = &t1
	cronLog.StartTime = &t1
	cronLog.IsCurrentScript = true
	isSuccess := true
	err := s.Daos.CloseOldCron(ctx, cronLog.Name)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Daos.InitCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
		isSuccess = false
		cronLog.ErrorMessage = "Init CronLog error - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	filter := new(models.PropertyPaymentFilter)
	filter.Status = append(filter.Status, []string{constants.PROPERTYPAYMENTCOMPLETED}...)
	filter.DateRange = new(models.DateRange)
	// t := time.Now().AddDate(0, -4, 0)
	// t := time.Now()

	//
	resFY, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		fmt.Println(err)
	}
	sd := time.Date(resFY.From.Year(), resFY.From.Month(), resFY.From.Day(), 0, 0, 0, 0, resFY.From.Location())
	ed := time.Date(resFY.To.Year(), resFY.To.Month(), resFY.To.Day(), 23, 59, 59, 0, resFY.To.Location())
	//
	filter.DateRange.From = &sd
	filter.DateRange.To = &ed
	if err := s.Daos.PropertyYearCollectionCron(ctx, filter); err != nil {
		log.Println("Error in Property Today Collection Cron - " + err.Error())
		isSuccess = false
		cronLog.ErrorMessage = "Error in Property Today Collection Cron - " + err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	if isSuccess {
		cronLog.Status = constants.CRONLOGSTATUSCOMPLETED
	}
	t2 := time.Now()
	cronLog.EndTime = &t2
	err = s.Daos.EndCronLog(ctx, cronLog)
	if err != nil {
		fmt.Println(err)
	}

}
