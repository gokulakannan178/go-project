package services

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//ImportExcelFileForFarmerLand :""
func (s *Service) FarmerLandUploadExcel(ctx *models.Context, file multipart.File) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN             = 16
		OMITROWS              = 0
		FARMERIDCOLUMN        = 0
		STATECOLUMN           = 1
		DISTRICTCOLUMN        = 2
		BLOCKCOLUMN           = 3
		GRAMPANCHAYATCOLUMN   = 4
		VILLAGECOLUMN         = 5
		CULTIVATIONCOLUMN     = 6
		AREAINACRECOLUMN      = 7
		CULTIVATEDAREACOLUMN  = 8
		VACANTAREACOLUMN      = 9
		IRRIGATIIONTYPECOLUMN = 10
		LANDPOSITIONCOLUMN    = 11
		LANDTYPECOLUMN        = 12
		OWNERSHIPCOLUMN       = 13
		PARCELNOCOLUMN        = 14
		SOILTYPECOLUMN        = 15
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		farmerLands := make([]models.FarmerLand, 0)
		rows := f.GetRows("Sheet1")

		for rowIndex, row := range rows {
			if rowIndex <= OMITROWS {
				continue
			}
			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			farmerLand := models.FarmerLand{}

			farmerLand.FarmerID = row[FARMERIDCOLUMN]
			farmer, err := s.Daos.GetSingleFarmerWithFarmerId(ctx, row[FARMERIDCOLUMN])
			if err != nil {
				return errors.New("farmer not found")
			}
			farmerLand.Farmer = farmer[0].ID
			if s, err := strconv.ParseFloat(row[AREAINACRECOLUMN], 64); err == nil {
				farmerLand.AreaInAcre = s
			}
			farmerLand.CultivationPractice = row[CULTIVATIONCOLUMN]
			if s, err := strconv.ParseFloat(row[CULTIVATEDAREACOLUMN], 64); err == nil {
				farmerLand.CultivatedArea = s
			}
			if s, err := strconv.ParseFloat(row[VACANTAREACOLUMN], 64); err == nil {
				farmerLand.VacantArea = s
			}
			farmerLand.IrrigationType = row[IRRIGATIIONTYPECOLUMN]
			farmerLand.LandPosition = row[LANDPOSITIONCOLUMN]
			farmerLand.LandType = row[LANDTYPECOLUMN]
			farmerLand.OwnerShip = row[OWNERSHIPCOLUMN]
			farmerLand.ParcelNumber = row[PARCELNOCOLUMN]
			if row[SOILTYPECOLUMN] != "" {
				resSoilType, err := s.Daos.GetSingleSoilTypeWithName(ctx, row[SOILTYPECOLUMN])
				if err != nil {
					return err
				}
				if len(resSoilType) > 0 {
					farmerLand.SoilType = resSoilType[0].ID

					fmt.Println("resSoilTypeName=========>", resSoilType[0].Name)
				}
			}
			if row[STATECOLUMN] != "" {
				resState, err := s.Daos.GetSingleStateWithUniqueID(ctx, row[STATECOLUMN])
				if err != nil {
					return err
				}
				if resState != nil {
					farmerLand.State = resState.ID
				}
			}
			if row[DISTRICTCOLUMN] != "" {
				resDistrict, err := s.Daos.GetSingleDistrictWithUniqueId(ctx, row[DISTRICTCOLUMN])
				if err != nil {
					return err
				}
				if resDistrict != nil {
					farmerLand.District = resDistrict.ID
				}
			}
			if row[BLOCKCOLUMN] != "" {
				resBlock, err := s.Daos.GetSingleBlockWithUnique(ctx, row[BLOCKCOLUMN])
				if err != nil {
					return err
				}
				if resBlock != nil {
					farmerLand.Block = resBlock.ID

				}
			}
			if row[GRAMPANCHAYATCOLUMN] != "" {
				resGramPanchayat, err := s.Daos.GetSingleGramPanchayatWithUniqueId(ctx, row[GRAMPANCHAYATCOLUMN])
				if err != nil {
					return err
				}
				if resGramPanchayat != nil {
					farmerLand.GramPanchayat = resGramPanchayat.ID

				}
			}
			if row[VILLAGECOLUMN] != "" {
				resVillage, err := s.Daos.GetSingleVillageWithUniqueId(ctx, row[VILLAGECOLUMN])
				if err != nil {
					return err
				}
				if resVillage != nil {
					farmerLand.Village = resVillage.ID
				}
			}
			farmerLand.Status = constants.FARMERLANDSTATUSACTIVE

			farmerLands = append(farmerLands, farmerLand)

		}
		for _, farmerland := range farmerLands {
			fmt.Println("farmerId====>", farmerland.FarmerID)
			// fmt.Println("SoilType====>", farmerland.SoilType)
		}
		fmt.Println("farmerlands====>", len(farmerLands))
		dberr := s.Daos.FarmerLandUploadExcel(ctx, farmerLands)
		if dberr != nil {
			return errors.New("Db Error" + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

// FarmerAggregationUploadExcel : ""
func (s *Service) FarmerAggregationUploadExcel(ctx *models.Context, file multipart.File) []models.FarmerUploadError {
	log.Println("transaction start")
	//Start Transaction
	orgRefMap := make(map[string]primitive.ObjectID)
	projectRefMap := make(map[string]primitive.ObjectID)
	stateRefMap := make(map[string]primitive.ObjectID)
	districtRefMap := make(map[string]primitive.ObjectID)
	blockRefMap := make(map[string]primitive.ObjectID)
	grampRefMap := make(map[string]primitive.ObjectID)
	villageRefMap := make(map[string]primitive.ObjectID)
	var errs []models.FarmerUploadError
	var farmererr models.FarmerUploadError
	if err := ctx.Session.StartTransaction(); err != nil {
		farmererr.Error = err.Error()
		errs = append(errs, farmererr)
		return errs
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN           = 13
		OMITROWS            = 0
		ORGANISATIONCOLUMN  = 0
		PROJECTCOLUMN       = 1
		NAMECOLUMN          = 2
		FATHERNAMECOLUMN    = 3
		USERNAMECOLUMN      = 4
		GENDERCOLUMN        = 5
		MOBILENOCOLUMN      = 6
		VILLAGECOLUMN       = 7
		GRAMPANCHAYATCOLUMN = 8
		BLOCKCOLUMN         = 9
		DISTRICTCOLUMN      = 10
		STATECOLUMN         = 11
		CREATEDDATECOLUMN   = 12
		CREATEDDATELAYOUT   = "02-January-2006"
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("started reading file")
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		farmers := make([]models.Farmer, 0)
		rows := f.GetRows("Sheet1")
		//var errors []string
		fmt.Println("started looping")

		for rowIndex, row := range rows {
			fmt.Println("row no === ", rowIndex)
			if rowIndex <= OMITROWS {
				continue
			}

			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			farmer := models.Farmer{}
			farmermobile, err := s.Daos.GetSingleFarmerWithMobileno(ctx, row[MOBILENOCOLUMN])
			if err != nil {
				if row[ORGANISATIONCOLUMN] != "" {
					_, ok := orgRefMap[row[ORGANISATIONCOLUMN]]
					if !ok {
						resOrganisation, err := s.Daos.GetSingleOrganisationWithUniqueID(ctx, row[ORGANISATIONCOLUMN])
						if err != nil {
							farmererr.MobileNumber = row[MOBILENOCOLUMN]
							farmererr.Error = err.Error()
							errs = append(errs, farmererr)
							//	farmer.FarmerOrg, _ = primitive.ObjectIDFromHex(err.Error())
							continue
						}
						orgRefMap[row[ORGANISATIONCOLUMN]] = resOrganisation.ID
						fmt.Println("resOrganisationName=========>", resOrganisation.Name)
						//if resOrganisation != nil {

					}
					farmer.FarmerOrg = orgRefMap[row[ORGANISATIONCOLUMN]]

				}
				if row[PROJECTCOLUMN] != "" {
					_, ok := projectRefMap[row[PROJECTCOLUMN]]
					if !ok {
						resProject, err := s.Daos.GetSingleProjectWithUniqueID(ctx, row[PROJECTCOLUMN])
						if err != nil {
							farmererr.MobileNumber = row[MOBILENOCOLUMN]
							farmererr.Error = err.Error()
							errs = append(errs, farmererr)
							//	farmer.ProjectID, _ = primitive.ObjectIDFromHex(err.Error())

							continue
						}
						if resProject != nil {
							projectRefMap[row[PROJECTCOLUMN]] = resProject.ID
							fmt.Println("resProjectName=========>", resProject.Name)

						}
					}
					farmer.ProjectID = projectRefMap[row[PROJECTCOLUMN]]

				}
				//}

				// if s, err := strconv.ParseFloat(row[AREAINACRECOLUMN], 64); err == nil {
				// 	farmerLand.AreaInAcre = s
				// }
				farmer.Name = row[NAMECOLUMN]
				farmer.FatherName = row[FATHERNAMECOLUMN]
				farmer.Gender = row[GENDERCOLUMN]
				farmer.MobileNumber = row[MOBILENOCOLUMN]
				farmer.Status = constants.FARMERSTATUSACTIVE
				farmer.FarmerID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFARMER)
				if row[STATECOLUMN] != "" {
					_, ok := stateRefMap[row[STATECOLUMN]]
					if !ok {
						resState, err := s.Daos.GetSingleStateWithUniqueID(ctx, row[STATECOLUMN])
						if err != nil {
							farmererr.MobileNumber = row[MOBILENOCOLUMN]
							farmererr.Error = err.Error()
							errs = append(errs, farmererr)
							//farmer.State, _ = primitive.ObjectIDFromHex(err.Error())
							continue
						}
						if resState != nil {

							stateRefMap[row[STATECOLUMN]] = resState.ID
						}
					}
					farmer.State = stateRefMap[row[STATECOLUMN]]

				}
				if row[DISTRICTCOLUMN] != "" {
					_, ok := districtRefMap[row[DISTRICTCOLUMN]]
					if !ok {
						resDistrict, err := s.Daos.GetSingleDistrictWithUniqueId(ctx, row[DISTRICTCOLUMN])
						if err != nil {
							farmererr.MobileNumber = row[MOBILENOCOLUMN]
							farmererr.Error = err.Error()
							errs = append(errs, farmererr)
							//farmer.District, _ = primitive.ObjectIDFromHex(err.Error())
							continue
						}
						if resDistrict != nil {
							districtRefMap[row[DISTRICTCOLUMN]] = resDistrict.ID

						}
					}
					farmer.District = districtRefMap[row[DISTRICTCOLUMN]]
				}
				//
				if row[BLOCKCOLUMN] != "" {
					_, ok := blockRefMap[row[BLOCKCOLUMN]]
					if !ok {
						resBlock, err := s.Daos.GetSingleBlockWithUnique(ctx, row[BLOCKCOLUMN])
						if err != nil {
							farmererr.MobileNumber = row[MOBILENOCOLUMN]
							farmererr.Error = err.Error()
							errs = append(errs, farmererr)
							//	farmer.Block, _ = primitive.ObjectIDFromHex(err.Error())
							continue
						}
						if resBlock != nil {
							blockRefMap[row[BLOCKCOLUMN]] = resBlock.ID

						}
					}
					farmer.Block = blockRefMap[row[BLOCKCOLUMN]]
				}
				if row[GRAMPANCHAYATCOLUMN] != "" {
					_, ok := grampRefMap[row[GRAMPANCHAYATCOLUMN]]
					if !ok {
						resGramPanchayat, err := s.Daos.GetSingleGramPanchayatWithUniqueId(ctx, row[GRAMPANCHAYATCOLUMN])
						if err != nil {
							farmererr.MobileNumber = row[MOBILENOCOLUMN]
							farmererr.Error = err.Error()
							errs = append(errs, farmererr)
							//	farmer.GramPanchayat, _ = primitive.ObjectIDFromHex(err.Error())
							continue
						}
						if resGramPanchayat != nil {
							grampRefMap[row[GRAMPANCHAYATCOLUMN]] = resGramPanchayat.ID

						}
					}
					farmer.GramPanchayat = grampRefMap[row[GRAMPANCHAYATCOLUMN]]
				}
				if row[VILLAGECOLUMN] != "" {
					_, ok := villageRefMap[row[GRAMPANCHAYATCOLUMN]]
					if !ok {
						resVillage, err := s.Daos.GetSingleVillageWithUniqueId(ctx, row[VILLAGECOLUMN])
						if err != nil {
							farmererr.MobileNumber = row[MOBILENOCOLUMN]
							farmererr.Error = err.Error()
							errs = append(errs, farmererr)
							//	farmer.Village, _ = primitive.ObjectIDFromHex(err.Error())
							continue
						}
						if resVillage != nil {
							villageRefMap[row[GRAMPANCHAYATCOLUMN]] = resVillage.ID

						}
					}
					farmer.Village = villageRefMap[row[GRAMPANCHAYATCOLUMN]]
				}
				//}

				//}

				//	}
				//}

				if row[CREATEDDATECOLUMN] != "" {
					layout := CREATEDDATELAYOUT
					t, err := time.Parse(layout, row[CREATEDDATECOLUMN])
					if err != nil {
						return err
					}
					farmer.CreatedDate = &t
				}
				farmers = append(farmers, farmer)

			}
			fmt.Println("finished arranging data")

			// if len(errs) != 0 {
			// 	farmererr.MobileNumber = row[MOBILENOCOLUMN]
			// 	farmererr.Error = "Success"
			// 	errs = append(errs, farmererr)
			// 	return errors.New("Success")
			// }
			if farmermobile != nil {
				farmererr.MobileNumber = row[MOBILENOCOLUMN]
				farmererr.Error = "already mobile no updated"
				errs = append(errs, farmererr)
				//return errors.New("already mobile no updated")
			}

		}
		if len(errs) == 0 {
			tempFarmers := make([]interface{}, 0)
			tempProjectFarmers := make([]interface{}, 0)
			for k, farmer := range farmers {

				fmt.Println("farmerId====>", farmer.FarmerID)
				// err := s.Daos.SaveFarmer(ctx, &farmer)
				// if err != nil {
				// 	return err
				// }
				tempFarmers = append(tempFarmers, farmer)
				projectFarmer := new(models.ProjectFarmer)
				projectFarmer.Farmer = farmer.ID
				projectFarmer.Project = farmer.ProjectID
				farmer.Status = constants.PROJECTFARMERSTATUSACTIVE
				t := time.Now()
				created := models.CreatedV2{}
				created.On = &t
				created.By = constants.SYSTEM
				projectFarmer.Created = &created
				// if err := s.Daos.SaveProjectFarmer(ctx, projectFarmer); err != nil {
				// 	return err
				// }
				tempProjectFarmers = append(tempProjectFarmers, projectFarmer)
				if k+1%1000 == 0 || k == len(farmers)-1 {
					fmt.Println((k+1)*1000, "time ===================================")
					_, err := s.Daos.SaveManyFarmers(ctx, tempFarmers)
					if err != nil {
						return err
					}
					err = s.Daos.SaveManyProjectFarmer(ctx, tempProjectFarmers)
					if err != nil {
						return err
					}
					tempFarmers = make([]interface{}, 0)
					tempProjectFarmers = make([]interface{}, 0)
				}
				//return errors.New("Success")
			}
		}
		// for k, _ := range farmers {
		// 	farmererr.MobileNumber = farmers[k].MobileNumber
		// 	farmererr.Error = "Success"
		// 	errs = append(errs, farmererr)
		// }
		fmt.Println("farmers====>", len(farmers))
		// dberr := s.Daos.FarmerAggregationUploadExcel(ctx, farmers)
		// if dberr != nil {
		// 	return errors.New("Db Error" + dberr.Error())
		// }

		return nil

	}); err != nil {
		farmererr.Error = err.Error()
		errs = append(errs, farmererr)
		return errs
	}
	return errs
}

// FarmerCasteUploadExcel : ""
func (s *Service) FarmerCasteUploadExcel(ctx *models.Context, file multipart.File) ([]models.ExcelUploadError, error) {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN = 3
		OMITROWS  = 0
		SNO       = 0
		FARMERID  = 1
		CASTROW   = 2
	)
	var errs []models.ExcelUploadError
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		farmers := make([]models.Farmer, 0)
		rows := f.GetRows("Sheet1")
		fmt.Printf("no of rows %v\n", len(rows))
		for rowIndex, row := range rows {
			var err models.ExcelUploadError
			if rowIndex <= OMITROWS {
				continue
			}
			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			farmer := models.Farmer{}
			//Farmers := models.RefFarmer{}
			//errs = append(errs, "farmer id is invaild")
			if row[FARMERID] == "" {
				err.ID = ""
				err.Message = fmt.Sprintf("farmer id is missing in column number %v", row[SNO])
				err.Error = ""
				errs = append(errs, err)
				continue

			}
			farmer.FarmerID = row[FARMERID]
			farmer.Cast = row[CASTROW]
			farmers = append(farmers, farmer)

		}
		for _, farmer := range farmers {
			var err models.ExcelUploadError
			err2 := s.Daos.UpdateCastFarmer(ctx, farmer.FarmerID, farmer.Cast)
			if err2 != nil {
				err.ID = farmer.FarmerID
				err.Message = "failed to update"
				err.Error = err2.Error()
				errs = append(errs, err)
				continue
			}
		}
		fmt.Println("farmers====>", len(farmers))
		return nil

	}); err != nil {
		return nil, err
	}
	return errs, nil

}

// FarmerSoilUploadExcel : ""
func (s *Service) FarmerSoilUploadExcel(ctx *models.Context, file multipart.File) ([]models.ExcelUploadError, error) {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	var excelerrs []models.ExcelUploadError
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN             = 11
		OMITROWS              = 0
		SNOCOLUMN             = 0
		FARMERCOLUMN          = 1
		LABNAMEECOLUMN        = 2
		ECVALUECOLUMN         = 3
		HUMIDITYCOLUMN        = 4
		PHCOLUMN              = 5
		SOILCOLLECTEDONCOLUMN = 6
		SOILTESTONCOLUMN      = 7
		VAILDFROMCOLUMN       = 8
		VAILDTOCOLUMN         = 9
		SOILSAMPLENOCOLUMN    = 10
		ORGANICCARBONCOLUMN   = 11
		CREATEDDATELAYOUT     = "02-January-2006"
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		f, err := excelize.OpenReader(file)
		if err != nil {

			return err
		}
		farmers := make([]models.FarmerSoilData, 0)
		rows := f.GetRows("Sheet1")

		for rowIndex, row := range rows {
			if rowIndex <= OMITROWS {
				continue
			}
			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			farmerSoilData := models.FarmerSoilData{}

			if row[FARMERCOLUMN] != "" {
				FarmerId, err := s.Daos.GetSingleFarmerWithFarmerId(ctx, row[FARMERCOLUMN])
				if err != nil {
					fmt.Println("open reader error")
					return err
				}
				if len(FarmerId) > 0 {
					farmerSoilData.Farmer = FarmerId[0].ID
					fmt.Println("FarmerIdName=========>", FarmerId[0].Name)

				}
			}
			// if s, err := strconv.ParseFloat(row[AREAINACRECOLUMN], 64); err == nil {
			// 	farmerLand.AreaInAcre = s
			// }
			farmerSoilData.Status = constants.SOILTYPESTATUSACTIVE
			if row[LABNAMEECOLUMN] != "" {
				lab, err := s.Daos.GetSingleAidlocationWithName(ctx, row[LABNAMEECOLUMN])
				if err != nil {
					return err
				}
				if len(lab) > 0 {
					farmerSoilData.LabName = lab[0].ID
					fmt.Println("labName=========>", lab[0].Name)

				}
			}

			b1, err := strconv.ParseFloat(row[ECVALUECOLUMN], 64)
			if err != nil {
				var excelerr models.ExcelUploadError
				excelerr.ID = row[SNOCOLUMN]
				excelerr.Message = "failed to convert EcValue "
				excelerr.Error = err.Error()
				excelerrs = append(excelerrs, excelerr)
				continue
			}
			farmerSoilData.EcValue = b1
			b2, err := strconv.ParseFloat(row[HUMIDITYCOLUMN], 64)
			if err != nil {
				var excelerr models.ExcelUploadError
				excelerr.ID = row[SNOCOLUMN]
				excelerr.Message = "failed to convert Humidity "
				excelerr.Error = err.Error()
				excelerrs = append(excelerrs, excelerr)
				continue
			}
			farmerSoilData.Humidity = b2
			b3, err := strconv.ParseFloat(row[PHCOLUMN], 64)
			if err != nil {
				var excelerr models.ExcelUploadError
				excelerr.ID = row[SNOCOLUMN]
				excelerr.Message = "failed to  convert PH "
				excelerr.Error = err.Error()
				excelerrs = append(excelerrs, excelerr)
				continue
			}
			farmerSoilData.PH = b3

			if row[SOILCOLLECTEDONCOLUMN] != "" {
				fmt.Println(row[SOILCOLLECTEDONCOLUMN])
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[SOILCOLLECTEDONCOLUMN])
				if err != nil {
					return err
				}
				farmerSoilData.SoilCollectedOn = &t
			}

			if row[SOILTESTONCOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[SOILTESTONCOLUMN])
				if err != nil {
					return err
				}
				farmerSoilData.SoilTestedOn = &t
			}
			if row[VAILDFROMCOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[VAILDFROMCOLUMN])
				if err != nil {
					return err
				}

				farmerSoilData.ValidFrom = &t
			}
			if row[VAILDTOCOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[VAILDTOCOLUMN])
				if err != nil {
					return err
				}
				farmerSoilData.ValidTo = &t
			}
			b4, err := strconv.ParseFloat(row[SOILSAMPLENOCOLUMN], 64)
			if err != nil {
				var excelerr models.ExcelUploadError
				excelerr.ID = row[SNOCOLUMN]
				excelerr.Message = "failed to convert  SoilSampleNo"
				excelerr.Error = err.Error()
				excelerrs = append(excelerrs, excelerr)
				continue
			}
			farmerSoilData.SoilSampleNo = b4
			b5, err := strconv.ParseFloat(row[ORGANICCARBONCOLUMN], 64)
			if err != nil {
				var excelerr models.ExcelUploadError
				excelerr.ID = row[SNOCOLUMN]
				excelerr.Message = "failed to convert OrganicCarbon"
				excelerr.Error = err.Error()
				excelerrs = append(excelerrs, excelerr)
				continue
			}
			farmerSoilData.OrganicCarbon = b5
			farmers = append(farmers, farmerSoilData)

		}
		for _, farmer := range farmers {
			fmt.Println("farmerId====>", farmer.Farmer)
			err := s.Daos.SaveFarmerSoilData(ctx, &farmer)
			if err != nil {
				return err
			}

		}
		fmt.Println("farmersSoild====>", len(farmers))
		// dberr := s.Daos.FarmerAggregationUploadExcel(ctx, farmers)
		// if dberr != nil {
		// 	return errors.New("Db Error" + dberr.Error())
		// }
		return nil

	}); err != nil {
		return nil, err
	}
	return excelerrs, nil

}

// FarmerCropUploadExcel : ""
func (s *Service) FarmerCropUploadExcel(ctx *models.Context, file multipart.File) ([]models.ExcelUploadError, error) {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	var excelerrs []models.ExcelUploadError
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN         = 16
		OMITROWS          = 0
		SNOCOLUMN         = 0
		AREACOLUMN        = 1
		CROPCOLUMN        = 2
		INTERCROPCOLUMN   = 3
		FARMERCOLUMN      = 4
		IRRIGATIONCOLUMN  = 5
		SEASONCOLUMN      = 6
		STARTDATECOLUMN   = 7
		UNITCOLUMN        = 8
		YEARCOLUMN        = 9
		VERIETYCOLUMN     = 10
		YEILDCOLUMN       = 11
		YEILDUNITCOLUMN   = 12
		INPUTCOSTCOLUMN   = 13
		YEILDVALUECOLUMN  = 14
		CONSUMATIONCOLUMN = 15
		COMPLETEDCOLUMN   = 16
		CREATEDDATELAYOUT = "02-January-2006"
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		f, err := excelize.OpenReader(file)
		if err != nil {

			return err
		}
		farmers := make([]models.FarmerCrop, 0)
		rows := f.GetRows("Sheet1")

		for rowIndex, row := range rows {
			if rowIndex <= OMITROWS {
				continue
			}
			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			farmerCrop := models.FarmerCrop{}
			farmerCrop.Area = row[AREACOLUMN]
			farmerCrop.Status = constants.FARMERCROPSTATUSWIP
			farmerCrop.ActiveStatus = true
			if row[CROPCOLUMN] != "" {
				CommodityId, err := s.Daos.GetSingleCommodityWithName(ctx, row[CROPCOLUMN])
				if err != nil {
					fmt.Println("open reader error")
					return err
				}
				if len(CommodityId) > 0 {
					farmerCrop.Crop = CommodityId[0].ID
					fmt.Println("cropName=========>", CommodityId[0].Name)

				}
			}
			if row[INTERCROPCOLUMN] != "" {
				InterCropId, err := s.Daos.GetSingleCommodityWithName(ctx, row[INTERCROPCOLUMN])
				if err != nil {
					fmt.Println("open reader error")
					return err
				}
				if len(InterCropId) > 0 {
					farmerCrop.InterCrop = InterCropId[0].ID
					fmt.Println("interCropName=========>", InterCropId[0].Name)

				}
			}

			if row[FARMERCOLUMN] != "" {
				FarmerId, err := s.Daos.GetSingleFarmerWithFarmerId(ctx, row[FARMERCOLUMN])
				if err != nil {
					fmt.Println("open reader error")
					return err
				}
				if len(FarmerId) > 0 {
					farmerCrop.Farmer = FarmerId[0].ID
					fmt.Println("FarmerIdName=========>", FarmerId[0].Name)

				}
			}
			farmerCrop.Irrigation = row[IRRIGATIONCOLUMN]
			if row[SEASONCOLUMN] != "" {
				SeanId, err := s.Daos.GetSingleCropseasonWithName(ctx, row[SEASONCOLUMN])
				if err != nil {
					fmt.Println("open reader error")
					return err
				}
				if len(SeanId) > 0 {
					farmerCrop.Season = SeanId[0].ID
					fmt.Println("SeasonName=========>", SeanId[0].Name)

				}
			}
			if row[STARTDATECOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[STARTDATECOLUMN])
				if err != nil {
					return err
				}
				farmerCrop.StartDate = &t
			} else {
				t2 := time.Now()
				farmerCrop.StartDate = &t2
			}
			farmerCrop.Unit = row[UNITCOLUMN]
			// if s, err := strconv.ParseFloat(row[AREAINACRECOLUMN], 64); err == nil {
			// 	farmerLand.AreaInAcre = s
			// }
			b1, err := strconv.ParseInt(row[YEARCOLUMN], 10, 64)
			if err != nil {
				var excelerr models.ExcelUploadError
				excelerr.ID = row[SNOCOLUMN]
				excelerr.Message = "failed to convert year "
				excelerr.Error = err.Error()
				excelerrs = append(excelerrs, excelerr)
				continue
			}
			farmerCrop.Year = int(b1)
			if row[VERIETYCOLUMN] != "" {
				VerityId, err := s.Daos.GetSingleCommodityVarietyWithName(ctx, row[VERIETYCOLUMN])
				if err != nil {
					fmt.Println("open reader error")
					return err
				}
				if len(VerityId) > 0 {
					farmerCrop.Season = VerityId[0].ID
					fmt.Println("VerietyName=========>", VerityId[0].Name)

				}
			}

			b2, err := strconv.ParseFloat(row[YEILDCOLUMN], 64)
			if err != nil {
				var excelerr models.ExcelUploadError
				excelerr.ID = row[SNOCOLUMN]
				excelerr.Message = "failed to convert yeild "
				excelerr.Error = err.Error()
				excelerrs = append(excelerrs, excelerr)
				continue
			}
			farmerCrop.Yeild = b2
			// b3, err := strconv.ParseInt(row[YEILDUNITCOLUMN], 10, 64)
			// if err != nil {
			// 	var excelerr models.ExcelUploadError
			// 	excelerr.ID = row[SNOCOLUMN]
			// 	excelerr.Message = "failed to convert yeidunit "
			// 	excelerr.Error = err.Error()
			// 	excelerrs = append(excelerrs, excelerr)
			// 	continue
			// }
			farmerCrop.YieldUnit = row[YEILDUNITCOLUMN]

			b4, err := strconv.ParseInt(row[INPUTCOSTCOLUMN], 10, 64)
			if err != nil {
				var excelerr models.ExcelUploadError
				excelerr.ID = row[SNOCOLUMN]
				excelerr.Message = "failed to convert inputcast "
				excelerr.Error = err.Error()
				excelerrs = append(excelerrs, excelerr)
				continue
			}
			farmerCrop.InputCost = int(b4)

			b5, err := strconv.ParseInt(row[YEILDVALUECOLUMN], 10, 64)
			if err != nil {
				var excelerr models.ExcelUploadError
				excelerr.ID = row[SNOCOLUMN]
				excelerr.Message = "failed to convert yeildvalue "
				excelerr.Error = err.Error()
				excelerrs = append(excelerrs, excelerr)
				continue
			}
			farmerCrop.YieldValue = int(b5)

			farmerCrop.Consumption = row[CONSUMATIONCOLUMN]

			if row[COMPLETEDCOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[COMPLETEDCOLUMN])
				if err != nil {
					return err
				}
				farmerCrop.CompletedDate = &t
			} else {
				t2 := time.Now()
				farmerCrop.CompletedDate = &t2
			}

			farmers = append(farmers, farmerCrop)

		}
		for _, farmer := range farmers {
			fmt.Println("farmerId====>", farmer.Farmer)
			err := s.Daos.SaveFarmerCrop(ctx, &farmer)
			if err != nil {
				return err
			}

		}
		fmt.Println("farmersCrop====>", len(farmers))
		// dberr := s.Daos.FarmerAggregationUploadExcel(ctx, farmers)
		// if dberr != nil {
		// 	return errors.New("Db Error" + dberr.Error())
		// }
		return nil

	}); err != nil {
		return nil, err
	}
	return excelerrs, nil

}

// FarmerAggregationUploadExcelWithNames : ""
func (s *Service) FarmerAggregationUploadExcelWithNames(ctx *models.Context, file multipart.File, searchPolicy bool, version string) []models.FarmerUploadError {
	log.Println("transaction start")
	t := time.Now()
	//Start Transaction
	orgRefMap := make(map[string]primitive.ObjectID)
	projectRefMap := make(map[string]primitive.ObjectID)
	//projectRefMap := make(map[string]primitive.ObjectID)
	stateRefMap := make(map[string]primitive.ObjectID)
	districtRefMap := make(map[string]primitive.ObjectID)
	blockRefMap := make(map[string]primitive.ObjectID)
	grampRefMap := make(map[string]primitive.ObjectID)
	villageRefMap := make(map[string]primitive.ObjectID)
	defer func() {
		for k := range orgRefMap {
			delete(orgRefMap, k)
		}
		for k := range projectRefMap {
			delete(projectRefMap, k)
		}
		for k := range stateRefMap {
			delete(stateRefMap, k)
		}
		for k := range districtRefMap {
			delete(districtRefMap, k)
		}
		for k := range blockRefMap {
			delete(blockRefMap, k)
		}
		for k := range grampRefMap {
			delete(grampRefMap, k)
		}
		for k := range villageRefMap {
			delete(villageRefMap, k)
		}
	}()
	var errs []models.FarmerUploadError
	var farmererr models.FarmerUploadError
	if err := ctx.Session.StartTransaction(); err != nil {
		farmererr.Error = err.Error()
		errs = append(errs, farmererr)
		return errs
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN = 10
		OMITROWS  = 0
		// ORGANISATIONCOLUMN  = 0
		// PROJECTCOLUMN       = 1
		FIRSTNAMECOLUMN   = 0
		FATHERNAMECOLUMN  = 1
		LASTNAMECOLUMN    = 2
		GENDERCOLUMN      = 3
		MOBILENOCOLUMN    = 4
		VILLAGECOLUMN     = 5
		BLOCKCOLUMN       = 6
		DISTRICTCOLUMN    = 7
		STATECOLUMN       = 8
		CREATEDDATECOLUMN = 9

		CREATEDDATELAYOUT = "02-January-2006"
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("started reading file")
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		farmers := make([]models.Farmer, 0)
		rows := f.GetRows("Sheet1")
		//var errors []string
		fmt.Println("started looping")
		var globalProjectId primitive.ObjectID
		var Orgid primitive.ObjectID
		var ProjectId primitive.ObjectID

		productconfig, err := s.Daos.GetactiveProductConfig(ctx, true)
		if err != nil {
			return err
		}
		if productconfig != nil {
			if &productconfig.Orgnisation != nil {
				//	farmer.FarmerOrg
				Orgid = productconfig.Orgnisation.OrgnisationID
				//farmer.ProjectID
			}
			if &productconfig.Project != nil {
				ProjectId = productconfig.Project.ProjectID

			}
		}
		for rowIndex, row := range rows {
			fmt.Println("row no === ", rowIndex)
			if rowIndex <= OMITROWS {
				continue
			}

			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			farmer := models.Farmer{}
			if row[MOBILENOCOLUMN] == "" {
				farmererr.MobileNumber = row[MOBILENOCOLUMN]
				farmererr.Error = "mobile no cannot be empty"
				errs = append(errs, farmererr)
				continue
			}
			farmer.FarmerOrg = Orgid
			farmer.ProjectID = ProjectId
			farmermobile, err := s.Daos.GetSingleFarmerWithMobilenoAndOrg(ctx, farmer.FarmerOrg.Hex(), row[MOBILENOCOLUMN])
			if err != nil {
				if err.Error() != "farmer not found" {
					farmererr.MobileNumber = row[MOBILENOCOLUMN]
					farmererr.Error = "error in fetching user - " + err.Error()
					errs = append(errs, farmererr)
					continue
					//return errors.New("already mobile no updated")
				} else {
					fmt.Println("farmer not found")
				}

			}
			if farmermobile != nil {
				if farmermobile != nil {
					farmererr.MobileNumber = row[MOBILENOCOLUMN]
					farmererr.Error = "already mobile no updated"
					errs = append(errs, farmererr)
					continue
					//return errors.New("already mobile no updated")
				}
			}
			/*
				Mpkv and Jnkv Not required Organisation and Project it have only one organisation and Project

			*/
			// if row[ORGANISATIONCOLUMN] != "" {
			// 	fmt.Println("org=", row[ORGANISATIONCOLUMN])
			// 	_, ok := orgRefMap[row[ORGANISATIONCOLUMN]]
			// 	if !ok {
			// 		resOrganisation, err := s.Daos.GetSingleOrganisationWithNameV2(ctx, row[ORGANISATIONCOLUMN], searchPolicy)
			// 		if err != nil {
			// 			farmererr.MobileNumber = row[MOBILENOCOLUMN]
			// 			farmererr.Error = err.Error()
			// 			errs = append(errs, farmererr)
			// 			//	farmer.FarmerOrg, _ = primitive.ObjectIDFromHex(err.Error())
			// 			continue
			// 		}
			// 		orgRefMap[row[ORGANISATIONCOLUMN]] = resOrganisation.ID
			// 		fmt.Println("resOrganisationName=========>", resOrganisation.Name)
			// 		//if resOrganisation != nil {

			// 	}
			// 	farmer.FarmerOrg = orgRefMap[row[ORGANISATIONCOLUMN]]

			// }

			// if row[PROJECTCOLUMN] != "" {
			// 	_, ok := projectRefMap[row[PROJECTCOLUMN]]
			// 	fmt.Println("project=", projectRefMap[row[PROJECTCOLUMN]])
			// 	if !ok {
			// 		resProject, err := s.Daos.GetSingleprojectWithName(ctx, row[PROJECTCOLUMN], farmer.FarmerOrg, searchPolicy)
			// 		if err != nil {
			// 			farmererr.MobileNumber = row[MOBILENOCOLUMN]
			// 			farmererr.Error = err.Error()
			// 			errs = append(errs, farmererr)
			// 			//	farmer.ProjectID, _ = primitive.ObjectIDFromHex(err.Error())

			// 			continue
			// 		}
			// 		if resProject != nil {
			// 			projectRefMap[row[PROJECTCOLUMN]] = resProject.ID
			// 			fmt.Println("resProjectName=========>", resProject.Name)

			// 		}
			// 	}
			// 	globalProjectId = projectRefMap[row[PROJECTCOLUMN]]
			// 	farmer.ProjectID = projectRefMap[row[PROJECTCOLUMN]]

			// }
			//}
			// if s, err := strconv.ParseFloat(row[AREAINACRECOLUMN], 64); err == nil {
			// 	farmerLand.AreaInAcre = s
			// }

			farmer.Name = row[FATHERNAMECOLUMN] + row[LASTNAMECOLUMN]
			farmer.FatherName = row[FATHERNAMECOLUMN]
			farmer.Gender = row[GENDERCOLUMN]
			farmer.MobileNumber = row[MOBILENOCOLUMN]
			farmer.Status = constants.FARMERSTATUSACTIVE
			farmer.ActiveStatus = true
			farmer.UploadVersion = row[CREATEDDATECOLUMN] + "-" + s.Daos.GetUniqueID(ctx, constants.COLLECTIONFARMERUPLOAD)
			createrid, err := primitive.ObjectIDFromHex("5a068710e4b093427c5af853")
			if err != nil {
				return err
			}
			farmer.CreatedBy = createrid
			farmer.UserName = "FAR_" + farmer.MobileNumber

			// farmer.PinCode = row[PINCODECOLUMN]
			farmer.FarmerID, farmer.UniqueId = s.Daos.GetUniqueIDV2(ctx, constants.COLLECTIONFARMER)
			if row[STATECOLUMN] != "" {
				_, ok := stateRefMap[row[STATECOLUMN]]
				fmt.Println("state==>", stateRefMap[row[STATECOLUMN]])
				if !ok {
					resState, err := s.Daos.GetSingleStateWithNameV2(ctx, row[STATECOLUMN], searchPolicy)
					if err != nil {
						farmererr.MobileNumber = row[MOBILENOCOLUMN]
						farmererr.Error = err.Error()
						errs = append(errs, farmererr)
						//farmer.State, _ = primitive.ObjectIDFromHex(err.Error())
						continue
					}
					// if resState == nil {
					// 	state := new(models.State)
					// 	state.Name = row[STATECOLUMN]
					// 	state.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSTATE)
					// 	state.Status = constants.STATESTATUSACTIVE
					// 	state.ActiveStatus = true
					// 	t := time.Now()
					// 	created := models.Created{}
					// 	created.On = &t
					// 	created.By = constants.SYSTEM
					// 	log.Println("b4 state.created")
					// 	state.Created = created
					// 	err := s.Daos.SaveState(ctx, state)
					// 	if err != nil {
					// 		return err
					// 	}
					// 	stateRefMap[row[STATECOLUMN]] = state.ID
					// 	continue
					// }
					if resState != nil {
						fmt.Println("state==>", resState.Name)
						stateRefMap[row[STATECOLUMN]] = resState.ID
					}
				}
				farmer.State = stateRefMap[row[STATECOLUMN]]

			}
			if row[DISTRICTCOLUMN] != "" {
				_, ok := districtRefMap[row[DISTRICTCOLUMN]]
				fmt.Println("district==>", districtRefMap[row[DISTRICTCOLUMN]])

				if !ok {
					resDistrict, err := s.Daos.GetSingleDistrictWithNameV2(ctx, row[DISTRICTCOLUMN], farmer.State, searchPolicy)
					if err != nil {
						farmererr.MobileNumber = row[MOBILENOCOLUMN]
						farmererr.Error = err.Error()
						errs = append(errs, farmererr)
						//farmer.District, _ = primitive.ObjectIDFromHex(err.Error())
						continue
					}
					if resDistrict != nil {
						fmt.Println("district==>", resDistrict.Name)
						districtRefMap[row[DISTRICTCOLUMN]] = resDistrict.ID

					}
				}
				farmer.District = districtRefMap[row[DISTRICTCOLUMN]]
			}
			//
			if row[BLOCKCOLUMN] != "" {
				_, ok := blockRefMap[row[BLOCKCOLUMN]]
				fmt.Println("block==>", blockRefMap[row[BLOCKCOLUMN]])
				if !ok {
					resBlock, err := s.Daos.GetSingleBlockWithNameV2(ctx, row[BLOCKCOLUMN], farmer.District, searchPolicy)
					if err != nil {
						farmererr.MobileNumber = row[MOBILENOCOLUMN]
						farmererr.Error = err.Error()
						errs = append(errs, farmererr)
						//	farmer.Block, _ = primitive.ObjectIDFromHex(err.Error())
						continue
					}
					if resBlock != nil {
						fmt.Println("block==>", resBlock.Name)
						blockRefMap[row[BLOCKCOLUMN]] = resBlock.ID

					}
				}
				farmer.Block = blockRefMap[row[BLOCKCOLUMN]]
			}
			// if row[GRAMPANCHAYATCOLUMN] != "" {
			// 	_, ok := grampRefMap[row[GRAMPANCHAYATCOLUMN]]
			// 	fmt.Println("grampanchat==>", blockRefMap[row[BLOCKCOLUMN]])
			// 	if !ok {
			// 		resGramPanchayat, err := s.Daos.GetSingleGrampanchayatWithNameV2(ctx, row[GRAMPANCHAYATCOLUMN], farmer.Block, searchPolicy)
			// 		if err != nil {
			// 			farmererr.MobileNumber = row[MOBILENOCOLUMN]
			// 			farmererr.Error = err.Error()
			// 			errs = append(errs, farmererr)
			// 			//	farmer.GramPanchayat, _ = primitive.ObjectIDFromHex(err.Error())
			// 			continue
			// 		}
			// 		if resGramPanchayat != nil {
			// 			fmt.Println("grampanchat==>", resGramPanchayat.Name)
			// 			grampRefMap[row[GRAMPANCHAYATCOLUMN]] = resGramPanchayat.ID

			// 		}
			// 	}
			// 	farmer.GramPanchayat = grampRefMap[row[GRAMPANCHAYATCOLUMN]]
			// }
			if row[VILLAGECOLUMN] != "" {
				_, ok := villageRefMap[row[VILLAGECOLUMN]]
				fmt.Println("village==>", villageRefMap[row[VILLAGECOLUMN]])
				if !ok {
					resVillage, err := s.Daos.GetSingleVillageWithNameV3(ctx, row[VILLAGECOLUMN], searchPolicy)
					if err != nil {
						farmererr.MobileNumber = row[MOBILENOCOLUMN]
						farmererr.Error = err.Error()
						errs = append(errs, farmererr)
						//	farmer.Village, _ = primitive.ObjectIDFromHex(err.Error())
						continue
					}
					if resVillage != nil {
						fmt.Println("village==>", resVillage.Name)
						villageRefMap[row[VILLAGECOLUMN]] = resVillage.ID
						grampRefMap[row[VILLAGECOLUMN]] = resVillage.Ref.GramPanchayat.ID

					}
				}
				farmer.Village = villageRefMap[row[VILLAGECOLUMN]]
				farmer.GramPanchayat = grampRefMap[row[VILLAGECOLUMN]]
			}
			//}

			//}

			//	}
			//}

			if row[CREATEDDATECOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t2, err := time.Parse(layout, row[CREATEDDATECOLUMN])
				if err != nil {
					farmer.CreatedDate = &t
				}
				farmer.CreatedDate = &t2
			} else {
				farmer.CreatedDate = &t

			}
			farmererr.MobileNumber = row[MOBILENOCOLUMN]
			farmererr.Error = "Successfully Registered"
			errs = append(errs, farmererr)
			farmers = append(farmers, farmer)

		}
		fmt.Println("finished arranging data")

		// if len(errs) != 0 {
		// 	farmererr.MobileNumber = row[MOBILENOCOLUMN]
		// 	farmererr.Error = "Success"
		// 	errs = append(errs, farmererr)
		// 	return errors.New("Success")
		// }

		// if len(errs) == 0 {
		tempFarmers := make([]interface{}, 0)
		// tempProjectFarmers := make([]interface{}, 0)
		// farmerProjectUpdate := make(map[primitive.ObjectID][]primitive.ObjectID)

		for k, farmer := range farmers {

			fmt.Println("farmerId====>", farmer.FarmerID)
			// err := s.Daos.SaveFarmer(ctx, &farmer)
			// if err != nil {
			// 	return err
			// }
			tempFarmers = append(tempFarmers, farmer)
			switch version {
			case "v2":
				// farmerProjectUpdate[farmer.ProjectID] = append(farmerProjectUpdate[farmer.ProjectID], farmer.ID)
			case "v3":
				// projectFarmer := new(models.ProjectFarmer)
				// projectFarmer.Farmer = farmer.ID
				// projectFarmer.Project = farmer.ProjectID
				// farmer.Status = constants.PROJECTFARMERSTATUSACTIVE
				// t := time.Now()
				// created := models.CreatedV2{}
				// created.On = &t
				// created.By = constants.SYSTEM
				// projectFarmer.Created = &created
				// // if err := s.Daos.SaveProjectFarmer(ctx, projectFarmer); err != nil {
				// // 	return err
				// // }
				// tempProjectFarmers = append(tempProjectFarmers, projectFarmer)
			}

			if k+1%1000 == 0 || k == len(farmers)-1 {
				fmt.Println((k+1)*1000, "time ===================================")
				farmerIds, err := s.Daos.SaveManyFarmers(ctx, tempFarmers)
				if err != nil {
					return errors.New("Errorin updating farmer  - " + err.Error())
				}
				switch version {
				case "v2":

					if err := s.Daos.LegacyUpdateProjectUsers(ctx, globalProjectId, farmerIds); err != nil {
						return errors.New("Errorin updating farmer projects v2 - " + err.Error())
					}

				case "v3":
					tempProjectFarmers := make([]interface{}, 0)
					for _, insertedFarmerID := range farmerIds {
						projectFarmer := new(models.ProjectFarmer)
						projectFarmer.Farmer = insertedFarmerID.(primitive.ObjectID)
						projectFarmer.Project = globalProjectId
						projectFarmer.Status = constants.PROJECTFARMERSTATUSACTIVE
						farmer.Status = constants.PROJECTFARMERSTATUSACTIVE
						t := time.Now()
						created := models.CreatedV2{}
						created.On = &t
						created.By = constants.SYSTEM
						projectFarmer.Created = &created
						// if err := s.Daos.SaveProjectFarmer(ctx, projectFarmer); err != nil {
						// 	return err
						// }
						tempProjectFarmers = append(tempProjectFarmers, projectFarmer)
					}

					err = s.Daos.SaveManyProjectFarmer(ctx, tempProjectFarmers)
					if err != nil {
						return errors.New("Errorin updating farmer projects v3 - " + err.Error())
					}
				}

				tempFarmers = make([]interface{}, 0)
				// farmerProjectUpdate = make(map[primitive.ObjectID][]primitive.ObjectID)

			}
			//return errors.New("Success")
		}

		// for k, _ := range farmers {
		// 	farmererr.MobileNumber = farmers[k].MobileNumber
		// 	farmererr.Error = "Success"
		// 	errs = append(errs, farmererr)
		// }
		fmt.Println("farmers====>", len(farmers))
		// dberr := s.Daos.FarmerAggregationUploadExcel(ctx, farmers)
		// if dberr != nil {
		// 	return errors.New("Db Error" + dberr.Error())
		// }

		return nil

	}); err != nil {
		farmererr.Error = err.Error()
		errs = append(errs, farmererr)
		return errs
	}
	return errs
}

// FarmerAggregationUploadExcelWithNamesV2 : ""
func (s *Service) FarmerAggregationUploadExcelWithNamesV2(ctx *models.Context, file multipart.File, searchPolicy bool, version string) []models.FarmerUploadError {
	log.Println("transaction start")
	t := time.Now()
	//Start Transaction
	orgRefMap := make(map[string]primitive.ObjectID)
	projectRefMap := make(map[string]primitive.ObjectID)
	//projectRefMap := make(map[string]primitive.ObjectID)
	stateRefMap := make(map[string]primitive.ObjectID)
	districtRefMap := make(map[string]primitive.ObjectID)
	blockRefMap := make(map[string]primitive.ObjectID)
	grampRefMap := make(map[string]primitive.ObjectID)
	villageRefMap := make(map[string]primitive.ObjectID)
	defer func() {
		for k := range orgRefMap {
			delete(orgRefMap, k)
		}
		for k := range projectRefMap {
			delete(projectRefMap, k)
		}
		for k := range stateRefMap {
			delete(stateRefMap, k)
		}
		for k := range districtRefMap {
			delete(districtRefMap, k)
		}
		for k := range blockRefMap {
			delete(blockRefMap, k)
		}
		for k := range grampRefMap {
			delete(grampRefMap, k)
		}
		for k := range villageRefMap {
			delete(villageRefMap, k)
		}
	}()
	var errs []models.FarmerUploadError
	var farmererr models.FarmerUploadError
	if err := ctx.Session.StartTransaction(); err != nil {
		farmererr.Error = err.Error()
		errs = append(errs, farmererr)
		return errs
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN = 8
		OMITROWS  = 0
		// ORGANISATIONCOLUMN  = 0
		// PROJECTCOLUMN       = 1
		FIRSTNAMECOLUMN = 0
		// FATHERNAMECOLUMN  = 1
		// LASTNAMECOLUMN    = 2
		GENDERCOLUMN      = 1
		MOBILENOCOLUMN    = 2
		VILLAGECOLUMN     = 3
		BLOCKCOLUMN       = 4
		DISTRICTCOLUMN    = 5
		STATECOLUMN       = 6
		CREATEDDATECOLUMN = 7

		CREATEDDATELAYOUT = "02-January-2006"
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("started reading file")
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		farmers := make([]models.Farmer, 0)
		rows := f.GetRows("Sheet1")
		//var errors []string
		fmt.Println("started looping")
		var globalProjectId primitive.ObjectID
		var Orgid primitive.ObjectID
		var ProjectId primitive.ObjectID

		productconfig, err := s.Daos.GetactiveProductConfig(ctx, true)
		if err != nil {
			return err
		}
		if productconfig != nil {
			if &productconfig.Orgnisation != nil {
				//	farmer.FarmerOrg
				Orgid = productconfig.Orgnisation.OrgnisationID
				//farmer.ProjectID
			}
			if &productconfig.Project != nil {
				ProjectId = productconfig.Project.ProjectID

			}
		}
		for rowIndex, row := range rows {
			fmt.Println("row no === ", rowIndex)
			if rowIndex <= OMITROWS {
				continue
			}

			if len(row) < MAXCOLUMN {
				return errors.New("Excel is not upto the format")
			}
			farmer := models.Farmer{}
			if row[MOBILENOCOLUMN] == "" {
				farmererr.MobileNumber = row[MOBILENOCOLUMN]
				farmererr.Error = "mobile no cannot be empty"
				errs = append(errs, farmererr)
				continue
			}
			farmer.FarmerOrg = Orgid
			farmer.ProjectID = ProjectId
			farmermobile, err := s.Daos.GetSingleFarmerWithMobilenoAndOrg(ctx, farmer.FarmerOrg.Hex(), row[MOBILENOCOLUMN])
			if err != nil {
				if err.Error() != "farmer not found" {
					farmererr.MobileNumber = row[MOBILENOCOLUMN]
					farmererr.Error = "error in fetching user - " + err.Error()
					errs = append(errs, farmererr)
					continue
					//return errors.New("already mobile no updated")
				} else {
					fmt.Println("farmer not found")
				}

			}
			if farmermobile != nil {
				if farmermobile != nil {
					farmererr.MobileNumber = row[MOBILENOCOLUMN]
					farmererr.Error = "already mobile no updated"
					errs = append(errs, farmererr)
					continue
					//return errors.New("already mobile no updated")
				}
			}
			/*
				Mpkv and Jnkv Not required Organisation and Project it have only one organisation and Project

			*/
			// if row[ORGANISATIONCOLUMN] != "" {
			// 	fmt.Println("org=", row[ORGANISATIONCOLUMN])
			// 	_, ok := orgRefMap[row[ORGANISATIONCOLUMN]]
			// 	if !ok {
			// 		resOrganisation, err := s.Daos.GetSingleOrganisationWithNameV2(ctx, row[ORGANISATIONCOLUMN], searchPolicy)
			// 		if err != nil {
			// 			farmererr.MobileNumber = row[MOBILENOCOLUMN]
			// 			farmererr.Error = err.Error()
			// 			errs = append(errs, farmererr)
			// 			//	farmer.FarmerOrg, _ = primitive.ObjectIDFromHex(err.Error())
			// 			continue
			// 		}
			// 		orgRefMap[row[ORGANISATIONCOLUMN]] = resOrganisation.ID
			// 		fmt.Println("resOrganisationName=========>", resOrganisation.Name)
			// 		//if resOrganisation != nil {

			// 	}
			// 	farmer.FarmerOrg = orgRefMap[row[ORGANISATIONCOLUMN]]

			// }

			// if row[PROJECTCOLUMN] != "" {
			// 	_, ok := projectRefMap[row[PROJECTCOLUMN]]
			// 	fmt.Println("project=", projectRefMap[row[PROJECTCOLUMN]])
			// 	if !ok {
			// 		resProject, err := s.Daos.GetSingleprojectWithName(ctx, row[PROJECTCOLUMN], farmer.FarmerOrg, searchPolicy)
			// 		if err != nil {
			// 			farmererr.MobileNumber = row[MOBILENOCOLUMN]
			// 			farmererr.Error = err.Error()
			// 			errs = append(errs, farmererr)
			// 			//	farmer.ProjectID, _ = primitive.ObjectIDFromHex(err.Error())

			// 			continue
			// 		}
			// 		if resProject != nil {
			// 			projectRefMap[row[PROJECTCOLUMN]] = resProject.ID
			// 			fmt.Println("resProjectName=========>", resProject.Name)

			// 		}
			// 	}
			// 	globalProjectId = projectRefMap[row[PROJECTCOLUMN]]
			// 	farmer.ProjectID = projectRefMap[row[PROJECTCOLUMN]]

			// }
			//}
			// if s, err := strconv.ParseFloat(row[AREAINACRECOLUMN], 64); err == nil {
			// 	farmerLand.AreaInAcre = s
			// }

			farmer.Name = row[FIRSTNAMECOLUMN]
			farmer.Gender = row[GENDERCOLUMN]
			farmer.MobileNumber = row[MOBILENOCOLUMN]
			farmer.Status = constants.FARMERSTATUSACTIVE
			farmer.ActiveStatus = true
			farmer.UploadVersion = row[CREATEDDATECOLUMN] + "-" + s.Daos.GetUniqueID(ctx, constants.COLLECTIONFARMERUPLOAD)
			createrid, err := primitive.ObjectIDFromHex("5a068710e4b093427c5af853")
			if err != nil {
				return err
			}
			farmer.CreatedBy = createrid
			farmer.UserName = "FAR_" + farmer.MobileNumber

			// farmer.PinCode = row[PINCODECOLUMN]
			farmer.FarmerID, farmer.UniqueId = s.Daos.GetUniqueIDV2(ctx, constants.COLLECTIONFARMER)
			if row[STATECOLUMN] != "" {
				_, ok := stateRefMap[row[STATECOLUMN]]
				fmt.Println("state==>", stateRefMap[row[STATECOLUMN]])
				if !ok {
					resState, err := s.Daos.GetSingleStateWithNameV2(ctx, row[STATECOLUMN], searchPolicy)
					if err != nil {
						farmererr.MobileNumber = row[MOBILENOCOLUMN]
						farmererr.Error = err.Error()
						errs = append(errs, farmererr)
						//farmer.State, _ = primitive.ObjectIDFromHex(err.Error())
						continue
					}
					// if resState == nil {
					// 	state := new(models.State)
					// 	state.Name = row[STATECOLUMN]
					// 	state.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSTATE)
					// 	state.Status = constants.STATESTATUSACTIVE
					// 	state.ActiveStatus = true
					// 	t := time.Now()
					// 	created := models.Created{}
					// 	created.On = &t
					// 	created.By = constants.SYSTEM
					// 	log.Println("b4 state.created")
					// 	state.Created = created
					// 	err := s.Daos.SaveState(ctx, state)
					// 	if err != nil {
					// 		return err
					// 	}
					// 	stateRefMap[row[STATECOLUMN]] = state.ID
					// 	continue
					// }
					if resState != nil {
						fmt.Println("state==>", resState.Name)
						stateRefMap[row[STATECOLUMN]] = resState.ID
					}
				}
				farmer.State = stateRefMap[row[STATECOLUMN]]

			}
			if row[DISTRICTCOLUMN] != "" {
				_, ok := districtRefMap[row[DISTRICTCOLUMN]]
				fmt.Println("district==>", districtRefMap[row[DISTRICTCOLUMN]])

				if !ok {
					resDistrict, err := s.Daos.GetSingleDistrictWithNameV2(ctx, row[DISTRICTCOLUMN], farmer.State, searchPolicy)
					if err != nil {
						farmererr.MobileNumber = row[MOBILENOCOLUMN]
						farmererr.Error = err.Error()
						errs = append(errs, farmererr)
						//farmer.District, _ = primitive.ObjectIDFromHex(err.Error())
						continue
					}
					if resDistrict != nil {
						fmt.Println("district==>", resDistrict.Name)
						districtRefMap[row[DISTRICTCOLUMN]] = resDistrict.ID

					}
				}
				farmer.District = districtRefMap[row[DISTRICTCOLUMN]]
			}
			//
			if row[BLOCKCOLUMN] != "" {
				_, ok := blockRefMap[row[BLOCKCOLUMN]]
				fmt.Println("block==>", blockRefMap[row[BLOCKCOLUMN]])
				if !ok {
					resBlock, err := s.Daos.GetSingleBlockWithNameV2(ctx, row[BLOCKCOLUMN], farmer.District, searchPolicy)
					if err != nil {
						farmererr.MobileNumber = row[MOBILENOCOLUMN]
						farmererr.Error = err.Error()
						errs = append(errs, farmererr)
						//	farmer.Block, _ = primitive.ObjectIDFromHex(err.Error())
						continue
					}
					if resBlock != nil {
						fmt.Println("block==>", resBlock.Name)
						blockRefMap[row[BLOCKCOLUMN]] = resBlock.ID

					}
				}
				farmer.Block = blockRefMap[row[BLOCKCOLUMN]]
			}
			// if row[GRAMPANCHAYATCOLUMN] != "" {
			// 	_, ok := grampRefMap[row[GRAMPANCHAYATCOLUMN]]
			// 	fmt.Println("grampanchat==>", blockRefMap[row[BLOCKCOLUMN]])
			// 	if !ok {
			// 		resGramPanchayat, err := s.Daos.GetSingleGrampanchayatWithNameV2(ctx, row[GRAMPANCHAYATCOLUMN], farmer.Block, searchPolicy)
			// 		if err != nil {
			// 			farmererr.MobileNumber = row[MOBILENOCOLUMN]
			// 			farmererr.Error = err.Error()
			// 			errs = append(errs, farmererr)
			// 			//	farmer.GramPanchayat, _ = primitive.ObjectIDFromHex(err.Error())
			// 			continue
			// 		}
			// 		if resGramPanchayat != nil {
			// 			fmt.Println("grampanchat==>", resGramPanchayat.Name)
			// 			grampRefMap[row[GRAMPANCHAYATCOLUMN]] = resGramPanchayat.ID

			// 		}
			// 	}
			// 	farmer.GramPanchayat = grampRefMap[row[GRAMPANCHAYATCOLUMN]]
			// }
			if row[VILLAGECOLUMN] != "" {
				_, ok := villageRefMap[row[VILLAGECOLUMN]]
				fmt.Println("village==>", villageRefMap[row[VILLAGECOLUMN]])
				if !ok {
					resVillage, err := s.Daos.GetSingleVillageWithNameV3(ctx, row[VILLAGECOLUMN], searchPolicy)
					if err != nil {
						farmererr.MobileNumber = row[MOBILENOCOLUMN]
						farmererr.Error = err.Error()
						errs = append(errs, farmererr)
						//	farmer.Village, _ = primitive.ObjectIDFromHex(err.Error())
						continue
					}
					if resVillage != nil {
						fmt.Println("village==>", resVillage.Name)
						villageRefMap[row[VILLAGECOLUMN]] = resVillage.ID
						grampRefMap[row[VILLAGECOLUMN]] = resVillage.Ref.GramPanchayat.ID

					}
				}
				farmer.Village = villageRefMap[row[VILLAGECOLUMN]]
				farmer.GramPanchayat = grampRefMap[row[VILLAGECOLUMN]]
			}
			//}

			//}

			//	}
			//}

			if row[CREATEDDATECOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t2, err := time.Parse(layout, row[CREATEDDATECOLUMN])
				if err != nil {
					farmer.CreatedDate = &t
				}
				farmer.CreatedDate = &t2
			} else {
				farmer.CreatedDate = &t

			}
			farmererr.MobileNumber = row[MOBILENOCOLUMN]
			farmererr.Error = "Successfully Registered"
			errs = append(errs, farmererr)
			farmers = append(farmers, farmer)

		}
		fmt.Println("finished arranging data")

		// if len(errs) != 0 {
		// 	farmererr.MobileNumber = row[MOBILENOCOLUMN]
		// 	farmererr.Error = "Success"
		// 	errs = append(errs, farmererr)
		// 	return errors.New("Success")
		// }

		// if len(errs) == 0 {
		tempFarmers := make([]interface{}, 0)
		// tempProjectFarmers := make([]interface{}, 0)
		// farmerProjectUpdate := make(map[primitive.ObjectID][]primitive.ObjectID)

		for k, farmer := range farmers {

			fmt.Println("farmerId====>", farmer.FarmerID)
			// err := s.Daos.SaveFarmer(ctx, &farmer)
			// if err != nil {
			// 	return err
			// }
			tempFarmers = append(tempFarmers, farmer)
			switch version {
			case "v2":
				// farmerProjectUpdate[farmer.ProjectID] = append(farmerProjectUpdate[farmer.ProjectID], farmer.ID)
			case "v3":
				// projectFarmer := new(models.ProjectFarmer)
				// projectFarmer.Farmer = farmer.ID
				// projectFarmer.Project = farmer.ProjectID
				// farmer.Status = constants.PROJECTFARMERSTATUSACTIVE
				// t := time.Now()
				// created := models.CreatedV2{}
				// created.On = &t
				// created.By = constants.SYSTEM
				// projectFarmer.Created = &created
				// // if err := s.Daos.SaveProjectFarmer(ctx, projectFarmer); err != nil {
				// // 	return err
				// // }
				// tempProjectFarmers = append(tempProjectFarmers, projectFarmer)
			}

			if k+1%1000 == 0 || k == len(farmers)-1 {
				fmt.Println((k+1)*1000, "time ===================================")
				farmerIds, err := s.Daos.SaveManyFarmers(ctx, tempFarmers)
				if err != nil {
					return errors.New("Errorin updating farmer  - " + err.Error())
				}
				switch version {
				case "v2":

					if err := s.Daos.LegacyUpdateProjectUsers(ctx, globalProjectId, farmerIds); err != nil {
						return errors.New("Errorin updating farmer projects v2 - " + err.Error())
					}

				case "v3":
					tempProjectFarmers := make([]interface{}, 0)
					for _, insertedFarmerID := range farmerIds {
						projectFarmer := new(models.ProjectFarmer)
						projectFarmer.Farmer = insertedFarmerID.(primitive.ObjectID)
						projectFarmer.Project = globalProjectId
						projectFarmer.Status = constants.PROJECTFARMERSTATUSACTIVE
						farmer.Status = constants.PROJECTFARMERSTATUSACTIVE
						t := time.Now()
						created := models.CreatedV2{}
						created.On = &t
						created.By = constants.SYSTEM
						projectFarmer.Created = &created
						// if err := s.Daos.SaveProjectFarmer(ctx, projectFarmer); err != nil {
						// 	return err
						// }
						tempProjectFarmers = append(tempProjectFarmers, projectFarmer)
					}

					err = s.Daos.SaveManyProjectFarmer(ctx, tempProjectFarmers)
					if err != nil {
						return errors.New("Errorin updating farmer projects v3 - " + err.Error())
					}
				}

				tempFarmers = make([]interface{}, 0)
				// farmerProjectUpdate = make(map[primitive.ObjectID][]primitive.ObjectID)

			}
			//return errors.New("Success")
		}

		// for k, _ := range farmers {
		// 	farmererr.MobileNumber = farmers[k].MobileNumber
		// 	farmererr.Error = "Success"
		// 	errs = append(errs, farmererr)
		// }
		fmt.Println("farmers====>", len(farmers))
		// dberr := s.Daos.FarmerAggregationUploadExcel(ctx, farmers)
		// if dberr != nil {
		// 	return errors.New("Db Error" + dberr.Error())
		// }

		return nil

	}); err != nil {
		farmererr.Error = err.Error()
		errs = append(errs, farmererr)
		return errs
	}
	return errs
}

//ImportExcelFileForFarmerLandV2 :""
func (s *Service) FarmerLandUploadExcelV2(ctx *models.Context, file multipart.File, searchPolicy bool) []models.FarmerUploadError {
	log.Println("transaction start")
	//Start Transaction
	var errs []models.FarmerUploadError
	var farmererr models.FarmerUploadError
	if err := ctx.Session.StartTransaction(); err != nil {
		farmererr.Error = err.Error()
		errs = append(errs, farmererr)
		return errs
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN             = 11
		OMITROWS              = 0
		FARMERIDCOLUMN        = 0
		VILLAGECOLUMN         = 1
		BLOCKCOLUMN           = 2
		DISTRICTCOLUMN        = 3
		STATECOLUMN           = 4
		AREAINACRECOLUMN      = 5
		CULTIVATEDAREACOLUMN  = 6
		VACANTAREACOLUMN      = 7
		IRRIGATIIONTYPECOLUMN = 8
		SOILTYPECOLUMN        = 9
		PARCELNOCOLUMN        = 10
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}

		farmerLands := make([]models.FarmerLand, 0)
		rows := f.GetRows("Sheet1")
		var Orgid primitive.ObjectID
		//var ProjectId primitive.ObjectID

		productconfig, err := s.Daos.GetactiveProductConfig(ctx, true)
		if err != nil {
			return err
		}
		if productconfig != nil {
			if &productconfig.Orgnisation != nil {
				//	farmer.FarmerOrg
				Orgid = productconfig.Orgnisation.OrgnisationID
				//farmer.ProjectID
			}

		}
		for rowIndex, row := range rows {
			if rowIndex <= OMITROWS {
				continue
			}
			if len(row) < MAXCOLUMN {
				return errors.New("Land Excel is not upto the format")
			}
			farmerLand := models.FarmerLand{}

			farmerLand.FarmerID = row[FARMERIDCOLUMN]
			//farmerLand.o = row[FARMERIDCOLUMN]
			farmer, err := s.Daos.GetSingleFarmerWithName(ctx, Orgid.Hex(), row[FARMERIDCOLUMN])
			if err != nil {
				farmererr.MobileNumber = row[FARMERIDCOLUMN]
				farmererr.Error = "farmer not found"
				errs = append(errs, farmererr)
				continue
			}
			farmerLand.Farmer = farmer.ID
			if s, err := strconv.ParseFloat(row[AREAINACRECOLUMN], 64); err == nil {
				farmerLand.AreaInAcre = s
			}

			if s, err := strconv.ParseFloat(row[VACANTAREACOLUMN], 64); err == nil {
				farmerLand.VacantArea = s
			}
			farmerLand.IrrigationType = row[IRRIGATIIONTYPECOLUMN]

			farmerLand.ParcelNumber = row[PARCELNOCOLUMN]
			if row[SOILTYPECOLUMN] != "" {
				resSoilType, err := s.Daos.GetSingleSoilTypeWithName(ctx, row[SOILTYPECOLUMN])
				if err != nil {
					farmererr.MobileNumber = row[FARMERIDCOLUMN]
					farmererr.Error = err.Error()
					errs = append(errs, farmererr)
					continue
				}
				if len(resSoilType) > 0 {
					farmerLand.SoilType = resSoilType[0].ID

					fmt.Println("resSoilTypeName=========>", resSoilType[0].Name)
				}
			}
			if row[STATECOLUMN] != "" {
				resState, err := s.Daos.GetSingleStateWithNameV2(ctx, row[STATECOLUMN], searchPolicy)
				if err != nil {
					return err
				}
				if resState != nil {
					farmerLand.State = resState.ID
				}
			}
			if row[DISTRICTCOLUMN] != "" {
				resDistrict, err := s.Daos.GetSingleDistrictWithNameV2(ctx, row[DISTRICTCOLUMN], farmerLand.State, searchPolicy)
				if err != nil {
					farmererr.MobileNumber = row[FARMERIDCOLUMN]
					farmererr.Error = err.Error()
					errs = append(errs, farmererr)
					continue
				}
				if resDistrict != nil {
					farmerLand.District = resDistrict.ID
				}
			}
			if row[BLOCKCOLUMN] != "" {
				resBlock, err := s.Daos.GetSingleBlockWithNameV2(ctx, row[BLOCKCOLUMN], farmerLand.District, searchPolicy)
				if err != nil {
					farmererr.MobileNumber = row[FARMERIDCOLUMN]
					farmererr.Error = err.Error()
					errs = append(errs, farmererr)
					continue
				}
				if resBlock != nil {
					farmerLand.Block = resBlock.ID

				}
			}
			// if row[GRAMPANCHAYATCOLUMN] != "" {
			// 	resGramPanchayat, err := s.Daos.GetSingleGramPanchayatWithUniqueId(ctx, row[GRAMPANCHAYATCOLUMN])
			// 	if err != nil {
			// 		return err
			// 	}
			// 	if resGramPanchayat != nil {
			// 		farmerLand.GramPanchayat = resGramPanchayat.ID

			// 	}
			// }
			if row[VILLAGECOLUMN] != "" {
				resVillage, err := s.Daos.GetSingleVillageWithNameV3(ctx, row[VILLAGECOLUMN], searchPolicy)
				if err != nil {
					farmererr.MobileNumber = row[FARMERIDCOLUMN]
					farmererr.Error = err.Error()
					errs = append(errs, farmererr)
					continue
				}
				if resVillage != nil {
					farmerLand.Village = resVillage.ID
					farmerLand.GramPanchayat = resVillage.Ref.GramPanchayat.ID
				}
			}
			farmerLand.Status = constants.FARMERLANDSTATUSACTIVE

			farmerLands = append(farmerLands, farmerLand)
			farmererr.MobileNumber = row[FARMERIDCOLUMN]
			farmererr.Error = "Successfully Upload"
			errs = append(errs, farmererr)
		}
		for _, farmerland := range farmerLands {
			fmt.Println("farmerId====>", farmerland.FarmerID)
			// fmt.Println("SoilType====>", farmerland.SoilType)
		}
		//farmererr.MobileNumber = row[MOBILENOCOLUMN]

		//	farmers = append(farmers, farmer)
		fmt.Println("farmerlands====>", len(farmerLands))
		dberr := s.Daos.FarmerLandUploadExcel(ctx, farmerLands)
		if dberr != nil {
			return errors.New("Db Error" + dberr.Error())
		}

		return nil

	}); err != nil {
		return errs
	}
	return errs
}
