package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDistrictWeatherAlert :""
func (s *Service) SaveDistrictWeatherAlert(ctx *models.Context, DistrictWeatherAlert *models.DistrictWeatherAlert) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	DistrictWeatherAlert.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
	DistrictWeatherAlert.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DistrictWeatherAlert.created")
	DistrictWeatherAlert.Created = &created
	log.Println("b4 DistrictWeatherAlert.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDistrictWeatherAlert(ctx, DistrictWeatherAlert)
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

//UpdateDistrictWeatherAlert : ""
func (s *Service) UpdateDistrictWeatherAlert(ctx *models.Context, DistrictWeatherAlert *models.DistrictWeatherAlert) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDistrictWeatherAlert(ctx, DistrictWeatherAlert)
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
func (s *Service) UpdateDistrictWeatherAlertMasterwithWeatheralert(ctx *models.Context, DistrictWeatherAlert *models.UpdateDistrictWeatherAlert) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.UpdateDistrictWeatherAlertMasterwithWeatheralert(ctx, DistrictWeatherAlert)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		if DistrictWeatherAlert.IsUpdateMode == "Yes" {
			err := s.Daos.IsUpdateDistrictWeatherAlertMaster(ctx, DistrictWeatherAlert)
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

//EnableDistrictWeatherAlert : ""
func (s *Service) EnableDistrictWeatherAlert(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDistrictWeatherAlert(ctx, UniqueID)
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

//DisableDistrictWeatherAlert : ""
func (s *Service) DisableDistrictWeatherAlert(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDistrictWeatherAlert(ctx, UniqueID)
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

//DeleteDistrictWeatherAlert : ""
func (s *Service) DeleteDistrictWeatherAlert(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDistrictWeatherAlert(ctx, UniqueID)
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

//GetSingleDistrictWeatherAlert :""
func (s *Service) GetSingleDistrictWeatherAlert(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherAlert, error) {
	DistrictWeatherAlert, err := s.Daos.GetSingleDistrictWeatherAlert(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return DistrictWeatherAlert, nil
}

//FilterDistrictWeatherAlert :""
func (s *Service) FilterDistrictWeatherAlert(ctx *models.Context, DistrictWeatherAlertfilter *models.DistrictWeatherAlertFilter, pagination *models.Pagination) (DistrictWeatherAlert []models.RefDistrictWeatherAlert, err error) {
	return s.Daos.FilterDistrictWeatherAlert(ctx, DistrictWeatherAlertfilter, pagination)
}

// func (s *Service) SendDistrictWeatherAlertCron() {
// 	c := context.TODO()
// 	ctx := app.GetApp(c, s.Daos)
// 	defer ctx.Client.Disconnect(c)
// 	//Dissemination := new(models.Dissemination)
// 	districtWeatherAlert, err := s.Daos.GetTodayActiveDistrictWeatherAlert(ctx)
// 	if err != nil {
// 		log.Println("dissemination not found" + err.Error())
// 	}
// 	fmt.Println("no.of districtWeatherAlert==>", len(districtWeatherAlert))
// 	for _, v := range districtWeatherAlert {
// 		_, err := s.DisseminationWeatherAlert(ctx, &v, true)
// 		if err != nil {
// 			log.Println("not Weather Alert Not send----" + v.ParameterId.Name + "-" + v.SeverityType.Name + "" + err.Error())
// 			continue
// 		}
// 	}
// }
func (s *Service) GetSingleDistrictWeatherDataWithCureentDate(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherData, error) {
	districtweatherdata, err := s.Daos.GetSingleDistrictWeatherDataWithCureentDate(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return districtweatherdata, nil
}
