package services

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"

	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveHolidays :""
func (s *Service) SaveHolidays(ctx *models.Context, Holidays *models.Holidays) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	currentTime := Holidays.Date
	Day := currentTime.Day()
	Month := currentTime.Month()
	Year := currentTime.Year()
	strDay := strconv.Itoa(Day)
	strMonth := Month.String()
	strYear := strconv.Itoa(Year)
	UniqueID := strDay + strMonth + strYear
	Holidays.UniqueID = UniqueID
	Holidays.Status = constants.HOLIDAYSSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Holidays.created")
	Holidays.Created = created
	log.Println("b4 Holidays.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveHolidays(ctx, Holidays)
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

//GetSingleHolidays :""
func (s *Service) GetSingleHolidays(ctx *models.Context, UniqueID string) (*models.RefHolidays, error) {
	Holidays, err := s.Daos.GetSingleHolidaysWithOutStatus(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Holidays, nil
}

//UpdateHolidays : ""
func (s *Service) UpdateHolidays(ctx *models.Context, Holidays *models.Holidays) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateHolidays(ctx, Holidays)
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

//EnableHolidays : ""
func (s *Service) EnableHolidays(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableHolidays(ctx, UniqueID)
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

//DisableHolidays : ""
func (s *Service) DisableHolidays(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableHolidays(ctx, UniqueID)
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

//DeleteHolidays : ""
func (s *Service) DeleteHolidays(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteHolidays(ctx, UniqueID)
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

//FilterHolidays :""
func (s *Service) FilterHolidays(ctx *models.Context, holidaysfilter *models.FilterHolidays, pagination *models.Pagination) ([]models.RefHolidays, error) {
	err := s.HolidayDataAccess(ctx, holidaysfilter)
	if err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterHolidays(ctx, holidaysfilter, pagination)

}
func (s *Service) HolidayDataAccess(ctx *models.Context, filter *models.FilterHolidays) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationID = append(filter.OrganisationID, v.UniqueID)
				}
			}

		}

	}
	return err
}
func (s *Service) GetHolidaysWeek(ctx *models.Context, Holidays *models.FilterHolidays) ([]models.HolidaysList, error) {
	sd, ed, err := s.FindWeekStartAndEndDate(ctx, Holidays.Date.StartDate)
	if err != nil {
		return nil, err
	}
	Holidays.Date.StartDate = sd
	Holidays.Date.EndDate = ed
	var Holiday []models.HolidaysList
	// Holiday, err := s.Daos.GetHolidaysWithDays(ctx, Holidays)
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println("Holiday=====>", len(Holiday))
	hh := new(models.HolidaysList)
	//var va time.Time
	if len(Holiday) <= 0 {
		//va = *sd

		for rd := rangeDate(*sd, *ed); ; {
			date := rd()
			if date.IsZero() {
				break
			}
			hh.Date = &date
			Holiday = append(Holiday, *hh)
			fmt.Println(date.Format("2006-01-02"))
		}
		fmt.Println("final====>", Holiday)
	}
	for k, v := range Holiday {
		var hh models.HolidaysList
		fmt.Println("Holiday[k].Date======>", hh.Date)
		date := fmt.Sprintf("%v%v%v", v.Date.Day(), v.Date.Month(), v.Date.Year())
		hoildays, err := s.Daos.GetSingleArrayHolidays(ctx, date)
		if err != nil {
			return nil, err
		}
		fmt.Println("len hoildays======>", len(hoildays))

		brithemployee, err := s.Daos.GetBrithdayEmployees(ctx, nil, v.Date)
		if err != nil {
			return nil, err
		}
		anniemployee, err := s.Daos.GetAnnivesaryEmployee(ctx, nil, v.Date)
		if err != nil {
			return nil, err
		}
		Holiday[k].Holidays = hoildays

		Holiday[k].BrithdateEmployee = brithemployee
		Holiday[k].AnniversaryEmployee = anniemployee
		Holiday[k].Day = v.Date.Weekday().String()
	}
	return Holiday, nil
}
func (s *Service) HoildayUploadExcel(ctx *models.Context, file multipart.File) []models.EmployeeUploadError {
	log.Println("transaction start")
	//Start Transaction
	// orgRefMap := make(map[string]primitive.ObjectID)
	// projectRefMap := make(map[string]primitive.ObjectID)
	// stateRefMap := make(map[string]primitive.ObjectID)
	// districtRefMap := make(map[string]primitive.ObjectID)
	// blockRefMap := make(map[string]primitive.ObjectID)
	// grampRefMap := make(map[string]primitive.ObjectID)
	// villageRefMap := make(map[string]primitive.ObjectID)
	var errs []models.EmployeeUploadError
	var employeeerr models.EmployeeUploadError
	if err := ctx.Session.StartTransaction(); err != nil {
		employeeerr.Error = err.Error()
		errs = append(errs, employeeerr)
		return errs
	}
	defer ctx.Session.EndSession(ctx.CTX)
	const (
		MAXCOLUMN          = 3
		OMITROWS           = 0
		ORGANISATIONCOLUMN = 0
		DATECOLUMN         = 1
		NAMECOLUMN         = 2
		CREATEDDATELAYOUT  = "02-January-2006"
	)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("started reading file")
		f, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		holidays := make([]models.Holidays, 0)
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
			holiday := new(models.Holidays)

			organisation, _ := s.Daos.GetSingleActiveOrganisationWithName(ctx, row[ORGANISATIONCOLUMN])
			if organisation == nil {
				employeeerr.Name = row[NAMECOLUMN]
				employeeerr.UserName = row[DATECOLUMN]
				employeeerr.Error = "organisation Not Found"
				errs = append(errs, employeeerr)
				continue
			}
			if organisation != nil {
				holiday.OrganisationID = organisation.UniqueID
			}

			if row[DATECOLUMN] != "" {
				layout := CREATEDDATELAYOUT
				t, err := time.Parse(layout, row[DATECOLUMN])
				if err != nil {
					return err
				}
				holiday.Date = &t
			}
			holiday.Name = row[NAMECOLUMN]

			currentTime := holiday.Date
			Day := currentTime.Day()
			Month := currentTime.Month()
			Year := currentTime.Year()
			strDay := strconv.Itoa(Day)
			strMonth := Month.String()
			strYear := strconv.Itoa(Year)
			UniqueID := strDay + strMonth + strYear
			holiday.UniqueID = UniqueID
			holiday.Status = constants.HOLIDAYSSTATUSACTIVE
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 Holidays.created")
			holiday.Created = created
			log.Println("b4 Holidays.created")
			err = s.Daos.SaveHolidays(ctx, holiday)
			if err != nil {
				employeeerr.Name = row[NAMECOLUMN]
				employeeerr.UserName = row[DATECOLUMN]
				employeeerr.Error = err.Error()
				errs = append(errs, employeeerr)
				continue
			}
			//	employees = append(employees, *employee)
			if err == nil {
				employeeerr.Name = row[NAMECOLUMN]
				employeeerr.UserName = holiday.UniqueID
				employeeerr.Error = "Sucess"
				errs = append(errs, employeeerr)
				continue
			}
		}
		fmt.Println("no.of.employee==>", len(holidays))
		return nil

	}); err != nil {
		employeeerr.Error = err.Error()
		errs = append(errs, employeeerr)
		return errs
	}
	return errs
}
