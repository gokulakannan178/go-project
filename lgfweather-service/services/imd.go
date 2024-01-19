package services

// func (s *Service) LoadIMDDistrictWeather() {
// 	config := config.NewConfig("districtimd", "config")
// 	maxLength := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxlength)

// 	// stateNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDstate)
// 	// districtNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDdistrict)
// 	latNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDlat)
// 	longNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDlong)
// 	altoNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDalto)
// 	pcodNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDpcod)
// 	dayNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDday)
// 	monthNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmonth)
// 	yearNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDyear)
// 	msplNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmspl)
// 	icidNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDicid)
// 	rainfallNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDrainfall)
// 	maxtempNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxtemp)
// 	mintempNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmintemp)
// 	maxrelhumNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxrelhum)
// 	minremhumNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDminremhum)
// 	windspeedNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDwindspeed)
// 	winddirectionNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDwinddirection)
// 	client, err := ftp.Dial("103.215.208.49:21")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	if err := client.Login("imdgfs", "imdgfs@2012"); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	var date string
// 	var param1 string
// 	imdurl := "DIST_BLOCK_FT1534/2020/dfcst/"
// 	param1 = "madhya-pradesh"
// 	file := "/dfcst2000z"
// 	var t time.Time
// 	t = time.Date(2022, 8, 1, 0, 0, 0, 0, t.Location())
// 	month := s.Shared.GetMonthInt(t.Month().String())
// 	yearstr := fmt.Sprintf("%v", t.Year())
// 	years := strings.Split(yearstr, "20")
// 	year := years[1]
// 	if t.Day() < 10 {
// 		date = fmt.Sprintf("%v%v0%v", year, month, t.Day())
// 	} else {
// 		date = fmt.Sprintf("%v%v%v", year, month, t.Day())
// 	}
// 	gokul := fmt.Sprintf("%v%v%v%v", imdurl, param1, file, date)

// 	fmt.Println("URL===>", gokul)
// 	//	r, err := client.Retr("andaman-and-nicobar-220408")

// 	//r, err := client.Retr("DIST_BLOCK_FT1534/2020/dfcst/andaman-and-nicobar-islands/dfcst00z220408")
// 	r, err := client.Retr(gokul)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer r.Close()

// 	buf, err := ioutil.ReadAll(r)
// 	// println(string(buf))
// 	//var lbwr []models.LoadBlockWeatherReport
// 	lines := strings.Split(string(buf), "\n")

// 	for _, v := range lines {
// 		var err error
// 		words := strings.Fields(v)
// 		if len(words) != maxLength {
// 			fmt.Println("length not proper")
// 			continue
// 		}
// 		dwd := models.DistrictWeatherDataV2{}
// 		// state := words[stateNo]
// 		// district := words[districtNo]
// 		latStr := words[longNo]
// 		lat, err := strconv.ParseFloat(latStr, 64)
// 		if err != nil {
// 			log.Println("Error in geting latStr", err.Error())
// 			err = nil
// 		}
// 		longStr := words[latNo]
// 		long, err := strconv.ParseFloat(longStr, 64)
// 		if err != nil {
// 			log.Println("Error in geting longStr", err.Error())
// 			err = nil
// 		}
// 		dwd.Location = models.Location{
// 			Type:        "point",
// 			Coordinates: []float64{long, lat},
// 		}
// 		alto := words[altoNo]
// 		dwd.Alto, err = strconv.ParseFloat(alto, 64)
// 		if err != nil {
// 			log.Println("Error in geting alto", err.Error())
// 			err = nil
// 		}
// 		dwd.Pcod = words[pcodNo]

// 		yearStr := words[yearNo]
// 		year, err := strconv.ParseInt(yearStr[0:4], 10, 0)
// 		if err != nil {
// 			log.Println("Error in geting year", err.Error())
// 			err = nil
// 		}
// 		monthStr := words[monthNo]
// 		month, err := strconv.ParseInt(monthStr, 10, 0)
// 		if err != nil {
// 			log.Println("Error in geting year", err.Error())
// 			err = nil
// 		}
// 		dayStr := words[dayNo]
// 		day, err := strconv.ParseInt(dayStr, 10, 0)
// 		if err != nil {
// 			log.Println("Error in geting year", err.Error())
// 			err = nil
// 		}
// 		loc, _ := time.LoadLocation("Asia/Kolkata")

// 		t := time.Date(int(year), time.Month(int(month)), int(day), 0, 0, 0, 0, loc)
// 		dwd.Date = &t
// 		mspl := words[msplNo]
// 		dwd.Mspl, err = strconv.ParseFloat(mspl, 64)
// 		if err != nil {
// 			log.Println("Error in geting mspl", err.Error())
// 			err = nil
// 		}
// 		icid := words[icidNo]
// 		dwd.Icid, err = strconv.ParseFloat(icid, 64)
// 		if err != nil {
// 			log.Println("Error in geting icid", err.Error())
// 			err = nil
// 		}
// 		rainfall := words[rainfallNo]
// 		dwd.Rain, err = strconv.ParseFloat(rainfall, 64)
// 		if err != nil {
// 			log.Println("Error in geting rainfall", err.Error())
// 			err = nil
// 		}
// 		maxtemp := words[maxtempNo]
// 		dwd.Temp.Max, err = strconv.ParseFloat(maxtemp, 64)
// 		if err != nil {
// 			log.Println("Error in geting maxtemp", err.Error())
// 			err = nil
// 		}
// 		mintemp := words[mintempNo]
// 		dwd.Temp.Min, err = strconv.ParseFloat(mintemp, 64)
// 		if err != nil {
// 			log.Println("Error in geting mintemp", err.Error())
// 			err = nil
// 		}
// 		maxrelhum := words[maxrelhumNo]
// 		dwd.HumidityMax, err = strconv.ParseFloat(maxrelhum, 64)
// 		if err != nil {
// 			log.Println("Error in geting maxrelhum", err.Error())
// 			err = nil
// 		}
// 		minremhum := words[minremhumNo]
// 		dwd.HumidityMin, err = strconv.ParseFloat(minremhum, 64)
// 		if err != nil {
// 			log.Println("Error in geting minremhum", err.Error())
// 			err = nil
// 		}
// 		windspeed := words[windspeedNo]
// 		dwd.Wind_speed, err = strconv.ParseFloat(windspeed, 64)
// 		if err != nil {
// 			log.Println("Error in geting windspeed", err.Error())
// 			err = nil
// 		}
// 		winddirection := words[winddirectionNo]
// 		dwd.Wind_deg, err = strconv.ParseFloat(winddirection, 64)
// 		if err != nil {
// 			log.Println("Error in geting winddirection", err.Error())
// 			err = nil
// 		}
// 		c := context.TODO()
// 		ctx := app.GetApp(c, s.Daos)
// 		defer ctx.Client.Disconnect(c)
// 		fmt.Println("=================", dwd)

// 	}
// 	//	err = s.Daos.LoadBlockWeatherReport(ctx)
// 	if err != nil {
// 		return
// 	}
// 	//	fmt.Println("weather data===>", r)
// 	if err := client.Quit(); err != nil {
// 		log.Fatal(err)
// 	}
// }
