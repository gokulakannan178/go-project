package daos

import (
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
)

func (d *Daos) DataAccess(ctx *models.Context, data *models.DataAccessRequest) (*models.DataAccess, error) {
	dataAccess := new(models.DataAccess)
	if data == nil {

	}
	if !data.IsAccess {

		dataAccess.Organisation = make([]models.RefOrganisation, 0)
		return dataAccess, nil

	}
	user, err := d.GetSingleUserWithUserName(ctx, data.UserName)
	if err != nil {
		return nil, errors.New("getsingle user not username error" + err.Error())
	}
	if user == nil {
		return nil, errors.New("loging user not vailable")
	}
	dataAccess.User = user
	if user.Type == constants.USERTYPESUPERADMIN {
		dataAccess.Organisation = make([]models.RefOrganisation, 0)
		dataAccess.SuperAdmin = true
		return dataAccess, nil
	} else if user.Type == constants.USERTYPEHR {
		//	dataAccess.Organisation = make([]models.RefOrganisation, 0)
		dataAccess.SuperAdmin = true
		return dataAccess, nil
	}
	fmt.Println("getting organisation", user.OrganisationID)
	organization, err := d.GetSingleOrganisation(ctx, user.OrganisationID)
	if err != nil {
		return nil, errors.New("getsingle user not organisation error" + err.Error())
	}
	dataAccess.Organisation = append(dataAccess.Organisation, *organization)
	if user.Type == constants.USERTYPEEMPLOYEE {
		employee, err := d.GetSingleEmployee(ctx, user.EmployeeId)
		if err != nil {
			return nil, err
		}
		dataAccess.Branch = employee.BranchID
		dataAccess.Department = employee.DepartmentID
		dataAccess.Designation = employee.DesignationID
		return dataAccess, nil
	}
	// dataAccess.AccessStates = user.Ref.AccessStates
	// dataAccess.AccessDistricts = user.Ref.AccessDistricts
	// dataAccess.AccessBlocks = user.Ref.AccessBlocks
	// dataAccess.AccessVillages = user.Ref.AccessVillages
	// dataAccess.AccessGrampanchayats = user.Ref.AccessGrampanchayats
	// fmt.Println("user.Ref.AccessDistricts====>", user.Ref.AccessDistricts)

	return dataAccess, nil
}
