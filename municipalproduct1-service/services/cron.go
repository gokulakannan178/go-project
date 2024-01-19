package services

import (
	"context"
	"fmt"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"sync"
	"time"
)

//PropertyCalcCron :
func (s *Service) PropertyCalcCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	// GetAllPropertyIds
	IDs, err := s.Daos.GetAllPropertyIds(ctx)

	if err != nil {
		fmt.Println(err)
	}
	IdsArray := s.Shared.SplitPropertyIds(IDs, 100)

	for _, v := range IdsArray {
		err = s.SavePropertyDemandForAll(v)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(5000)

	}
	fmt.Println("-------------script completed------------------")
	// err = s.CalcualteTotalCollectionForAllPayments()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, v := range IdsArray {
	// 	err = s.SavePropertyCollectionForAll(v)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

}

//ShopRentCalcCron :
func (s *Service) ShopRentCalcCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	// GetAllPropertyIds
	IDs, err := s.Daos.GetAllShopRentIds(ctx)

	if err != nil {
		fmt.Println(err)
	}
	IdsArray := s.Shared.SplitPropertyIds(IDs, 100)

	for _, v := range IdsArray {
		err = s.SaveShopRentDemandForAll(v)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(5000)

	}
	fmt.Println("-------------script completed------------------")
}

//SampleCron :
func (s *Service) SampleCron() {
	t := time.Now()
	fmt.Println("Every Five ", time.Since(t))
}

func (s *Service) PropertyCalc() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)

	IDs, err := s.Daos.GetAllPropertyIds(ctx)

	if err != nil {
		fmt.Println(err)
	}
	IdsArray := s.Shared.SplitPropertyIds(IDs, 500)
	for _, v := range IdsArray {
		err = s.SaveOverallPropertyDemandForAllV2(v)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//PropertyDemandSummaryCalc
func (s *Service) PropertyDemandSummaryCalc() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)

	IDs, err := s.Daos.GetAllPropertyIds(ctx)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Total Properties = ", len(IDs))
	IdsArray := s.Shared.SplitPropertyIds(IDs, 100)
	for k, v := range IdsArray {
		t := time.Now()
		var wg sync.WaitGroup
		c2 := context.TODO()
		ctx2 := app.GetApp(c2, s.Daos)
		for _, v2 := range v {

			filter := new(models.PropertyDemandFilter)
			filter.PropertyID = v2
			wg.Add(1)
			go s.PropertyDemandSummaryCalcForWG(ctx2, &wg, filter, "")
		}
		wg.Wait()
		log.Println(k, "Completed waiting")

		ctx2.Client.Disconnect(c)

		log.Println(k, "Disconnected out of", len(IdsArray))
		duration := time.Since(t)
		log.Println(" Time taken ===> ", duration.Minutes(), "m")
		time.Sleep(3 * time.Second)

	}

	fmt.Println("ALL PROPERTIES COMPLETED")
}

func (s *Service) PropertyDemandSummaryCalcForWG(ctx *models.Context, wg *sync.WaitGroup, filter *models.PropertyDemandFilter, collectionName string) {

	_, err := s.GetPropertyDemandCalc(ctx, filter, "")
	// err = s.SaveOverallPropertyDemandForAllV2(v)
	wg.Done()
	if err != nil {
		fmt.Println(err)
	}
}
