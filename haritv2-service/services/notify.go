package services

import (
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"time"
)

// NotifyForUpdateLocation : ""
func (s *Service) NotifyForUpdateLocation(ctx *models.Context, notificationType string) {

	res, err := s.Daos.GutULBNotUpdatedLocation(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	tempmsg := "Hi %v nodal officer of %v Kindly update your location. from URNCC"
	// ulbIds := []string{}
	for _, v := range res {
		// ulbIds = append(ulbIds, v.UniqueID)
		msg := fmt.Sprintf(tempmsg, v.NodalOfficer.Name, v.Name)
		switch notificationType {
		case "SMS":
			err = s.SendSMS(v.NodalOfficer.MobileNo, msg)
			if err != nil {
				fmt.Println(err)
			}
		case "NOTIFY":
			refTokens, err := s.Daos.GetReGTokenOfSingleULB(ctx, v.UniqueID)
			if err != nil {
				fmt.Println(err)
			}

			title := "Location Update Remainder"
			body := msg
			image := ""
			fmt.Println(msg)
			if refTokens == nil {
				fmt.Println("no token available")
				continue
			}
			err2 := s.SendNotification(title, body, image, refTokens.RegistrationToken)
			if err != nil {
				fmt.Println(err2)
			}
		}

	}

}

// NotifyForUpdateLocation : ""
func (s *Service) NotifyForUpdateProfile(ctx *models.Context, notificationType string) {

	res, err := s.Daos.GutULBNotUpdatedProfile(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	tempmsg := "Hi %v nodal officer of %v Kindly update your location. from URNCC"
	// ulbIds := []string{}
	for _, v := range res {
		// ulbIds = append(ulbIds, v.UniqueID)
		msg := fmt.Sprintf(tempmsg, v.NodalOfficer.Name, v.Name)
		switch notificationType {
		case "SMS":
			err = s.SendSMS(v.NodalOfficer.MobileNo, msg)
			if err != nil {
				fmt.Println(err)
			}
		case "NOTIFY":
			refTokens, err := s.Daos.GetReGTokenOfSingleULB(ctx, v.UniqueID)
			if err != nil {
				fmt.Println(err)
			}

			title := "Profile Update Remainder"
			body := msg
			image := ""
			fmt.Println(msg)
			if refTokens == nil {
				fmt.Println("no token available")
				continue
			}
			err2 := s.SendNotification(title, body, image, refTokens.RegistrationToken)
			if err != nil {
				fmt.Println(err2)
			}
		}

	}

}

// ULBInventoryUpdateMessage : ""
func (s *Service) NotifyForULBInventoryUpdate(ctx *models.Context, notifyType string) {
	filter := new(models.ULBInventoryUpdateMessageFilter)
	t := time.Now()
	t = t.AddDate(0, 0, -2)
	filter.Date = &t
	res, err := s.Daos.GetULBNotUpdatedInventory(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res == nil {
		fmt.Println("ALL ulb updated inventory")
		return
	}
	// mobileNo := ""
	for _, v := range res {
		// fmt.Println("MobileNo = ", v.NoMobile)
		// if k == len(res.ULBData)-1 {
		// 	mobileNo = mobileNo + v.NoMobile
		// } else {
		// 	mobileNo = mobileNo + v.NoMobile + ","
		// }
		// tempmsg := "Hi %v nodal officer of %v Kindly update your inventory for %v %v from URNCC"
		tempmsg := "Hi %v nodal officer of %v Kindly update your inventory for %v %v from URNCC Thank You URNCCH"
		// msg := fmt.Sprintf(`Hi Nodal Officer , Your ulb has not updated inventory for the month %v %v.        Kindly update the stock  sale details of compost of in your ULB.    Thanks and Regards        URNCCH,URNCCS`, t.Month(), t.Year())
		msg := fmt.Sprintf(tempmsg, v.NodalOfficer.Name, v.Name, t.Month(), t.Year())
		switch notifyType {
		case constants.NOTIFICATIONTYPESMS:
			if v.NodalOfficer.MobileNo == "" {
				continue
			}
			err = s.SendSMS(v.NodalOfficer.MobileNo, msg)
			if err != nil {
				fmt.Println(err)
			}
		case constants.NOTIFICATIONTYPENOTIFY:
			refTokens, err := s.Daos.GetReGTokenOfSingleULB(ctx, v.UniqueID)
			if err != nil {
				fmt.Println(err)
			}
			title := "Inventory Update Remainder"
			body := msg
			image := ""
			fmt.Println(msg)
			if refTokens == nil || len(refTokens.RegistrationToken) == 0 {
				fmt.Println("no token available")
				continue
			}
			err2 := s.SendNotification(title, body, image, refTokens.RegistrationToken)
			if err2 != nil {
				fmt.Println(err2)
			}
		}

	}

	// fmt.Println("message = ", msg)
	// fmt.Println("mobile no = ", mobileNo)
	// mobileNo = mobileNo + ",7299424027"
	//mobileNo = "7299424027"
	//msg = "Your password is 9999 NICESM"

}

// NotifyForULBInventoryUpdateV2 : "For Custom Month"
func (s *Service) NotifyForULBInventoryUpdateV2(ctx *models.Context, filter *models.ULBInventoryUpdateMessageFilterV2) {

	res, err := s.Daos.GetULBNotUpdatedInventoryV2(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res == nil {
		fmt.Println("ALL ulb updated inventory")
		return
	}
	// mobileNo := ""
	for _, v := range res {

		tempmsg := "Hi %v nodal officer of %v Kindly update your inventory for %v %v from URNCC Thank You URNCCH"

		msg := fmt.Sprintf(tempmsg, v.NodalOfficer.Name, v.Name, filter.Month, filter.Year)
		switch filter.NotifyType {
		case constants.NOTIFICATIONTYPESMS:
			if v.NodalOfficer.MobileNo == "" {
				continue
			}
			err = s.SendSMS(v.NodalOfficer.MobileNo, msg)
			if err != nil {
				fmt.Println(err)
			}
		case constants.NOTIFICATIONTYPENOTIFY:
			refTokens, err := s.Daos.GetReGTokenOfSingleULB(ctx, v.UniqueID)
			if err != nil {
				fmt.Println(err)
			}
			title := "Inventory Update Remainder"
			body := msg
			image := ""
			fmt.Println(msg)
			if refTokens == nil || len(refTokens.RegistrationToken) == 0 {
				fmt.Println("no token available")
				continue
			}
			err2 := s.SendNotification(title, body, image, refTokens.RegistrationToken)
			if err2 != nil {
				fmt.Println(err2)
			}
		case constants.NOTIFICATIONTYPESMSV2:
			msg = fmt.Sprintf(constants.COMMONTEMPLATE, v.NodalOfficer.Name, "URNCCH", "Inventory Update of "+v.Name+" ULB", "Last date for updating inventory for "+filter.Month+" "+filter.Year+" is "+filter.LastDate, "http://urncc.org")
			fmt.Println(msg)
			if v.NodalOfficer.MobileNo == "" {
				continue
			}
			err = s.SendSMS(v.NodalOfficer.MobileNo, msg)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
