package services

import (
	"bpms-service/models"
	"fmt"
)

//ACLAccess :
func (s *Service) ACLAccess(ctx *models.Context, userTypeID string) (*models.ACLAccess, error) {

	var aclAccess = new(models.ACLAccess)
	utma, err := s.GetSingleModuleUserType(ctx, userTypeID)
	if err != nil {
		return nil, err
	}
	// aclAccess.Module = utma

	if utma != nil {
		if len(utma.Modules) > 0 {
			for _, v := range utma.Modules {
				var utmenua *models.UserTypeMenuAccess
				aclAccess.ModuleAccess = append(aclAccess.ModuleAccess, v)

				utmenua, err := s.GetSingleUserTypeMenuAccess(ctx, userTypeID, v.UniqueID)
				if err == nil {
					// aclAccess.Menu = append(aclAccess.Menu, *utmenua)
					if len(utmenua.Module.Menus) > 0 {
						for _, v2 := range utmenua.Module.Menus {
							aclAccess.MenuAccess = append(aclAccess.MenuAccess, v2)

						}
					} else {
						aclAccess.MenuAccess = make([]models.MenuAccess, 0)
					}
				}
				if err != nil {
					fmt.Println("err in menu find - menu - ", v.UniqueID, "user type")
				}

				uttaba, err := s.GetSingleUserTypeTabAccess(ctx, userTypeID, v.UniqueID)
				if err == nil {

					// aclAccess.Tab = append(aclAccess.Tab, *uttaba)

					if len(uttaba.Module.Tabs) > 0 {
						for _, v2 := range uttaba.Module.Tabs {
							aclAccess.TabAccess = append(aclAccess.TabAccess, v2)
						}
					} else {
						aclAccess.TabAccess = make([]models.TabAccess, 0)
					}
				}
				if err != nil {
					fmt.Println("err in menu find - tab - ", v.UniqueID, "user type", userTypeID)
				}

				utfeaturea, err := s.GetSingleUserTypeFeatureAccess(ctx, userTypeID, v.UniqueID)
				if err == nil {
					// aclAccess.Feature = append(aclAccess.Feature, *utfeaturea)
					if len(utfeaturea.Module.Features) > 0 {
						for _, v2 := range utfeaturea.Module.Features {
							aclAccess.FeatureAccess = append(aclAccess.FeatureAccess, v2)
						}
					} else {
						aclAccess.FeatureAccess = make([]models.FeatureAccess, 0)
					}
				}
				if err != nil {
					fmt.Println("err in menu find - feature - ", v.UniqueID, "user type", userTypeID)
				}
			}
		} else {
			aclAccess.ModuleAccess = make([]models.ModuleAccess, 0)
		}
	}
	return aclAccess, nil
}
