package services

// FarmerAggregationUploadExcel : ""
// func (s *Service) FarmerAggregationUploadExcel(ctx *models.Context, file multipart.File) []models.PropertyUploadError {
// 	log.Println("transaction start")
// 	//Start Transaction
// 	orgRefMap := make(map[string]primitive.ObjectID)
// 	projectRefMap := make(map[string]primitive.ObjectID)
// 	stateRefMap := make(map[string]primitive.ObjectID)
// 	districtRefMap := make(map[string]primitive.ObjectID)
// 	blockRefMap := make(map[string]primitive.ObjectID)
// 	grampRefMap := make(map[string]primitive.ObjectID)
// 	villageRefMap := make(map[string]primitive.ObjectID)
// 	var errs []models.PropertyUploadError
// 	var propertyErr models.PropertyUploadError
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		propertyErr.Error = err.Error()
// 		errs = append(errs, propertyErr)
// 		return errs
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	const (
// 		MAXCOLUMN                 = 13
// 		OMITROWS                  = 0
// 		ORGANISATIONCOLUMN        = 0
// 		PROPERTYNOCOLUMN          = 1
// 		OLDHOLDINGNOCOLUMN        = 2
// 		PROPERTYTYPECOLUMN        = 3
// 		ISMATCHEDCOLUMN           = 4
// 		ROADONWHICHLOCATIONCOLUMN = 5
// 		DOACOLUMN                 = 6
// 		AREAOFPLOTCOLUMN          = 7
// 		BUILDUPAREACOLUMN         = 8
// 		HOUSENOCOLUMN             = 9
// 		PLOTNOCOLUMN              = 10
// 		KHATANOCOLUMN             = 11
// 		AL1COLUMN                 = 12
// 		AL2COLUMN                 = 13
// 		STATECOLUMN               = 14
// 		DISTRICTCOLUMN            = 15
// 		CITYVILLAGECOLUMN         = 16
// 		ZONECOLUMN                = 17
// 		WARDCOLUMN                = 18
// 		PINCODECOLUMN             = 19
// 		CREATEDDATELAYOUT         = "02-January-2006"
// 	)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		fmt.Println("started reading file")
// 		f, err := excelize.OpenReader(file)
// 		if err != nil {
// 			return err
// 		}
// 		propertys := make([]models.Property, 0)
// 		rows := f.GetRows("Sheet1")
// 		//var errors []string
// 		fmt.Println("started looping")

// 		for rowIndex, row := range rows {
// 			fmt.Println("row no === ", rowIndex)
// 			if rowIndex <= OMITROWS {
// 				continue
// 			}

// 			if len(row) < MAXCOLUMN {
// 				return errors.New("Excel is not upto the format")
// 			}
// 			property := models.Property{}
// 			farmermobile, err := s.Daos.GetSinglePropertyOwnerWithMobileNo(ctx, row[MOBILENOCOLUMN])
// 			if err != nil {
// 				if row[ORGANISATIONCOLUMN] != "" {
// 					_, ok := orgRefMap[row[ORGANISATIONCOLUMN]]
// 					if !ok {
// 						resOrganisation, err := s.Daos.GetSingleOrganisationWithUniqueID(ctx, row[ORGANISATIONCOLUMN])
// 						if err != nil {
// 							farmererr.MobileNumber = row[MOBILENOCOLUMN]
// 							farmererr.Error = err.Error()
// 							errs = append(errs, farmererr)
// 							//	farmer.FarmerOrg, _ = primitive.ObjectIDFromHex(err.Error())
// 							continue
// 						}
// 						orgRefMap[row[ORGANISATIONCOLUMN]] = resOrganisation.ID
// 						fmt.Println("resOrganisationName=========>", resOrganisation.Name)
// 						//if resOrganisation != nil {

// 					}
// 					farmer.FarmerOrg = orgRefMap[row[ORGANISATIONCOLUMN]]

// 				}
// 				if row[PROJECTCOLUMN] != "" {
// 					_, ok := projectRefMap[row[PROJECTCOLUMN]]
// 					if !ok {
// 						resProject, err := s.Daos.GetSingleProjectWithUniqueID(ctx, row[PROJECTCOLUMN])
// 						if err != nil {
// 							farmererr.MobileNumber = row[MOBILENOCOLUMN]
// 							farmererr.Error = err.Error()
// 							errs = append(errs, farmererr)
// 							//	farmer.ProjectID, _ = primitive.ObjectIDFromHex(err.Error())

// 							continue
// 						}
// 						if resProject != nil {
// 							projectRefMap[row[PROJECTCOLUMN]] = resProject.ID
// 							fmt.Println("resProjectName=========>", resProject.Name)

// 						}
// 					}
// 					farmer.ProjectID = projectRefMap[row[PROJECTCOLUMN]]

// 				}
// 				//}

// 				// if s, err := strconv.ParseFloat(row[AREAINACRECOLUMN], 64); err == nil {
// 				// 	farmerLand.AreaInAcre = s
// 				// }
// 				farmer.Name = row[NAMECOLUMN]
// 				farmer.FatherName = row[FATHERNAMECOLUMN]
// 				farmer.Gender = row[GENDERCOLUMN]
// 				farmer.MobileNumber = row[MOBILENOCOLUMN]
// 				farmer.Status = constants.FARMERSTATUSACTIVE
// 				farmer.FarmerID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFARMER)
// 				if row[STATECOLUMN] != "" {
// 					_, ok := stateRefMap[row[STATECOLUMN]]
// 					if !ok {
// 						resState, err := s.Daos.GetSingleStateWithUniqueID(ctx, row[STATECOLUMN])
// 						if err != nil {
// 							farmererr.MobileNumber = row[MOBILENOCOLUMN]
// 							farmererr.Error = err.Error()
// 							errs = append(errs, farmererr)
// 							//farmer.State, _ = primitive.ObjectIDFromHex(err.Error())
// 							continue
// 						}
// 						if resState != nil {

// 							stateRefMap[row[STATECOLUMN]] = resState.ID
// 						}
// 					}
// 					farmer.State = stateRefMap[row[STATECOLUMN]]

// 				}
// 				if row[DISTRICTCOLUMN] != "" {
// 					_, ok := districtRefMap[row[DISTRICTCOLUMN]]
// 					if !ok {
// 						resDistrict, err := s.Daos.GetSingleDistrictWithUniqueId(ctx, row[DISTRICTCOLUMN])
// 						if err != nil {
// 							farmererr.MobileNumber = row[MOBILENOCOLUMN]
// 							farmererr.Error = err.Error()
// 							errs = append(errs, farmererr)
// 							//farmer.District, _ = primitive.ObjectIDFromHex(err.Error())
// 							continue
// 						}
// 						if resDistrict != nil {
// 							districtRefMap[row[DISTRICTCOLUMN]] = resDistrict.ID

// 						}
// 					}
// 					farmer.District = districtRefMap[row[DISTRICTCOLUMN]]
// 				}
// 				//
// 				if row[BLOCKCOLUMN] != "" {
// 					_, ok := blockRefMap[row[BLOCKCOLUMN]]
// 					if !ok {
// 						resBlock, err := s.Daos.GetSingleBlockWithUnique(ctx, row[BLOCKCOLUMN])
// 						if err != nil {
// 							farmererr.MobileNumber = row[MOBILENOCOLUMN]
// 							farmererr.Error = err.Error()
// 							errs = append(errs, farmererr)
// 							//	farmer.Block, _ = primitive.ObjectIDFromHex(err.Error())
// 							continue
// 						}
// 						if resBlock != nil {
// 							blockRefMap[row[BLOCKCOLUMN]] = resBlock.ID

// 						}
// 					}
// 					farmer.Block = blockRefMap[row[BLOCKCOLUMN]]
// 				}
// 				if row[GRAMPANCHAYATCOLUMN] != "" {
// 					_, ok := grampRefMap[row[GRAMPANCHAYATCOLUMN]]
// 					if !ok {
// 						resGramPanchayat, err := s.Daos.GetSingleGramPanchayatWithUniqueId(ctx, row[GRAMPANCHAYATCOLUMN])
// 						if err != nil {
// 							farmererr.MobileNumber = row[MOBILENOCOLUMN]
// 							farmererr.Error = err.Error()
// 							errs = append(errs, farmererr)
// 							//	farmer.GramPanchayat, _ = primitive.ObjectIDFromHex(err.Error())
// 							continue
// 						}
// 						if resGramPanchayat != nil {
// 							grampRefMap[row[GRAMPANCHAYATCOLUMN]] = resGramPanchayat.ID

// 						}
// 					}
// 					farmer.GramPanchayat = grampRefMap[row[GRAMPANCHAYATCOLUMN]]
// 				}
// 				if row[VILLAGECOLUMN] != "" {
// 					_, ok := villageRefMap[row[GRAMPANCHAYATCOLUMN]]
// 					if !ok {
// 						resVillage, err := s.Daos.GetSingleVillageWithUniqueId(ctx, row[VILLAGECOLUMN])
// 						if err != nil {
// 							farmererr.MobileNumber = row[MOBILENOCOLUMN]
// 							farmererr.Error = err.Error()
// 							errs = append(errs, farmererr)
// 							//	farmer.Village, _ = primitive.ObjectIDFromHex(err.Error())
// 							continue
// 						}
// 						if resVillage != nil {
// 							villageRefMap[row[GRAMPANCHAYATCOLUMN]] = resVillage.ID

// 						}
// 					}
// 					farmer.Village = villageRefMap[row[GRAMPANCHAYATCOLUMN]]
// 				}
// 				//}

// 				//}

// 				//	}
// 				//}

// 				if row[CREATEDDATECOLUMN] != "" {
// 					layout := CREATEDDATELAYOUT
// 					t, err := time.Parse(layout, row[CREATEDDATECOLUMN])
// 					if err != nil {
// 						return err
// 					}
// 					farmer.CreatedDate = &t
// 				}
// 				farmers = append(farmers, farmer)

// 			}
// 			fmt.Println("finished arranging data")

// 			// if len(errs) != 0 {
// 			// 	farmererr.MobileNumber = row[MOBILENOCOLUMN]
// 			// 	farmererr.Error = "Success"
// 			// 	errs = append(errs, farmererr)
// 			// 	return errors.New("Success")
// 			// }
// 			if farmermobile != nil {
// 				farmererr.MobileNumber = row[MOBILENOCOLUMN]
// 				farmererr.Error = "already mobile no updated"
// 				errs = append(errs, farmererr)
// 				//return errors.New("already mobile no updated")
// 			}

// 		}
// 		if len(errs) == 0 {
// 			tempFarmers := make([]interface{}, 0)
// 			tempProjectFarmers := make([]interface{}, 0)
// 			for k, farmer := range farmers {

// 				fmt.Println("farmerId====>", farmer.FarmerID)
// 				// err := s.Daos.SaveFarmer(ctx, &farmer)
// 				// if err != nil {
// 				// 	return err
// 				// }
// 				tempFarmers = append(tempFarmers, farmer)
// 				projectFarmer := new(models.ProjectFarmer)
// 				projectFarmer.Farmer = farmer.ID
// 				projectFarmer.Project = farmer.ProjectID
// 				farmer.Status = constants.PROJECTFARMERSTATUSACTIVE
// 				t := time.Now()
// 				created := models.CreatedV2{}
// 				created.On = &t
// 				created.By = constants.SYSTEM
// 				projectFarmer.Created = &created
// 				// if err := s.Daos.SaveProjectFarmer(ctx, projectFarmer); err != nil {
// 				// 	return err
// 				// }
// 				tempProjectFarmers = append(tempProjectFarmers, projectFarmer)
// 				if k+1%1000 == 0 || k == len(farmers)-1 {
// 					fmt.Println((k+1)*1000, "time ===================================")
// 					_, err := s.Daos.SaveManyFarmers(ctx, tempFarmers)
// 					if err != nil {
// 						return err
// 					}
// 					err = s.Daos.SaveManyProjectFarmer(ctx, tempProjectFarmers)
// 					if err != nil {
// 						return err
// 					}
// 					tempFarmers = make([]interface{}, 0)
// 					tempProjectFarmers = make([]interface{}, 0)
// 				}
// 				//return errors.New("Success")
// 			}
// 		}
// 		// for k, _ := range farmers {
// 		// 	farmererr.MobileNumber = farmers[k].MobileNumber
// 		// 	farmererr.Error = "Success"
// 		// 	errs = append(errs, farmererr)
// 		// }
// 		fmt.Println("farmers====>", len(farmers))
// 		// dberr := s.Daos.FarmerAggregationUploadExcel(ctx, farmers)
// 		// if dberr != nil {
// 		// 	return errors.New("Db Error" + dberr.Error())
// 		// }

// 		return nil

// 	}); err != nil {
// 		farmererr.Error = err.Error()
// 		errs = append(errs, farmererr)
// 		return errs
// 	}
// 	return errs
// }
