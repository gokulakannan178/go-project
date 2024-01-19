package models

import (
	"encoding/json"
	"fmt"
	"math"
	"municipalproduct1-service/constants"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PropertyDemand : ""
type PropertyDemand struct {
	Property          `bson:",inline"`
	PropertyDemandLog `bson:",inline"`
	FYs               []FinancialYearDemand `json:"fys,omitempty" bson:"fys,omitempty"`
	NoteFy            []string              `json:"noteFy,omitempty" bson:"noteFy,omitempty"`

	ProductConfiguration *RefProductConfiguration `json:"-" bson:"productConfiguration,omitempty"`
	AllDemand            bool                     `json:"allDemand,omitempty" bson:"allDemand,omitempty"`
	Summary              PropertyDemandSummary    `json:"summary,omitempty" bson:"summary,omitempty"`
	FYV2Summary          []FinancialYearDemandV2  `json:"fyV2Summary,omitempty" bson:"fyV2Summary,omitempty"`

	CTX *Context
}

func (pd *PropertyDemand) Inc(i int) int {
	return i + 1
}

type PropertyDemandLog struct {
	PropertyID                string                `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	PercentAreaBuildup        float64               `json:"percentAreaBuildup,omitempty" bson:"percentAreaBuildup,omitempty"`
	TaxableVacantLand         float64               `json:"taxableVacantLand,omitempty" bson:"taxableVacantLand,omitempty"`
	ServiceCharge             float64               `json:"serviceCharge" bson:"serviceCharge"`
	IsServiceChargeApplicable bool                  `json:"isServiceChargeApplicable" bson:"isServiceChargeApplicable"`
	PropertyConfig            PropertyConfiguration `json:"propertyConfig" bson:"propertyConfig"`
	OtherDemand               float64               `json:"otherDemand" bson:"otherDemand"`
	OtherCharges              *OtherCharges         `json:"otherCharges" bson:"otherCharges"`
	FYTax                     float64               `json:"fyTax" bson:"fyTax"`
	FlTax                     float64               `json:"flTax" bson:"flTax"`
	VlTax                     float64               `json:"vlTax" bson:"vlTax"`
	Tax                       float64               `json:"tax" bson:"tax"`
	LateSubmissionCharge      float64               `json:"lateSubmissionCharge" bson:"lateSubmissionCharge"`
	BoreCharge                float64               `json:"boreCharge" bson:"boreCharge"`
	FormFee                   float64               `json:"formFee" bson:"formFee"`
	TotalTax                  float64               `json:"totalTax" bson:"totalTax"`
	Current                   float64               `json:"current" bson:"current"`
	PenalCharge               float64               `json:"penalCharge" bson:"penalCharge"`
	Arrear                    float64               `json:"arrear" bson:"arrear"`
	AdvanceReceived           float64               `json:"advanceReceived" bson:"advanceReceived"`
	CompositeTax              float64               `json:"compositeTax" bson:"compositeTax"`
	Ecess                     float64               `json:"ecess" bson:"ecess"`
	PanelCh                   float64               `json:"panelCh" bson:"panelCh"`
	// Rebate                    float64               `json:"rebate" bson:"rebate"`
	OverallPropertyDemand OverallPropertyDemand `json:"overallPropertyDemand" bson:"-"`
	AlreadyPayedMain      struct {
		BoreCharge float64 `json:"boreCharge" bson:"boreCharge"`
		FormFee    float64 `json:"formFee" bson:"formFee"`
	} `json:"alreadyPayedMain" bson:"alreadyPayedMain"`
	Ref struct {
		PropertyOwner   []RefPropertyOwner   `json:"propertyOwner" bson:"propertyOwner,omitempty"`
		Floors          []RefPropertyFloor   `json:"floors" bson:"floors,omitempty"`
		EstimatedFloors []RefEstimatedFloors `json:"estimatedFloors" bson:"estimatedFloors,omitempty"`
		Wallet          []RefPropertyWallet  `json:"wallet" bson:"wallet,omitempty"`
		Address         RefAddress           `json:"address" bson:"address,omitempty"`
		PenalCharges    PenalCharge          `json:"penalCharges" bson:"penalCharges,omitempty"`
		PropertyType    *RefPropertyType     `json:"propertyType" bson:"propertyType,omitempty"`
		YOA             *RefFinancialYear    `json:"yoa" bson:"yoa,omitempty"`
		MunicipalType   *RefMunicipalType    `json:"municipalType" bson:"municipalType,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
	NewPropertyID string `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID string `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
}

// PenalCharge : ""
type PenalCharge struct {
	Name           string     `json:"name" bson:"name,omitempty"`
	Tax            float64    `json:"tax" bson:"tax,omitempty"`
	PropertyTypeID string     `json:"propertyTypeId" bson:"propertyTypeId,omitempty"`
	DOE            *time.Time `json:"doe" bson:"doe,omitempty"`
}

type UpdateDemand struct {
	TotalTax       float64 `json:"totalTax" bson:"totalTax"`
	Current        float64 `json:"current" bson:"current"`
	Arrear         float64 `json:"arrear" bson:"arrear"`
	TotalPenalty   float64 `json:"totalPenalty" bson:"totalPenalty"`
	CurrentPenalty float64 `json:"currentPenalty" bson:"currentPenalty"`
	ArrearPenalty  float64 `json:"arrearPenalty" bson:"arrearPenalty"`
	FromYear       struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}

type UpdateCollection struct {
	TotalTax       float64 `json:"totalTax" bson:"totalTax"`
	Current        float64 `json:"current" bson:"current"`
	Arrear         float64 `json:"arrear" bson:"arrear"`
	Other          float64 `json:"other" bson:"other"`
	Penalty        float64 `json:"penalty" bson:"penalty"`
	TotalPenalty   float64 `json:"totalPenalty" bson:"totalPenalty"`
	CurrentPenalty float64 `json:"currentPenalty" bson:"currentPenalty"`
	ArrearPenalty  float64 `json:"arrearPenalty" bson:"arrearPenalty"`
	FromYear       struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}
type UpdatePenalty struct {
	Total   float64 `json:"total" bson:"total"`
	Current float64 `json:"current" bson:"current"`
	Arrear  float64 `json:"arrear" bson:"arrear"`
}

type UpdateRebate struct {
	Total   float64 `json:"total" bson:"total"`
	Current float64 `json:"current" bson:"current"`
	Arrear  float64 `json:"arrear" bson:"arrear"`
}
type PropertyUpdateCollection struct {
	Data UpdateCollection `json:"data" bson:"data"`
}
type PropertyDemandRef struct {
	PropertyOwner []RefPropertyOwner `json:"propertyOwner" bson:"propertyOwner,omitempty"`
	Floors        []RefPropertyFloor `json:"floors" bson:"floors,omitempty"`
	Address       RefAddress         `json:"address" bson:"address,omitempty"`
	PropertyType  *RefPropertyType   `json:"propertyType" bson:"propertyType,omitempty"`
	YOA           *RefFinancialYear  `json:"yoa" bson:"yoa,omitempty"`
	MunicipalType *RefMunicipalType  `json:"municipalType" bson:"municipalType,omitempty"`
}

// PropertyDemandFilter : ""
type PropertyDemandFilter struct {
	PropertyID    string         `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	Status        []string       `json:"status"`
	Address       *AddressSearch `json:"address"`
	PropertyIDs   []string       `json:"propertyIDs"`
	Fys           []string       `json:"fys,omitempty" bson:"fys,omitempty"`
	IsOmitPaidFys bool           `json:"isOmitFys,omitempty" bson:"isOmitFys,omitempty"`
	PartPayment   struct {
		IS     bool    `json:"is,omitempty" bson:"is,omitempty"`
		Amount float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	} `json:"partPayment,omitempty" bson:"partPayment,omitempty"`
	AllDemand bool `json:"allDemand,omitempty" bson:"allDemand,omitempty"`
}

// FinancialYearDemand : ""
type FinancialYearDemand struct {
	PropertyID    string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	FinancialYear `bson:",inline"`
	VacantLandTax float64    `json:"vacantLandTax" bson:"vacantLandTax"`
	Floors        []FloorTax `json:"floors" bson:"floors"`
	SumARV        float64    `json:"sumARV" bson:"sumARV"`
	SumFloorTax   float64    `json:"sumFloorTax" bson:"sumFloorTax"`
	Rebate        float64    `json:"rebate" bson:"rebate"`
	RebateRate    float64    `json:"rebateRate" bson:"rebateRate"`

	PanelCharge   float64 `json:"panelCharge" bson:"panelCharge"`
	Penalty       float64 `json:"penanty" bson:"penanty"`
	PenaltyMonths float64 `json:"penaltyMonths" bson:"penaltyMonths"`
	Tax           float64 `json:"tax" bson:"tax"`
	PaidTotalTax  float64 `json:"paidTotalTax" bson:"paidTotalTax"`
	PaidTax       float64 `json:"paidTax" bson:"paidTax"`
	PaidPenalty   float64 `json:"paidPenalty" bson:"paidPenalty"`
	CompositeTax  float64 `json:"compositeTax" bson:"compositeTax"`

	ActualTotalTax            float64 `json:"actualTotalTax" bson:"actualTotalTax"`
	ActualTax                 float64 `json:"actualTax" bson:"actualTax"`
	ActualFlTax               float64 `json:"actualFlTax" bson:"actualFlTax"`
	ActualVlTax               float64 `json:"actualVlTax" bson:"actualVlTax"`
	ActualPenalty             float64 `json:"actualPenalty" bson:"actualPenalty"`
	PaidPartPaymentPercentage float64 `json:"paidPartPaymentPercentage" bson:"paidPartPaymentPercentage"`

	PaidRate   float64 `json:"paidRate" bson:"paidRate"`
	TotalTax   float64 `json:"totalTax" bson:"totalTax"`
	OverallTax struct {
		Total        float64 `json:"total" bson:"total"`
		FYTax        float64 `json:"fyTax" bson:"fyTax"`
		VLTax        float64 `json:"vlTax" bson:"vlTax"`
		CompositeTax float64 `json:"compositeTax" bson:"compositeTax"`
		Ecess        float64 `json:"ecess" bson:"ecess"`
		Penalty      float64 `json:"penalty" bson:"penalty"`
		Rebate       float64 `json:"rebate" bson:"rebate"`
		PanelCharge  float64 `json:"panelCharge" bson:"panelCharge"`
		ToBePaid     float64 `json:"toBePaid" bson:"toBePaid"`
		PayableTax   float64 `json:"payableTax" bson:"payableTax"`
		OtherDemand  float64 `json:"otherDemand" bson:"otherDemand"`
	} `json:"overallTax" bson:"overallTax"`
	AlreadyPayed struct {
		Amount       float64 `json:"amount" bson:"amount"`
		FYTax        float64 `json:"fyTax" bson:"fyTax"`
		VLTax        float64 `json:"vlTax" bson:"vlTax"`
		CompositeTax float64 `json:"compositeTax" bson:"compositeTax"`
		Ecess        float64 `json:"ecess" bson:"ecess"`
		Penalty      float64 `json:"penalty" bson:"penalty"`
		Rebate       float64 `json:"rebate" bson:"rebate"`
		PanelCharge  float64 `json:"panelCharge" bson:"panelCharge"`
		PaidTax      float64 `json:"paidTax" bson:"paidTax"`
		OtherDemand  float64 `json:"otherDemand" bson:"otherDemand"`
	} `json:"alreadyPayed" bson:"alreadyPayed"`

	ServiceCharge                float64                `json:"serviceCharge" bson:"serviceCharge"`
	OtherDemand                  float64                `json:"otherDemand" bson:"otherDemand"`
	OtherDemandAdditionalPenalty PropertyOtherDemand    `json:"otherDemandAdditionalPenalty" bson:"otherDemandAdditionalPenalty"`
	ConstructedArea              float64                `json:"constructedArea" bson:"constructedArea"`
	Ref                          FinancialYearDemandRef `json:"ref" bson:"ref"`
	VLR                          VacantLandRate         `json:"vlr" bson:"vlr"`
	Legacy                       RefLegacyPropertyFy    `json:"legacy" bson:"legacy"`
	FixedArv                     *PropertyFixedArv      `json:"fixedArv" bson:"fixedArv"`
	FixedDemand                  *PropertyFixedDemand   `json:"propertyFixedDemand" bson:"propertyFixedDemand"`
	FloorBuildupArea             struct {
		Area float64 `json:"area" bson:"area"`
	} `json:"floorBuildupArea" bson:"floorBuildupArea"`
	CompositeTaxRate      CompositeTaxRateMaster `json:"compositeTaxRate" bson:"compositeTaxRate,omitempty"`
	PanelChargeRatemaster PanelChargeRateMaster  `json:"panelChargeRatemaster" bson:"panelChargeRatemaster,omitempty"`
	Ecess                 float64                `json:"ecess" bson:"ecess,omitempty"`
	FormFee               float64                `json:"formFee" bson:"formFee,omitempty"`
	AVRRanges             []AVRRange             `json:"avrRanges" bson:"avrRanges,omitempty"`
	AVRRange              AVRRange               `json:"avrRange" bson:"avrRange,omitempty"`

	EcessRateMaster RefEcessRateMaster `json:"ecessRateMaster" bson:"ecessRateMaster,omitempty"`
	CG              struct {
		SumARV                        float64 `json:"sumarv" bson:"sumarv"`
		MaintenanceDiscountPercentage float64 `json:"maintenanceDiscountPercentage" bson:"maintenanceDiscountPercentage"`
		MaintenanceDiscount           float64 `json:"maintenanceDiscount" bson:"maintenanceDiscount"`
		MaintenanceDiscountedARV      float64 `json:"maintenanceDiscountedARV" bson:"maintenanceDiscountedARV"`
		ResDiscountPercentage         float64 `json:"resDiscountPercentage" bson:"resDiscountPercentage"`
		ResDiscount                   float64 `json:"resDiscount" bson:"resDiscount"`
		ResDiscountedARV              float64 `json:"resDiscountedARV" bson:"resDiscountedARV"`
		// CompositeTax                  float64 `json:"compositeTax" bson:"compositeTax"`
		// Ecess                         float64 `json:"ecess" bson:"ecess"`
		TaxRate float64 `json:"taxRate" bson:"taxRate"`
	} `json:"cg" bson:"cg"`
	NewPropertyID string                 `json:"newPropertyId" bson:"newPropertyId,omitempty"`
	OldPropertyID string                 `json:"oldPropertyId" bson:"oldPropertyId,omitempty"`
	FYv2          *FinancialYearDemandV2 `json:"fyv2" bson:"fyv2,omitempty"`
}

type FixedARVDemand struct {
	Total float64 `json:"total" bson:"total"`
}

func (fyd *FinancialYearDemand) Inc(i int) int {
	return i + 1
}

func (fyd *FinancialYearDemand) Sum(a, b float64) float64 {
	return a + b
}

func (fyd *FinancialYearDemand) PendingTaxCalc(a, b, c, d float64) float64 {
	x := a + b
	y := c + d
	z := x - y
	return z
}

type FinancialYearDemandRef struct {
	PropertyTax PropertyTax `json:"propertyTax" bson:"propertyTax"`
	Penalty     Penalty     `json:"penalty" bson:"penalty"`
}

// FloorTax : ""
type FloorTax struct {
	RefPropertyFloor `bson:",inline"`
	CG               struct {
		Rate                    float64 `json:"rate" bson:"rate"`
		RateDiscount            float64 `json:"rateDiscount" bson:"rateDiscount"`
		FLoorDiscountPercentage float64 `json:"fLoorDiscountPercentage" bson:"fLoorDiscountPercentage"`
		DiscountedRate          float64 `json:"discountedRate" bson:"discountedRate"`
		ARV                     float64 `json:"arv" bson:"arv"`
	} `json:"cg" bson:"cg"`
	ARV           float64 `json:"arv" bson:"arv"`
	APTR          float64 `json:"aptr" bson:"aptr"`
	CompostiteTax float64 `json:"CompostiteTax" bson:"compostiteTax"`
}

// DemandCalculation : ""
func (pd *PropertyDemand) DemandCalculation() PropertyDemand {
	var legacyAmount float64
	pd.FYTax = 0
	pds := PropertyDemandSummary{}
	pdfyv2 := make([]FinancialYearDemandV2, 0)

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
		var groundFloorBuildUpArea float64
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
			fmt.Println("pd.ProductConfiguration.LocationID=========> ", pd.ProductConfiguration.LocationID)
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
			if v1.No == "16" {
				groundFloorBuildUpArea = v1.BuildUpArea
			}
		}
		fmt.Println("GroundFloorBuildUpArea =", groundFloorBuildUpArea)
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
			pd.FYs[k].Ecess = 0
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
		if pd.ProductConfiguration != nil {
			if pd.ProductConfiguration.RemoveBuildUpAreaRestriction == "Yes" {
				maxConstructedArea = groundFloorBuildUpArea
			}
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
		// if pd.FYs[k].FixedDemand != nil {
		// 	pd.FYs[k].VacantLandTax = 0
		// 	pd.FYs[k].Tax = pd.FYs[k].FixedDemand.Total
		// }
		pd.FYs[k].OverallTax.FYTax = pd.FYs[k].Tax
		pd.FYs[k].OverallTax.VLTax = pd.FYs[k].VacantLandTax
		pd.FYs[k].OverallTax.CompositeTax = pd.FYs[k].CompositeTax
		pd.FYs[k].OverallTax.Ecess = pd.FYs[k].Ecess
		pd.FYs[k].OverallTax.ToBePaid = pd.FYs[k].Tax + pd.FYs[k].VacantLandTax + pd.FYs[k].CompositeTax + pd.FYs[k].Ecess + pd.FYs[k].OtherDemand
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
			//By Baskar(baskar.i@logikoof.in)
			//Commenting Rebate -comment and uncomment following 2 lines

			if pd.IsRainWaterHarvesting == "Yes" {
				pd.FYs[k].RebateRate = pd.FYs[k].RebateRate + 5

				pd.FYs[k].Rebate = pd.FYs[k].Rebate + (((pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax) - pd.FYs[k].Legacy.TaxAmount) * 5) / 100)
			}
			//if math.Ceil(pd.FYs[k].AlreadyPayed.Amount) < (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax) {
			pd.FYs[k].Penalty = (((pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax) - pd.FYs[k].Legacy.TaxAmount) * pd.FYs[k].PenaltyRate) / 100) * float64(months+1)

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
				// pd.FYs[k].Penalty = (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax)) * (pd.FYs[k].PenaltyRate / 100)
				pd.FYs[k].Penalty = (pd.FYs[k].Tax + pd.FYs[k].VacantLandTax - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax) - pd.FYs[k].Legacy.TaxAmount) * (pd.FYs[k].PenaltyRate / 100)
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
		if pd.ParkPenalty {
			pd.FYs[k].Penalty = 0
		}
		// if pd.IsRainWaterHarvesting == "Yes" {
		// 	pd.FYs[k].Rebate = pd.FYs[k].Rebate + ((pd.FYs[k].Tax * 5) / 100)
		// }
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

		if pd.ProductConfiguration.LocationID == "CGBhilai" {
			if fytax > 0 {
				if !pd.FYs[k].IsCurrent {

					pd.FYs[k].PanelCharge = pd.FYs[k].PanelChargeRatemaster.Rate
				}
				// pd.FYs[k].FormFee = 3

			}
			fytax = fytax + pd.FYs[k].Ecess + pd.FYs[k].FormFee
		}

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
			pd.FYs[k].TotalTax = fytax - (pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax + pd.FYs[k].AlreadyPayed.CompositeTax + pd.FYs[k].AlreadyPayed.Ecess + pd.FYs[k].AlreadyPayed.OtherDemand + pd.FYs[k].Legacy.TaxAmount)
			if !pd.FYs[k].IsCurrent {
				if pd.FYs[k].TotalTax > 0 {
					pd.FYs[k].TotalTax = pd.FYs[k].TotalTax + pd.FYs[k].PanelCharge
				}
			}
			if pd.FYs[k].TotalTax <= 1 {
				pd.FYs[k].TotalTax = 0
			}
			pd.FYTax = pd.FYTax + pd.FYs[k].TotalTax

			fmt.Println("After adding- ", pd.FYTax, pd.FYs[k].TotalTax)

		}
		fmt.Println("Current FY tax - ", pd.FYTax, pd.FYs[k].TotalTax)
		if pd.FYs[k].IsCurrent {
			pd.Current = pd.Current + pd.FYs[k].TotalTax + pd.FYs[k].ServiceCharge
		} else {
			pd.Arrear = pd.Arrear + pd.FYs[k].TotalTax + pd.FYs[k].ServiceCharge
		}

		pd.FlTax = pd.FYs[k].Tax
		pd.VlTax = pd.FYs[k].VacantLandTax
		pd.Tax = pd.FYs[k].Tax + pd.FYs[k].VacantLandTax

		if pd.FYs[k].IsCurrent {

			pd.OverallPropertyDemand.Current.VacantLandTax = pd.OverallPropertyDemand.Current.VacantLandTax + pd.FYs[k].VacantLandTax
			pd.OverallPropertyDemand.Current.Rebate = pd.OverallPropertyDemand.Current.Rebate + pd.FYs[k].Rebate
			pd.OverallPropertyDemand.Current.Penalty = pd.OverallPropertyDemand.Current.Penalty + pd.FYs[k].Penalty
			pd.OverallPropertyDemand.Current.Tax = pd.OverallPropertyDemand.Current.Tax + pd.FYs[k].Tax
			pd.OverallPropertyDemand.Current.CompositeTax = pd.OverallPropertyDemand.Current.CompositeTax + pd.FYs[k].CompositeTax
			pd.OverallPropertyDemand.Current.Ecess = pd.OverallPropertyDemand.Current.Ecess + pd.FYs[k].Ecess
			pd.OverallPropertyDemand.Current.PanelCh = pd.OverallPropertyDemand.Current.PanelCh + pd.FYs[k].PanelCharge

			//Actual Overall Demand
			pd.OverallPropertyDemand.Actual.Current.VacantLandTax = pd.OverallPropertyDemand.Actual.Current.VacantLandTax + pd.FYs[k].OverallTax.VLTax
			pd.OverallPropertyDemand.Actual.Current.Tax = pd.OverallPropertyDemand.Actual.Current.Tax + pd.FYs[k].OverallTax.FYTax
			pd.OverallPropertyDemand.Actual.Current.TotalTax = pd.OverallPropertyDemand.Actual.Current.TotalTax + pd.FYs[k].OverallTax.Total

		} else {
			pd.OverallPropertyDemand.Arrear.VacantLandTax = pd.OverallPropertyDemand.Arrear.VacantLandTax + pd.FYs[k].VacantLandTax
			pd.OverallPropertyDemand.Arrear.Rebate = pd.OverallPropertyDemand.Arrear.Rebate + pd.FYs[k].Rebate
			pd.OverallPropertyDemand.Arrear.Penalty = pd.OverallPropertyDemand.Arrear.Penalty + pd.FYs[k].Penalty
			pd.OverallPropertyDemand.Arrear.Tax = pd.OverallPropertyDemand.Arrear.Tax + pd.FYs[k].Tax
			pd.OverallPropertyDemand.Arrear.CompositeTax = pd.OverallPropertyDemand.Arrear.CompositeTax + pd.FYs[k].CompositeTax
			pd.OverallPropertyDemand.Arrear.Ecess = pd.OverallPropertyDemand.Arrear.Ecess + pd.FYs[k].Ecess
			pd.OverallPropertyDemand.Arrear.PanelCh = pd.OverallPropertyDemand.Arrear.PanelCh + pd.FYs[k].PanelCharge

			//Actual Overall Demand
			pd.OverallPropertyDemand.Actual.Arrear.VacantLandTax = pd.OverallPropertyDemand.Actual.Arrear.VacantLandTax + pd.FYs[k].OverallTax.VLTax
			pd.OverallPropertyDemand.Actual.Arrear.Tax = pd.OverallPropertyDemand.Actual.Arrear.Tax + pd.FYs[k].OverallTax.FYTax
			pd.OverallPropertyDemand.Actual.Arrear.TotalTax = pd.OverallPropertyDemand.Actual.Arrear.TotalTax + pd.FYs[k].OverallTax.Total

		}
		//Create new FY Log by Solomon Arumugam 13-Nov-2023
		if pd.FYs[k].FYv2 == nil {
			pd.FYs[k].FYv2 = new(FinancialYearDemandV2)
			pd.FYs[k].FYv2.PropertyID = pd.UniqueID
			pd.FYs[k].FYv2.FinancialYearId = v.UniqueID
			pd.FYs[k].FYv2.IsCurrent = v.IsCurrent
			// [A] - CALCULATING TOTAL DEMAND
			pd.FYs[k].FYv2.Demand.FLTax = pd.FYs[k].Tax
			pd.FYs[k].FYv2.Demand.VLTax = pd.FYs[k].VacantLandTax
			pd.FYs[k].FYv2.Demand.Tax = pd.FYs[k].Tax + pd.FYs[k].VacantLandTax
			pd.FYs[k].FYv2.Demand.OtherDemand = pd.FYs[k].OtherDemand

			//Since usercharge is not calculated with demand - it is given ZERO for now
			pd.FYs[k].FYv2.Demand.UserCharge = 0

			/*
				This is overall Demand
				So it does not have
				1) Penalty
				2) Rebate
				So giving ZERO
			*/
			pd.FYs[k].FYv2.Demand.Penalty = 0
			pd.FYs[k].FYv2.Demand.Rebate = 0

			//TotalTax = Tax+OD+UC
			pd.FYs[k].FYv2.Demand.TotalTax = pd.FYs[k].FYv2.Demand.Tax + pd.FYs[k].FYv2.Demand.UserCharge + pd.FYs[k].FYv2.Demand.OtherDemand

			////Total = TotalTax+Penalty-Rebate
			pd.FYs[k].FYv2.Demand.Total = pd.FYs[k].FYv2.Demand.TotalTax + pd.FYs[k].FYv2.Demand.Penalty - pd.FYs[k].FYv2.Demand.Rebate
			//END OF [A] - CALCULATING TOTAL DEMAND

			// [B] - CALCULATING PAID COLLECTIONS
			pd.FYs[k].FYv2.Collections.FlTax = pd.FYs[k].AlreadyPayed.FYTax
			pd.FYs[k].FYv2.Collections.VlTax = pd.FYs[k].AlreadyPayed.VLTax
			pd.FYs[k].FYv2.Collections.Tax = pd.FYs[k].AlreadyPayed.FYTax + pd.FYs[k].AlreadyPayed.VLTax

			pd.FYs[k].FYv2.Collections.OtherDemand = pd.FYs[k].AlreadyPayed.OtherDemand

			//Since usercharge is not calculated with demand - it is given ZERO for now
			pd.FYs[k].FYv2.Collections.UserCharge = 0

			pd.FYs[k].FYv2.Collections.Penalty = pd.FYs[k].AlreadyPayed.Penalty
			pd.FYs[k].FYv2.Collections.Rebate = pd.FYs[k].AlreadyPayed.Rebate
			pd.FYs[k].FYv2.Collections.BoringCharge = pd.AlreadyPayedMain.BoreCharge
			pd.FYs[k].FYv2.Collections.FormFee = pd.AlreadyPayedMain.FormFee

			//TotalTax = Tax+OD+UC
			pd.FYs[k].FYv2.Collections.TotalTax = pd.FYs[k].FYv2.Collections.Tax + pd.FYs[k].FYv2.Collections.UserCharge + pd.FYs[k].FYv2.Collections.OtherDemand

			////Total = TotalTax+Penalty-Rebate+Formfee+BoringCharge
			pd.FYs[k].FYv2.Collections.Total = pd.FYs[k].FYv2.Collections.TotalTax + pd.FYs[k].FYv2.Collections.Penalty - pd.FYs[k].FYv2.Collections.Rebate
			// END OF [B] - CALCULATING PAID COLLECTIONS

			// [C] - CALCULATING PENDING DEMAND
			pd.FYs[k].FYv2.ToPay.FlTax = pd.FYs[k].Tax - pd.FYs[k].AlreadyPayed.FYTax
			pd.FYs[k].FYv2.ToPay.VlTax = pd.FYs[k].VacantLandTax - pd.FYs[k].AlreadyPayed.VLTax
			pd.FYs[k].FYv2.ToPay.Tax = pd.FYs[k].FYv2.ToPay.FlTax + pd.FYs[k].FYv2.ToPay.VlTax
			pd.FYs[k].FYv2.ToPay.OtherDemand = pd.FYs[k].OtherDemand - pd.FYs[k].AlreadyPayed.OtherDemand
			pd.FYs[k].FYv2.ToPay.UserCharge = 0

			pd.FYs[k].FYv2.ToPay.Penalty = pd.FYs[k].Penalty
			pd.FYs[k].FYv2.ToPay.PenaltyRate = pd.FYs[k].PenaltyRate

			pd.FYs[k].FYv2.ToPay.Rebate = pd.FYs[k].Rebate
			pd.FYs[k].FYv2.ToPay.RebateRate = pd.FYs[k].RebateRate
			//TotalTax = Tax+OD+UC
			pd.FYs[k].FYv2.ToPay.TotalTax = pd.FYs[k].FYv2.ToPay.Tax + pd.FYs[k].FYv2.ToPay.UserCharge + pd.FYs[k].FYv2.ToPay.OtherDemand

			////Total = TotalTax+Penalty-Rebate
			pd.FYs[k].FYv2.ToPay.Total = pd.FYs[k].FYv2.ToPay.TotalTax + pd.FYs[k].FYv2.ToPay.Penalty - pd.FYs[k].FYv2.ToPay.Rebate

			// END OF [C] - CALCULATING PENDING DEMAND
			// [D] - SAVING SUMMARY TO PROPERTY
			if pd.FYs[k].IsCurrent {
				//For Demand
				pds.Demand.Current.Tax = pds.Demand.Current.Tax + pd.FYs[k].FYv2.Demand.Tax
				pds.Demand.Current.OtherDemand = pds.Demand.Current.OtherDemand + pd.FYs[k].FYv2.Demand.OtherDemand
				pds.Demand.Current.UserCharge = pds.Demand.Current.UserCharge + pd.FYs[k].FYv2.Demand.UserCharge
				pds.Demand.Current.TotalTax = pds.Demand.Current.TotalTax + pd.FYs[k].FYv2.Demand.TotalTax

				//For Collections
				pds.Collections.Current.Tax = pds.Collections.Current.Tax + pd.FYs[k].FYv2.Collections.Tax
				pds.Collections.Current.OtherDemand = pds.Collections.Current.OtherDemand + pd.FYs[k].FYv2.Collections.OtherDemand
				pds.Collections.Current.UserCharge = pds.Collections.Current.UserCharge + pd.FYs[k].FYv2.Collections.UserCharge
				pds.Collections.Current.TotalTax = pds.Collections.Current.TotalTax + pd.FYs[k].FYv2.Collections.TotalTax
				pds.Collections.Current.Rebate = pds.Collections.Current.Rebate + pd.FYs[k].FYv2.Collections.Rebate
				pds.Collections.Current.Penalty = pds.Collections.Current.Penalty + pd.FYs[k].FYv2.Collections.Penalty
				pds.Collections.Current.Total = pds.Collections.Current.Total + pd.FYs[k].FYv2.Collections.Total

				//To Pay
				pds.ToPay.Current.Tax = pds.ToPay.Current.Tax + pd.FYs[k].FYv2.ToPay.Tax
				pds.ToPay.Current.OtherDemand = pds.ToPay.Current.OtherDemand + pd.FYs[k].FYv2.ToPay.OtherDemand
				pds.ToPay.Current.UserCharge = pds.ToPay.Current.UserCharge + pd.FYs[k].FYv2.ToPay.UserCharge
				pds.ToPay.Current.TotalTax = pds.ToPay.Current.TotalTax + pd.FYs[k].FYv2.ToPay.TotalTax
				pds.ToPay.Current.Rebate = pds.ToPay.Current.Rebate + pd.FYs[k].FYv2.ToPay.Rebate
				pds.ToPay.Current.Penalty = pds.ToPay.Current.Penalty + pd.FYs[k].FYv2.ToPay.Penalty
				pds.ToPay.Current.Total = pds.ToPay.Current.Total + pd.FYs[k].FYv2.ToPay.Total

			} else {
				//For Demand
				pds.Demand.Arrear.Tax = pds.Demand.Arrear.Tax + pd.FYs[k].FYv2.Demand.Tax
				pds.Demand.Arrear.OtherDemand = pds.Demand.Arrear.OtherDemand + pd.FYs[k].FYv2.Demand.OtherDemand
				pds.Demand.Arrear.UserCharge = pds.Demand.Arrear.UserCharge + pd.FYs[k].FYv2.Demand.UserCharge
				pds.Demand.Arrear.TotalTax = pds.Demand.Arrear.TotalTax + pd.FYs[k].FYv2.Demand.TotalTax

				//For Collections
				pds.Collections.Arrear.Tax = pds.Collections.Arrear.Tax + pd.FYs[k].FYv2.Collections.Tax
				pds.Collections.Arrear.OtherDemand = pds.Collections.Arrear.OtherDemand + pd.FYs[k].FYv2.Collections.OtherDemand
				pds.Collections.Arrear.UserCharge = pds.Collections.Arrear.UserCharge + pd.FYs[k].FYv2.Collections.UserCharge
				pds.Collections.Arrear.TotalTax = pds.Collections.Arrear.TotalTax + pd.FYs[k].FYv2.Collections.TotalTax
				pds.Collections.Arrear.Rebate = pds.Collections.Arrear.Rebate + pd.FYs[k].FYv2.Collections.Rebate
				pds.Collections.Arrear.Penalty = pds.Collections.Arrear.Penalty + pd.FYs[k].FYv2.Collections.Penalty
				pds.Collections.Arrear.Total = pds.Collections.Arrear.Total + pd.FYs[k].FYv2.Collections.Total

				//To Pay
				pds.ToPay.Arrear.Tax = pds.ToPay.Arrear.Tax + pd.FYs[k].FYv2.ToPay.Tax
				pds.ToPay.Arrear.OtherDemand = pds.ToPay.Arrear.OtherDemand + pd.FYs[k].FYv2.ToPay.OtherDemand
				pds.ToPay.Arrear.UserCharge = pds.ToPay.Arrear.UserCharge + pd.FYs[k].FYv2.ToPay.UserCharge
				pds.ToPay.Arrear.TotalTax = pds.ToPay.Arrear.TotalTax + pd.FYs[k].FYv2.ToPay.TotalTax
				pds.ToPay.Arrear.Rebate = pds.ToPay.Arrear.Rebate + pd.FYs[k].FYv2.ToPay.Rebate
				pds.ToPay.Arrear.Penalty = pds.ToPay.Arrear.Penalty + pd.FYs[k].FYv2.ToPay.Penalty
				pds.ToPay.Arrear.Total = pds.ToPay.Arrear.Total + pd.FYs[k].FYv2.ToPay.Total

			}
			//Totaling Property Demand Summmary
			//For Demand
			pds.Demand.Total.Tax = pds.Demand.Total.Tax + pd.FYs[k].FYv2.Demand.Tax
			pds.Demand.Total.OtherDemand = pds.Demand.Total.OtherDemand + pd.FYs[k].FYv2.Demand.OtherDemand
			pds.Demand.Total.UserCharge = pds.Demand.Total.UserCharge + pd.FYs[k].FYv2.Demand.UserCharge
			pds.Demand.Total.TotalTax = pds.Demand.Total.TotalTax + pd.FYs[k].FYv2.Demand.TotalTax

			//For Collections
			pds.Collections.Total.Tax = pds.Collections.Total.Tax + pd.FYs[k].FYv2.Collections.Tax
			pds.Collections.Total.OtherDemand = pds.Collections.Total.OtherDemand + pd.FYs[k].FYv2.Collections.OtherDemand
			pds.Collections.Total.UserCharge = pds.Collections.Total.UserCharge + pd.FYs[k].FYv2.Collections.UserCharge
			pds.Collections.Total.TotalTax = pds.Collections.Total.TotalTax + pd.FYs[k].FYv2.Collections.TotalTax
			pds.Collections.Total.Rebate = pds.Collections.Total.Rebate + pd.FYs[k].FYv2.Collections.Rebate
			pds.Collections.Total.Penalty = pds.Collections.Total.Penalty + pd.FYs[k].FYv2.Collections.Penalty
			pds.Collections.Total.BoringCharge = pd.AlreadyPayedMain.BoreCharge
			pds.Collections.Total.FormFee = pd.AlreadyPayedMain.FormFee
			pds.Collections.Total.Total = pds.Collections.Total.Total + pd.FYs[k].FYv2.Collections.Total

			//To Pay
			pds.ToPay.Total.Tax = pds.ToPay.Total.Tax + pd.FYs[k].FYv2.ToPay.Tax
			pds.ToPay.Total.OtherDemand = pds.ToPay.Total.OtherDemand + pd.FYs[k].FYv2.ToPay.OtherDemand
			pds.ToPay.Total.UserCharge = pds.ToPay.Total.UserCharge + pd.FYs[k].FYv2.ToPay.UserCharge
			pds.ToPay.Total.TotalTax = pds.ToPay.Total.TotalTax + pd.FYs[k].FYv2.ToPay.TotalTax
			pds.ToPay.Total.Rebate = pds.ToPay.Total.Rebate + pd.FYs[k].FYv2.ToPay.Rebate
			pds.ToPay.Total.Penalty = pds.ToPay.Total.Penalty + pd.FYs[k].FYv2.ToPay.Penalty
			pds.ToPay.Total.Total = pds.ToPay.Total.Total + pd.FYs[k].FYv2.ToPay.Total

			//Saving FY to demand
			fmt.Println(pd.FYs[k].FYv2.FinancialYearId)
			opts := options.Update().SetUpsert(true)
			query := bson.M{"propertyId": pd.FYs[k].FYv2.PropertyID, "financialyearId": pd.FYs[k].FYv2.FinancialYearId}
			update := bson.M{"$set": pd.FYs[k].FYv2}
			res, err := pd.CTX.DB.Collection(constants.COLLECTIONOSTOREDPROPERTYDEMANDFYS).UpdateOne(pd.CTX.CTX, query, update, opts)
			if err != nil {
				fmt.Println("Error in saving FY Demand - ", err.Error())
			}
			fmt.Println(res)

		}
		pdfyv2 = append(pdfyv2, *pd.FYs[k].FYv2)
		fmt.Printf("####################### End Of %v #########################", v.Name)

	}
	//Form fee and boring charge are added only once
	pds.Collections.Total.Total = pds.Collections.Total.Total + pds.Collections.Total.BoringCharge + pds.Collections.Total.FormFee
	//Saving Property Demand Summary
	opts := options.Update().SetUpsert(true)
	query2 := bson.M{"uniqueId": pd.UniqueID}
	update2 := bson.M{"$set": bson.M{"summary": pds}}
	res2, err2 := pd.CTX.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(pd.CTX.CTX, query2, update2, opts)
	if err2 != nil {
		fmt.Println("Error in saving Property Demand Summary - ", err2.Error())
	}
	fmt.Println(res2)
	pd.Summary = pds
	pd.FYV2Summary = pdfyv2
	if !pd.AllDemand {
		tempFYdemand := make([]FinancialYearDemand, 0)
		for _, v := range pd.FYs {
			if v.TotalTax > 0 {
				tempFYdemand = append(tempFYdemand, v)
			}
		}
		pd.FYs = tempFYdemand
	}

	if len(pd.FYs) > 0 {
		if pd.OtherCharges.FormFeeParking == constants.PARKFORMFEENO {
			if pd.IsFormFeePayed {
				pd.FormFee = 0

			} else {
				pd.FormFee = pd.OtherCharges.FormFeeCharges
			}

		}
	} else {
		pd.FormFee = 0
		pd.BoreCharge = 0
	}

	if pd.ProductConfiguration.LocationID == "CGBhilai" {
		pd.FormFee = 3
	}
	if pd.ProductConfiguration.LocationID == "Bhagalpur" {
		if pd.FYTax == 0 {
			pd.FormFee = 0
		} else {
			pd.FormFee = 5
		}
	}
	pd.OverallPropertyDemand.Other.FormFee = pd.FormFee
	pd.OverallPropertyDemand.Other.BoreCharge = pd.BoreCharge
	pd.OverallPropertyDemand.Other.OtherDemand = pd.OtherDemand

	// pd.TotalTax = pd.FYTax + pd.ServiceCharge + pd.BoreCharge + pd.FormFee
	pd.TotalTax = pd.FYTax + pd.ServiceCharge + pd.BoreCharge + pd.FormFee

	fmt.Println("pd.FYTax =========>", pd.FYTax)
	fmt.Println("pd.ServiceCharge =========>", pd.ServiceCharge)
	fmt.Println("pd.BoreCharge =========>", pd.BoreCharge)
	fmt.Println("pd.OtherDemand =========>", pd.OtherDemand)
	fmt.Println("pd.TotalTax =========>", pd.TotalTax)
	if !pd.PreviousCollection.IsCalculated {
		pd.TotalTax = pd.TotalTax - pd.PreviousCollection.Amount
	}

	//Adding Advance
	if pd.TotalTax < 0 {
		pd.AdvanceReceived = pd.TotalTax
		pd.TotalTax = 0
	}

	return *pd
}

func (pd *PropertyDemand) GetCGPropertyTaxRate(ctx *Context, arv float64, date *time.Time) (*RefAVRRange, error) {
	query := []bson.M{
		bson.M{"$match": bson.M{
			"from": bson.M{"$lte": arv}, "to": bson.M{"$gte": arv}, "doe": bson.M{"$lte": date},
		}},
		bson.M{"$sort": bson.M{"doe": -1}},
	}
	b, err1 := json.Marshal(query)
	fmt.Println("err1", err1, string(b))
	fmt.Println("query = ", query)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).Aggregate(ctx.CTX, query, nil)
	if err != nil {
		return nil, err
	}
	var avrRanges []RefAVRRange
	var avrRange *RefAVRRange
	if err = cursor.All(ctx.CTX, &avrRanges); err != nil {
		return nil, err
	}
	if len(avrRanges) > 0 {
		avrRange = &avrRanges[0]
	} else {
		avrRange = new(RefAVRRange)
	}
	return avrRange, nil
}

func monthsCountSince(createdAtTime time.Time) int {
	now := time.Now()
	months := 0
	month := createdAtTime.Month()
	for createdAtTime.Before(now) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextMonth := createdAtTime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

// Demand : ""
type Demand struct {
	Current float64 `json:"current" bson:"current"`
	Arrear  float64 `json:"arrear" bson:"arrear"`
	Total   float64 `json:"total" bson:"total"`
}
type PropertyDemandSummary struct {
	Demand struct {
		Current struct {
			Tax         float64 `json:"tax" bson:"tax"`
			OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
			UserCharge  float64 `json:"userCharge" bson:"userCharge"`
			TotalTax    float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
		}

		Arrear struct {
			Tax         float64 `json:"tax" bson:"tax"`
			OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
			UserCharge  float64 `json:"userCharge" bson:"userCharge"`
			TotalTax    float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
		}

		Total struct {
			Tax         float64 `json:"tax" bson:"tax"`
			OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
			UserCharge  float64 `json:"userCharge" bson:"userCharge"`
			TotalTax    float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
		}
	} `json:"demand,omitempty" bson:"demand,omitempty"`
	Collections struct {
		Current struct {
			Tax          float64 `json:"tax" bson:"tax"`
			OtherDemand  float64 `json:"otherDemand" bson:"otherDemand"`
			UserCharge   float64 `json:"userCharge" bson:"userCharge"`
			TotalTax     float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
			Rebate       float64 `json:"rebate" bson:"rebate"`
			Penalty      float64 `json:"penalty" bson:"penalty"`
			BoringCharge float64 `json:"boringCharge" bson:"boringCharge"`
			FormFee      float64 `json:"Formfee" bson:"Formfee"`
			Total        float64 `json:"total" bson:"total"` //totaltax-rebate+penalty+Boringcharge+formfee
		}

		Arrear struct {
			Tax         float64 `json:"tax" bson:"tax"`
			OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
			UserCharge  float64 `json:"userCharge" bson:"userCharge"`
			TotalTax    float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
			Rebate      float64 `json:"rebate" bson:"rebate"`
			Penalty     float64 `json:"penalty" bson:"penalty"`
			// BoringCharge float64 `json:"boringCharge" bson:"boringCharge"`
			// FormFee      float64 `json:"Formfee" bson:"Formfee"`
			Total float64 `json:"total" bson:"total"` //totaltax-rebate+penalty+Boringcharge+formfee
		}

		Total struct {
			Tax          float64 `json:"tax" bson:"tax"`
			OtherDemand  float64 `json:"otherDemand" bson:"otherDemand"`
			UserCharge   float64 `json:"userCharge" bson:"userCharge"`
			TotalTax     float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
			Rebate       float64 `json:"rebate" bson:"rebate"`
			Penalty      float64 `json:"penalty" bson:"penalty"`
			BoringCharge float64 `json:"boringCharge" bson:"boringCharge"`
			FormFee      float64 `json:"Formfee" bson:"Formfee"`
			Total        float64 `json:"total" bson:"total"` //totaltax-rebate+penalty+Boringcharge+formfee
		}
	} `json:"collection,omitempty" bson:"collection,omitempty"`
	ToPay struct {
		Current struct {
			Tax         float64 `json:"tax" bson:"tax"`
			OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
			UserCharge  float64 `json:"userCharge" bson:"userCharge"`
			TotalTax    float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
			Rebate      float64 `json:"rebate" bson:"rebate"`
			Penalty     float64 `json:"penalty" bson:"penalty"`
			Total       float64 `json:"total" bson:"total"` //totaltax-rebate+penalty
		}
		Arrear struct {
			Tax         float64 `json:"tax" bson:"tax"`
			OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
			UserCharge  float64 `json:"userCharge" bson:"userCharge"`
			TotalTax    float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
			Rebate      float64 `json:"rebate" bson:"rebate"`
			Penalty     float64 `json:"penalty" bson:"penalty"`
			Total       float64 `json:"total" bson:"total"` //totaltax-rebate+penalty
		}
		Total struct {
			Tax         float64 `json:"tax" bson:"tax"`
			OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
			UserCharge  float64 `json:"userCharge" bson:"userCharge"`
			TotalTax    float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
			Rebate      float64 `json:"rebate" bson:"rebate"`
			Penalty     float64 `json:"penalty" bson:"penalty"`
			Total       float64 `json:"total" bson:"total"` //totaltax-rebate+penalty
		}
	} `json:"toPay,omitempty" bson:"toPay,omitempty"`
}
type FinancialYearDemandV2 struct {
	PropertyID      string `json:"propertyId,omitempty" bson:"propertyId,omitempty"`
	FinancialYearId string ` json:"financialyearId" bson:",financialyearId"`
	IsCurrent       bool   ` json:"isCurrent" bson:",isCurrent"`
	Demand          struct {
		VLTax       float64 `json:"vLTax" bson:"vLTax"`
		FLTax       float64 `json:"fLTax" bson:"fLTax"`
		Tax         float64 `json:"tax" bson:"tax"`
		OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
		UserCharge  float64 `json:"userCharge" bson:"userCharge"`
		Rebate      float64 `json:"rebate" bson:"rebate"`
		Penalty     float64 `json:"penalty" bson:"penalty"`
		TotalTax    float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
		Total       float64 `json:"total" bson:"total"`       //TotalTax+Penalty-Rebate
		Others      float64 `json:"others" bson:"others"`
	} `json:"demand,omitempty" bson:"demand,omitempty"`
	Collections struct {
		VlTax        float64 `json:"vlTax" bson:"vlTax"`
		FlTax        float64 `json:"flTax" bson:"flTax"`
		Tax          float64 `json:"tax" bson:"tax"`
		OtherDemand  float64 `json:"otherDemand" bson:"otherDemand"`
		UserCharge   float64 `json:"userCharge" bson:"userCharge"`
		TotalTax     float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
		Rebate       float64 `json:"rebate" bson:"rebate"`
		Penalty      float64 `json:"penalty" bson:"penalty"`
		Total        float64 `json:"total" bson:"total"` //TotalTax+Penalty-Rebate
		Others       float64 `json:"others" bson:"others"`
		BoringCharge float64 `json:"boringCharge" bson:"boringCharge"`
		FormFee      float64 `json:"Formfee" bson:"Formfee"`
	} `json:"collections,omitempty" bson:"collections,omitempty"`
	ToPay struct {
		VlTax       float64 `json:"vlTax" bson:"vlTax"`
		FlTax       float64 `json:"flTax" bson:"flTax"`
		Tax         float64 `json:"tax" bson:"tax"`
		OtherDemand float64 `json:"otherDemand" bson:"otherDemand"`
		UserCharge  float64 `json:"userCharge" bson:"userCharge"`
		TotalTax    float64 `json:"totalTax" bson:"totalTax"` //Tax+OD+UC
		Penalty     float64 `json:"penalty" bson:"penalty"`
		Rebate      float64 `json:"rebate" bson:"rebate"`
		PenaltyRate float64 `json:"penaltyRate" bson:"penaltyRate"`
		RebateRate  float64 `json:"rebateRate" bson:"rebateRate"`
		Total       float64 `json:"total" bson:"total"` //TotalTax+Penalty-Rebate
		Others      float64 `json:"others" bson:"others"`
	} `json:"toPay,omitempty" bson:"toPay	"`
	Ref struct {
		Fy                     FinancialYear `json:"fy" bson:"fy,omitempty"`
		EarlyPaymentRebate     Rebate        `json:"earlyPaymentRebate" bson:"earlyPaymentRebate,omitempty"`
		RainWaterHarvestRebate Rebate        `json:"rainWaterHarvestRebate" bson:"rainWaterHarvestRebate,omitempty"`
	} `json:"-" bson:"ref,omitempty"`
}

type StoredPropertyDemand struct {
	Fys []FinancialYearDemandV2 ` json:"fys" bson:",fys"`
}

type DemandV3 struct {
	StoredPropertyDemand `bson:",inline"`
}

type SavedDemand struct {
}

// DemandCalculation : ""
func (pd *PropertyDemand) PreSaveDemandCalculation() *StoredPropertyDemand {

	storeFys := make([]FinancialYearDemandV2, 0)
	//pds := new(PropertyDemandSummary)

	fmt.Println("pd fyslen", len(pd.FYs))
	for k, v := range pd.FYs {
		storeFy := FinancialYearDemandV2{}
		var sumARV, sumAPTR float64
		var maxConstructedArea float64
		maxConstructedAreaMap := make(map[string]float64)
		for k1, v1 := range v.Floors {

			maxConstructedAreaMap[v1.No] = maxConstructedAreaMap[v1.No] + v1.BuildUpArea

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
			floorarv := buildupArea * v1.Ref.AVR.Rate
			v1.ARV = floorarv * v1.Ref.OccupancyType.Factor
			v1.APTR = (v1.ARV * v.Ref.PropertyTax.Rate) / 100

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
		//Calculating vacant land
		pd.FYs[k].ConstructedArea = maxConstructedArea
		pd.PercentAreaBuildup = (maxConstructedArea / pd.AreaOfPlot) * 100
		pd.PercentAreaBuildup = math.Ceil(pd.PercentAreaBuildup)
		pd.TaxableVacantLand = pd.AreaOfPlot - (maxConstructedArea * pd.PropertyConfig.TaxableVacantLandConfig)
		if pd.PercentAreaBuildup < pd.PropertyConfig.VacantLandRatePercentage {
			pd.FYs[k].VacantLandTax = pd.TaxableVacantLand * v.VLR.Rate
		}

		pd.FYs[k].SumARV = sumARV
		pd.FYs[k].SumFloorTax = sumAPTR
		pd.FYs[k].Tax = pd.FYs[k].Tax + sumAPTR
		storeFy.PropertyID = pd.Property.UniqueID
		storeFy.FinancialYearId = v.FinancialYear.UniqueID
		storeFy.Demand.FLTax = pd.FYs[k].Tax
		storeFy.Demand.VLTax = pd.FYs[k].VacantLandTax
		storeFy.Demand.Total = pd.FYs[k].Tax + pd.FYs[k].VacantLandTax

		storeFys = append(storeFys, storeFy)
	}
	fmt.Println("storeFys fyslen", len(storeFys))

	spd := new(StoredPropertyDemand)
	spd.Fys = storeFys

	return spd
}
