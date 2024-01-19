package services

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveState :""
func (s *Service) SaveState(ctx *models.Context, state *models.State) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	state.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSTATE)
	state.Status = constants.STATESTATUSACTIVE
	state.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 state.created")
	state.Created = created
	log.Println("b4 state.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveState(ctx, state)
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

//UpdateState : ""
func (s *Service) UpdateState(ctx *models.Context, state *models.State) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateState(ctx, state)
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

//EnableState : ""
func (s *Service) EnableState(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableState(ctx, UniqueID)
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

//DisableState : ""
func (s *Service) DisableState(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableState(ctx, UniqueID)
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

//DeleteState : ""
func (s *Service) DeleteState(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteState(ctx, UniqueID)
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

//GetSingleState :""
func (s *Service) GetSingleState(ctx *models.Context, UniqueID string) (*models.RefState, error) {
	state, err := s.Daos.GetSingleState(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return state, nil
}

//FilterState :""
func (s *Service) FilterState(ctx *models.Context, statefilter *models.StateFilter, pagination *models.Pagination) (state []models.RefState, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	if statefilter != nil {

		dataaccess, err := s.Daos.DataAccess(ctx, &statefilter.DataAccess)
		if err != nil {
			return nil, err
		}
		if dataaccess != nil {
			if len(dataaccess.AccessStates) > 0 {
				for _, v := range dataaccess.AccessStates {
					statefilter.ID = append(statefilter.ID, v.ID)
				}
			}

		}
	}

	return s.Daos.FilterState(ctx, statefilter, pagination)

}
func (s *Service) GeoDetatilsReport(ctx *models.Context, filter *models.StateFilter) ([]models.GeoDetailsReport2, error) {

	res, err := s.Daos.GeoDetatilsReportV2(ctx, filter)
	if err != nil {
		return nil, err
	}
	// var DS []models.GeoDetailsReport2
	// var V *models.GeoDetailsReport2
	// for _, V := range DS {
	// 	res.Districts = append(res.Districts, V.Districts...)
	// 	res.Block = append(res.Block, V.Block...)
	// 	res.GramPanchayat = append(res.GramPanchayat, V.GramPanchayat...)
	// 	res.Village = append(res.Village, V.Village...)
	// }

	return res, nil
}

func (s *Service) GeoDetatilsReportExcel(ctx *models.Context, filter *models.StateFilter) (*excelize.File, error) {
	data, err := s.GeoDetatilsReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Geo Details Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "J1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "state")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "stateCode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "distric")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "districCode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "block")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "blockCode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "grampanchayat")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "grampanchayatCode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "village")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "villageCode")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		if len(v.Districts) == 0 {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
			rowNo++
			continue
		}

		for _, v2 := range v.Districts {
			if len(v2.Block) == 0 {
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.Name)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.UniqueID)
				rowNo++
				continue
			}

			for _, v3 := range v2.Block {
				if len(v3.GramPanchayat) == 0 {
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.Name)
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.UniqueID)
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v3.Name)
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v3.UniqueID)
					rowNo++
					continue
				}
				for _, v4 := range v3.GramPanchayat {
					if len(v4.Village) == 0 {
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.UniqueID)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v3.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v3.UniqueID)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v4.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v4.UniqueID)
						rowNo++
						continue
					}
					for _, v5 := range v4.Village {

						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.UniqueID)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v3.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v3.UniqueID)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v4.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v4.UniqueID)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v5.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v5.UniqueID)
						rowNo++
					}
				}
			}
		}

	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}
func (s *Service) GeoDetatilsReportExcelV2(ctx *models.Context, filter *models.StateFilter) (*excelize.File, error) {
	data, err := s.GeoDetatilsReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Geo Details Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "E1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "state")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "distric")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "block")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "grampanchayat")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "village")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		if len(v.Districts) == 0 {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
			rowNo++
			continue
		}

		for _, v2 := range v.Districts {
			if len(v2.Block) == 0 {
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.Name)
				rowNo++
				continue
			}

			for _, v3 := range v2.Block {
				if len(v3.GramPanchayat) == 0 {
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.Name)
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v3.Name)
					rowNo++
					continue
				}
				for _, v4 := range v3.GramPanchayat {
					if len(v4.Village) == 0 {
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v3.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v4.Name)
						rowNo++
						continue
					}
					for _, v5 := range v4.Village {
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v3.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v4.Name)
						excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v5.Name)
						rowNo++
					}
				}
			}
		}

	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}
func (s *Service) GeoUploadExcelWithNames(ctx *models.Context, file multipart.File, isRegex bool) error {
	log.Println("transaction start")
	//t := time.Now()
	//Start Transaction
	stateRefMap := make(map[string]primitive.ObjectID)
	districtRefMap := make(map[string]primitive.ObjectID)
	blockRefMap := make(map[string]primitive.ObjectID)
	grampRefMap := make(map[string]primitive.ObjectID)
	villageRefMap := make(map[string]primitive.ObjectID)
	var err error

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN           = 4
		OMITROWS            = 0
		STATECOLUMN         = 0
		DISTRICTCOLUMN      = 1
		BLOCKCOLUMN         = 2
		GRAMPANCHAYATCOLUMN = 3
		VILLAGECOLUMN       = 4
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("started reading file")
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
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
			//fmt.Println("isregex===>", isRegex)

			if row[STATECOLUMN] != "" {
				_, ok := stateRefMap[row[STATECOLUMN]]
				if !ok {
					resState, err := s.Daos.GetSingleStateWithNameV2(ctx, row[STATECOLUMN], isRegex)
					if err != nil {
						state := new(models.State)
						state.Name = row[STATECOLUMN]
						state.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSTATE)
						state.Status = constants.STATESTATUSACTIVE
						state.ActiveStatus = true
						t := time.Now()
						created := models.Created{}
						created.On = &t
						created.By = constants.SYSTEM
						state.Created = created
						err := s.Daos.SaveState(ctx, state)
						if err != nil {
							return errors.New("state not save")

						}
						//fmt.Println("State Id====>", state.ID)
						stateRefMap[row[STATECOLUMN]] = state.ID
						fmt.Println("state err====>", err)

					}
					if resState != nil {
						stateRefMap[row[STATECOLUMN]] = resState.ID
					}
				}
				fmt.Println("stateId====>", stateRefMap[row[STATECOLUMN]])

			}
			if row[DISTRICTCOLUMN] != "" {
				_, ok := districtRefMap[row[DISTRICTCOLUMN]]
				if !ok {
					resDistrict, _ := s.Daos.GetSingleDistrictWithNameV2(ctx, row[DISTRICTCOLUMN], stateRefMap[row[STATECOLUMN]], isRegex)
					if resDistrict == nil {
						distric := new(models.District)
						distric.Name = row[DISTRICTCOLUMN]
						distric.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDISTRICT)
						distric.Status = constants.STATESTATUSACTIVE
						distric.ActiveStatus = true
						distric.State = stateRefMap[row[STATECOLUMN]]
						t := time.Now()
						created := models.Created{}
						created.On = &t
						created.By = constants.SYSTEM
						log.Println("b4 distric.created")
						distric.Created = created
						err := s.Daos.SaveDistrict(ctx, distric)
						if err != nil {
							return errors.New("distric not save")
						}

						districtRefMap[row[DISTRICTCOLUMN]] = distric.ID

					}
					if resDistrict != nil {
						districtRefMap[row[DISTRICTCOLUMN]] = resDistrict.ID

					}

				}
				//farmer.District = districtRefMap[row[DISTRICTCOLUMN]]
				fmt.Println("districId====>", districtRefMap[row[DISTRICTCOLUMN]])

			}
			//
			if row[BLOCKCOLUMN] != "" {
				_, ok := blockRefMap[row[BLOCKCOLUMN]]
				if !ok {
					fmt.Println("block====>", row[BLOCKCOLUMN])
					resBlock, _ := s.Daos.GetSingleBlockWithNameV2(ctx, row[BLOCKCOLUMN], districtRefMap[row[DISTRICTCOLUMN]], isRegex)
					if resBlock == nil {
						Block := new(models.Block)
						Block.Name = row[BLOCKCOLUMN]
						Block.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBLOCK)
						Block.Status = constants.STATESTATUSACTIVE
						Block.ActiveStatus = true
						Block.District = districtRefMap[row[DISTRICTCOLUMN]]
						t := time.Now()
						created := models.CreatedV2{}
						created.On = &t
						created.By = constants.SYSTEM
						log.Println("b4 block.created")
						Block.Created = created
						err := s.Daos.SaveBlock(ctx, Block)
						if err != nil {
							return errors.New("block not save")
						}
						blockRefMap[row[BLOCKCOLUMN]] = Block.ID

					}
					if resBlock != nil {
						blockRefMap[row[BLOCKCOLUMN]] = resBlock.ID

					}

				}
				//farmer.Block = blockRefMap[row[BLOCKCOLUMN]]
				fmt.Println("blockId====>", blockRefMap[row[BLOCKCOLUMN]])

			}
			if row[GRAMPANCHAYATCOLUMN] != "" {
				_, ok := grampRefMap[row[GRAMPANCHAYATCOLUMN]]
				if !ok {
					fmt.Println("grampancat====>", row[GRAMPANCHAYATCOLUMN])
					resGramPanchayat, _ := s.Daos.GetSingleGrampanchayatWithNameV2(ctx, row[GRAMPANCHAYATCOLUMN], blockRefMap[row[BLOCKCOLUMN]], isRegex)
					if resGramPanchayat == nil {
						GramPanchayat := new(models.GramPanchayat)
						GramPanchayat.Name = row[GRAMPANCHAYATCOLUMN]
						GramPanchayat.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONGRAMPANCHAYAT)
						GramPanchayat.Status = constants.STATESTATUSACTIVE
						GramPanchayat.ActiveStatus = true
						GramPanchayat.Block = blockRefMap[row[BLOCKCOLUMN]]
						t := time.Now()
						created := models.CreatedV2{}
						created.On = &t
						created.By = constants.SYSTEM
						log.Println("b4 grampancat.created")
						GramPanchayat.Created = created
						err := s.Daos.SaveGramPanchayat(ctx, GramPanchayat)
						if err != nil {
							return errors.New("grampancht not save")
						}
						grampRefMap[row[GRAMPANCHAYATCOLUMN]] = GramPanchayat.ID

					}
					if resGramPanchayat != nil {
						grampRefMap[row[GRAMPANCHAYATCOLUMN]] = resGramPanchayat.ID

					}

				}
				fmt.Println("grampancatId====>", grampRefMap[row[GRAMPANCHAYATCOLUMN]])

				//farmer.GramPanchayat = grampRefMap[row[GRAMPANCHAYATCOLUMN]]
			}
			if row[VILLAGECOLUMN] != "" {
				_, ok := villageRefMap[row[VILLAGECOLUMN]]
				if !ok {
					fmt.Println("village====>", row[VILLAGECOLUMN])
					resVillage, _ := s.Daos.GetSingleVillageWithNameV2(ctx, row[VILLAGECOLUMN], grampRefMap[row[GRAMPANCHAYATCOLUMN]], isRegex)
					if resVillage == nil {
						village := new(models.Village)
						village.Name = row[VILLAGECOLUMN]
						village.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONGRAMPANCHAYAT)
						village.Status = constants.STATESTATUSACTIVE
						village.GramPanchayat = grampRefMap[row[GRAMPANCHAYATCOLUMN]]
						village.ActiveStatus = true
						t := time.Now()
						created := models.CreatedV2{}
						created.On = &t
						created.By = constants.SYSTEM
						log.Println("b4 village.created")
						village.Created = created
						err := s.Daos.SaveVillage(ctx, village)
						if err != nil {
							return errors.New("village not save")
						}
						villageRefMap[row[VILLAGECOLUMN]] = village.ID

					}
					if resVillage != nil {
						villageRefMap[row[VILLAGECOLUMN]] = resVillage.ID

					}

				}
				fmt.Println("villageId====>", villageRefMap[row[VILLAGECOLUMN]])
				//farmer.Village = villageRefMap[row[GRAMPANCHAYATCOLUMN]]
			}

		}
		fmt.Println("finished arranging data")
		return err

	}); err != nil {

		return err
	}
	return err
}
func (s *Service) GetWeatherDataWithSeverityType(ctx *models.Context, statefilter *models.StateWeatherAlertMasterFilterv2, pagination *models.Pagination) (state []models.GetStateLeveWeatherDataAlert, err error) {
	return s.Daos.GetWeatherDataWithSeverityType(ctx, statefilter, pagination)

}
