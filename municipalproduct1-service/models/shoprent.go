package models

import "time"

//ShopRent : ""
type ShopRent struct {
	UniqueID             string                   `json:"uniqueId" bson:"uniqueId,omitempty"`
	ShopCategoryID       string                   `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
	ShopSubCategoryID    string                   `json:"shopSubCategoryId" bson:"shopSubCategoryId,omitempty"`
	Sqft                 float64                  `json:"sqft" bson:"sqft,omitempty"`
	ShopRentArea         float64                  `json:"shopRentArea" bson:"shopRentArea,omitempty"`
	DateFrom             *time.Time               `json:"dateFrom" bson:"dateFrom,omitempty"`
	DateTo               *time.Time               `json:"dateTo" bson:"dateTo,omitempty"`
	Status               string                   `json:"status" bson:"status,omitempty"`
	IndividualRateMaster string                   `json:"individualRateMaster" bson:"individualRateMaster,omitempty"`
	Created              *Created                 `json:"created,omitempty"  bson:"created,omitempty"`
	Updated              []Updated                `json:"updated,omitempty"  bson:"updated,omitempty"`
	RentAmount           float64                  `json:"rentAmount,omitempty"  bson:"rentAmount,omitempty"`
	MobileNo             string                   `json:"mobileNo" bson:"mobileNo,omitempty"`
	OwnerName            string                   `json:"ownerName" bson:"ownerName,omitempty"`
	GuardianName         string                   `json:"guardianName" bson:"guardianName,omitempty"`
	Address              Address                  `json:"address" bson:"address,omitempty"`
	Demand               ShopRentTotalDemand      `json:"demand" bson:"demand,omitempty"`
	Collections          ShopRentTotalCollection  `json:"collection" bson:"collection,omitempty"`
	PendingCollections   ShopRentTotalCollection  `json:"pendingCollections" bson:"pendingCollections,omitempty"`
	OutStanding          ShopRentTotalOutStanding `json:"outstanding" bson:"outstanding,omitempty"`
}

//ShopRentFilter : ""
type ShopRentFilter struct {
	Address           *AddressSearch `json:"address"`
	UniqueID          []string       `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	ShopCategoryID    []string       `json:"shopCategoryId" bson:"shopCategoryId,omitempty"`
	ShopSubCategoryID []string       `json:"shopSubCategoryId" bson:"shopSubCategoryId,omitempty"`
	Status            []string       `json:"status,omitempty" bson:"status,omitempty"`
	SearchText        struct {
		MobileNo     string `json:"mobileNo" bson:"mobileNo,omitempty"`
		OwnerName    string `json:"ownerName" bson:"ownerName,omitempty"`
		UniqueID     string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
		GuardianName string `json:"guardianName" bson:"guardianName,omitempty"`
	} `json:"searchText"`
	FromDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"fromDateRange"`
	ToDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"toDateRange"`
}

//RefShopRent : ""
type RefShopRent struct {
	ShopRent `bson:",inline"`
	Ref      struct {
		Address                 RefAddress              `json:"address" bson:"address,omitempty"`
		ShopRentShopCategory    ShopRentShopCategory    `json:"shopRentShopCategory,omitempty" bson:"shopRentShopCategory,omitempty"`
		ShopRentShopSubCategory ShopRentShopSubCategory `json:"shopRentShopSubCategory,omitempty" bson:"shopRentShopSubCategory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

func (res *RefShopRent) Inc(a int) int {
	return a + 1
}

type ShopRentOverallDemandReport struct {
	ShopRents []RefShopRent    `json:"shoprents" bson:"shoprents,omitempty"`
	CFY       RefFinancialYear `json:"cfy" bson:"cfy,omitempty"`
}

type ShopRentTotalDemand struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	FromYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}

type ShopRentTotalCollection struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	FromYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}

type ShopRentTotalOutStanding struct {
	Current struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"current" bson:"current,omitempty"`
	Arrear struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"arrear" bson:"arrear,omitempty"`
	Total struct {
		Tax     float64 `json:"tax" bson:"tax,omitempty"`
		Penalty float64 `json:"penalty" bson:"penalty,omitempty"`
		Other   float64 `json:"other" bson:"other,omitempty"`
		Total   float64 `json:"total" bson:"total,omitempty"`
	} `json:"total" bson:"total,omitempty"`
	FromYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"fromYear" bson:"fromYear"`
	ToYear struct {
		Name string `json:"name" bson:"name"`
		ID   string `json:"id" bson:"id"`
	} `json:"toYear" bson:"toYear"`
}
