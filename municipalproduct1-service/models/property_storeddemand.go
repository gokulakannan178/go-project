package models

import (
	"fmt"
	"math"
	"municipalproduct1-service/constants"
	"time"
)

//DemandCalculation : ""
func (pd *PropertyDemand) StoredPropertyDemandCalculation() PropertyDemand {
	var legacyAmount float64
	pd.FYTax = 0

	for k, v := range pd.FYs {
		fmt.Printf("####################### Start Of %v #######################", v.Name)
		// this section is updated for legacy payment

		//Adjusting Legacy Amount
		if pd.ProductConfiguration != nil {
			if pd.ProductConfiguration.ProductConfiguration.ISLegacy == "Yes" {
				fmt.Println(pd.FYs[k].Name, " before pd.FYs[k].OverallTax.Total =====>", pd.FYs[k].OverallTax.Total)
				// pd.FYs[k].Tax = (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax + pd.FYs[k].CompositeTax + pd.FYs[k].Ecess) - pd.FYs[k].Legacy.TaxAmount
				pd.FYs[k].OverallTax.Total = pd.FYs[k].OverallTax.Total - pd.FYs[k].Legacy.TaxAmount
				legacyAmount = legacyAmount + pd.FYs[k].Legacy.TaxAmount
				fmt.Println(pd.FYs[k].Name, " after pd.FYs[k].OverallTax.Total =====>", pd.FYs[k].OverallTax.Total)
			}
		}

		// this section is updated for legacy payment
		var sumARV, sumAPTR float64
		var maxConstructedArea float64
		maxConstructedAreaMap := make(map[string]float64)
		floorARV := make(map[string]float64)
		// floorARVForEcess := make(map[string]float64)
		for k1, v1 := range v.Floors {
			// if maxConstructedArea < v1.BuildUpArea {
			// 	maxConstructedArea = v1.BuildUpArea
			// }
			fmt.Println("Statt of FLoor - ", v1.No)
			maxConstructedAreaMap[v1.No] = maxConstructedAreaMap[v1.No] + v1.BuildUpArea

			if v1.Ref.NonResUsageType != nil {
				if v1.Ref.NonResUsageType.IsServiceCharge == "Yes" {
					pd.IsServiceChargeApplicable = true
				}
			}
			if v1.Ref.FloorRatableArea == nil {
				v1.Ref.FloorRatableArea = new(FloorRatableArea)
			}
			if v1.Ref.AVR == nil {
				v1.Ref.AVR = new(AVR)
			}
			if v1.Ref.OccupancyType == nil {
				v1.Ref.OccupancyType = new(OccupancyType)
			}
			if v1.Ref.PropertyTax == nil {
				v1.Ref.PropertyTax = new(PropertyTax)
			}
			buildupArea := (v1.BuildUpArea * v1.Ref.FloorRatableArea.Rate) / 100
			if v1.ConstructionType == "OTHERS" && buildupArea < 500 {
				v1.Ref.AVR = new(AVR)
			}

			floorarv := buildupArea * v1.Ref.AVR.Rate
			pd.FYs[k].Floors[k1].CG.Rate = v1.Ref.AVR.Rate
			// if v1.No == "VACANT" {
			// 	floorarv = buildupArea * v1.Ref.VLR.Rate
			// 	v1.Ref.AVR = new(AVR)
			// 	// pd.FYs[k].Floors[k1].ARV = 0
			// }
			v1.ARV = floorarv * v1.Ref.OccupancyType.Factor
			if pd.ProductConfiguration.LocationID == "CGBhilai" {
				fmt.Println("floor discount", pd.FYs[k].Floors[k1].Ref.FloorNo.Discount)
				fmt.Println("rate = ", pd.FYs[k].Floors[k1].Ref.AVR.Rate)
				pd.FYs[k].Floors[k1].CG.Rate = pd.FYs[k].Floors[k1].Ref.AVR.Rate
				pd.FYs[k].Floors[k1].CG.FLoorDiscountPercentage = pd.FYs[k].Floors[k1].Ref.FloorNo.Discount

				arvDiscount := (pd.FYs[k].Floors[k1].Ref.AVR.Rate * pd.FYs[k].Floors[k1].Ref.FloorNo.Discount) / 100
				pd.FYs[k].Floors[k1].CG.RateDiscount = arvDiscount

				pd.FYs[k].Floors[k1].Ref.AVR.Rate = pd.FYs[k].Floors[k1].Ref.AVR.Rate - arvDiscount
				pd.FYs[k].Floors[k1].CG.DiscountedRate = pd.FYs[k].Floors[k1].Ref.AVR.Rate

				// v1.ARV = v1.ARV - arvDiscount
				floorarv = buildupArea * pd.FYs[k].Floors[k1].Ref.AVR.Rate
				v1.ARV = floorarv * v1.Ref.OccupancyType.Factor
				pd.FYs[k].Floors[k1].CG.ARV = v1.ARV

				// pd.FYs[k].Floors[k1].ARV = v1.ARV
				fmt.Println("discounted rate = ", pd.FYs[k].Floors[k1].Ref.AVR.Rate, "discounted arv = ", v1.ARV)

				pd.FYs[k].Floors[k1].ARV = v1.ARV
				floorARV[pd.FYs[k].Floors[k1].ConstructionType] = floorARV[pd.FYs[k].Floors[k1].No] + v1.ARV
			}

			v1.APTR = (v1.ARV * v.Ref.PropertyTax.Rate) / 100
			// v1.CompostiteTax = v1.Ref.CompositeTax.Rate
			pd.FYs[k].Floors[k1].ARV = v1.ARV
			pd.FYs[k].Floors[k1].APTR = v1.APTR

			fmt.Println("FLoor Calc = **************")
			fmt.Println("Floor no = ", v1.Ref.FloorNo.Name)
			fmt.Println("buildupArea = ", v1.BuildUpArea)
			fmt.Println("ratable buildupArea = ", buildupArea)
			fmt.Println("arv rate = ", v1.Ref.AVR.Rate)
			fmt.Println("OccupancyType.Factor rate = ", v1.Ref.OccupancyType.Factor)
			fmt.Println("floorarv = ", floorarv)
			fmt.Println("ARV = ", pd.FYs[k].Floors[k1].ARV)
			fmt.Println("Floor Tax = ", pd.FYs[k].Floors[k1].APTR)
			pd.FYs[k].Floors[k1].Ref.RatableArea = buildupArea

			sumARV = sumARV + v1.ARV
			sumAPTR = sumAPTR + v1.APTR
			fmt.Println("sumARV = ", sumARV)
			fmt.Println("sumAPTR = ", sumAPTR)

			// sumCompositeTax = sumCompositeTax + v1.CompostiteTax
			fmt.Println("-- -- -- -- -- End of FLoor - ", v1.No, " -- -- -- -- -- -- ")
		}

		fmt.Println("After Floor Calculation")
		fmt.Println("GOT SUMMED ARV = ", sumARV)
		if pd.ProductConfiguration.LocationID == "CGBhilai" {
			fmt.Println("UPDATING ARV RANGE FOR CG ")
			sumAPTR = 0
			// var floorAPTR, vlAPTR float64
			// // discountablesumAPTR := 0
			// for frvk, frvv := range floorARV {

			// 	pd.FYs[k].Ecess = pd.FYs[k].Ecess + ((frvv * 2) / 100)
			// 	arvRange, err := pd.GetCGPropertyTaxRate(pd.CTX, frvv, v.FinancialYear.To)
			// 	if err != nil {
			// 		fmt.Println("err in geting ar range ", err)
			// 		arvRange = new(RefAVRRange)
			// 	}
			// 	fmt.Println("got rate = ", arvRange.AVRRange.Rate)
			// 	tempAPTR := (frvv * arvRange.AVRRange.Rate) / 100
			// 	if frvk == "VACANT" {
			// 		vlAPTR = vlAPTR + tempAPTR
			// 	} else {
			// 		floorAPTR = floorAPTR + tempAPTR
			// 	}
			// 	// sumAPTR = sumAPTR + tempAPTR
			// 	fmt.Println("floor CT ", frvk, "floor sum arv", frvv, "sumAPTR", sumAPTR)
			// }
			// fmt.Println("FLoor APTR = ", floorAPTR)
			// fmt.Println("VL APTR = ", vlAPTR)

			//10 % of maintenance - floor type / construction type != VACANT
			isVacant := false
			for _, isVacantFloors := range pd.FYs[k].Floors {
				if isVacantFloors.No == "VACANT" {
					isVacant = true
					break
				}
			}
			if !isVacant {
				pd.FYs[k].CG.SumARV = sumARV
				pd.FYs[k].CG.MaintenanceDiscountPercentage = 10

				fmt.Println("Since no vacant land giving flat 10 % discount")
				arvDisc := (sumARV * pd.FYs[k].CG.MaintenanceDiscountPercentage / 100)
				pd.FYs[k].CG.MaintenanceDiscount = arvDisc

				sumARV = sumARV - arvDisc
				pd.FYs[k].CG.MaintenanceDiscountedARV = sumARV
				fmt.Println("again 10 % =", arvDisc, "discounted ARV = ", sumARV)
			} else {
				fmt.Println("Since der is a  vacant land NO 10 % discount")
			}
			finalARV := sumARV

			arvRange, err := pd.GetCGPropertyTaxRate(pd.CTX, sumARV, v.FinancialYear.To)
			if err != nil {
				fmt.Println("err in geting ar range ", err)
				arvRange = new(RefAVRRange)
			}

			pd.FYs[k].CG.TaxRate = arvRange.AVRRange.Rate
			sumAPTR = (sumARV * arvRange.AVRRange.Rate) / 100
			fmt.Println("Property Tax Slab - from -> ", arvRange.From, " to -> ", arvRange.To, " Percentage =>", arvRange.Rate, "%")
			fmt.Println("Property Tax =", sumAPTR)
			//50% Discount on Residential Building (Only for Self Accommodation)
			//residential id  = 1  and self id = 1
			isResidential := true
			isSelf := true

			for _, disFLoorv := range pd.FYs[k].Floors {
				if disFLoorv.UsageType != "1" {
					isResidential = false
					break
				}
				if disFLoorv.OccupancyType != "1" {
					isSelf = false
					break
				}
			}
			if isResidential && isSelf && !isVacant {
				pd.FYs[k].CG.ResDiscountPercentage = 50

				//make discount
				sumAPTR = sumAPTR / 2
				pd.FYs[k].CG.ResDiscount = sumAPTR
				pd.FYs[k].CG.ResDiscountedARV = sumAPTR
				fmt.Println("Since Residential and Self Building Discounting Property Tax - 50 %")
				fmt.Println("After 50% discounting Property Tax = ", sumAPTR)
			} else {
				fmt.Println("NO 50 % discount  isResidential", isResidential, " isSelf", isSelf, " isVacant", isVacant)
			}

			//Ecess = 2%
			fmt.Println("Ecess master ", pd.FYs[k].EcessRateMaster.ON)
			if pd.FYs[k].EcessRateMaster.ON == "ARV" {
				pd.FYs[k].Ecess = pd.FYs[k].Ecess + ((finalARV * 2) / 100)
				fmt.Println("Ecess on SARV = ", pd.FYs[k].Ecess)
			}
			if pd.FYs[k].EcessRateMaster.ON == "PT" {
				pd.FYs[k].Ecess = pd.FYs[k].Ecess + ((sumAPTR * 2) / 100)
				fmt.Println("Ecess on PT = ", pd.FYs[k].Ecess)
			}
			// sumAPTR = floorAPTR + vlAPTR
			// fmt.Println("DIscounted FLoor APTR = ", floorAPTR)
			// fmt.Println("VL APTR = ", vlAPTR)
			// fmt.Println("FInal SMPTR = ", sumAPTR, "ecess", pd.FYs[k].Ecess)
			/*
				//get ARV Range
				arvRange, err := pd.GetCGPropertyTaxRate(pd.CTX, sumARV, v.FinancialYear.To)
				if err != nil {
					fmt.Println("err in geting ar range ", err)
					arvRange = new(RefAVRRange)
				}
				pd.FYs[k].AVRRange = arvRange.AVRRange

				// v.Ref.PropertyTax.Rate = arvRange.AVRRange.Rate
				sumAPTR = (sumARV * arvRange.AVRRange.Rate) / 100
			*/
		}
		// fmt.Println("FINAL AVRRANGE = ", pd.FYs[k].AVRRange)
		// fmt.Println("FINAL sumAPTR = ", sumAPTR)

		for k1, v1 := range maxConstructedAreaMap {
			fmt.Println("fy", v.FinancialYear.Name,
				"floor id =", k1,
				"constructed area =", v1)
			if maxConstructedArea < v1 {
				maxConstructedArea = v1
			}
			fmt.Println("==============")

		}

		// Calculating vacant land
		pd.FYs[k].ConstructedArea = maxConstructedArea
		pd.PercentAreaBuildup = (maxConstructedArea / pd.AreaOfPlot) * 100
		pd.PercentAreaBuildup = math.Ceil(pd.PercentAreaBuildup)
		fmt.Println("Area of Plot ====================>", pd.AreaOfPlot)
		fmt.Println("Max Constructed Area ===================>", maxConstructedArea)
		fmt.Println("Taxable Vacant Land ==============> ", pd.PropertyConfig.TaxableVacantLandConfig)
		pd.TaxableVacantLand = pd.AreaOfPlot - (maxConstructedArea * pd.PropertyConfig.TaxableVacantLandConfig)
		if pd.PercentAreaBuildup < pd.PropertyConfig.VacantLandRatePercentage {
			pd.FYs[k].VacantLandTax = pd.TaxableVacantLand * v.VLR.Rate
			if pd.FYs[k].CommonVLR == "Yes" {
				var multiplyingFactor float64
				multiplyingFactor = math.Ceil(pd.AreaOfPlot / 720)
				fmt.Println("multiplyingFactor =============> ", multiplyingFactor)
				pd.FYs[k].VacantLandTax = v.VLR.Rate * multiplyingFactor
			}
		}
		if pd.ProductConfiguration.LocationID == "CGBhilai" {
			pd.FYs[k].VacantLandTax = 0
		}
		pd.FYs[k].SumARV = sumARV
		pd.FYs[k].SumFloorTax = sumAPTR
		pd.FYs[k].Tax = pd.FYs[k].Tax + sumAPTR
		pd.FYs[k].CompositeTax = pd.FYs[k].CompositeTaxRate.Rate

		calculatePenalty := true

		// if pd.OtherCharges != nil {
		// 	//Checking Boring charge parking
		// 	if pd.OtherCharges.BoringChargeParking == constants.PARKBORINGCHARGENO {
		// 		fmt.Println("Calculating boring charge")
		// 		if !pd.IsBoringChargePayed {
		// 			if pd.Property.MunicipalityWaterConnection == constants.MUNICIPALITYWATERCONNECTIONYES {
		// 				pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeWithWaterConnection
		// 			} else {
		// 				pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeWithoutWaterConnection
		// 			}
		// 		}
		// 	}
		// 	//Checking penalty parking
		// 	if pd.OtherCharges.PenaltyParking == constants.PARKPENALTYYES || pd.ParkPenalty {
		// 		calculatePenalty = false
		// 	}
		// }
		// for munger
		if pd.OtherCharges != nil {
			//Checking Boring charge parking
			if pd.OtherCharges.BoringChargeParking == constants.PARKBORINGCHARGENO {
				fmt.Println("Calculating boring charge")
				if !pd.IsBoringChargePayed {
					if pd.Property.MunicipalityWaterConnection == constants.MUNICIPALITYWATERCONNECTIONYES {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeWithWaterConnection
					} else if pd.Property.MunicipalityWaterConnection == constants.MUNICIPALITYWATERCONNECTIONSUPPLYOWN {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeWithWaterConnectionSupplyAndOwn
					} else if pd.Property.MunicipalityWaterConnection == constants.MUNICIPALITYWATERCONNECTIONNOTAPPLICABLE {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeNotApplicable
					} else if pd.Property.MunicipalityWaterConnection == constants.MUNICIPALITYWATERCONNECTIONALREADYPAID {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeAlreadyPaied
					} else {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeWithoutWaterConnection
					}
				}
			}
			//Checking penalty parking
			if pd.OtherCharges.PenaltyParking == constants.PARKPENALTYYES || pd.ParkPenalty {
				calculatePenalty = false
			}
		}
		//
		fmt.Println("is current", pd.FYs[k].IsCurrent)
		if pd.FYs[k].FixedArv != nil {
			pd.FYs[k].VacantLandTax = 0
			if pd.ProductConfiguration.FixedARVTaxCalc == "PropertyTaxPercentage" {
				pd.FYs[k].Tax = (pd.FYs[k].FixedArv.ARV / 100) * pd.FYs[k].Ref.PropertyTax.Rate
			}
			if pd.ProductConfiguration.FixedARVTaxCalc == "IndividualTaxPercentage" {
				pd.FYs[k].Tax = pd.FYs[k].FixedArv.Total
			}
		}
		pd.FYs[k].OverallTax.FYTax = pd.FYs[k].Tax
		pd.FYs[k].OverallTax.VLTax = pd.FYs[k].VacantLandTax
		pd.FYs[k].OverallTax.CompositeTax = pd.FYs[k].CompositeTax
		pd.FYs[k].OverallTax.Ecess = pd.FYs[k].Ecess
		pd.FYs[k].OverallTax.ToBePaid = pd.FYs[k].Tax + pd.FYs[k].VacantLandTax + pd.FYs[k].CompositeTax + pd.FYs[k].Ecess
		pd.FYs[k].OverallTax.PanelCharge = pd.FYs[k].PanelCharge
		pd.FYs[k].OverallTax.Penalty = pd.FYs[k].Penalty
		pd.FYs[k].OverallTax.Total = pd.FYs[k].Tax + pd.FYs[k].VacantLandTax + pd.FYs[k].CompositeTax + pd.FYs[k].Ecess
		pd.FYs[k].OverallTax.PayableTax = pd.FYs[k].OverallTax.ToBePaid - (pd.FYs[k].AlreadyPayed.PaidTax + pd.FYs[k].Legacy.TaxAmount)
		// if pd.ProductConfiguration.LocationID == "CGBhilai" {
		// 	pd.FYs[k].Ecess = (pd.FYs[k].SumARV / 100) * 2
		// }

		if pd.FYs[k].IsCurrent {

			ct := time.Now()
			d := *pd.FYs[k].LastDate
			fmt.Println("CURR FY YEAR PENALTY CALC", d)
			months := monthsCountSince(d)
			fmt.Println("CURR FY YEAR PENALTY CALC Months", months, pd.FYs[k].AlreadyPayed.Amount)

			pd.FYs[k].PenaltyMonths = float64(months + 1)
			//COmmenting Rebate
			// pd.FYs[k].Rebate = pd.FYs[k].Rebate + (((pd.FYs[k].Tax + pd.FYs[k].VacantLandTax) * 5) / 100)
			//if math.Ceil(pd.FYs[k].AlreadyPayed.Amount) < (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax) {
			pd.FYs[k].Penalty = (((pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax)) * pd.FYs[k].PenaltyRate) / 100) * float64(months+1)

			fmt.Println("calculated penalty", (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - pd.FYs[k].AlreadyPayed.Amount), "*", pd.FYs[k].PenaltyRate, "/100 *", float64(months+1), "=", pd.FYs[k].Penalty)
			// 	if calculatePenalty {

			// 	pd.FYs[k].Penalty = (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax + pd.FYs[k].Rebate)) * ((pd.FYs[k].PenaltyRate * float64(months+1)) / 100)
			// 	fmt.Println("calculated penalty", (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - pd.FYs[k].AlreadyPayed.Amount), "*", pd.FYs[k].PenaltyRate, "/100 *", float64(months+1), "=", pd.FYs[k].Penalty)
			// }
			//}
			// pd.FYs[k].Penalty = (((pd.FYs[k].Tax + pd.FYs[k].VacantLandTax) * pd.FYs[k].PenaltyRate) / 100)
			if ct.After(d) {
				// months := monthsCountSince(d)
				// pd.FYs[k].Penalty = ((pd.FYs[k].Tax * pd.FYs[k].Ref.Penalty.Rate) / 100) * float64(months)
				//Penalty cacculation for current year
				// if calculatePenalty {
				// 	pd.FYs[k].Penalty = (((pd.FYs[k].Tax + pd.FYs[k].VacantLandTax) * pd.FYs[k].PenaltyRate) / 100)

				// }
			} else {
				// fmt.Println("calculating rebate for early payment")
				// pd.FYs[k].Rebate = pd.FYs[k].Rebate + ((pd.FYs[k].Tax * 5) / 100)
			}
			if pd.ProductConfiguration.LocationID == "CGBhilai" {
				pd.FYs[k].Penalty = 0
				t := time.Now()
				m := int(t.Month())
				if m == 12 || m == 1 || m == 2 || m == 3 {
					pd.FYs[k].Rebate = ((pd.FYs[k].Tax * 0) / 100)
				}
				if m == 4 || m == 5 {
					pd.FYs[k].Rebate = ((pd.FYs[k].Tax * 6.25) / 100)

				}
				if m == 6 || m == 7 {
					pd.FYs[k].Rebate = ((pd.FYs[k].Tax * 5) / 100)

				}
				if m == 8 || m == 9 {
					pd.FYs[k].Rebate = ((pd.FYs[k].Tax * 4) / 100)

				}
				if m == 10 || m == 11 {
					pd.FYs[k].Rebate = ((pd.FYs[k].Tax * 2) / 100)

				}
				//pd.FYs[k].Rebate
			}

		} else {

			// months := 12
			// pd.FYs[k].Penalty = ((pd.FYs[k].Tax * pd.FYs[k].Ref.Penalty.Rate) / 100) * float64(months)
			//Penalty cacculation for previous years
			if calculatePenalty {
				// if math.Ceil(pd.FYs[k].AlreadyPayed.Amount) < math.Ceil(pd.FYs[k].Tax+pd.FYs[k].VacantLandTax) {
				// pd.FYs[k].Penalty = (((pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - pd.FYs[k].AlreadyPayed.Amount) * pd.FYs[k].PenaltyRate) / 100)
				fmt.Printf("calc penalty (%v + %v - (%v + %v))",
					pd.FYs[k].Tax, pd.FYs[k].VacantLandTax, pd.FYs[k].CompositeTax, pd.FYs[k].Ecess, pd.FYs[k].AlreadyPayed.FYTax, pd.FYs[k].AlreadyPayed.VLTax,
				)
				fmt.Println()
				pd.FYs[k].Penalty = (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax)) * (pd.FYs[k].PenaltyRate / 100)
				if pd.ProductConfiguration.LocationID == "CGBhilai" {
					// pd.FYs[k].Penalty = (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax + pd.FYs[k].CompositeTax + pd.FYs[k].Ecess - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax + pd.FYs[k].AlreadyPayed.CompositeTax + pd.FYs[k].AlreadyPayed.Ecess)) * (pd.FYs[k].PenaltyRate / 100)
					pd.FYs[k].Penalty = (pd.FYs[k].OverallTax.PayableTax) * (pd.FYs[k].PenaltyRate / 100)

					if pd.FYs[k].FloorBuildupArea.Area <= 500 {
						// pd.FYs[k].Penalty = (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax + pd.FYs[k].CompositeTax + pd.FYs[k].Ecess - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax + pd.FYs[k].AlreadyPayed.CompositeTax + pd.FYs[k].AlreadyPayed.Ecess)) * (pd.FYs[k].PenaltyRate / 100)
						pd.FYs[k].Penalty = (pd.FYs[k].OverallTax.PayableTax) * (pd.FYs[k].PenaltyRate / 100)
					}
				}
				// }
			}
		}
		if pd.IsRainWaterHarvesting == "Yes" {
			pd.FYs[k].Rebate = pd.FYs[k].Rebate + ((pd.FYs[k].Tax * 5) / 100)
		}
		fmt.Println("calculating fy tax ", pd.FYs[k].Tax, "- ", pd.FYs[k].Rebate, " +", pd.FYs[k].Penalty)

		pd.FYs[k].Tax = math.Ceil(pd.FYs[k].Tax)
		pd.FYs[k].Rebate = math.Ceil(pd.FYs[k].Rebate)
		pd.FYs[k].Penalty = math.Ceil(pd.FYs[k].Penalty)
		pd.FYs[k].VacantLandTax = math.Ceil(pd.FYs[k].VacantLandTax)
		pd.FYs[k].CompositeTax = math.Ceil(pd.FYs[k].CompositeTax)
		pd.FYs[k].OtherDemand = math.Ceil(pd.FYs[k].OtherDemand)

		fytax := pd.FYs[k].Tax - pd.FYs[k].Rebate + pd.FYs[k].Penalty

		//Adjusting Legacy Amount
		// if pd.ProductConfiguration != nil {
		// 	if pd.ProductConfiguration.ProductConfiguration.ISLegacy == "Yes" {
		// 		fytax = fytax - pd.FYs[k].Legacy.TaxAmount
		// 		legacyAmount = legacyAmount + pd.FYs[k].Legacy.TaxAmount
		// 	}
		// } this is commented for legacy payment and overwritten in

		fmt.Println("calculating fy tax and vacant lan tax and comp tax ", fytax, " +", pd.FYs[k].VacantLandTax, " + ", pd.FYs[k].CompositeTax)

		fytax = fytax + pd.FYs[k].VacantLandTax + pd.FYs[k].CompositeTax + pd.FYs[k].OtherDemand

		if pd.IsServiceChargeApplicable {
			fmt.Println("service charge calculated")
			pd.FYs[k].ServiceCharge = fytax
			pd.ServiceCharge = pd.ServiceCharge + pd.FYs[k].ServiceCharge*(pd.PropertyConfig.ServiceCharge/100)
		} else {
			fmt.Println("Before adding- ")
			fmt.Println("pd.FYs[k].TotalTax- ", pd.FYs[k].TotalTax)
			fmt.Println("fytax= ", fytax)
			fmt.Println("pd.FYs[k].AlreadyPayed.Amount- ", pd.FYs[k].AlreadyPayed.Amount)
			fmt.Println("pd.FYs[k].Rebate- ", pd.FYs[k].Rebate)
			fmt.Println("pd.FYs[k].AlreadyPayed.FYTax- ", pd.FYs[k].AlreadyPayed.FYTax)
			fmt.Println("pd.FYs[k].AlreadyPayed.VLTax- ", pd.FYs[k].AlreadyPayed.VLTax)
			fmt.Println("pd.FYTax= ", pd.FYTax)

			// pd.FYs[k].TotalTax = fytax
			pd.FYs[k].TotalTax = fytax - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax + pd.FYs[k].AlreadyPayed.CompositeTax + pd.FYs[k].AlreadyPayed.Ecess)
			if !pd.FYs[k].IsCurrent {
				if pd.FYs[k].TotalTax > 0 {
					pd.FYs[k].TotalTax = pd.FYs[k].TotalTax + pd.FYs[k].PanelCharge
				}
			}
			pd.FYTax = pd.FYTax + pd.FYs[k].TotalTax

			fmt.Println("After adding- ", pd.FYTax, pd.FYs[k].TotalTax)

		}
		fmt.Println("Current FY tax - ", pd.FYTax, pd.FYs[k].TotalTax)
		if pd.FYs[k].IsCurrent {
			pd.Current = pd.Current + pd.FYs[k].TotalTax
		} else {
			pd.Arrear = pd.Arrear + pd.FYs[k].TotalTax
		}

		// pd.FlTax = pd.FYs[k].Tax
		// pd.VlTax = pd.FYs[k].VacantLandTax
		// pd.Tax = pd.FYs[k].Tax + pd.FYs[k].VacantLandTax

		fmt.Printf("####################### End Of %v", v.Name, " #########################")

	}
	// if !pd.AllDemand {
	// 	tempFYdemand := make([]FinancialYearDemand, 0)
	// 	for _, v := range pd.FYs {
	// 		if v.TotalTax > 0 {
	// 			tempFYdemand = append(tempFYdemand, v)
	// 		}
	// 	}
	// 	pd.FYs = tempFYdemand
	// }

	// if len(pd.FYs) > 0 {
	// 	if pd.OtherCharges.FormFeeParking == constants.PARKFORMFEENO {
	// 		if pd.IsFormFeePayed {
	// 			pd.FormFee = 0

	// 		} else {
	// 			pd.FormFee = pd.OtherCharges.FormFeeCharges
	// 		}

	// 	}
	// } else {
	// 	pd.FormFee = 0
	// 	pd.BoreCharge = 0
	// }

	// if pd.ProductConfiguration.LocationID == "CGBhilai" {
	// 	pd.FormFee = 3
	// }
	// if pd.ProductConfiguration.LocationID == "Bhagalpur" {
	// 	pd.FormFee = 5
	// }

	pd.TotalTax = pd.FYTax + pd.ServiceCharge + pd.BoreCharge + pd.FormFee

	// if !pd.PreviousCollection.IsCalculated {
	// 	pd.TotalTax = pd.TotalTax - pd.PreviousCollection.Amount
	// }

	//Adding Advance
	// if pd.TotalTax < 0 {
	// 	pd.AdvanceReceived = pd.TotalTax
	// 	pd.TotalTax = 0
	// }

	return *pd
}

func (pd *PropertyDemand) StoredPropertyDemandCalculationV2() ([]StoredCalculationDemandfy, *StoredCalculationDemand) {
	var legacyAmount float64
	pd.FYTax = 0
	var StoredCalculationDemandfyArr []StoredCalculationDemandfy
	StoredCalculationDemand := new(StoredCalculationDemand)

	for k, v := range pd.FYs {
		fmt.Printf("####################### Start Of %v #######################", v.Name)
		if pd.ProductConfiguration != nil {
			if pd.ProductConfiguration.ProductConfiguration.ISLegacy == "Yes" {
				legacyAmount = legacyAmount + pd.FYs[k].Legacy.TaxAmount
			}
		}
		var sumARV, sumAPTR float64
		var maxConstructedArea float64
		maxConstructedAreaMap := make(map[string]float64)
		//Calculating floor tax
		for k1, v1 := range v.Floors {
			fmt.Println("Statt of FLoor - ", v1.No)
			maxConstructedAreaMap[v1.No] = maxConstructedAreaMap[v1.No] + v1.BuildUpArea

			// if v1.Ref.NonResUsageType != nil {
			// 	if v1.Ref.NonResUsageType.IsServiceCharge == "Yes" {
			// 		pd.IsServiceChargeApplicable = true
			// 	}
			// }
			if v1.Ref.FloorRatableArea == nil {
				v1.Ref.FloorRatableArea = new(FloorRatableArea)
			}
			if v1.Ref.AVR == nil {
				v1.Ref.AVR = new(AVR)
			}
			if v1.Ref.OccupancyType == nil {
				v1.Ref.OccupancyType = new(OccupancyType)
			}
			if v1.Ref.PropertyTax == nil {
				v1.Ref.PropertyTax = new(PropertyTax)
			}
			buildupArea := (v1.BuildUpArea * v1.Ref.FloorRatableArea.Rate) / 100
			if v1.ConstructionType == "OTHERS" && buildupArea < 500 {
				v1.Ref.AVR = new(AVR)
			}

			floorarv := buildupArea * v1.Ref.AVR.Rate
			v1.ARV = floorarv * v1.Ref.OccupancyType.Factor

			v1.APTR = (v1.ARV * v.Ref.PropertyTax.Rate) / 100
			// v1.CompostiteTax = v1.Ref.CompositeTax.Rate
			pd.FYs[k].Floors[k1].ARV = v1.ARV
			pd.FYs[k].Floors[k1].APTR = v1.APTR

			fmt.Println("FLoor Calc = **************")
			fmt.Println("Floor no = ", v1.Ref.FloorNo.Name)
			fmt.Println("buildupArea = ", v1.BuildUpArea)
			fmt.Println("ratable buildupArea = ", buildupArea)
			fmt.Println("arv rate = ", v1.Ref.AVR.Rate)
			fmt.Println("OccupancyType.Factor rate = ", v1.Ref.OccupancyType.Factor)
			fmt.Println("floorarv = ", floorarv)
			fmt.Println("ARV = ", pd.FYs[k].Floors[k1].ARV)
			fmt.Println("Floor Tax = ", pd.FYs[k].Floors[k1].APTR)
			pd.FYs[k].Floors[k1].Ref.RatableArea = buildupArea

			sumARV = sumARV + v1.ARV
			sumAPTR = sumAPTR + v1.APTR
			fmt.Println("sumARV = ", sumARV)
			fmt.Println("sumAPTR = ", sumAPTR)

			// sumCompositeTax = sumCompositeTax + v1.CompostiteTax
			fmt.Println("-- -- -- -- -- End of FLoor - ", v1.No, " -- -- -- -- -- -- ")
		}
		for k1, v1 := range maxConstructedAreaMap {
			fmt.Println("fy", v.FinancialYear.Name,
				"floor id =", k1,
				"constructed area =", v1)
			if maxConstructedArea < v1 {
				maxConstructedArea = v1
			}
			fmt.Println("==============")

		}
		// Calculating vacant land
		pd.FYs[k].ConstructedArea = maxConstructedArea
		pd.PercentAreaBuildup = (maxConstructedArea / pd.AreaOfPlot) * 100
		pd.PercentAreaBuildup = math.Ceil(pd.PercentAreaBuildup)
		fmt.Println("Area of Plot ====================>", pd.AreaOfPlot)
		fmt.Println("Max Constructed Area ===================>", maxConstructedArea)
		fmt.Println("Taxable Vacant Land ==============> ", pd.PropertyConfig.TaxableVacantLandConfig)
		pd.TaxableVacantLand = pd.AreaOfPlot - (maxConstructedArea * pd.PropertyConfig.TaxableVacantLandConfig)
		if pd.PercentAreaBuildup < pd.PropertyConfig.VacantLandRatePercentage {
			pd.FYs[k].VacantLandTax = pd.TaxableVacantLand * v.VLR.Rate
			if pd.FYs[k].CommonVLR == "Yes" {
				var multiplyingFactor float64
				multiplyingFactor = math.Ceil(pd.AreaOfPlot / 720)
				fmt.Println("multiplyingFactor =============> ", multiplyingFactor)
				pd.FYs[k].VacantLandTax = v.VLR.Rate * multiplyingFactor
			}
		}
		pd.FYs[k].SumARV = sumARV
		pd.FYs[k].SumFloorTax = sumAPTR
		pd.FYs[k].Tax = pd.FYs[k].Tax + sumAPTR
		pd.FYs[k].CompositeTax = pd.FYs[k].CompositeTaxRate.Rate

		if pd.OtherCharges != nil {
			//Checking Boring charge parking
			if pd.OtherCharges.BoringChargeParking != constants.PARKBORINGCHARGEYES {
				fmt.Println("Calculating boring charge")
				if !pd.IsBoringChargePayed {
					if pd.Property.MunicipalityWaterConnection == constants.MUNICIPALITYWATERCONNECTIONYES {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeWithWaterConnection
					} else if pd.Property.MunicipalityWaterConnection == constants.MUNICIPALITYWATERCONNECTIONSUPPLYOWN {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeWithWaterConnectionSupplyAndOwn
					} else if pd.Property.MunicipalityWaterConnection == constants.MUNICIPALITYWATERCONNECTIONNOTAPPLICABLE {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeNotApplicable
					} else if pd.Property.MunicipalityWaterConnection == constants.MUNICIPALITYWATERCONNECTIONALREADYPAID {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeAlreadyPaied
					} else {
						pd.BoreCharge = pd.OtherCharges.OneTimeBoringChargeWithoutWaterConnection
					}
				}
			}
			//Checking penalty parking
			if pd.OtherCharges.PenaltyParking == constants.PARKPENALTYYES || pd.ParkPenalty {
				// calculatePenalty = false
			}
		}
		fmt.Println("is current", pd.FYs[k].IsCurrent)
		if pd.FYs[k].FixedArv != nil {
			pd.FYs[k].VacantLandTax = 0
			if pd.ProductConfiguration.FixedARVTaxCalc == "PropertyTaxPercentage" {
				pd.FYs[k].Tax = (pd.FYs[k].FixedArv.ARV / 100) * pd.FYs[k].Ref.PropertyTax.Rate
			}
			if pd.ProductConfiguration.FixedARVTaxCalc == "IndividualTaxPercentage" {
				pd.FYs[k].Tax = pd.FYs[k].FixedArv.Total
			}
		}
		if pd.FYs[k].IsCurrent {
			if pd.IsRainWaterHarvesting == "Yes" {
				pd.FYs[k].Rebate = pd.FYs[k].Rebate + ((pd.FYs[k].Tax * 5) / 100)
			}
		}

		pd.FYs[k].Tax = math.Ceil(pd.FYs[k].Tax)
		pd.FYs[k].Rebate = math.Ceil(pd.FYs[k].Rebate)
		pd.FYs[k].Penalty = math.Ceil(pd.FYs[k].Penalty)
		pd.FYs[k].VacantLandTax = math.Ceil(pd.FYs[k].VacantLandTax)
		pd.FYs[k].CompositeTax = math.Ceil(pd.FYs[k].CompositeTax)
		pd.FYs[k].OtherDemand = math.Ceil(pd.FYs[k].OtherDemand)
		pd.FYs[k].Ecess = math.Ceil(pd.FYs[k].Ecess)
		pd.FYs[k].AlreadyPayed.VLTax = math.Ceil(pd.FYs[k].AlreadyPayed.VLTax)
		pd.FYs[k].AlreadyPayed.FYTax = math.Ceil(pd.FYs[k].AlreadyPayed.FYTax)
		StoredCalculationDemandfy := new(StoredCalculationDemandfy)
		StoredCalculationDemandfy.PropertyID = pd.UniqueID
		StoredCalculationDemandfy.FyId = pd.FYs[k].Name
		StoredCalculationDemandfy.Status = constants.STOREDCALCSTATUSACTIVE
		StoredCalculationDemandfy.UniqueID = fmt.Sprintf("%v_%v", StoredCalculationDemandfy.PropertyID, StoredCalculationDemandfy.FyId)
		t := time.Now()
		created := Created{}
		created.On = &t
		created.By = constants.SYSTEM
		StoredCalculationDemandfy.Created = created
		StoredCalculationDemandfy.Actual.VacantLandTax = pd.FYs[k].VacantLandTax
		StoredCalculationDemandfy.Actual.FloorTax = pd.FYs[k].Tax
		StoredCalculationDemandfy.Actual.CompositeTax = pd.FYs[k].CompositeTax
		StoredCalculationDemandfy.Actual.OtherDemand = pd.FYs[k].OtherDemand
		StoredCalculationDemandfy.Actual.Ecess = pd.FYs[k].Ecess
		StoredCalculationDemandfy.IsCurrent = pd.FYs[k].IsCurrent
		StoredCalculationDemandfy.Actual.Total = StoredCalculationDemandfy.Actual.VacantLandTax + StoredCalculationDemandfy.Actual.FloorTax + StoredCalculationDemandfy.Actual.CompositeTax + StoredCalculationDemandfy.Actual.OtherDemand + StoredCalculationDemandfy.Actual.Ecess
		StoredCalculationDemandfy.Actual.PenaltyChargeableTax = StoredCalculationDemandfy.Actual.VacantLandTax + StoredCalculationDemandfy.Actual.FloorTax
		StoredCalculationDemandfy.Actual.PenaltyNonChargeable = StoredCalculationDemandfy.Actual.OtherDemand
		StoredCalculationDemandfy.Paid.AmountWithPenalty = pd.FYs[k].AlreadyPayed.Amount
		StoredCalculationDemandfy.Paid.AmountWithOutPenalty = pd.FYs[k].AlreadyPayed.PaidTax
		StoredCalculationDemandfy.Paid.PenaltyChargeableTax = pd.FYs[k].AlreadyPayed.VLTax + pd.FYs[k].AlreadyPayed.FYTax
		StoredCalculationDemandfy.Paid.PenaltyNonChargeable = pd.FYs[k].AlreadyPayed.OtherDemand
		StoredCalculationDemandfy.Paid.Penalty = pd.FYs[k].AlreadyPayed.Penalty
		StoredCalculationDemandfy.Paid.PanelCharge = pd.FYs[k].AlreadyPayed.PanelCharge
		StoredCalculationDemandfy.Paid.Rebate = pd.FYs[k].AlreadyPayed.Rebate
		StoredCalculationDemandfy.Pending.PenaltyChargeableTax = StoredCalculationDemandfy.Actual.PenaltyChargeableTax - StoredCalculationDemandfy.Paid.PenaltyChargeableTax
		StoredCalculationDemandfy.Pending.PenaltyNonChargeable = StoredCalculationDemandfy.Actual.PenaltyNonChargeable - StoredCalculationDemandfy.Paid.PenaltyNonChargeable
		StoredCalculationDemandfy.Pending.Penalty = math.Ceil((StoredCalculationDemandfy.Pending.PenaltyChargeableTax * pd.FYs[k].PenaltyRate) / 100)
		StoredCalculationDemandfy.Pending.TotalWithPenalty = StoredCalculationDemandfy.Pending.PenaltyChargeableTax + StoredCalculationDemandfy.Pending.PenaltyNonChargeable + StoredCalculationDemandfy.Pending.Penalty
		StoredCalculationDemandfy.Pending.TotalWithOutPenalty = StoredCalculationDemandfy.Pending.PenaltyChargeableTax + StoredCalculationDemandfy.Pending.PenaltyNonChargeable
		StoredCalculationDemandfyArr = append(StoredCalculationDemandfyArr, *StoredCalculationDemandfy)

		StoredCalculationDemand.PropertyID = pd.PropertyID
		if pd.FYs[k].IsCurrent {
			StoredCalculationDemand.Actual.Current = StoredCalculationDemandfy.Actual.Total
			StoredCalculationDemand.Paid.Current.Amount = pd.FYs[k].AlreadyPayed.Amount
			StoredCalculationDemand.Paid.Current.VLTax = pd.FYs[k].AlreadyPayed.VLTax
			StoredCalculationDemand.Paid.Current.FYTax = pd.FYs[k].AlreadyPayed.FYTax
			StoredCalculationDemand.Paid.Current.CompositeTax = pd.FYs[k].AlreadyPayed.CompositeTax
			StoredCalculationDemand.Paid.Current.Ecess = pd.FYs[k].AlreadyPayed.Ecess
			StoredCalculationDemand.Paid.Current.Penalty = pd.FYs[k].AlreadyPayed.Penalty
			StoredCalculationDemand.Paid.Current.Rebate = pd.FYs[k].AlreadyPayed.Rebate
			StoredCalculationDemand.Paid.Current.PanelCharge = pd.FYs[k].AlreadyPayed.PanelCharge
			StoredCalculationDemand.Paid.Current.PaidTax = pd.FYs[k].AlreadyPayed.PaidTax
			StoredCalculationDemand.Paid.Current.OtherDemand = pd.FYs[k].AlreadyPayed.OtherDemand
		} else {
			StoredCalculationDemand.Actual.Arrear = StoredCalculationDemand.Actual.Arrear + StoredCalculationDemandfy.Actual.Total
			StoredCalculationDemand.Paid.Arrear.Amount = StoredCalculationDemand.Paid.Arrear.Amount + pd.FYs[k].AlreadyPayed.Amount
			StoredCalculationDemand.Paid.Arrear.VLTax = StoredCalculationDemand.Paid.Arrear.VLTax + pd.FYs[k].AlreadyPayed.VLTax
			StoredCalculationDemand.Paid.Arrear.FYTax = StoredCalculationDemand.Paid.Arrear.FYTax + pd.FYs[k].AlreadyPayed.FYTax
			StoredCalculationDemand.Paid.Arrear.CompositeTax = StoredCalculationDemand.Paid.Arrear.CompositeTax + pd.FYs[k].AlreadyPayed.CompositeTax
			StoredCalculationDemand.Paid.Arrear.Ecess = StoredCalculationDemand.Paid.Arrear.Ecess + pd.FYs[k].AlreadyPayed.Ecess
			StoredCalculationDemand.Paid.Arrear.Penalty = StoredCalculationDemand.Paid.Arrear.Penalty + pd.FYs[k].AlreadyPayed.Penalty
			StoredCalculationDemand.Paid.Arrear.Rebate = StoredCalculationDemand.Paid.Arrear.Rebate + pd.FYs[k].AlreadyPayed.Rebate
			StoredCalculationDemand.Paid.Arrear.PanelCharge = StoredCalculationDemand.Paid.Arrear.PanelCharge + pd.FYs[k].AlreadyPayed.PanelCharge
			StoredCalculationDemand.Paid.Arrear.PaidTax = StoredCalculationDemand.Paid.Arrear.PaidTax + pd.FYs[k].AlreadyPayed.PaidTax
			StoredCalculationDemand.Paid.Arrear.OtherDemand = StoredCalculationDemand.Paid.Arrear.OtherDemand + pd.FYs[k].AlreadyPayed.OtherDemand
		}

		StoredCalculationDemand.PropertyID = StoredCalculationDemandfy.PropertyID
		StoredCalculationDemand.UniqueID = fmt.Sprintf("%v_%v", StoredCalculationDemandfy.PropertyID, t.Format("2006-Jan-02"))
		StoredCalculationDemand.Created = created
	}
	fmt.Println("StoredCalculationDemandfyArr====>", len(StoredCalculationDemandfyArr))
	pd.FormFee = pd.OtherCharges.FormFeeCharges
	StoredCalculationDemand.Actual.FormFee = pd.FormFee
	StoredCalculationDemand.Actual.BoreCharge = pd.BoreCharge
	StoredCalculationDemand.Status = constants.STOREDCALCSTATUSACTIVE
	StoredCalculationDemand.Paid.Total.Total = StoredCalculationDemand.Paid.Arrear.Amount + StoredCalculationDemand.Paid.Current.Amount
	StoredCalculationDemand.Paid.Total.Penalty = StoredCalculationDemand.Paid.Arrear.Penalty + StoredCalculationDemand.Paid.Current.Penalty
	StoredCalculationDemand.Paid.Total.Rebate = StoredCalculationDemand.Paid.Arrear.Rebate + StoredCalculationDemand.Paid.Current.Rebate
	StoredCalculationDemand.Paid.Total.PanelCharge = StoredCalculationDemand.Paid.Arrear.PanelCharge + StoredCalculationDemand.Paid.Current.PanelCharge
	StoredCalculationDemand.Paid.Total.PaidTax = StoredCalculationDemand.Paid.Arrear.PaidTax + StoredCalculationDemand.Paid.Current.PaidTax
	StoredCalculationDemand.Paid.Total.OtherDemand = StoredCalculationDemand.Paid.Arrear.OtherDemand + StoredCalculationDemand.Paid.Current.OtherDemand

	return StoredCalculationDemandfyArr, StoredCalculationDemand
}
