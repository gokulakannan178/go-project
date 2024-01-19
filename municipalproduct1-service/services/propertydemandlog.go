package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePropertyDemand : ""
func (s *Service) SavePropertyDemand(ctx *models.Context, propertyID string) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		var property *models.Property
		refProp, err := s.Daos.GetSingleProperty(ctx, propertyID)
		if err != nil {
			return errors.New("Err in geting property - " + err.Error())
		}

		property = &refProp.Property
		propertyDemandFilter := new(models.PropertyDemandFilter)
		propertyDemandFilter.PropertyID = property.UniqueID
		fmt.Println("property id - ", propertyDemandFilter)
		demand, err := s.GetPropertyDemandCalc(ctx, propertyDemandFilter, "")
		if err != nil {
			return errors.New("Not able to prepare demand - " + err.Error())
		}
		if demand == nil {
			return errors.New("Not able to prepare demand - nil demand")
		}
		demand.PropertyID = property.UniqueID
		/*
			if err := s.Daos.SavePropertyDemandLog(ctx, &demand.PropertyDemandLog); err != nil {
				return errors.New("Not able to save demand - " + err.Error())
			}*/

		for k := range demand.FYs {
			demand.FYs[k].PropertyID = property.UniqueID
		}
		/*
			if err := s.Daos.SavePropertyFyDemandLog(ctx, property.UniqueID, demand.FYs); err != nil {
				return errors.New("Not able to save demand  FY- " + err.Error())
			}
		*/
		UpdateDemand := new(models.UpdateDemand)
		UpdateDemand.TotalTax = demand.TotalTax
		UpdateDemand.Current = demand.Current
		UpdateDemand.Arrear = demand.Arrear
		if err := s.Daos.UpdateDemandToProperty(ctx, property.UniqueID, UpdateDemand); err != nil {
			return errors.New("Err in updating demand in property- " + err.Error())
		}
		demand.OverallPropertyDemand.PropertyID = demand.UniqueID
		if err := s.Daos.UpdateOverallPropertyDemand(ctx, &demand.OverallPropertyDemand); err != nil {
			return errors.New("Err in updating overallpropertydemand - " + err.Error())
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

func (s *Service) SavePropertyDemandForAll(IDs []string) error {

	for k, v := range IDs {
		fmt.Println("k------>", k)
		c := context.TODO()
		ctx := app.GetApp(c, s.Daos)
		defer ctx.Client.Disconnect(c)

		err := s.SavePropertyDemand(ctx, v)
		fmt.Println(v, err)
	}
	return nil
}

func (s *Service) SaveOverAllPropertyDemand(ctx *models.Context, propertyID string) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		var property *models.Property
		refProp, err := s.Daos.GetSingleProperty(ctx, propertyID)
		if err != nil {
			return errors.New("Err in geting property - " + err.Error())
		}

		property = &refProp.Property
		propertyDemandFilter := new(models.PropertyDemandFilter)
		propertyDemandFilter.PropertyID = property.UniqueID
		propertyDemandFilter.IsOmitPaidFys = false
		fmt.Println("property id - ", propertyDemandFilter)

		//Calculate Demand
		demand, err := s.GetPropertyDemandCalc(ctx, propertyDemandFilter, "")

		if err != nil {
			return errors.New("Not able to prepare demand - " + err.Error())
		}
		if demand == nil {
			return errors.New("Not able to prepare demand - nil demand")
		}
		for _, v := range demand.FYs {
			if v.IsCurrent {
				property.NDemand.Current.Tax = property.NDemand.Current.Tax + v.Tax + v.VacantLandTax
				property.NDemand.Current.Penalty = property.NDemand.Current.Penalty + v.Penalty
				property.NDemand.Current.Rebate = property.NDemand.Current.Rebate + v.Rebate

			} else {
				property.NDemand.Arrear.Tax = property.NDemand.Arrear.Tax + v.Tax + v.VacantLandTax
				property.NDemand.Arrear.Penalty = property.NDemand.Arrear.Penalty + v.Penalty
				property.NDemand.Arrear.Rebate = property.NDemand.Arrear.Rebate + v.Rebate
			}
		}
		property.NDemand.Total.Tax = property.NDemand.Current.Tax + property.NDemand.Arrear.Tax
		property.NDemand.Total.Penalty = property.NDemand.Current.Penalty + property.NDemand.Arrear.Penalty
		property.NDemand.Total.Rebate = property.NDemand.Current.Rebate + property.NDemand.Arrear.Rebate
		if err := s.Daos.UpdateNewDemandToProperty(ctx, property.UniqueID, &property.NDemand); err != nil {
			return errors.New("Err in updating demand in property- " + err.Error())
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

func (s *Service) SaveOverAllPropertyDemandForAll(IDs []string) error {

	for k, v := range IDs {
		fmt.Println("k------>", k)
		c := context.TODO()
		ctx := app.GetApp(c, s.Daos)
		defer ctx.Client.Disconnect(c)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Got panic")
				time.Sleep(10000)
				ctx = app.GetApp(c, s.Daos)
			}
		}()
		err := s.SaveOverAllPropertyDemand(ctx, v)
		fmt.Println(v, err)
	}
	return nil
}

//Updated

func (s *Service) SaveOverAllPropertyDemandToProperty(ctx *models.Context, propertyID string) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		var property *models.Property
		refProp, err := s.Daos.GetSingleProperty(ctx, propertyID)
		if err != nil {
			return errors.New("Err in geting property - " + err.Error())
		}

		property = &refProp.Property
		propertyDemandFilter := new(models.PropertyDemandFilter)
		propertyDemandFilter.PropertyID = property.UniqueID
		propertyDemandFilter.AllDemand = false
		// propertyDemandFilter.IsOmitPaidFys = true
		fmt.Println("property id - ", propertyDemandFilter)
		demand, err := s.GetPropertyDemandCalc(ctx, propertyDemandFilter, "")

		if err != nil {
			return errors.New("Not able to prepare demand - " + err.Error())
		}
		if demand == nil {
			return errors.New("Not able to prepare demand - nil demand")
		}
		demand.PropertyID = property.UniqueID
		demand.OverallPropertyDemand.PropertyID = demand.UniqueID
		//demand.OverallPropertyDemand
		if err := s.Daos.UpdateOverallPropertyDemand(ctx, &demand.OverallPropertyDemand); err != nil {
			return errors.New("Err in updating overallpropertydemand - " + err.Error())
		}
		fmt.Println("OVERALL PROPERTY DEMAND UPDATED")
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

func (s *Service) SaveOverallPropertyDemandForAllV2(IDs []string) error {

	for k, v := range IDs {
		fmt.Println("k------>", k)
		c := context.TODO()
		ctx := app.GetApp(c, s.Daos)
		defer ctx.Client.Disconnect(c)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Got panic")
				time.Sleep(10000)
				ctx = app.GetApp(c, s.Daos)
			}
		}()
		err := s.SaveOverAllPropertyDemandToProperty(ctx, v)
		fmt.Println(v, err)
	}
	return nil
}

func (s *Service) GetOverAllPropertyDemandToProperty(ctx *models.Context, propertyID string) (*models.PropertyDemand, error) {
	log.Println("transaction start")

	var property *models.Property
	refProp, err := s.Daos.GetSingleProperty(ctx, propertyID)
	if err != nil {
		return nil, errors.New("Err in geting property - " + err.Error())
	}

	property = &refProp.Property
	propertyDemandFilter := new(models.PropertyDemandFilter)
	propertyDemandFilter.PropertyID = property.UniqueID
	propertyDemandFilter.AllDemand = false
	// propertyDemandFilter.IsOmitPaidFys = true
	fmt.Println("property id - ", propertyDemandFilter)
	demand, err := s.GetPropertyDemandCalc(ctx, propertyDemandFilter, "")

	if err != nil {
		return nil, errors.New("Not able to prepare demand - " + err.Error())
	}
	if demand == nil {
		return nil, errors.New("Not able to prepare demand - nil demand")
	}
	return demand, nil

}

// UpdatePropertyDemandLogPropertyID : ""
func (s *Service) UpdatePropertyDemandLogPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
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
			err = s.Daos.UpdatePropertyDemandLogPropertyID(ctx, uniqueIds)
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
