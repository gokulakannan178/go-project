package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/app"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveStateWeatherAlert :""
func (s *Service) SaveStateWeatherAlert(ctx *models.Context, StateWeatherAlert *models.StateWeatherAlert) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	StateWeatherAlert.Status = constants.STATEWEATHERALERTSTATUSACTIVE
	StateWeatherAlert.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 StateWeatherAlert.created")
	StateWeatherAlert.Created = &created
	log.Println("b4 StateWeatherAlert.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveStateWeatherAlert(ctx, StateWeatherAlert)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//UpdateStateWeatherAlert : ""
func (s *Service) UpdateStateWeatherAlert(ctx *models.Context, StateWeatherAlert *models.StateWeatherAlert) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateStateWeatherAlert(ctx, StateWeatherAlert)
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
func (s *Service) UpdateWeatherAlertMaster(ctx *models.Context, StateWeatherAlert *models.UpdateStateWeatherAlert) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.UpdateWeatherAlertMaster(ctx, StateWeatherAlert)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		if StateWeatherAlert.IsUpdateMode == "Yes" {
			err := s.Daos.IsUpdateStateWeatherAlertMaster(ctx, StateWeatherAlert)
			if err != nil {
				if err = ctx.Session.AbortTransaction(sc); err != nil {
					return errors.New("Transaction Aborted with error" + err.Error())
				}
				return errors.New("Transaction Aborted - " + err.Error())
			}
		}

		return nil

	}); err != nil {
		return err
	}
	return nil
}

//EnableStateWeatherAlert : ""
func (s *Service) EnableStateWeatherAlert(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableStateWeatherAlert(ctx, UniqueID)
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

//DisableStateWeatherAlert : ""
func (s *Service) DisableStateWeatherAlert(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableStateWeatherAlert(ctx, UniqueID)
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

//DeleteStateWeatherAlert : ""
func (s *Service) DeleteStateWeatherAlert(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteStateWeatherAlert(ctx, UniqueID)
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

//GetSingleStateWeatherAlert :""
func (s *Service) GetSingleStateWeatherAlert(ctx *models.Context, UniqueID string) (*models.RefStateWeatherAlert, error) {
	StateWeatherAlert, err := s.Daos.GetSingleStateWeatherAlert(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return StateWeatherAlert, nil
}

//FilterStateWeatherAlert :""
func (s *Service) FilterStateWeatherAlert(ctx *models.Context, StateWeatherAlertfilter *models.StateWeatherAlertFilter, pagination *models.Pagination) (StateWeatherAlert []models.RefStateWeatherAlert, err error) {
	return s.Daos.FilterStateWeatherAlert(ctx, StateWeatherAlertfilter, pagination)
}
func (s *Service) SendStateWeatherAlertCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	//Dissemination := new(models.Dissemination)
	stateWeatherAlert, err := s.Daos.GetTodayActiveStateWeatherAlert(ctx)
	if err != nil {
		log.Println("dissemination not found" + err.Error())
	}
	fmt.Println("no.of stateWeatherAlert==>", len(stateWeatherAlert))
	for _, v := range stateWeatherAlert {
		_, err := s.DisseminationWeatherAlert(ctx, &v, true)
		if err != nil {
			log.Println("not Weather Alert Not send----" + v.ParameterId.Name + "-" + v.SeverityType.Name + "" + err.Error())
			continue
		}
	}
}
