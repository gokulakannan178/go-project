package services

import (
	"context"
	"fmt"
	"lgfweather-service/app"
	"lgfweather-service/config"
	"log"
)

func (s *Service) LoadIMDDistrictWeatherV2() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	// config := config.NewConfig("districtimd", "config")
	// maxLength := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxlength)
	lines, err := s.GetWeatherDataWithImd(ctx)
	if err != nil {
		log.Println(err)
	}
	err = s.SaveDistrictWeatherDataWithImd(ctx, lines)
	if err != nil {
		log.Println(err)
	}
}
func (s *Service) LoadIMDDistrictWeatherWithState() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	State, err := s.Daos.GetActiveState(ctx)
	if err != nil {
		log.Println("State Not Found")
	}
	fmt.Println("No.Of.State===>", len(State))
	for _, v := range State {
		config := config.NewConfig("districtimd", "config")
		fmt.Println(config)
		lines, err := s.GetWeatherDataWithImdWithState(ctx, v.ImdStateName, v.ImdFileName)
		if err != nil {
			log.Println(err)
		}

		state := v.ID
		if len(lines) > 0 {
			for _, v := range lines {
				if len(v) > 0 {
					err = s.SaveDistrictWeatherDataWithImdWithState(ctx, v, state.Hex())
					if err != nil {
						log.Println(err)
					}

				}
			}
		}

	}
}

// func (s *Service) LoadIMDBlockWeatherV2() {
// 	c := context.TODO()
// 	ctx := app.GetApp(c, s.Daos)
// 	defer ctx.Client.Disconnect(c)
// 	// config := config.NewConfig("districtimd", "config")
// 	// maxLength := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxlength)
// 	lines, err := s.GetWeatherDataWithImd(ctx)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	err = s.SaveBlockWeatherDataWithImd(ctx, lines)
// 	if err != nil {
// 		log.Println(err)
// 	}

// }
func (s *Service) LoadIMDBlockWeatherWithState() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	State, err := s.Daos.GetActiveState(ctx)
	if err != nil {
		log.Println("State Not Found")
	}
	fmt.Println("No.Of.State===>", len(State))
	for _, v := range State {
		config := config.NewConfig("blockimd", "config")
		fmt.Println(config)
		lines, err := s.GetBlockWeatherDataWithImdWithState(ctx, v.ImdStateName, v.ImdFileName)
		if err != nil {
			log.Println(err)
		}

		state := v.ID
		if len(lines) > 0 {
			for _, v := range lines {
				if len(v) > 0 {
					err = s.SaveBlockWeatherDataWithImdWithState(ctx, v, state.Hex())
					if err != nil {
						log.Println(err)
					}

				}
			}
		}

	}
}

// func (s *Service) LoadIMDBlockWeatherWithDistric() {
// 	c := context.TODO()
// 	ctx := app.GetApp(c, s.Daos)
// 	defer ctx.Client.Disconnect(c)
// 	state, err := s.Daos.GetActiveState(ctx)
// 	if err != nil {
// 		log.Println("State Not Found")
// 	}
// 	fmt.Println("No.Of.State===>", len(state))
// 	for _, v := range state {
// 		fmt.Println("StateName===>", v.ImdStateName)
// 		fmt.Println("StateFile===>", v.ImdBlockFileName)
// 		fmt.Println("StateId===>", v.ID)
// 		lines, err := s.GetBlockWeatherDataWithImdWithState(ctx, v.ImdStateName, v.ImdBlockFileName)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		state := v.ID.Hex()
// 		if len(lines) > 0 {
// 			for _, v := range lines {
// 				words := strings.Fields(v)
// 				fmt.Println("words===>", len(words))
// 				if len(words) > 0 {
// 					err = s.SaveBlockWeatherDataWithImdWithState(ctx, words, state)
// 					if err != nil {
// 						log.Println(err)
// 					}
// 				}
// 			}

// 		}

// 	}
// }
