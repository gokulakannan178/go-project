package models

type Product struct {
	UniqueID      string   `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name          string   `json:"name" bson:"name,omitempty"`
	Collection    string   `json:"collection" bson:"collection,omitempty"`
	Tags          string   `json:"tags" bson:"tags,omitempty"`
	Desc          string   `json:"desc" bson:"desc,omitempty"`
	Status        string   `json:"status" bson:"status,omitempty"`
	Image         []string `json:"image" bson:"image,omitempty"`
	CategoryID    string   `json:"categoryId" bson:"categoryId,omitempty"`
	SubCategoryID string   `json:"subCategoryId" bson:"subCategoryId,omitempty"`
}

type ProductFilter struct {
	Status        []string `json:"status" bson:"status,omitempty"`
	CategoryID    []string `json:"categoryId" bson:"categoryId,omitempty"`
	SubCategoryID []string `json:"subCategoryId" bson:"subCategoryId,omitempty"`
	SearchText    struct {
		ProductName string `json:"productName" bson:"productName,omitempty"`
		UniqueID    string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	} `json:"searchText"`
}

type RefProduct struct {
	Product `bson:",inline"`
	Ref     struct {
		Category    RefCategory    `json:"category,omitempty" bson:"category,omitempty"`
		SubCategory RefSubCategory `json:"subCategory,omitempty" bson:"subCategory,omitempty"`
		Inventory   []RefInventory `json:"inventory,omitempty" bson:"inventory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
