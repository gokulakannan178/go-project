package handlers

// // ULBLogin : ""
// func (h *Handler) ULBLogin(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())

// 	user := new(models.Login)
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
// 		return
// 	}

// 	token, stat, err := h.Service.ULBLogin(ctx, user)
// 	log.Println("stat ==>", stat)
// 	//	log.Println("err ==>", err.Error())
// 	log.Println("TOKEN==>", token)
// 	if err != nil {
// 		if err.Error() == constants.NOTFOUND {
// 			response.With403mV2(w, "Invalid User", platform)
// 			return
// 		}
// 		response.With500mV2(w, err.Error(), platform)
// 		return
// 	}
// 	if !stat {
// 		response.With403mV2(w, "Invalid Username or Password", platform)
// 		return
// 	}
// 	role, respUser, err := h.Service.GetSingleUserAndRolewithUniqueID(ctx, user.UserName)
// 	if err != nil {
// 		log.Println("err=>", err.Error())
// 	}
// 	m := make(map[string]interface{})
// 	m["token"] = token
// 	m["user"] = respUser
// 	m["role"] = role
// 	response.With200V2(w, "Success", m, platform)
// }
