package daos

import (
	"errors"
	"fmt"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
)

func (d *Daos) DataAccess(ctx *models.Context, data *models.DataAccessRequest) (*models.DataAccess, error) {
	dataAccess := new(models.DataAccess)
	if data == nil {

	}
	if !data.IsAccess {

		dataAccess.Organisation = make([]models.RefOrganisation, 0)
		dataAccess.Projects = make([]models.RefProjectUser, 0)
		dataAccess.AccessStates = make([]models.State, 0)
		dataAccess.AccessDistricts = make([]models.District, 0)
		dataAccess.AccessBlocks = make([]models.Block, 0)
		dataAccess.AccessVillages = make([]models.Village, 0)
		dataAccess.AccessGrampanchayats = make([]models.GramPanchayat, 0)

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
		dataAccess.Projects = make([]models.RefProjectUser, 0)
		dataAccess.AccessStates = make([]models.State, 0)
		dataAccess.AccessDistricts = make([]models.District, 0)
		dataAccess.AccessBlocks = make([]models.Block, 0)
		dataAccess.AccessVillages = make([]models.Village, 0)
		dataAccess.AccessGrampanchayats = make([]models.GramPanchayat, 0)
		return dataAccess, nil

	}
	fmt.Println("getting organisation", user.UserOrg.Hex())
	organization, err := d.GetSingleOrganisation(ctx, user.UserOrg.Hex())
	if err != nil {
		return nil, errors.New("getsingle user not organisation error" + err.Error())
	}
	dataAccess.Organisation = append(dataAccess.Organisation, *organization)
	projectFilter := new(models.ProjectUserFilter)
	projectFilter.User = append(projectFilter.User, user.ID)

	filterProject, err := d.FilterProjectUser(ctx, projectFilter, nil)
	if err != nil {
		return nil, errors.New("getsingle user not Project error" + err.Error())
	}
	dataAccess.Projects = filterProject
	dataAccess.AccessStates = user.Ref.AccessStates
	dataAccess.AccessDistricts = user.Ref.AccessDistricts
	dataAccess.AccessBlocks = user.Ref.AccessBlocks
	dataAccess.AccessVillages = user.Ref.AccessVillages
	dataAccess.AccessGrampanchayats = user.Ref.AccessGrampanchayats
	fmt.Println("user.Ref.AccessDistricts====>", user.Ref.AccessDistricts)

	return dataAccess, nil
}
