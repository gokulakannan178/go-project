package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"municipalproduct1-service/app"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"os"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

//SavePropertyV2 :""
func (s *Service) SavePropertyV2(ctx *models.Context, property *models.Property) error {

	client := s.Daos.GetDBV3(context.TODO())
	defer client.Disconnect(context.TODO())
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	database := client.Database("municipalproduct1")

	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())
	log.Println("transaction start")

	if err := mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		//Start Transaction
		// wc := writeconcern.New(writeconcern.WMajority())
		// rc := readconcern.Snapshot()
		// txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
		if err := session.StartTransaction(txnOpts); err != nil {
			return err
		}
		ctx.CTX = sc

		t := time.Now()
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM
		//property.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTY)
		property.UniqueID, property.SortOrder = s.Daos.GetUniqueIDV2(ctx, constants.COLLECTIONPROPERTY)
		property.ApplicationNo = "APPNO-" + property.UniqueID

		//Save Owners
		if property.Owner != nil {
			for k := range property.Owner {
				property.Owner[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYOWNER)
				property.OwnerID = append(property.OwnerID, property.Owner[k].UniqueID)
				property.Owner[k].Status = constants.PROPERTYSTATUSACTIVE
				property.Owner[k].PropertyID = property.UniqueID
				property.Owner[k].Created = created
			}
			dberr := s.Daos.SavePropertyOwnerV2(ctx, database, &sc, property.Owner)
			if dberr != nil {
				return errors.New("Transaction Aborted <property owner> - " + dberr.Error())
			}
		}

		//Save Property
		property.Status = constants.PROPERTYSTATUSPENDING
		property.Created = created
		dberr := s.Daos.SavePropertyV2(ctx, database, &sc, property)
		if dberr != nil {
			return errors.New("Transaction Aborted  <property> - " + dberr.Error())
		}

		//Save Floors
		if property.Floors != nil {
			for k := range property.Floors {
				property.Floors[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFLOOR)
				property.Floors[k].Status = constants.PROPERTYFLOORSTATUSACTIVE
				property.Floors[k].Created = created
			}
			dberr = s.Daos.SavePropertyFloorsV2(ctx, database, &sc, property.Floors)
			if dberr != nil {
				return errors.New("Transaction Aborted <property floor> - " + dberr.Error())
			}
		}
		if property.PropertyDocument != nil {
			for v := range property.PropertyDocument {
				fmt.Println("Updating doc")
				property.PropertyDocument[v].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYDOCUMENT)
				property.PropertyDocument[v].Status = constants.PROPERTYDOCUMENTSTATUSACTIVE
				property.PropertyDocument[v].Created = created
				property.PropertyDocument[v].PropertyID = property.UniqueID
			}
			dberr = s.Daos.SavePropertyDocumentv2(ctx, database, &sc, property.PropertyDocument)
			if dberr != nil {
				return errors.New("Transaction Aborted <property floor> - " + dberr.Error())
			}
		}
		if err := session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//SaveProperty :""
func (s *Service) SaveProperty(ctx *models.Context, property *models.Property, collectionName string) error {
	log.Println("transaction start")
	// contextCreated := context.TODO()
	// c, db := s.Daos.GetDBV4(contextCreated)
	// ctx = new(models.Context)
	// ctx.DB = db
	// ctx.CTX = contextCreated
	// session, err := c.StartSession()
	// if err != nil {
	// 	panic(err)
	// }
	defer func() {
		log.Println("transaction completed")
		ctx.Session.EndSession(ctx.CTX)
	}()
	client := s.Daos.GetDBV3(context.TODO())
	defer client.Disconnect(context.TODO())

	database := client.Database(s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + ".database_name"))
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		ctx.SC = sc
		//Start Transaction
		wc := writeconcern.New(writeconcern.WMajority())
		rc := readconcern.Snapshot()
		txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
		if err := ctx.Session.StartTransaction(txnOpts); err != nil {
			return err
		}

		// ctx.CTX = sc

		t := time.Now()
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM

		if property.Created.By != "" {
			created.By = property.Created.By
		}
		//
		resPC, err := s.GetSingleProductConfiguration(ctx)
		if err != nil {
			return errors.New("error in getting the product configutaion" + err.Error())
		}
		if collectionName == "estimatedpropertydemand" {
			property.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONESTIMATEDPROPERTYDEMAND)
			//Save Property
			property.Status = constants.ESTIMATEDPROPERTYDEMANDSTATUSACTIVE

			property.Created = created
			dberr := s.Daos.SaveEstimatedPropertyDemand(ctx, property, collectionName)
			if dberr != nil {
				return errors.New("Transaction Aborted  <property> - " + dberr.Error())
			}
			// Save Estimated Floors
			if property.Floors != nil {
				if len(property.Floors) > 0 {
					for k := range property.Floors {
						property.Floors[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONESTIMATEDPROPERTYFLOOR)
						property.Floors[k].PropertyID = property.UniqueID
						property.Floors[k].Status = constants.ESTIMATEDPROPERTYFLOORSTATUSACTIVE
						property.Floors[k].Created = created
					}
					dberr = s.Daos.SaveEstimatedFloors(ctx, property.Floors)
					if dberr != nil {
						return errors.New("Transaction Aborted <property floor> - " + dberr.Error())
					}
				}
			}
		} else {
			//

			if resPC.LocationID == "Bhagalpur" || resPC.LocationID == "Bakhtiyarpur" {
				resWards, err := s.GetSingleWard(ctx, property.Address.WardCode)
				if err != nil {
					return errors.New("error in getting the ward" + err.Error())
				}
				var tempBuildUpArea float64
				tempConstructionType := "15"
				for _, v := range property.Floors {
					if v.No == "16" {
						if v.BuildUpArea > tempBuildUpArea {
							tempBuildUpArea = v.BuildUpArea
							tempConstructionType = v.ConstructionType
						}
					}
				}
				resConstruction, err := s.GetSingleConstructionType(ctx, tempConstructionType)
				if err != nil {
					return errors.New("error in getting the constructed type" + err.Error())
				}
				resRoadType, err := s.GetSingleRoadType(ctx, property.RoadTypeID)
				if err != nil {
					return errors.New("error in getting the road type" + err.Error())
				}
				resPropertyType, err := s.GetSinglePropertyType(ctx, property.PropertyTypeID)
				if err != nil {
					return errors.New("error in getting the property type" + err.Error())
				}
				fmt.Println("Ward Name =========>", resWards.Ward.Name)
				fmt.Println("Property RoadType ======>", resRoadType.Label)
				fmt.Println("construction Type ======>", resConstruction.Label)
				fmt.Println("No of Floors ======>", len(property.Floors))
				fmt.Println("property Type ======>", resPropertyType.Label)
				fmt.Println("HouseHold No ======>", property.UniqueID)
				tempFloors := len(property.Floors) - 1
				tempUniqueId := s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTY)
				var varPrefix string
				if resPC.LocationID == "Bhagalpur" {
					varPrefix = "BMC"
				} else if resPC.LocationID == "Bakhtiyarpur" {
					varPrefix = "BNP"
				}
				property.UniqueID = fmt.Sprintf("%v%v%v%v%d%v%v", varPrefix, resWards.Ward.Name,
					resRoadType.Label,
					resConstruction.Label,
					tempFloors,
					resPropertyType.Label,
					tempUniqueId,
				)

			} else {
				//
				property.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTY)
			}
			// property.UniqueID = fmt.Sprintf("BMC%v%v%d%v", property.Address.WardCode,
			// 	property.RoadTypeID,
			// 	// property.Construct
			// 	len(property.Floors),
			// 	tempUniqueId)
			property.ApplicationNo = "APPNO-" + property.UniqueID
			//Save Owners
			if property.Owner != nil {
				for k := range property.Owner {
					property.Owner[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYOWNER)
					property.OwnerID = append(property.OwnerID, property.Owner[k].UniqueID)
					property.Owner[k].Status = constants.PROPERTYOWNERSTATUSACTIVE
					property.Owner[k].PropertyID = property.UniqueID
					property.Owner[k].Created = created
				}
			}

			dberr := s.Daos.SavePropertyOwner(ctx, property.Owner)
			if dberr != nil {
				return errors.New("Transaction Aborted <property owner> - " + dberr.Error())
			}

			//Save Property
			property.Status = constants.PROPERTYSTATUSPENDING

			property.Created = created

			dberr = s.Daos.SaveProperty(ctx, property, "")
			if dberr != nil {
				return errors.New("Transaction Aborted  <property> - " + dberr.Error())
			}

			//Save Floors
			if property.Floors != nil {
				if len(property.Floors) > 0 {
					for k := range property.Floors {
						property.Floors[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFLOOR)
						property.Floors[k].PropertyID = property.UniqueID
						property.Floors[k].Status = constants.PROPERTYFLOORSTATUSACTIVE
						property.Floors[k].Created = created
					}
					dberr = s.Daos.SavePropertyFloors(ctx, property.Floors)
					if dberr != nil {
						// if err1 := sc.AbortTransaction(sc); err1 != nil {
						// 	log.Println("err in abort  <property floor>")
						// 	return errors.New("Transaction Aborted <property floor> with error" + err1.Error())
						// }
						// log.Println("err in abort out <property floor>")
						return errors.New("Transaction Aborted <property floor> - " + dberr.Error())
					}
				}
			}

			// Save Legacy
			if property.Legacy.IsLegacy {
				if property.Legacy.LegacyProperty != nil {
					property.Legacy.LegacyProperty.LegacyProperty.PropertyID = property.UniqueID
					s.PreSaveLegacy(ctx, property.Legacy.LegacyProperty)

					dberr2 := s.Daos.SaveLegacy(ctx, property.Legacy.LegacyProperty)
					if dberr2 != nil {
						return dberr2
					}
				}
			}

			// Save Mobile Tower
			// if property.MobileTower.IsMobileTower {
			// 	if property.MobileTower.PropertyMobileTower != nil {
			// 		property.MobileTower.PropertyMobileTower.PropertyID = property.UniqueID

			// 		s.PreSaveMobileTower(ctx, property.MobileTower.PropertyMobileTower)

			// 		dberr2 := s.Daos.SavePropertyMobileTower(ctx, property.MobileTower.PropertyMobileTower)
			// 		if dberr2 != nil {
			// 			return dberr2
			// 		}
			// 	}
			// }
			if property.PropertyDocument != nil {
				for v := range property.PropertyDocument {
					fmt.Println("Updating doc")
					property.PropertyDocument[v].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYDOCUMENT)
					property.PropertyDocument[v].Status = constants.PROPERTYDOCUMENTSTATUSACTIVE
					property.PropertyDocument[v].Created = created
					property.PropertyDocument[v].PropertyID = property.UniqueID
					dberr = s.Daos.SavePropertyDocumentv2(ctx, database, &sc, property.PropertyDocument)
					if dberr != nil {
						return errors.New("Transaction Aborted <property floor> - " + dberr.Error())
					}
				}

			}
		}
		if resPC.LocationID == "Bhagalpur" {
			resDetails, err := s.CheckWhetherPenalChargeApplies(ctx, property.DOA)
			if err != nil {
				return err
			}
			if resDetails.PenalChargeStatus == "Yes" {
				t := time.Now()
				pod := new(models.PropertyOtherDemand)
				pod.FyID = resDetails.FyID
				pod.PropertyID = property.UniqueID
				pod.Created.By = constants.SYSTEM
				pod.Reason = "PENAL CHARGES"
				pod.OneTimePenalCharges = "Yes"
				pod.Created.On = &t
				pod.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYOTHERDEMAND)
				pod.Status = constants.PROPERTYOTHERDEMANDSTATUSACTIVE
				pod.PaymentStatus = constants.PROPERTYOTHERDEMANDPAYMENTSTATUSNOTPAID
				Created := new(models.CreatedV2)
				Created.On = &t
				pod.Created = *Created

				switch property.PropertyTypeID {
				case "1":
					pod.Amount = 2000
				case "3":
					pod.Amount = 5000
				case "2":
					pod.Amount = 2000
				case "15":
					pod.Amount = 5000

				}

				err = s.Daos.SavePropertyOtherDemand(ctx, pod)
				if err != nil {
					return errors.New("Transaction Aborted <property floor> - " + err.Error())
				}
			}

		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil

	}); err != nil {
		log.Println("Transaction start aborting" + err.Error())
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//UpdateProperty : ""
func (s *Service) UpdateProperty(ctx *models.Context, property *models.Property) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	// client := s.Daos.GetDBV3(context.TODO())
	// defer client.Disconnect(context.TODO())

	// database := client.Database("municipalproduct1")

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateProperty(ctx, property)
		if err != nil {
			return err

		}
		if len(property.Floors) > 0 {
			for k := range property.Floors {
				if property.Floors[k].UniqueID == "" {
					property.Floors[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFLOOR)
					fmt.Println("THIS IS FLOOR UNIQUEID=====>", property.Floors[k].UniqueID)
					property.Floors[k].PropertyID = property.UniqueID
					property.Floors[k].Status = constants.PROPERTYFLOORSTATUSACTIVE
					property.Floors[k].Created = created

				}
			}
			if err := s.Daos.UpdatePropertyFloorV2(ctx, property.Floors); err != nil {
				return err
			}
			// deletePropertyFLoors := []string{}
			// for _,v:= range property.Floors

		}

		if len(property.Owner) > 0 {
			for k := range property.Owner {
				if property.Owner[k].UniqueID == "" {
					property.Owner[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYOWNER)
					property.Owner[k].PropertyID = property.UniqueID
					property.Owner[k].Status = constants.PROPERTYOWNERSTATUSACTIVE
					property.Owner[k].Created = created
				}
			}
			if err := s.Daos.UpdatePropertyOwnerV2(ctx, property.Owner); err != nil {
				return err
			}
		}
		// if property.PropertyDocument != nil {
		// 	for v := range property.PropertyDocument {
		// 		fmt.Println("Updating doc")
		// 		property.PropertyDocument[v].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYDOCUMENT)
		// 		property.PropertyDocument[v].Status = constants.PROPERTYDOCUMENTSTATUSACTIVE
		// 		property.PropertyDocument[v].Created = created
		// 		property.PropertyDocument[v].PropertyID = property.UniqueID
		// 	}
		// 	dberr := s.Daos.SavePropertyDocumentv2(ctx, database, &sc, property.PropertyDocument)
		// 	if dberr != nil {
		// 		return errors.New("Transaction Aborted <property floor> - " + dberr.Error())
		// 	}
		// }
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

// UpdatePropertyWithOutTransaction : ""
func (s *Service) UpdatePropertyWithOutTransaction(ctx *models.Context, property *models.Property) error {

	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	client := s.Daos.GetDBV3(context.TODO())
	defer client.Disconnect(context.TODO())

	database := client.Database("municipalproduct1")

	err := s.Daos.UpdateProperty(ctx, property)
	if err != nil {
		return err

	}
	if len(property.Floors) > 0 {
		for k := range property.Floors {
			if property.Floors[k].UniqueID == "" {
				property.Floors[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFLOOR)
				fmt.Println("THIS IS FLOOR UNIQUEID=====>", property.Floors[k].UniqueID)
				property.Floors[k].PropertyID = property.UniqueID
				property.Floors[k].Status = constants.PROPERTYFLOORSTATUSACTIVE
				property.Floors[k].Created = created

			}
		}
		if err := s.Daos.UpdatePropertyFloorV2(ctx, property.Floors); err != nil {
			return err
		}

	}

	if len(property.Owner) > 0 {
		for k := range property.Owner {
			if property.Owner[k].UniqueID == "" {
				property.Owner[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYOWNER)
				property.Owner[k].PropertyID = property.UniqueID
				property.Owner[k].Status = constants.PROPERTYOWNERSTATUSACTIVE
				property.Owner[k].Created = created
			}
		}
		if err := s.Daos.UpdatePropertyOwnerV2(ctx, property.Owner); err != nil {
			return err
		}
	}
	if property.PropertyDocument != nil {
		for v := range property.PropertyDocument {
			fmt.Println("Updating doc")
			property.PropertyDocument[v].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYDOCUMENT)
			property.PropertyDocument[v].Status = constants.PROPERTYDOCUMENTSTATUSACTIVE
			property.PropertyDocument[v].Created = created
			property.PropertyDocument[v].PropertyID = property.UniqueID
		}
		dberr := s.Daos.SavePropertyDocumentv2(ctx, database, nil, property.PropertyDocument)
		if dberr != nil {
			return errors.New("Transaction Aborted <property documentv2> - " + dberr.Error())
		}
	}

	return nil

}

//EnableProperty : ""
func (s *Service) UpdatePropertyPreviousYrCollection(ctx *models.Context, ppyc *models.PropertyPreviousYrCollection) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyPreviousYrCollection(ctx, ppyc)
		if err != nil {
			return err
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//EnableProperty : ""
func (s *Service) EnableProperty(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableProperty(ctx, UniqueID)
		if err != nil {
			return err
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//DisableProperty : ""
func (s *Service) DisableProperty(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableProperty(ctx, UniqueID)
		if err != nil {
			return err
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//DeleteProperty : ""
func (s *Service) DeleteProperty(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteProperty(ctx, UniqueID)
		if err != nil {
			return err
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//ActivateProperty : ""
func (s *Service) ActivateProperty(ctx *models.Context, req *models.ActivateProperty) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		data := make(map[string]string)
		data["status"] = constants.PROPERTYSTATUSACTIVE
		req.PropertyTimeline.Type = constants.PROPERTYSTATUSACTIVE
		t := time.Now()
		req.PropertyTimeline.On = &t

		err := s.Daos.PropertyTimelineUpdate(ctx, req.PropertyID, data, req.PropertyTimeline)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//RejectProperty : ""
func (s *Service) RejectProperty(ctx *models.Context, req *models.RejectProperty) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		data := make(map[string]string)
		data["status"] = constants.PROPERTYSTATUSREJECTED
		req.PropertyTimeline.Type = constants.PROPERTYSTATUSREJECTED
		t := time.Now()
		req.PropertyTimeline.On = &t
		err := s.Daos.PropertyTimelineUpdate(ctx, req.PropertyID, data, req.PropertyTimeline)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//GetSingleProperty :""
func (s *Service) GetSingleProperty(ctx *models.Context, UniqueID string) (*models.RefProperty, error) {
	property, err := s.Daos.GetSingleProperty(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return property, nil
}

//FilterProperty :""
func (s *Service) FilterProperty(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) (property []models.RefProperty, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.FilterProperty(ctx, propertyfilter, pagination)

}

// PropertyWiseDemandandCollectionJSON : ""
func (s *Service) PropertyWiseDemandandCollectionJSON(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) (property []models.RefProperty, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.PropertyWiseDemandandCollectionJSON(ctx, propertyfilter, pagination)

}

// ZoneAndWardWiseFilter :""
func (s *Service) ZoneAndWardWiseReport(ctx *models.Context, filter *models.ZoneAndWardWiseReportFilter, pagination *models.Pagination) (property []models.ZoneAndWardWiseReport, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.ZoneAndWardWiseReport(ctx, filter, pagination)
}

// UpdatePropertyGISTagging : ""
func (s *Service) UpdatePropertyGISTagging(ctx *models.Context, UniqueID string, gis *models.PropertyGISTagging) error {
	err := s.Daos.UpdatePropertyGISTagging(ctx, UniqueID, gis)
	if err != nil {
		return err
	}
	return nil
}

// BasicUpdateProperty : ""
func (s *Service) BasicUpdateProperty(ctx *models.Context, bpu *models.BasicPropertyUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		// err := s.Daos.BasicUpdateProperty(ctx, bpu)
		// if err != nil {
		// 	return errors.New("Error in upddating Property" + err.Error())
		// }
		oldPropertyData, err := s.Daos.GetSingleProperty(ctx, bpu.PropertyID)
		if err != nil {
			return errors.New("Error in geting old Property" + err.Error())

		}
		if oldPropertyData == nil {
			return errors.New("property Not Found")
		}
		bpul := new(models.BasicPropertyUpdateLog)
		bpul.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBASICPROPERTYUPDATELOG)
		bpul.PropertyID = bpu.PropertyID
		bpul.Previous.Address = oldPropertyData.Property.Address
		if len(oldPropertyData.Ref.PropertyOwner) > 0 {
			bpul.Previous.Owner = oldPropertyData.Ref.PropertyOwner[0]
		}
		bpul.New.Address = bpu.Address
		bpul.New.Owner = bpu.Owner
		bpul.UserName = bpu.UserName
		bpul.UserType = bpu.UserType
		t := time.Now()
		bpul.Requester = models.Updated{
			On:       &t,
			By:       bpu.UserName,
			Scenario: "BasicUpdate",
			ByType:   bpu.UserType,
			Remarks:  bpu.Remarks,
		}
		bpul.Proof = bpu.Proof
		bpul.Status = constants.PROPERTYBASICUPDATELOGINIT
		err = s.Daos.SaveBasicPropertyUpdateLog(ctx, bpul)
		if err != nil {
			return errors.New("Error in upddating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}

		templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
		//html template path
		templateID := templatePathStart + "PropertyUpdateRequestEmail.html"
		templateID = "templates/PropertyUpdateRequestEmail.html"

		//sending email
		if err := s.SendEmailWithTemplate("Property Update Request - holding no 1111", []string{"solomon2261993@gmail.com"}, templateID, nil); err != nil {
			log.Println("email not sent - ", err.Error())
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

// AcceptBasicPropertyUpdate : ""
func (s *Service) AcceptBasicPropertyUpdate(ctx *models.Context, req *models.AcceptBasicPropertyUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//Finding the request
		rbpu, err := s.Daos.GetSingleBasicPropertyUpdateLog(ctx, req.UniqueID)
		if err != nil {
			return errors.New("Not able to find the request" + err.Error())
		}
		if rbpu == nil {
			return errors.New("Request in nil")
		}
		//updating property
		bpu := new(models.BasicPropertyUpdate)
		bpu.PropertyID = rbpu.PropertyID
		bpu.Address = rbpu.New.Address
		bpu.Owner = rbpu.New.Owner
		bpu.UserName = rbpu.Requester.By
		bpu.UserType = rbpu.Requester.ByType
		bpu.Proof = rbpu.Proof
		bpu.Remarks = rbpu.Requester.Remarks
		err = s.Daos.BasicUpdateProperty(ctx, bpu)
		if err != nil {
			return errors.New("Error in upddating Property" + err.Error())
		}
		//Updating Receipts
		//get current Financial year

		// cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
		// if err != nil {
		// 	return errors.New("Error in getting current financial year " + err.Error())
		// }
		// if cfy == nil {
		// 	return errors.New("current financial year is nil")
		// }
		// ppf := new(models.PropertyPaymentFilter)
		// ppf.PropertyIds = append(ppf.PropertyIds, rbpu.PropertyID)
		// ppf.Status = append(ppf.Status, constants.PROPERTYPAYMENTCOMPLETED)
		// ppf.DateRange = new(models.DateRange)
		// ppf.DateRange.From = cfy.From
		// ppf.DateRange.To = cfy.To
		// refPayments, err := s.Daos.FilterPropertyPayment(ctx, ppf, nil)
		// if err != nil {
		// 	return errors.New("Error in getting payments " + err.Error())
		// }
		// if len(refPayments) > 0 {
		// 	fmt.Println("payments updated - ", len(refPayments))
		// 	tnxIds := []string{}
		// 	for _,v := range refPayments {
		// 		tnxIds=append(tnxIds, v.)
		// 	}
		// }
		//updating the request
		err = s.Daos.AcceptBasicPropertyUpdate(ctx, req)
		if err != nil {
			return nil
		}
		err = s.Daos.BasicPropertyUpdateToPayments(ctx, rbpu)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// RejectBasicPropertyUpdate : ""
func (s *Service) RejectBasicPropertyUpdate(ctx *models.Context, req *models.RejectBasicPropertyUpdate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		data := make(map[string]string)
		data["status"] = constants.BASICPROPERTYUPDATESTATUSDISABLED

		err := s.Daos.RejectBasicPropertyUpdate(ctx, req)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//FilterBasicPropertyUpdate :""
func (s *Service) FilterBasicPropertyUpdate(ctx *models.Context, filter *models.FilterBasicPropertyUpdate, pagination *models.Pagination) (property []models.RefBasicPropertyUpdateLog, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.FilterBasicPropertyUpdate(ctx, filter, pagination)
}

// FilterPropertyExcel : ""
func (s *Service) FilterPropertyExcel(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterProperty(ctx, filter, pagination)
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
	sheet1 := "Property"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "P3")
		excel.MergeCell(sheet1, "A4", "P5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "P3")
		excel.MergeCell(sheet1, "C4", "P5")
	}
	style, err := excel.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
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
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Property List")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property List")
	}

	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "P", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Application No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Zone")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "District")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Owner")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Mobile Number")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Application Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Status")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Old Holding Number")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Is Matched")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Assessment By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), "Activated By")
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Plot Area")

	fmt.Println("'res length==>'", len(res))
	for i, v := range res {
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.ApplicationNo)
		if v.Ref.Address.Zone != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Ref.Address.Zone.Name)
		}
		if v.Ref.Address.District != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Ref.Address.District.Name)
		}
		if v.Ref.Address.Ward != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Ref.Address.Ward.Name)
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Address.AL1+v.Address.Al2)
		if v.Created.On != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Created.On.Format("2006-01-02"))
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Status)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Ref.PropertyType.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.OldHoldingNumber)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), func() string {
			if v.IsMatched != "" {
				return v.IsMatched
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), func() string {
			if v.Ref.User != nil {
				// if v.Ref.User.Name != "" {
				return v.Ref.User.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), func() string {
			if v.Ref.Activator.Name != "" {
				return v.Ref.Activator.Name
			}
			return "NA"
		}())
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), v.AreaOfPlot)

	}

	return excel, nil
}

// FilterPropertyExcelV2 : ""
func (s *Service) FilterPropertyExcelV2(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterProperty(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Property"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "L3")
	excel.MergeCell(sheet1, "C4", "L5")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property List")
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Guardian Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "From Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "To Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Arrear Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Current Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Total Demand")

	var arrearDemand, currentDemand, totalDemand float64

	fmt.Println("'res length==>'", len(res))
	for i, v := range res {
		arrearDemand = arrearDemand + v.Demand.Arrear
		currentDemand = currentDemand + v.Demand.Current
		totalDemand = totalDemand + v.Demand.TotalTax
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].FatherRpanRhusband
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Demand.FromYear.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Demand.ToYear.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Demand.TotalPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Demand.Arrear)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Demand.Current)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Demand.TotalTax)

	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style4)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%.2f", arrearDemand))

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%.2f", currentDemand))

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%.2f", totalDemand))

	return excel, nil
}

// FilterPropertyExcelV3 : ""
func (s *Service) FilterPropertyExcelV3(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterProperty(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Property"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "L3")
	excel.MergeCell(sheet1, "C4", "L5")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property List")
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Floors")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Paid")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Dues")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Total Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "New Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Old Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Address")

	var arrearDemand, currentDemand, totalDemand float64

	fmt.Println("'res length==>'", len(res))
	for i, v := range res {
		arrearDemand = arrearDemand + v.Demand.Arrear
		currentDemand = currentDemand + v.Demand.Current
		totalDemand = totalDemand + v.Demand.TotalTax
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), len(v.Ref.Floors))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() float64 {
			return (v.Collection.TotalTax - v.Collection.Penalty)
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() float64 {
			return (v.Ref.Demand.Total.TotalTax - (v.Collection.TotalTax - v.Collection.Penalty))
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Ref.Demand.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.OldHoldingNumber)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Address.AL1+" "+v.Address.Al2)

	}
	rowNo++
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style4)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style4)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style4)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style4)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style4)

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%.2f", arrearDemand))

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%.2f", currentDemand))

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%.2f", totalDemand))

	return excel, nil
}

// FilterPropertyPdf : ""
func (s *Service) FilterPropertyPdf(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) ([]byte, error) {
	r := NewRequestPdfV2("", "Landscape")
	res, err := s.FilterProperty(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "samplepdf.html"
	var fm = template.FuncMap{
		"add": func(a int) int {
			return a + 1

		},
	}
	// for k := range res {
	// 	res[k].Ref.Inc = func(a int) int {
	// 		return a + 1

	// 	}
	// }
	err = r.ParseTemplatev2(templatePath, fm, res)
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

// PropertyWiseDemandandCollectionExcel: ""
func (s *Service) PropertyWiseDemandandCollectionExcel(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.PropertyWiseDemandandCollectionJSON(ctx, filter, pagination)
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
	sheet1 := "Property Wise Demand and Collection"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "X3")
		excel.MergeCell(sheet1, "A4", "X5")
		excel.MergeCell(sheet1, "A6", "X6")
		excel.MergeCell(sheet1, "A7", "X7")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "X3")
		excel.MergeCell(sheet1, "C4", "X5")
		excel.MergeCell(sheet1, "A6", "X6")
		excel.MergeCell(sheet1, "A7", "X7")
	}

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++
	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Property Wise Demand And Collection")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Wise Demand And Collection")
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")

	reportFromMsg2 := "Report"
	if filter != nil {
		if filter.AppliedRange != nil {
			if filter.AppliedRange.From != nil && filter.AppliedRange.To == nil {
				reportFromMsg2 = reportFromMsg2 + " on " + fmt.Sprintf("%v-%v-%v", filter.AppliedRange.From.Day(), filter.AppliedRange.From.Month(), filter.AppliedRange.From.Year())
			}
			if filter.AppliedRange.From != nil && filter.AppliedRange.To != nil {
				reportFromMsg2 = reportFromMsg2 + " From " + fmt.Sprintf("%v-%v-%v", filter.AppliedRange.From.Day(), filter.AppliedRange.From.Month(), filter.AppliedRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.AppliedRange.To.Day(), filter.AppliedRange.To.Month(), filter.AppliedRange.To.Year())
			}
			if filter.AppliedRange.From == nil && filter.AppliedRange.To == nil {
				fmt.Println("everything is nil")
			}
		}
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "X", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "SAF No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Father/Husband Name/PAN")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Road Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Property Type")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Usage Type")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Property Category")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Plot Area")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Arrear Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Current Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Form Fee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), "Advance")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Total Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), "Arrear Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), "Current Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), "Rebate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), "Total Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), "Arrear Outstanding Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), "Current Outstanding Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), "Total Outstanding Demand")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Total Demand")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Arrear Collection")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), "Current Collection")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Total Collection")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), "Penalty")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), "Rebate")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), "Advance")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), "Arrear Outstanding Demand")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), "Current Outstanding Demand")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), "Total Outstanding Demand")

	fmt.Println("'res length==>'", len(res))
	var arrearDemand, currentDemand, totalDemand float64
	var arrearCollection, currentCollection, totalCollection float64
	var arrearOutstandingDemand, currentOutstandingDemand, totalOutstandingDemand float64

	for i, v := range res {

		v.Ref.Demand.Total.TotalTax = v.Ref.Demand.Total.TotalTax - v.Ref.Demand.Total.Ecess
		v.Ref.Collections.TotalTax = v.Ref.Collections.TotalTax + v.Ref.PropertyPayments.FormFee
		v.Ref.Demand.Current.TotalTax = v.Ref.Demand.Current.TotalTax + v.Ref.PropertyPayments.Rebate

		arrearDemand = arrearDemand + v.Ref.Demand.Arrear.TotalTax
		currentDemand = currentDemand + v.Ref.Demand.Current.TotalTax
		totalDemand = totalDemand + v.Ref.Demand.Total.TotalTax
		arrearCollection = arrearCollection + v.Ref.Collections.ArrearTax
		currentCollection = currentCollection + v.Ref.Collections.CurrentTax
		totalCollection = totalCollection + v.Ref.Collections.TotalTax
		arrearOutstandingDemand = arrearOutstandingDemand + (v.Ref.Demand.Arrear.TotalTax - v.Ref.Collections.ArrearTax - v.Ref.Collections.ArrearPenalty)
		currentOutstandingDemand = currentOutstandingDemand + (v.Ref.Demand.Current.TotalTax - v.Ref.Collections.CurrentTax - v.Ref.Collections.CurrentPenalty)
		totalOutstandingDemand = totalOutstandingDemand + (v.Ref.Demand.Total.TotalTax - v.Ref.Collections.TotalTax)

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Property.ApplicationNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Property.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				if len(v.Ref.PropertyOwner) > 0 {
					return v.Ref.PropertyOwner[0].Name
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				if len(v.Ref.PropertyOwner) > 0 {
					return v.Ref.PropertyOwner[0].FatherRpanRhusband
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				if len(v.Ref.PropertyOwner) > 0 {
					return v.Ref.PropertyOwner[0].Mobile
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Ref.RoadType != nil {
				return v.Ref.RoadType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if v.Ref.PropertyType != nil {
				return v.Ref.PropertyType.Name
			}
			return "NA"
		}())
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Usage Type")
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Property Category")
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Address.AL1+v.Address.Al2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.AreaOfPlot)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Demand.Arrear)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Demand.Current)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.Demand.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Ref.Demand.Arrear.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Ref.Demand.Current.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.Ref.PropertyPayments.FormFee)

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), (v.Ref.Collections.ArrearPenalty + v.Ref.Collections.CurrentPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), v.Advance)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), v.Ref.Demand.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), v.Ref.Collections.ArrearTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), v.Ref.Collections.CurrentTax)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), (v.Ref.Collections.ArrearRebate + v.Ref.Collections.CurrentRebate))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), v.Ref.PropertyPayments.Rebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), v.Ref.Collections.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), (v.Ref.Demand.Arrear.TotalTax - v.Ref.Collections.ArrearTax - v.Ref.Collections.ArrearPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), (v.Ref.Demand.Current.TotalTax - v.Ref.Collections.CurrentTax - v.Ref.Collections.CurrentPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), (v.Ref.Demand.Total.TotalTax - v.Ref.Collections.TotalTax))

	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "R", rowNo), fmt.Sprintf("%v%v", "R", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "S", rowNo), fmt.Sprintf("%v%v", "S", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "T", rowNo), fmt.Sprintf("%v%v", "T", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "U", rowNo), fmt.Sprintf("%v%v", "U", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "V", rowNo), fmt.Sprintf("%v%v", "V", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "W", rowNo), fmt.Sprintf("%v%v", "W", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "X", rowNo), fmt.Sprintf("%v%v", "X", rowNo), style4)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%.2f", arrearDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%.2f", currentDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), fmt.Sprintf("%v%v", "Q", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), fmt.Sprintf("%.2f", totalDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "R", rowNo), fmt.Sprintf("%v%v", "R", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), fmt.Sprintf("%.2f", arrearCollection))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "S", rowNo), fmt.Sprintf("%v%v", "S", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), fmt.Sprintf("%.2f", currentCollection))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "U", rowNo), fmt.Sprintf("%v%v", "U", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), fmt.Sprintf("%.2f", totalCollection))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "V", rowNo), fmt.Sprintf("%v%v", "V", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), fmt.Sprintf("%.2f", arrearOutstandingDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "W", rowNo), fmt.Sprintf("%v%v", "W", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), fmt.Sprintf("%.2f", currentOutstandingDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "X", rowNo), fmt.Sprintf("%v%v", "X", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), fmt.Sprintf("%.2f", totalOutstandingDemand))
	return excel, nil
}

// WardwiseDemand : ""
func (s *Service) WardwiseDemand(ctx *models.Context, propertyfilter *models.PropertyWardwiseDemandFilter, pagination *models.Pagination) ([]models.WardwiseDemandandCollection, error) {
	return s.Daos.WardwiseDemandandCollection(ctx, propertyfilter, pagination)
}

// WardwiseDemandExcel : ""
func (s *Service) WardwiseDemandExcel(ctx *models.Context, propertyfilter *models.PropertyWardwiseDemandFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.WardwiseDemand(ctx, propertyfilter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	// create an excel file
	excel := excelize.NewFile()
	sheet1 := "Ward Wise Demand"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "E3")
	excel.MergeCell(sheet1, "C4", "E5")
	excel.MergeCell(sheet1, "A6", "E6")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Ward Wise Demand")
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Arrear")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Current")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Total")

	var arrearDemand, currentDemand, totalDemand float64

	for i, v := range res {
		arrearDemand = arrearDemand + v.Properties.TotalDemandArrear
		currentDemand = currentDemand + v.Properties.TotalDemandCurrent
		totalDemand = totalDemand + v.Properties.TotalDemandTax

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Properties.TotalDemandArrear)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Properties.TotalDemandCurrent)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Properties.TotalDemandTax)
	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%.2f", arrearDemand))

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%.2f", currentDemand))

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%.2f", totalDemand))
	return excel, nil
}

// WardwiseDemand : ""
func (s *Service) WardwiseCollection(ctx *models.Context, propertyfilter *models.PropertyWardwiseDemandFilter, pagination *models.Pagination) ([]models.WardwiseDemandandCollection, error) {
	return s.Daos.WardwiseDemandandCollection(ctx, propertyfilter, pagination)
}

// WardwiseCollectionExcel : ""
func (s *Service) WardwiseCollectionExcel(ctx *models.Context, propertyfilter *models.PropertyWardwiseDemandFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.WardwiseDemand(ctx, propertyfilter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	// create an excel file
	excel := excelize.NewFile()
	sheet1 := "Ward Wise Collection"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "E3")
	excel.MergeCell(sheet1, "C4", "E5")

	style, err := excel.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Ward Wise Collection")
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Arrear")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Current")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Total")

	for i, v := range res {
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Properties.TotalCollectionArrear)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Properties.TotalCollectionCurrent)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Properties.TotalCollectionTax)
	}

	return excel, nil
}

// FilterDemandExcel : ""
func (s *Service) PropertyDemandExcel(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterProperty(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Property"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "L3")
	excel.MergeCell(sheet1, "C4", "L5")
	excel.MergeCell(sheet1, "A6", "D6")
	excel.MergeCell(sheet1, "E6", "H6")
	excel.MergeCell(sheet1, "I6", "L6")
	excel.MergeCell(sheet1, "A7", "D8")
	excel.MergeCell(sheet1, "E7", "H8")
	excel.MergeCell(sheet1, "I7", "L8")

	style, err := excel.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property List")
	rowNo++
	rowNo++

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total Arrear")
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Total Current")
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Total Demand")
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style)

	rowNo++
	totalRowNumber := rowNo
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Arrear Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Current Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Total Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Application No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "District")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Owner")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Application Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Type")

	fmt.Println("'res length==>'", len(res))
	var arrear, current, totalTax float64
	for i, v := range res {
		arrear = arrear + v.Demand.Arrear
		current = current + v.Demand.Current
		totalTax = totalTax + v.Demand.TotalTax
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Demand.Arrear)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Demand.Current)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Demand.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.ApplicationNo)
		if v.Ref.Address.District != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Ref.Address.District.Name)
		}
		if v.Ref.Address.Ward != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Ref.Address.Ward.Name)
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if len(v.Owner) > 0 {
				return v.Owner[0].Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Address.AL1+v.Address.Al2)
		if v.Created.On != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Created.On.Format("2006-02-01"))
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Ref.PropertyType.Name)

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", totalRowNumber), fmt.Sprintf("%v%v", "L", totalRowNumber), style)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalRowNumber), arrear)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", totalRowNumber), current)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", totalRowNumber), totalTax)

	return excel, nil
}

//GetPropertyDemandCalc : ""
func (s *Service) GetPropertyDemandCalc(ctx *models.Context, filter *models.PropertyDemandFilter, collectionName string) (*models.PropertyDemand, error) {
	var data *models.PropertyDemand
	var err error
	if collectionName == constants.COLLECTIONESTIMATEDPROPERTYDEMAND {
		data, err = s.Daos.GetPropertyDemandCalc(ctx, filter, collectionName)
		if err != nil {
			return nil, err
		}
	} else {
		data, err = s.Daos.GetPropertyDemandCalc(ctx, filter, "")
		if err != nil {
			return nil, err
		}
	}
	if data != nil {
		if filter != nil {
			if filter.AllDemand {
				data.AllDemand = true
			}
		}
		pdc, err := s.Daos.GetSingleProductConfiguration(ctx, "1")
		if err != nil {
			return nil, errors.New("Error in geting prod config - " + err.Error())
		}
		if pdc == nil {
			return nil, errors.New("Prod config  is nil ")
		}
		data.ProductConfiguration = pdc
		data.CTX = ctx
		data.DemandCalculation()
	}
	data.OverallPropertyDemand.Current.Ecess = 0
	data.OverallPropertyDemand.Arrear.Ecess = 0
	data.OverallPropertyDemand.Total.VacantLandTax = data.OverallPropertyDemand.Current.VacantLandTax + data.OverallPropertyDemand.Arrear.VacantLandTax
	data.OverallPropertyDemand.Total.Rebate = data.OverallPropertyDemand.Current.Rebate + data.OverallPropertyDemand.Arrear.Rebate
	data.OverallPropertyDemand.Total.Penalty = data.OverallPropertyDemand.Current.Penalty + data.OverallPropertyDemand.Arrear.Penalty
	data.OverallPropertyDemand.Total.Tax = data.OverallPropertyDemand.Current.Tax + data.OverallPropertyDemand.Arrear.Tax
	data.OverallPropertyDemand.Total.CompositeTax = data.OverallPropertyDemand.Current.CompositeTax + data.OverallPropertyDemand.Arrear.CompositeTax
	data.OverallPropertyDemand.Total.Ecess = data.OverallPropertyDemand.Current.Ecess + data.OverallPropertyDemand.Arrear.Ecess
	data.OverallPropertyDemand.Total.PanelCh = data.OverallPropertyDemand.Current.PanelCh + data.OverallPropertyDemand.Arrear.PanelCh

	data.OverallPropertyDemand.Total.Other = data.OverallPropertyDemand.Other.BoreCharge + data.OverallPropertyDemand.Other.FormFee
	data.OverallPropertyDemand.Total.TotalTax = data.OverallPropertyDemand.Total.VacantLandTax + data.OverallPropertyDemand.Total.Tax + data.OverallPropertyDemand.Total.Penalty + data.OverallPropertyDemand.Total.Other - data.OverallPropertyDemand.Total.Rebate + data.OverallPropertyDemand.Total.CompositeTax + data.OverallPropertyDemand.Total.Ecess + data.OverallPropertyDemand.Total.PanelCh
	data.OverallPropertyDemand.Current.TotalTax = data.OverallPropertyDemand.Current.VacantLandTax + data.OverallPropertyDemand.Current.Tax + data.OverallPropertyDemand.Current.Penalty - data.OverallPropertyDemand.Current.Rebate + data.OverallPropertyDemand.Current.CompositeTax + data.OverallPropertyDemand.Current.Ecess + data.OverallPropertyDemand.Current.PanelCh
	data.OverallPropertyDemand.Arrear.TotalTax = data.OverallPropertyDemand.Arrear.VacantLandTax + data.OverallPropertyDemand.Arrear.Tax + data.OverallPropertyDemand.Arrear.Penalty - data.OverallPropertyDemand.Arrear.Rebate + data.OverallPropertyDemand.Arrear.CompositeTax + data.OverallPropertyDemand.Arrear.Ecess + data.OverallPropertyDemand.Arrear.PanelCh

	data.OverallPropertyDemand.Actual.Total.VacantLandTax = data.OverallPropertyDemand.Actual.Total.VacantLandTax + (data.OverallPropertyDemand.Actual.Current.VacantLandTax + data.OverallPropertyDemand.Actual.Arrear.VacantLandTax)
	data.OverallPropertyDemand.Actual.Total.Tax = data.OverallPropertyDemand.Actual.Total.Tax + (data.OverallPropertyDemand.Actual.Current.Tax + data.OverallPropertyDemand.Actual.Arrear.Tax)
	data.OverallPropertyDemand.Actual.Total.TotalTax = data.OverallPropertyDemand.Actual.Total.TotalTax + (data.OverallPropertyDemand.Actual.Current.TotalTax + data.OverallPropertyDemand.Actual.Arrear.TotalTax)
	return data, nil
}

// GetPropertyDemandCalcNotify : ""
func (s *Service) GetPropertyDemandCalcNotify(ctx *models.Context, filter *models.PropertyDemandFilter, notifyType string) error {
	resPropertyDemand, err := s.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		return err
	}

	if resPropertyDemand == nil {
		return errors.New("property demand is nil")
	}
	if len(resPropertyDemand.Ref.PropertyOwner) == 0 {
		return errors.New("property owner is nil")
	}
	if len(resPropertyDemand.Ref.PropertyOwner) > 0 {
		if resPropertyDemand.Ref.PropertyOwner[0].Mobile == "" {
			return errors.New("mobile number is empty")
		}
	}
	resFYs, err := s.GetCurrentFinancialYear(ctx)
	if err != nil {
		return err
	}
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return errors.New("Error in getting product config" + err.Error())
	}

	// msg := fmt.Sprintf(constants.PROPERTYTAXDEMANDCONTENT, math.Ceil(resPropertyDemand.TotalTax), resPropertyDemand.UniqueID, resFYs.Name)
	msg := fmt.Sprintf(constants.PROPERTYTAXDEMANDCONTENT, math.Ceil(resPropertyDemand.Demand.TotalTax), resPropertyDemand.UniqueID, resFYs.Name)
	fmt.Println("notifyType", notifyType)

	switch notifyType {
	case "SMS":
		if len(resPropertyDemand.Ref.PropertyOwner) > 0 {
			err = s.SendSMS(resPropertyDemand.Ref.PropertyOwner[0].Mobile, fmt.Sprintf(constants.COMMONSMSTEMPLATE, resPropertyDemand.Ref.PropertyOwner[0].Name, productConfig.Name, "Property Tax Demand", msg, "BRMNCP", productConfig.UIURL))
			if err != nil {
				fmt.Println("error in sending message to customer" + err.Error())
			}
		}

	case "EMAIL":
		fmt.Println("Case Email", len(resPropertyDemand.Ref.PropertyOwner))

		if len(resPropertyDemand.Ref.PropertyOwner) > 0 {
			fmt.Println("============> coming inside the condition")
			// err = s.SendEmail(fmt.Sprintf(constants.PROPERTYDEMANDEMAILSUBJECTTOCUSTOMER, resPropertyDemand.UniqueID),
			// 	[]string{resPropertyDemand.Ref.PropertyOwner[0].Email},

			// 	fmt.Sprintf(constants.COMMONEMAILBODYTOCUSTOMER, resPropertyDemand.Ref.PropertyOwner[0].Name, productConfig.Name, "Property Tax Demand", math.Ceil(resPropertyDemand.TotalTax), resPropertyDemand.UniqueID, resFYs.Name, "BRMNCP", productConfig.UIURL))
			// html template path
			d := make(map[string]interface{})
			d["Name"] = resPropertyDemand.Ref.PropertyOwner[0].Name
			d["ProductConfigName"] = productConfig.Name
			d["Regarding"] = "Property Tax Demand"
			d["PropertyNo"] = resPropertyDemand.UniqueID
			d["DemandAmount"] = math.Ceil(resPropertyDemand.Demand.TotalTax)
			d["HoldingNo"] = resPropertyDemand.UniqueID
			d["FY"] = resFYs.Name
			d["BRMNCP"] = "BRMNCP"
			d["Contact"] = productConfig.UIURL
			d["logo"] = productConfig.APIURL + productConfig.Logo
			d["Copyrights"] = productConfig.Copyrights + " " + productConfig.Rights
			t := time.Now().Format("2006-02-01")
			d["Date"] = &t
			templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
			templateID := templatePathStart + "email_template_for_demand.html"
			err := s.SendEmailWithTemplate(fmt.Sprintf(constants.PROPERTYDEMANDEMAILSUBJECTTOCUSTOMER, resPropertyDemand.UniqueID),
				[]string{resPropertyDemand.Ref.PropertyOwner[0].Email},
				templateID,
				d,
			)

			if err != nil {
				fmt.Println("error in sending email to customer" + err.Error())
			}
		}
	case "SMSEMAIL":
		if len(resPropertyDemand.Ref.PropertyOwner) > 0 {
			err = s.SendSMS(resPropertyDemand.Ref.PropertyOwner[0].Mobile, fmt.Sprintf(constants.COMMONSMSTEMPLATE, resPropertyDemand.Ref.PropertyOwner[0].Name, productConfig.Name, "Property Tax Demand", msg, "BRMNCP", productConfig.UIURL))
			if err != nil {
				fmt.Println("error in sending message to customer" + err.Error())
			}
		}
		if len(resPropertyDemand.Ref.PropertyOwner) > 0 {
			// old email format =======>
			// err = s.SendEmail(fmt.Sprintf(constants.PROPERTYDEMANDEMAILSUBJECTTOCUSTOMER, resPropertyDemand.UniqueID),
			// 	[]string{resPropertyDemand.Ref.PropertyOwner[0].Email},
			// 	fmt.Sprintf(constants.COMMONEMAILBODYTOCUSTOMER, resPropertyDemand.Ref.PropertyOwner[0].Name, productConfig.Name, "Property Tax Demand", resFYs.Name, "BRMNCP", productConfig.UIURL))

			// Updated email format with template
			d := make(map[string]interface{})
			d["Name"] = resPropertyDemand.Ref.PropertyOwner[0].Name
			d["ProductConfigName"] = productConfig.Name
			d["Regarding"] = "Property Tax Demand"
			d["PropertyNo"] = resPropertyDemand.UniqueID
			d["DemandAmount"] = math.Ceil(resPropertyDemand.Demand.TotalTax)
			d["HoldingNo"] = resPropertyDemand.UniqueID
			d["FY"] = resFYs.Name
			d["BRMNCP"] = "BRMNCP"
			d["Contact"] = productConfig.UIURL
			d["logo"] = productConfig.APIURL + productConfig.Logo
			d["Copyrights"] = productConfig.Copyrights + " " + productConfig.Rights
			t := time.Now().Format("2006-02-01")
			d["Date"] = &t
			templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
			templateID := templatePathStart + "email_template_for_demand.html"
			err := s.SendEmailWithTemplate(fmt.Sprintf(constants.PROPERTYDEMANDEMAILSUBJECTTOCUSTOMER, resPropertyDemand.UniqueID),
				[]string{resPropertyDemand.Ref.PropertyOwner[0].Email},
				templateID,
				d,
			)
			if err != nil {
				fmt.Println("error in sending email to customer" + err.Error())
			}
			fmt.Println("email body ====> ", constants.COMMONEMAILBODYTOCUSTOMER)
		}

	}

	return nil
}

// func (s *Service) GetPropertyDemandCalcEmail(ctx *models.Context, filter *models.PropertyDemandFilter) error {
// 	return nil, nil
// }

// GetPropertyDemandCalcEmail : ""
func (s *Service) GetPropertyDemandCalcEmail(ctx *models.Context, filter *models.PropertyDemandFilter) error {
	// var err error
	resPropertyDemand, err := s.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		return err
	}
	if resPropertyDemand == nil {
		return errors.New("property demand is nil")
	}
	if len(resPropertyDemand.Ref.PropertyOwner) == 0 {
		return errors.New("property owner is nil")
	}
	if len(resPropertyDemand.Ref.PropertyOwner) > 0 {
		if resPropertyDemand.Ref.PropertyOwner[0].Email == "" {
			return errors.New("email is empty")
		}
	}
	resFYs, err := s.GetSingleFinancialYear(ctx, resPropertyDemand.UniqueID)
	if err != nil {
		return err
	}
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return errors.New("Error in getting product config" + err.Error())
	}
	// msg := fmt.Sprintf(constants.PROPERTYTAXDEMANDCONTENT, resPropertyDemand.TotalTax, resPropertyDemand.UniqueID, resFYs.Name)
	if len(resPropertyDemand.Ref.PropertyOwner) > 0 {
		err = s.SendEmail(fmt.Sprintf(constants.PROPERTYDEMANDEMAILSUBJECTTOCUSTOMER, resPropertyDemand.UniqueID),
			[]string{resPropertyDemand.Ref.PropertyOwner[0].Email},
			fmt.Sprintf(constants.COMMONEMAILBODYTOCUSTOMER, resPropertyDemand.Ref.PropertyOwner[0].Name, productConfig.Name, "Property Tax Demand", resFYs.Name, "BRMNCP", productConfig.UIURL))

		if err != nil {
			fmt.Println("error in sending email to customer" + err.Error())
		}
	}
	return nil
}

// GetPropertyDemandCalcV2 : ""
func (s *Service) GetPropertyDemandCalcV2(ctx *models.Context, filter *models.PropertyDemandFilter) (*models.PropertyDemand, error) {
	var data = new(models.PropertyDemand)
	refProp, err := s.Daos.GetSingleProperty(ctx, filter.PropertyID)
	if err != nil {
		log.Println("error in geting propertyid values in property", err)
		return nil, err
	}
	prop, err := s.Daos.GetSinglePropertyDemandLog(ctx, filter.PropertyID)
	if err != nil {
		log.Println("error in geting Property values in property", err)
		return nil, err
	}
	fy, err := s.Daos.GetPropertyFyDemandLog(ctx, filter.PropertyID)
	if err != nil {
		log.Println("error in geting FYS values in property", err)
		return nil, err
	}
	for _, v := range fy {
		prop.PenalCharge = prop.PenalCharge + v.Penalty
	}
	fmt.Println(len(fy), "-------------------------")
	data.Property = refProp.Property
	data.PropertyDemandLog = *prop
	data.FYs = fy
	return data, nil
}

func monthsCountSince(createdAtTime time.Time) int {
	now := time.Now()
	months := 0
	month := createdAtTime.Month()
	for createdAtTime.Before(now) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextMonth := createdAtTime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

//DemandCalc : ""
func (s *Service) DemandCalc(ctx *models.Context, filter *models.PropertyDemandFilter) (models.Demand, error) {

	var demand models.Demand
	cacheDemand := s.Redis.GetValue(constants.CACHEDASHBOARDDEMANDOVERVIEW)
	if cacheDemand != nil {
		var tempDemand *models.Demand
		err := json.Unmarshal([]byte(cacheDemand.(string)), &tempDemand)
		if err != nil {
			log.Println("json Un-Marshal Cache Data err - ", err)
		}
		if tempDemand != nil {
			log.Println("Geting from Cache")
			return *tempDemand, nil
		}
	}
	properties, err := s.Daos.DemandCalc(ctx, filter)
	if err != nil {
		return demand, err
	}
	if len(properties) > 0 {
		for k, v := range properties {
			v.CTX = ctx
			properties[k] = v.DemandCalculation()
			demand.Current = demand.Current + properties[k].Current
			demand.Arrear = demand.Arrear + properties[k].Arrear
		}
	}
	demand.Total = demand.Current + demand.Arrear
	jsonData, err := json.Marshal(demand)
	if err != nil {
		log.Println("error in convarting data to print json")
	}
	log.Println("Updating Cache from Cache")
	if err := s.Redis.SetValue(constants.CACHEDASHBOARDDEMANDOVERVIEW, jsonData, 900); err != nil {
		log.Println("Error in seting cache - " + err.Error())
	}
	return demand, nil
}

//DemandCalc : ""
func (s *Service) DemandCalcV2(ctx *models.Context, filter *models.PropertyDemandFilter) (models.Demand, error) {
	if err := ctx.Session.StartTransaction(); err != nil {
		return models.Demand{}, err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	var demand *models.Demand
	txnOpts := options.Transaction().SetReadPreference(readpref.SecondaryPreferred())
	if _, err := ctx.Session.WithTransaction(ctx.CTX, func(sc mongo.SessionContext) (interface{}, error) {
		properties, err := s.Daos.DemandCalc(ctx, filter)
		if err != nil {
			return nil, err
		}
		if len(properties) > 0 {
			for k, v := range properties {
				v.CTX = ctx
				properties[k] = v.DemandCalculation()
				demand.Current = demand.Current + properties[k].Current
				demand.Arrear = demand.Arrear + properties[k].Arrear
			}
		}
		demand.Total = demand.Current + demand.Arrear
		return nil, nil
	}, txnOpts); err != nil {
		return models.Demand{}, err
	}
	return *demand, nil
}

//DashboardPropertyStatus : ""
func (s *Service) DashboardPropertyStatus(ctx *models.Context, filter *models.DashboardPropertyStatusFilter) (*models.DashboardPropertyStatus, error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.DashboardPropertyStatus(ctx, filter)
}

//GetMultiplePropertyDemandCalc : ""
func (s *Service) GetMultiplePropertyDemandCalc(ctx *models.Context, filter *models.PropertyDemandFilter, pagination *models.Pagination) ([]models.PropertyDemand, error) {
	fmt.Println("inside GetMultiplePropertyDemandCalc is working")
	data, err := s.Daos.GetMultiplePropertyDemandCalc(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	resPD, err := s.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		return nil, err
	}

	if data != nil {
		for k, v := range data {
			v.CTX = ctx
			v.ProductConfiguration = resPD
			data[k] = v.DemandCalculation()
		}

	}
	return data, nil
}

//GetPropertyDemandCalcPDF : ""
func (s *Service) GetPropertyDemandCalcPDF(ctx *models.Context, filter *models.PropertyDemandFilter) ([]byte, error) {
	demand, err := s.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		return nil, err
	}

	_, d := math.Modf(demand.TotalTax)
	if d >= 0.5 {
		demand.TotalTax = math.Ceil(demand.TotalTax)
	} else {
		demand.TotalTax = math.Floor(demand.TotalTax)
	}
	if len(demand.FYs) > 0 {
		for i := range demand.FYs {
			// demand.FYs[i].VacantLandTax = math.Ceil(demand.FYs[i].VacantLandTax)
			// demand.FYs[i].Tax = math.Ceil(demand.FYs[i].Tax)
			// demand.FYs[i].Rebate = math.Ceil(demand.FYs[i].Rebate)
			// demand.FYs[i].Penalty = math.Ceil(demand.FYs[i].Penalty)
			// demand.FYs[i].AlreadyPayed.Amount = math.Ceil(demand.FYs[i].AlreadyPayed.Amount)
			// demand.FYs[i].TotalTax = math.Ceil(demand.FYs[i].TotalTax)
			_, d := math.Modf(demand.FYs[i].TotalTax)
			if d >= 0.5 {
				demand.FYs[i].TotalTax = math.Ceil(demand.FYs[i].TotalTax)
			} else {
				demand.FYs[i].TotalTax = math.Floor(demand.FYs[i].TotalTax)
			}

		}
	}

	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = demand
	fmt.Println(demand.Ref.PropertyOwner)

	state, err := s.Daos.GetSingleState(ctx, demand.Address.StateCode)
	if state != nil {
		m2["state"] = &state.State
	}
	fmt.Println(err)
	district, err := s.Daos.GetSingleDistrict(ctx, demand.Address.DistrictCode)
	if district != nil {
		m2["district"] = &district.District
	}
	fmt.Println(err)
	village, err := s.Daos.GetSingleVillage(ctx, demand.Address.VillageCode)
	if village != nil {
		m2["village"] = &village.Village
	}
	fmt.Println(err)
	zone, err := s.Daos.GetSingleZone(ctx, demand.Address.ZoneCode)
	if zone != nil {
		m2["zone"] = &zone.Zone
	}
	fmt.Println(err)
	ward, err := s.Daos.GetSingleWard(ctx, demand.Address.WardCode)
	if ward != nil {
		m2["ward"] = &ward.Ward
	}
	fmt.Println(err)
	fy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if fy != nil {
		m2["cfy"] = &fy.FinancialYear
	}

	m2["currentDate"] = time.Now()
	m["extraRef"] = m2
	m["Inc"] = func(a int) int {
		return a + 1
	}

	m3["validityDate"] = s.Shared.EndOfMonth(time.Now())
	m["date"] = m3

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "Demand_Page.html"

	//path for download pdf
	// t := time.Now()
	// outputPath := fmt.Sprintf("storage/SampleTemplate%v.pdf", t.Unix())
	fmt.Println(m)
	if err := r.ParseTemplate(templatePath, m); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}
	return nil, nil
}

//GetPropertyDemandCalcPDF : ""
func (s *Service) GetPropertyDemandCalcPDFV2(ctx *models.Context, filter *models.PropertyDemandFilter) ([]byte, error) {
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]string)
	demand, err := s.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		return nil, err
	}
	_, d := math.Modf(demand.TotalTax)
	if d >= 0.5 {
		demand.TotalTax = math.Ceil(demand.TotalTax)
	} else {
		demand.TotalTax = math.Floor(demand.TotalTax)
	}
	m3["IsOneTimePenalCharges"] = "No"
	m2["OneTimePenalCharges"] = 0
	if len(demand.FYs) > 0 {
		for i := range demand.FYs {
			// demand.FYs[i].VacantLandTax = math.Ceil(demand.FYs[i].VacantLandTax)
			// demand.FYs[i].Tax = math.Ceil(demand.FYs[i].Tax)
			// demand.FYs[i].Rebate = math.Ceil(demand.FYs[i].Rebate)
			// demand.FYs[i].Penalty = math.Ceil(demand.FYs[i].Penalty)
			// demand.FYs[i].AlreadyPayed.Amount = math.Ceil(demand.FYs[i].AlreadyPayed.Amount)
			// demand.FYs[i].TotalTax = math.Ceil(demand.FYs[i].TotalTax)

			_, d := math.Modf(demand.FYs[i].TotalTax)
			if d >= 0.5 {
				demand.FYs[i].TotalTax = math.Ceil(demand.FYs[i].TotalTax)
			} else {
				demand.FYs[i].TotalTax = math.Floor(demand.FYs[i].TotalTax)
			}

		}

		for i := range demand.FYs {
			if demand.FYs[i].OtherDemandAdditionalPenalty.OneTimePenalCharges == "Yes" {
				m3["IsOneTimePenalCharges"] = "Yes"
				m2["OneTimePenalCharges"] = demand.FYs[i].OtherDemandAdditionalPenalty.Amount
				break
			}

		}
	}

	m["demand"] = demand
	fmt.Println(demand.Ref.PropertyOwner)

	state, err := s.Daos.GetSingleState(ctx, demand.Address.StateCode)
	if state != nil {
		m2["state"] = &state.State
	}
	fmt.Println(err)
	district, err := s.Daos.GetSingleDistrict(ctx, demand.Address.DistrictCode)
	if district != nil {
		m2["district"] = &district.District
	}
	fmt.Println(err)
	village, err := s.Daos.GetSingleVillage(ctx, demand.Address.VillageCode)
	if village != nil {
		m2["village"] = &village.Village
	}
	fmt.Println(err)
	zone, err := s.Daos.GetSingleZone(ctx, demand.Address.ZoneCode)
	if zone != nil {
		m2["zone"] = &zone.Zone
	}
	fmt.Println(err)
	ward, err := s.Daos.GetSingleWard(ctx, demand.Address.WardCode)
	if ward != nil {
		m2["ward"] = &ward.Ward
	}
	fmt.Println(err)
	fy, err := s.Daos.GetCurrentFinancialYear(ctx)
	fmt.Println("err in fy = ", err)
	if fy != nil {
		m2["cfy"] = &fy.FinancialYear
	}
	fmt.Println("property demand Demand.NoteFY =======================>", len(demand.NoteFy))
	if len(demand.NoteFy) > 0 {
		if len(demand.NoteFy) == 1 {
			m2["arrearStartFy"] = demand.NoteFy[0]
			m2["arrearEndFy"] = demand.NoteFy[0]
		}
		if len(demand.NoteFy) >= 2 {
			// m2["arrearStartFy"] = demand.NoteFy[len(demand.NoteFy)-1]
			if (len(demand.FYs)) > 0 {
				m2["arrearStartFy"] = demand.FYs[0].FinancialYear.Name
			} else {
				m2["arrearStartFy"] = ""
			}
			// fmt.Println("property demand arrearStartFy =======================>", demand.FYs[0].FinancialYear.Name)
			// fmt.Println("property demand arrearEndFy =======================>", demand.NoteFy[0])

			m2["arrearEndFy"] = demand.NoteFy[0]
		}
	}
	var arrearYearTax, arrearYearRebate, arrearYearPenalty, arrearAlreadyPaid, arrearYearTotal, arrearYearOtherDemand float64
	var currentYearTax, currentYearRebate, currentYearPenalty, currentAlreadyPaid, currentYearTotal, currentYearOtherDemand float64
	for _, v := range demand.FYs {
		if v.IsCurrent {
			currentYearTax = currentYearTax + v.Tax + v.VacantLandTax + v.CompositeTax + v.Ecess
			currentYearRebate = currentYearRebate + v.Rebate
			currentYearPenalty = currentYearPenalty + v.Penalty
			currentAlreadyPaid = currentAlreadyPaid + v.AlreadyPayed.Amount
			currentYearTotal = currentYearTotal + v.TotalTax
			currentYearOtherDemand = currentYearOtherDemand + v.OtherDemand

		} else {
			arrearYearTax = arrearYearTax + v.Tax + v.VacantLandTax + v.CompositeTax + v.Ecess
			arrearYearRebate = arrearYearRebate + v.Rebate
			arrearYearPenalty = arrearYearPenalty + v.Penalty
			arrearAlreadyPaid = arrearAlreadyPaid + v.AlreadyPayed.Amount
			arrearYearTotal = arrearYearTotal + v.TotalTax
			arrearYearOtherDemand = arrearYearOtherDemand + v.OtherDemand
		}
	}

	m2["currentYearTax"] = currentYearTax
	m2["currentYearRebate"] = currentYearRebate
	m2["currentYearPenalty"] = currentYearPenalty
	m2["currentAlreadyPaid"] = currentAlreadyPaid
	m2["currentYearTotal"] = currentYearTotal
	m2["currentYearOtherDemand"] = currentYearOtherDemand

	m2["arrearYearTax"] = arrearYearTax
	m2["arrearYearRebate"] = arrearYearRebate
	m2["arrearYearPenalty"] = arrearYearPenalty
	m2["arrearAlreadyPaid"] = arrearAlreadyPaid
	m2["arrearYearTotal"] = arrearYearTotal
	m2["arrearYearOtherDemand"] = arrearYearOtherDemand

	//fmt.Println(now.AddDate(-1, 0, 0))

	m2["currentDate"] = time.Now()

	m2["Inc"] = func(a int) int {
		return a + 1
	}

	m2["validityDate"] = s.Shared.EndOfMonth(time.Now())

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "Demand_Page_minimal.html"

	//path for download pdf
	// t := time.Now()
	// outputPath := fmt.Sprintf("storage/SampleTemplate%v.pdf", t.Unix())
	fmt.Println(m)

	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	pdfdata.RefDataStr = m3
	pdfdata.Config = productConfig.ProductConfiguration
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}
	return nil, nil
}

func (s *Service) GetPaymentReceiptsPDF(ctx *models.Context, ID string) ([]byte, error) {

	r := NewRequestPdf("")

	data, err := s.GetSinglePropertyPaymentTxtID(ctx, ID)
	if err != nil {
		return nil, err
	}

	// remainingAmount := data.Demand.TotalTax - data.Details.Amount
	// m := make(map[string]interface{})

	// m["remainingAmount"] = remainingAmount
	// var pdfdata models.PDFData
	// pdfdata.Data = m
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	var templatePath string
	if data.Type == "PartPayment" {
		templatePath = templatePathStart + "partpayment_receipt_page.html"
	} else {
		templatePath = templatePathStart + "Receipt_page.html"
	}

	// partpayment_receipt_page
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

// GetPaymentReceiptsPDFV2 : ""
func (s *Service) GetPaymentReceiptsPDFV2(ctx *models.Context, ID string) ([]byte, error) {
	data, err := s.GetSinglePropertyPaymentTxtID(ctx, ID)
	if err != nil {
		return nil, err
	}
	m2 := make(map[string]interface{})

	m2["arrearYears"] = ""
	m2["currentYear"] = ""
	var arrearYearTax, arrearYearRebate, arrearYearPenalty, arrearYearTotal, arrearYearOtherDemand float64
	if data != nil {
		if len(data.Fys) > 0 {
			if data.Fys[0].FY.IsCurrent {
				m2["currentYear"] = data.Fys[0].FY.Name
				m2["currentYearTax"] = data.Fys[0].FY.Tax + data.Fys[0].FY.VacantLandTax + data.Fys[0].FY.CompositeTax + data.Fys[0].FY.Ecess
				m2["currentYearRebate"] = data.Fys[0].FY.Rebate
				m2["currentYearPenalty"] = data.Fys[0].FY.Penalty
				m2["currentYearTotal"] = data.Fys[0].FY.TotalTax
				m2["currentYearOtherDemand"] = data.Fys[0].FY.OtherDemand

				if len(data.Fys) == 2 {
					m2["arrearYears"] = data.Fys[1].FY.Name
				} else if len(data.Fys) >= 3 {
					m2["arrearYears"] = data.Fys[len(data.Fys)-1].FY.Name + " to " + data.Fys[1].FY.Name
				}
			} else {
				if len(data.Fys) == 1 {
					m2["arrearYears"] = data.Fys[0].FY.Name
				} else if len(data.Fys) > 1 {
					m2["arrearYears"] = data.Fys[len(data.Fys)-1].FY.Name + " to" + data.Fys[0].FY.Name
				}
			}

		}
		for _, v := range data.Fys {
			if v.FY.IsCurrent {
				continue
			}
			arrearYearTax = arrearYearTax + v.FY.Tax + v.FY.VacantLandTax + v.FY.CompositeTax + v.FY.Ecess
			arrearYearRebate = arrearYearRebate + v.FY.Rebate
			arrearYearPenalty = arrearYearPenalty + v.FY.Penalty
			arrearYearTotal = arrearYearTotal + v.FY.TotalTax
			arrearYearOtherDemand = arrearYearOtherDemand + v.FY.OtherDemand

		}
		m2["arrearYearTax"] = arrearYearTax
		m2["arrearYearRebate"] = arrearYearRebate
		m2["arrearYearPenalty"] = arrearYearPenalty
		m2["arrearYearTotal"] = arrearYearTotal
		m2["arrearYearOtherDemand"] = arrearYearOtherDemand
	}
	remainingAmount := data.Demand.TotalTax - data.Details.Amount
	m := make(map[string]interface{})
	m["receipt"] = data
	m["remainingAmount"] = remainingAmount
	m2["currentDate"] = time.Now()
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "Receipt_page_minimal.html"

	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// GetPaymentReceiptsPDFV2 : ""
func (s *Service) SavePaymentReceiptsPDFV2(ctx *models.Context, ID string) error {
	data, err := s.GetSinglePropertyPaymentTxtID(ctx, ID)
	if err != nil {
		return err
	}
	m2 := make(map[string]interface{})

	m2["arrearYears"] = ""
	m2["currentYear"] = ""
	var arrearYearTax, arrearYearRebate, arrearYearPenalty, arrearYearTotal, arrearYearOtherDemand float64
	if data != nil {
		if len(data.Fys) > 0 {
			if data.Fys[0].FY.IsCurrent {
				m2["currentYear"] = data.Fys[0].FY.Name
				m2["currentYearTax"] = data.Fys[0].FY.Tax + data.Fys[0].FY.VacantLandTax + data.Fys[0].FY.CompositeTax + data.Fys[0].FY.Ecess
				m2["currentYearRebate"] = data.Fys[0].FY.Rebate
				m2["currentYearPenalty"] = data.Fys[0].FY.Penalty
				m2["currentYearTotal"] = data.Fys[0].FY.TotalTax
				m2["currentYearOtherDemand"] = data.Fys[0].FY.OtherDemand

				if len(data.Fys) == 2 {
					m2["arrearYears"] = data.Fys[1].FY.Name
				} else if len(data.Fys) >= 3 {
					m2["arrearYears"] = data.Fys[len(data.Fys)-1].FY.Name + " to " + data.Fys[1].FY.Name
				}
			} else {
				if len(data.Fys) == 1 {
					m2["arrearYears"] = data.Fys[0].FY.Name
				} else if len(data.Fys) > 1 {
					m2["arrearYears"] = data.Fys[len(data.Fys)-1].FY.Name + " to" + data.Fys[0].FY.Name
				}
			}

		}
		for _, v := range data.Fys {
			if v.FY.IsCurrent {
				continue
			}
			arrearYearTax = arrearYearTax + v.FY.Tax + v.FY.VacantLandTax + v.FY.CompositeTax + v.FY.Ecess
			arrearYearRebate = arrearYearRebate + v.FY.Rebate
			arrearYearPenalty = arrearYearPenalty + v.FY.Penalty
			arrearYearTotal = arrearYearTotal + v.FY.TotalTax
			arrearYearOtherDemand = arrearYearOtherDemand + v.FY.OtherDemand

		}
		m2["arrearYearTax"] = arrearYearTax
		m2["arrearYearRebate"] = arrearYearRebate
		m2["arrearYearPenalty"] = arrearYearPenalty
		m2["arrearYearTotal"] = arrearYearTotal
		m2["arrearYearOtherDemand"] = arrearYearOtherDemand
	}
	remainingAmount := data.Demand.TotalTax - data.Details.Amount
	m := make(map[string]interface{})
	m["receipt"] = data
	m["remainingAmount"] = remainingAmount
	m2["currentDate"] = time.Now()
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "Receipt_page_minimal.html"

	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, file, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")

		// Create blank file
		saveLocation := "/root/project/munger/receipts/"
		filesave, err := os.Create(saveLocation + data.ReciptNo + ".pdf")
		if err != nil {
			log.Fatal(err)
		}
		r2 := bytes.NewReader(file)
		size, err := io.Copy(filesave, r2)
		if err != nil {
			return err
		}

		defer filesave.Close()

		fmt.Printf("Downloaded a file %s with size %d", data.ReciptNo, size)
		return nil

		return err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return errors.New("Error in parcing template - " + err.Error())
	}

}

func (s *Service) GetMultiplePropertyDemandCalcExcel(ctx *models.Context, filter *models.PropertyDemandFilter, pagination *models.Pagination) (error, error) {
	// data,err=s.GetMultiplePropertyDemandCalc(ctx,filter,pagination)
	// if err != nil{
	// 	retrun nil,err
	// }
	// f := excelize.NewFile()
	// // Create a new sheet.
	// ulbSheet := f.NewSheet("Report")
	// f.SetActiveSheet(ulbSheet)
	// startRowIndex := 1
	return nil, nil

}

// BasicUpdateTradeLicense : ""
func (s *Service) PropertyDemandEmailTemplate(ctx *models.Context, bmtlu *models.BasicMobileTowerUpdateData) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		oldMobileTowerData, err := s.Daos.GetSinglePreviousMobileTower(ctx, bmtlu.MobileTowerID)
		if err != nil {
			return errors.New("Error in geting old MobileTower" + err.Error())

		}
		if oldMobileTowerData == nil {
			return errors.New("mobile tower Not Found")
		}
		mtul := new(models.BasicMobileTowerUpdateLog)
		mtul.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERUPDATELOG)
		mtul.MobileTowerID = bmtlu.MobileTowerID
		mtul.Previous = *oldMobileTowerData
		mtul.New = bmtlu.UpdateData
		mtul.UserName = bmtlu.UserName
		mtul.UserType = bmtlu.UserType
		mtul.Proof = bmtlu.Proof
		mtul.Status = constants.MOBILETOWERBASICUPDATELOGINIT
		err = s.Daos.SaveBasicMobileTowerUpdate(ctx, mtul)
		if err != nil {
			return errors.New("Error in updating log" + err.Error())
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
		// html template path
		templateID := templatePathStart + "MobileTowerUpdateRequestEmail.html"
		templateID = "templates/MobileTowerUpdateRequestEmail.html"

		//sending email
		if err := s.SendEmailWithTemplate("Property Update Request - holding no 1111", []string{"solomon2261993@gmail.com"}, templateID, nil); err != nil {
			log.Println("email not sent - ", err.Error())
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

//PropertyParkPenaltyEnable : ""
func (s *Service) PropertyParkPenaltyEnable(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.PropertyParkPenaltyEnable(ctx, UniqueID)
		if err != nil {
			return err
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//PropertyParkPenaltyDisable : ""
func (s *Service) PropertyParkPenaltyDisable(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.PropertyParkPenaltyDisable(ctx, UniqueID)
		if err != nil {
			return err
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// UpdatePropertyLocation : ""
func (s *Service) UpdatePropertyLocation(ctx *models.Context, property *models.PropertyLocation) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyLocation(ctx, property)
		if err != nil {
			return err
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// UpdatePropertyPicture : ""
func (s *Service) UpdatePropertyPicture(ctx *models.Context, property *models.PropertyPicture) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyPicture(ctx, property)
		if err != nil {
			return err
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

//GetPaymentReceiptsPDFS : "Written for ALL Receipts to be saved IN Local"
func (s *Service) GetPaymentReceiptsPDFS(ctx *models.Context, ID string) error {

	r := NewRequestPdf("")

	data, err := s.GetSinglePropertyPaymentTxtID(ctx, ID)
	if err != nil {
		return err
	}

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "Receipt_page.html"
	err = r.ParseTemplate(templatePath, data)
	if err != nil {
		return err
	}
	ok, file, err := r.GeneratePDFAsFile()
	if err != nil {
		return err
	}
	fmt.Println(ok, "pdf generated successfully")

	// Create blank file
	saveLocation := "D:/ProjectData/SilaoReceipts/"
	filesave, err := os.Create(saveLocation + data.ReciptNo + ".pdf")
	if err != nil {
		log.Fatal(err)
	}
	r2 := bytes.NewReader(file)
	size, err := io.Copy(filesave, r2)
	if err != nil {
		return err
	}

	defer filesave.Close()

	fmt.Printf("Downloaded a file %s with size %d", data.ReciptNo, size)
	return nil
}

//GetPaymentReceiptsPDFSWithReceiptNo : "Written for ALL Receipts to be saved IN Local"
func (s *Service) GetPaymentReceiptsPDFSWithReceiptNo(ctx *models.Context, ID string) error {

	r := NewRequestPdf("")

	data, err := s.GetSinglePropertyPaymentReceiptNo(ctx, ID)
	if err != nil {
		return err
	}

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "Receipt_page.html"
	err = r.ParseTemplate(templatePath, data)
	if err != nil {
		return err
	}
	ok, file, err := r.GeneratePDFAsFile()
	if err != nil {
		return err
	}
	fmt.Println(ok, "pdf generated successfully")

	// Create blank file
	saveLocation := "D:/ProjectData/SilaoReceipts/"
	filesave, err := os.Create(saveLocation + data.ReciptNo + ".pdf")
	if err != nil {
		log.Fatal(err)
	}
	r2 := bytes.NewReader(file)
	size, err := io.Copy(filesave, r2)
	if err != nil {
		return err
	}

	defer filesave.Close()

	fmt.Printf("Downloaded a file %s with size %d", data.ReciptNo, size)
	return nil
}

//GetPaymentReceiptsPDfServiceLOOP : "Written for ALL Receipts to be saved IN Local"
func (s *Service) GetPaymentReceiptsPDfServiceLOOP(ctx *models.Context, IDs []string) error {

	for _, v := range IDs {
		err := s.GetPaymentReceiptsPDFS(ctx, v)
		if err != nil {

			fmt.Println(err.Error())
		}
	}
	return nil
}

//SavePaymentReceiptsPDfServiceLOOP : "Written for ALL Receipts to be saved IN Local"
func (s *Service) SavePaymentReceiptsPDfServiceLOOP(ctx *models.Context, IDs []string) error {

	for _, v := range IDs {
		err := s.SavePaymentReceiptsPDFV2(ctx, v)
		if err != nil {

			fmt.Println(err.Error())
		}
	}
	return nil
}

//GetPropertyDemandCalc : ""
func (s *Service) SaveStoredDemand(ctx *models.Context, filter *models.PropertyDemandFilter) (*models.PropertyDemand, error) {

	data, err := s.Daos.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		return nil, err
	}

	if data != nil {

		pdc, err := s.Daos.GetSingleProductConfiguration(ctx, "1")
		if err != nil {
			return nil, errors.New("Error in geting prod config - " + err.Error())
		}
		if pdc == nil {
			return nil, errors.New("Prod config  is nil ")
		}
		data.ProductConfiguration = pdc
		spd := data.PreSaveDemandCalculation()
		dberr := s.Daos.SaveStoredPropertyDemand(ctx, spd)
		if dberr != nil {
			return nil, dberr
		}
	}

	return data, nil
}

func (s *Service) GetDemandV3(ctx *models.Context, UniqueID string) (*models.StoredPropertyDemand, error) {
	demand, err := s.Daos.GetDemandV3(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	ct := time.Now()
	ct = ct.AddDate(0, -3, 0)
	if demand != nil {
		for k, v := range demand.Fys {
			demand.Fys[k].ToPay.Tax = v.Demand.Total - v.Collections.Tax
			if v.Ref.Fy.IsCurrent {
				t1 := v.Ref.Fy.RebateLastDate.AddDate(0, 0, 1)
				t2 := v.Ref.Fy.LastDate.AddDate(0, 0, 1)
				//rain water harvesting rebate
				demand.Fys[k].ToPay.Rebate = (demand.Fys[k].ToPay.Tax * v.Ref.RainWaterHarvestRebate.Rate / 100)
				if s.Shared.ChkDateWithinRange(*v.Ref.Fy.From, *v.Ref.Fy.RebateLastDate, ct) {
					//No penalty | Yes Rebate - Apr 1 - Jun 30
					fmt.Println("No penalty | Yes Rebate - Apr 1 - Jun 30")
					demand.Fys[k].ToPay.Rebate = demand.Fys[k].ToPay.Rebate + (demand.Fys[k].ToPay.Tax * v.Ref.EarlyPaymentRebate.Rate / 100)
					demand.Fys[k].ToPay.Penalty = 0
					demand.Fys[k].ToPay.Others = 0
				} else if s.Shared.ChkDateWithinRange(t1, *v.Ref.Fy.LastDate, ct) {
					//No penalty | No Rebate - Jul 1 - Oct 31
					fmt.Println("No penalty | No Rebate - Jul 1 - Oct 31")
					demand.Fys[k].ToPay.Rebate = 0
					demand.Fys[k].ToPay.Penalty = 0
					demand.Fys[k].ToPay.Others = 0
				} else if s.Shared.ChkDateWithinRange(t2, *v.Ref.Fy.To, ct) {
					//Yes penalty | No Rebate - Nov 1 - Mar 31
					fmt.Println("Yes penalty | No Rebate - Nov 1 - Mar 31")
					demand.Fys[k].ToPay.Rebate = 0
					demand.Fys[k].ToPay.Penalty = (demand.Fys[k].ToPay.Tax * v.Ref.Fy.PenaltyRate / 100)
					demand.Fys[k].ToPay.Others = 0
				}

			} else {
				//Yes penalty | No Rebate
				demand.Fys[k].ToPay.Rebate = 0
				demand.Fys[k].ToPay.Penalty = (demand.Fys[k].ToPay.Tax * v.Ref.Fy.PenaltyRate / 100)
				demand.Fys[k].ToPay.Others = 0
			}
			demand.Fys[k].ToPay.Total = (demand.Fys[k].ToPay.Tax - demand.Fys[k].ToPay.Rebate + demand.Fys[k].ToPay.Penalty + demand.Fys[k].ToPay.Others)

		}
	}
	return demand, nil
}

// DayWisePropertyCollectionReportExcel: ""
func (s *Service) PropertyUpdateLocationExcelReport(ctx *models.Context, filter *models.PropertyUpdateLocationFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterProperty(ctx, &filter.PropertyFilter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}
	excel := excelize.NewFile()
	var sheet1 string
	if filter.IsGeoTagged == "Yes" {
		sheet1 = "Location Updated Properties Report"
	}
	if filter.IsGeoTagged != "Yes" {
		sheet1 = "Location Not Updated Properties Report"
	}
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)

	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "E3")
		excel.MergeCell(sheet1, "A4", "E5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "E3")
		excel.MergeCell(sheet1, "C4", "E5")
	}

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
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding Number")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Mobile Number")
	rowNo++
	// var totalAmount float64
	for i, v := range data {

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())

		// totalAmount = totalAmount + v.TotalTax
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Mobile
			}
			return "NA"
		}())

		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalHoldingsRowNo), fmt.Sprintf("Total Properties - %v", "varen thirumba varen"))

	return excel, nil

}

// GetPropertyDemandCalcNotify : ""
func (s *Service) GetPropertyDemandCalcNotifyV2(ctx *models.Context, filter *models.PropertyDemandFilter, notifyType string) (string, error) {
	resPropertyDemand, err := s.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		return "", err
	}

	if resPropertyDemand == nil {
		return "", errors.New("property demand is nil")
	}
	if len(resPropertyDemand.Ref.PropertyOwner) == 0 {
		return "", errors.New("property owner is nil")
	}
	if len(resPropertyDemand.Ref.PropertyOwner) > 0 {
		if resPropertyDemand.Ref.PropertyOwner[0].Mobile == "" {
			return "", errors.New("mobile number is empty")
		}
	}
	resFYs, err := s.GetCurrentFinancialYear(ctx)
	if err != nil {
		return "", err
	}
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return "", errors.New("Error in getting product config" + err.Error())
	}

	msg := fmt.Sprintf(constants.PROPERTYTAXDEMANDCONTENT, math.Ceil(resPropertyDemand.Demand.TotalTax), resPropertyDemand.UniqueID, resFYs.Name)
	// http://aurangabadmunicipal.com/api/property/getdemand/pdf?id=AURAN-09818
	pdfURL := fmt.Sprintf(productConfig.UIURL + "api/property/getdemand/pdf?id=" + filter.PropertyID)
	fmt.Println("notifyType", notifyType)
	fmt.Println("pdfURL===>", pdfURL)
	var whatsappMsg string
	switch notifyType {
	case "WHATSAPP":
		if len(resPropertyDemand.Ref.PropertyOwner) > 0 {
			whatsappMsg = fmt.Sprintf(constants.COMMONWHATSAPPTEMPLATE, resPropertyDemand.Ref.PropertyOwner[0].Name, productConfig.Name, "Property Tax Demand", msg, pdfURL)
			if err != nil {
				fmt.Println("error in sending message to customer" + err.Error())
			}
		}

	}
	fmt.Println("whatsappMsg ====>", whatsappMsg)
	return whatsappMsg, nil
}

func (s *Service) EnableHoldingProperty(ctx *models.Context, property *models.Property) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableHoldingProperty(ctx, property)
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

//DisableProperty : ""
func (s *Service) DisableHoldingProperty(ctx *models.Context, propertyID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableHoldingProperty(ctx, propertyID)
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

// CheckWhetherPenalChargeApplies : "This api is used to check penal charge applies for particular property"
func (s *Service) CheckWhetherPenalChargeApplies(ctx *models.Context, date *time.Time) (*models.PenalChargeFYDetails, error) {

	details := new(models.PenalChargeFYDetails)

	t := time.Now()
	// fromDate := time.Date(constants.PENALCHARGEFROM, 04, 01, 0, 0, 0, 0, t.Location())
	toDate := time.Date(constants.PENALCHARGETO, 04, 01, 0, 0, 0, 0, t.Location())
	details.FyID = "GEN20122013"

	// sd := fromDate
	ed := toDate
	if date.Before(ed) {
		details.PenalChargeStatus = "Yes"
	} else {
		details.PenalChargeStatus = "No"
	}
	return details, nil
}

//GetPropertyDemandCalcWithStoredCalc : ""
func (s *Service) GetPropertyDemandCalcWithStoredCalc(ctx *models.Context, filter *models.PropertyDemandFilter, collectionName string) (*models.PropertyDemand, error) {
	var data *models.PropertyDemand
	var err error
	data, err = s.Daos.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		return nil, err
	}
	if data != nil {
		// if filter != nil {
		// 	if filter.AllDemand {
		// 		data.AllDemand = true
		// 	}
		// }
		pdc, err := s.Daos.GetSingleProductConfiguration(ctx, "1")
		if err != nil {
			return nil, errors.New("Error in geting prod config - " + err.Error())
		}
		if pdc == nil {
			return nil, errors.New("Prod config  is nil ")
		}
		data.ProductConfiguration = pdc
		data.CTX = ctx
		Storedfy, Storeddemand := data.StoredPropertyDemandCalculationV2()

		err = s.Daos.SaveManyStoredCalDemandFyWithUpsert(ctx, Storedfy)
		if err != nil {
			return nil, err
		}
		err = s.Daos.SaveStoredCalcDemandWithUpsert(ctx, Storeddemand)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (s *Service) StoredPeropertyDemandCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	var arrString []string
	arrString = []string{
		"AURAN-10383", "AURAN-10382", "AURAN-10381", "AURAN-10380", "AURAN-10379",
	}
	for _, v := range arrString {
		filter := new(models.PropertyDemandFilter)
		filter.PropertyID = v
		s.GetPropertyDemandCalcWithStoredCalc(ctx, filter, "")
	}
}

//
func (s *Service) GetAllPropertyDemandCalcReportExcel(ctx *models.Context, filter *models.PropertyDemandFilter, collectionName string, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.GetAllPropertyDemandCalc(ctx, filter, collectionName, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Demand Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "H3")
		excel.MergeCell(sheet1, "A4", "H5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "H3")
		excel.MergeCell(sheet1, "C4", "H5")
	}

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
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
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v", "S.No"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v", "Ward"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", "Holding No."))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v", "Owner Name"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v", "F.Y"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v", "Demand"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v", "Interest"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v", "Total Demand"))
	rowNo++

	var totalAmount float64
	for k, res := range data {
		for _, v := range res.FYs {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
				if res.Ref.Address.Ward != nil {
					return res.Ref.Address.Ward.Name
				}
				return "NA"
			}())

			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), res.UniqueID)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
				if len(res.Ref.PropertyOwner) > 0 {
					return res.Ref.PropertyOwner[0].Name
				}
				return "NA"
			}())
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.VacantLandTax+v.Tax)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Penalty)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.TotalTax)
			totalAmount = totalAmount + v.TotalTax
			rowNo++

		}
	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}
func (s *Service) GetAllPropertyDemandCalc(ctx *models.Context, filter *models.PropertyDemandFilter, collectionName string, pagination *models.Pagination) ([]models.PropertyDemand, error) {
	//var data *models.PropertyDemand
	var datas []models.PropertyDemand
	var err error

	datas, err = s.Daos.GetAllPropertyDemandCalc(ctx, filter, "", pagination)
	if err != nil {
		return nil, err
	}
	for k, data := range datas {

		pdc, err := s.Daos.GetSingleProductConfiguration(ctx, "1")
		if err != nil {
			return nil, errors.New("Error in geting prod config - " + err.Error())
		}
		if pdc == nil {
			return nil, errors.New("Prod config  is nil ")
		}
		data.ProductConfiguration = pdc
		data.CTX = ctx
		data.DemandCalculation()

		datas[k].OverallPropertyDemand.Total.VacantLandTax = data.OverallPropertyDemand.Current.VacantLandTax + data.OverallPropertyDemand.Arrear.VacantLandTax
		datas[k].OverallPropertyDemand.Total.Rebate = data.OverallPropertyDemand.Current.Rebate + data.OverallPropertyDemand.Arrear.Rebate
		datas[k].OverallPropertyDemand.Total.Penalty = data.OverallPropertyDemand.Current.Penalty + data.OverallPropertyDemand.Arrear.Penalty
		datas[k].OverallPropertyDemand.Total.Tax = data.OverallPropertyDemand.Current.Tax + data.OverallPropertyDemand.Arrear.Tax
		datas[k].OverallPropertyDemand.Total.CompositeTax = data.OverallPropertyDemand.Current.CompositeTax + data.OverallPropertyDemand.Arrear.CompositeTax
		datas[k].OverallPropertyDemand.Total.Ecess = data.OverallPropertyDemand.Current.Tax + data.OverallPropertyDemand.Arrear.Ecess
		datas[k].OverallPropertyDemand.Total.PanelCh = data.OverallPropertyDemand.Current.PanelCh + data.OverallPropertyDemand.Arrear.PanelCh

		datas[k].OverallPropertyDemand.Total.Other = data.OverallPropertyDemand.Other.BoreCharge + data.OverallPropertyDemand.Other.FormFee
		datas[k].OverallPropertyDemand.Total.TotalTax = data.OverallPropertyDemand.Total.VacantLandTax + data.OverallPropertyDemand.Total.Tax + data.OverallPropertyDemand.Total.Penalty + data.OverallPropertyDemand.Total.Other - data.OverallPropertyDemand.Total.Rebate + data.OverallPropertyDemand.Total.CompositeTax + data.OverallPropertyDemand.Total.Ecess + data.OverallPropertyDemand.Total.PanelCh
		datas[k].OverallPropertyDemand.Current.TotalTax = data.OverallPropertyDemand.Current.VacantLandTax + data.OverallPropertyDemand.Current.Tax + data.OverallPropertyDemand.Current.Penalty - data.OverallPropertyDemand.Current.Rebate + data.OverallPropertyDemand.Current.CompositeTax + data.OverallPropertyDemand.Current.Ecess + data.OverallPropertyDemand.Current.PanelCh
		datas[k].OverallPropertyDemand.Arrear.TotalTax = data.OverallPropertyDemand.Arrear.VacantLandTax + data.OverallPropertyDemand.Arrear.Tax + data.OverallPropertyDemand.Arrear.Penalty - data.OverallPropertyDemand.Arrear.Rebate + data.OverallPropertyDemand.Arrear.CompositeTax + data.OverallPropertyDemand.Arrear.Ecess + data.OverallPropertyDemand.Arrear.PanelCh

		datas[k].OverallPropertyDemand.Actual.Total.VacantLandTax = data.OverallPropertyDemand.Actual.Total.VacantLandTax + (data.OverallPropertyDemand.Actual.Current.VacantLandTax + data.OverallPropertyDemand.Actual.Arrear.VacantLandTax)
		datas[k].OverallPropertyDemand.Actual.Total.Tax = data.OverallPropertyDemand.Actual.Total.Tax + (data.OverallPropertyDemand.Actual.Current.Tax + data.OverallPropertyDemand.Actual.Arrear.Tax)
		datas[k].OverallPropertyDemand.Actual.Total.TotalTax = data.OverallPropertyDemand.Actual.Total.TotalTax + (data.OverallPropertyDemand.Actual.Current.TotalTax + data.OverallPropertyDemand.Actual.Arrear.TotalTax)
	}
	return datas, nil
}

// CheckWardWiseOldHoldingNoOfProperty : ""
func (s *Service) CheckWardWiseOldHoldingNoOfProperty(ctx *models.Context, ward string, oldHoldingNo string) (*models.RefProperty, error) {
	// property, err := s.Daos.CheckWardWiseOldHoldingNoOfProperty(ctx, ward, oldHoldingNo)
	// if err != nil {
	// 	return nil, err
	// }
	// if property != nil {
	// 	return nil, errors.New("old holding no already exists, please try another")
	// }
	property := new(models.RefProperty)
	return property, nil
}

// UpdatePropertyTotalDemand : ""
func (s *Service) UpdatePropertyTotalDemand(ctx *models.Context, property *models.UpdatePropertyTotalDemand) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyTotalDemand(ctx, property)
		if err != nil {
			return err
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// UpdatePropertyUniqueID : ""
func (s *Service) UpdatePropertyUniqueID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
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
			//
			resWards, err := s.GetSingleWard(ctx, resProperty.Address.WardCode)
			if err != nil {
				return errors.New("error in getting the ward" + err.Error())
			}
			resFloors, err := s.Daos.GetFloorsOfProperty(ctx, v)
			if err != nil {
				log.Println("Error in geting Floors - " + err.Error())
			}
			var tempBuildUpArea float64
			tempConstructionType := "15"
			for _, v := range resFloors {
				if v.No == "16" {
					if v.BuildUpArea > tempBuildUpArea {
						tempBuildUpArea = v.BuildUpArea
						tempConstructionType = v.ConstructionType
					}
				}
			}
			resConstruction, err := s.GetSingleConstructionType(ctx, tempConstructionType)
			if err != nil {
				return errors.New("error in getting the constructed type" + err.Error())
			}
			resRoadType, err := s.GetSingleRoadType(ctx, resProperty.RoadTypeID)
			if err != nil {
				return errors.New("error in getting the road type" + err.Error())
			}
			resPropertyType, err := s.GetSinglePropertyType(ctx, resProperty.PropertyTypeID)
			if err != nil {
				return errors.New("error in getting the property type" + err.Error())
			}
			fmt.Println("Ward Name =========>", resWards.Ward.Name)
			fmt.Println("Property RoadType ======>", resRoadType.Label)
			fmt.Println("construction Type ======>", resConstruction.Label)
			fmt.Println("No of Floors ======>", len(resFloors))
			fmt.Println("property Type ======>", resPropertyType.Label)
			fmt.Println("HouseHold No ======>", resProperty.UniqueID)
			tempFloors := len(resFloors) - 1

			splitString := strings.Split(resProperty.UniqueID, "-")
			tempUniqueId := splitString[1]
			varPrefix := "BNP"
			varDigit := "0"
			//
			uniqueIds.UniqueID = resProperty.UniqueID
			uniqueIds.OldUniqueID = resProperty.UniqueID
			uniqueIds.NewUniqueID = fmt.Sprintf("%v%v%v%v%d%v%v%v", varPrefix, resWards.Ward.Name,
				resRoadType.Label,
				resConstruction.Label,
				tempFloors,
				resPropertyType.Label,
				varDigit,
				tempUniqueId,
			)
			err = s.Daos.UpdatePropertyUniqueID(ctx, uniqueIds)
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

func (s *Service) CreateUserChargeForProperty(ctx *models.Context, PropertyUserCharge *models.PropertyUserCharge) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		floor, err := s.Daos.GetPropertyDAOUsingPropertyID(ctx, PropertyUserCharge.PropertyID)
		if err != nil {
			return err
		}
		PropertyUserCharge.DOA = floor.DateFrom
		PropertyUserCharge.IsUserCharge = "Yes"
		PropertyUserCharge.Createdby.On = &t
		err = s.Daos.CreateUserChargeForProperty(ctx, PropertyUserCharge)
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
