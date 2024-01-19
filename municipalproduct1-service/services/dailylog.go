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

func (s *Service) TodaysLog() error {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)

	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		dailyLog := new(models.DailyLog)
		dailyLog.Date = &t
		datestr, err := s.Shared.UniqueDateStr(&t)
		if err != nil {
			return err
		}
		dailyLog.Datestr = datestr
		err = s.Daos.InitiateDailylog(ctx, dailyLog)
		if err != nil {
			return err
		}
		err = s.Daos.DailylogGetTodaysCompletedPayments(ctx, dailyLog)
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

func (s *Service) UpdateTotalDeand() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	cronLog := new(models.CronLog)
	cronLog.Name = constants.UPDATEPROPERTYTAXCRON
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
		cronLog.ErrorMessage = err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	IDs, err := s.Daos.GetAllPropertyIds(ctx)
	if err != nil {
		fmt.Println(err)
		isSuccess = false
		cronLog.ErrorMessage = err.Error()
		cronLog.Status = constants.CRONLOGSTATUSFAILED
	}
	IdsArray := s.Shared.SplitPropertyIds(IDs, 100)
	fmt.Println("total Sprint = >", len(IdsArray))
	for k, v := range IdsArray {
		fmt.Println("Sprint = >", k+1)
		err = s.SaveOverAllPropertyDemandForAll(v)
		if err != nil {
			fmt.Println(err)
			isSuccess = false
			cronLog.ErrorMessage = err.Error()
			cronLog.Status = constants.CRONLOGSTATUSFAILED

		}
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
