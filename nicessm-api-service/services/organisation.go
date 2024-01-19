package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveOrganisation :""
func (s *Service) SaveOrganisation(ctx *models.Context, organisation *models.Organisation) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	organisation.Status = constants.ORGANISATIONOWNERSTATUSACTIVE
	organisation.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 organisation.created")
	organisation.Created = created
	log.Println("b4 organisation.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		found, _ := s.Daos.ChkCommonUniqueness(ctx, constants.COLLECTIONORGANISATION, "name", organisation.Name)
		if found {
			return errors.New("organisation name already Registered")
		}
		dberr := s.Daos.SaveOrganisation(ctx, organisation)
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

//UpdateOrganisation : ""
func (s *Service) UpdateOrganisation(ctx *models.Context, organisation *models.Organisation) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOrganisation(ctx, organisation)
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

//EnableOrganisation : ""
func (s *Service) EnableOrganisation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableOrganisation(ctx, UniqueID)
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

//DisableOrganisation : ""
func (s *Service) DisableOrganisation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableOrganisation(ctx, UniqueID)
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

//DeleteOrganisation : ""
func (s *Service) DeleteOrganisation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOrganisation(ctx, UniqueID)
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

//GetSingleOrganisation :""
func (s *Service) GetSingleOrganisation(ctx *models.Context, UniqueID string) (*models.RefOrganisation, error) {
	organisation, err := s.Daos.GetSingleOrganisation(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return organisation, nil
}

//FilterOrganisation :""
func (s *Service) FilterOrganisation(ctx *models.Context, organisationfilter *models.OrganisationFilter, pagination *models.Pagination) (organisation []models.RefOrganisation, err error) {
	// if organisationfilter != nil {
	// 	if organisationfilter.UserAccess.Is {
	// 		user, err := s.GetAccessPrivillege(ctx, organisationfilter.UserAccess)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		if user.Type != constants.USERTYPESUPERADMIN {
	// 			organisationfilter.ID = append(organisationfilter.ID, user.UserOrg)
	// 		}
	// 	}
	// }
	if organisationfilter != nil {

		dataaccess, err := s.Daos.DataAccess(ctx, &organisationfilter.DataAccess)
		if err != nil {
			return nil, err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					organisationfilter.ID = append(organisationfilter.ID, v.ID)
				}

			}

		}

	}
	return s.Daos.FilterOrganisation(ctx, organisationfilter, pagination)
}
func (s *Service) OrganistationProjectDetails(ctx *models.Context, filter *models.StateFilter) ([]models.OrganistationProjectDetails, error) {

	res, err := s.Daos.OrganistationProjectDetails(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *Service) OrganistationProjectDetailsExcel(ctx *models.Context, filter *models.StateFilter) (*excelize.File, error) {
	data, err := s.OrganistationProjectDetails(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "OrganisationProject Details Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "D1")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "organisation")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "organisationCode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "project")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "projectCode")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		if len(v.Projects) == 0 {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
			rowNo++
			continue
		}

		for _, v2 := range v.Projects {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.UniqueID)
			rowNo++
			continue

		}

	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}
func (s *Service) OrganistationProjectDetailsExcelV2(ctx *models.Context, filter *models.StateFilter) (*excelize.File, error) {
	data, err := s.OrganistationProjectDetails(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "OrganisationProject Details Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B1")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "organisation")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "project")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		if len(v.Projects) == 0 {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
			rowNo++
			continue
		}

		for _, v2 := range v.Projects {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.Name)
			rowNo++
			continue

		}

	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}
