package services

// //Login :
// func (s *Service) ULBLogin(ctx *models.Context, login *models.Login) (string, bool, error) {

// 	data, err := s.Daos.GetSingleUser(ctx, login.UserName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "dal err", false, err
// 	}
// 	if ok := data.Password == login.PassWord; !ok {
// 		log.Println("Data password ==>", data.Password)
// 		log.Println("login password ==>", login.PassWord)
// 		return "Passs false", false, nil
// 	}
// 	if data.Status == constants.USERSTATUSINIT {
// 		return "", false, errors.New("Awaiting Activation")
// 	}
// 	var auth models.Authentication
// 	auth.UserID = data.ID
// 	auth.UserName = data.UserName

// 	auth.Status = data.Status
// 	auth.Role = data.Role
// 	fmt.Println("auth user ==>", auth, data)

// 	token, err := CreateTokenV2(&auth)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return "", false, errors.New("Error in Generating Token - " + err.Error())
// 	}

// 	data.Token = token
// 	// data.CurrentLocation = login.Location
// 	err = s.Daos.UpdateUser(ctx, &data.User)
// 	if err != nil {
// 		log.Println("Error in saving token - " + err.Error())
// 		return "", false, errors.New(constants.INTERNALSERVERERROR)
// 	}
// 	return "", true, nil
// }
