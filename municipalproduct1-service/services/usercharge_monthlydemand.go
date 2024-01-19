package services

import (
	"errors"
	"fmt"
	"municipalproduct1-service/models"
)

func (s *Service) GetUserChargeDemand(ctx *models.Context, id string) (*models.UserChargeDemand, error) {
	ucmcf := new(models.UserChargeMonthlyCalcQueryFilter)
	ucmcf.UserChargeID = id
	return s.GetUserChargeDemandWithFilter(ctx, ucmcf)

}

func (s *Service) GetUserChargeDemandWithFilter(ctx *models.Context, ucmcf *models.UserChargeMonthlyCalcQueryFilter) (*models.UserChargeDemand, error) {
	demand, err := s.Daos.GetUserChargeDemand(ctx, ucmcf)
	if err != nil {
		return nil, err
	}
	if demand == nil {
		return nil, errors.New("Demand is nil")
	}
	fmt.Println("First demand.Demand.TotalTax ", demand.UCDemand.TotalTax)
	for k, v := range demand.Ref.Fy {
		for k2, v2 := range v.FyMonth {
			demand.Ref.Fy[k].FyMonth[k2].Tax = v2.Rate.Rate
			demand.Ref.Fy[k].FyMonth[k2].Total = demand.Ref.Fy[k].FyMonth[k2].Tax

			demand.Ref.Fy[k].FyMonth[k2].ToBePaid = demand.Ref.Fy[k].FyMonth[k2].Tax - demand.Ref.Fy[k].FyMonth[k2].Alreadypaid
			demand.Ref.Fy[k].FyMonth[k2].Penalty = (demand.Ref.Fy[k].FyMonth[k2].ToBePaid / 100) * 0
			demand.Ref.Fy[k].FyMonth[k2].TotalTaxToBePaid = demand.Ref.Fy[k].FyMonth[k2].ToBePaid + demand.Ref.Fy[k].FyMonth[k2].Penalty

			fmt.Println("demand.Ref.Fy[k].Penalty ", demand.Ref.Fy[k].CalculatedPenalty)
			fmt.Println("demand.Ref.Fy[k].Penalty", demand.Ref.Fy[k].FyMonth[k2].Penalty)
			fmt.Println("demand.Ref.Fy[k].Total", demand.Ref.Fy[k].Total)
			fmt.Println("demand.Ref.Fy[k].FyMonth[k2].Total", demand.Ref.Fy[k].FyMonth[k2].Total)
			demand.Ref.Fy[k].Tax = demand.Ref.Fy[k].Tax + demand.Ref.Fy[k].FyMonth[k2].Tax
			demand.Ref.Fy[k].Total = demand.Ref.Fy[k].Total + demand.Ref.Fy[k].FyMonth[k2].Total

			demand.Ref.Fy[k].CalculatedPenalty = demand.Ref.Fy[k].CalculatedPenalty + demand.Ref.Fy[k].FyMonth[k2].Penalty
			demand.Ref.Fy[k].Alreadypaid = demand.Ref.Fy[k].Alreadypaid + demand.Ref.Fy[k].FyMonth[k2].Alreadypaid
			demand.Ref.Fy[k].ToBePaid = demand.Ref.Fy[k].ToBePaid + demand.Ref.Fy[k].FyMonth[k2].ToBePaid
			demand.Ref.Fy[k].TotalTaxToBePaid = demand.Ref.Fy[k].TotalTaxToBePaid + demand.Ref.Fy[k].FyMonth[k2].TotalTaxToBePaid

		}

		if demand.Ref.Fy[k].IsCurrent {
			demand.UCDemand.CurrentTax = demand.Ref.Fy[k].ToBePaid
			demand.UCDemand.CurrentPenalty = demand.Ref.Fy[k].CalculatedPenalty
			demand.UCDemand.CurrentTotal = demand.Ref.Fy[k].TotalTaxToBePaid

			demand.UCDemand.Actual.CurrentTax = demand.Ref.Fy[k].Tax

		} else {
			fmt.Println("demand.Demand.ArrearPenalty", demand.UCDemand.ArrearPenalty)
			fmt.Println("demand.Ref.Fy[k].Penalty", demand.Ref.Fy[k].CalculatedPenalty)
			demand.UCDemand.ArrearTax = demand.UCDemand.ArrearTax + demand.Ref.Fy[k].ToBePaid
			demand.UCDemand.ArrearPenalty = demand.UCDemand.ArrearPenalty + demand.Ref.Fy[k].CalculatedPenalty
			demand.UCDemand.ArrearTotal = demand.UCDemand.ArrearTotal + demand.Ref.Fy[k].TotalTaxToBePaid

			demand.UCDemand.Actual.ArrearTax = demand.Ref.Fy[k].Tax
		}
		fmt.Println("demand.Demand.TotalTax ", demand.UCDemand.TotalTax)
		fmt.Println("demand.Ref.Fy[k].Total", demand.Ref.Fy[k].Total)
		demand.UCDemand.TotalTax = demand.UCDemand.TotalTax + demand.Ref.Fy[k].Total
		demand.UCDemand.Actual.TotalTax = demand.UCDemand.Actual.TotalTax + demand.Ref.Fy[k].Tax

	}

	var fy []models.UserChargeDemandFY
	for _, v := range demand.Ref.Fy {
		var fyMonth []models.UserChargeDemandFyMonth
		if v.TotalTaxToBePaid == 0 {
			continue
		}
		for _, v2 := range v.FyMonth {
			if v2.TotalTaxToBePaid > 0 {
				fyMonth = append(fyMonth, v2)
			}
		}

		v.FyMonth = fyMonth
		fy = append(fy, v)
	}

	demand.Ref.Fy = fy
	return demand, err

}
