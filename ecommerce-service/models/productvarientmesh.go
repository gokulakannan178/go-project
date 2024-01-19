package models

type ProductVariantMesh struct {
	UniqueID         string `json:"uniqueId" bson:"uniqueId,omitempty"`
	ProductVariantID string `json:"productVariantID" bson:"productVariantID,omitempty"`
	ProductID        string `json:"productId" bson:"productId,omitempty"`
	VariantTypeID    string `json:"variantTypeID" bson:"variantTypeID,omitempty"`
	VariantTypeName  string `json:"variantTypeName" bson:"variantTypeName,omitempty"`
	Value            string `json:"value" bson:"value,omitempty"`
	Status           string `json:"status" bson:"status,omitempty"`
}

type ProductVariantMeshFilter struct {
	Status []string `json:"status" bson:"status,omitempty"`
}

type RefProductVariantMesh struct {
	ProductVariantMesh `bson:",inline"`
}
