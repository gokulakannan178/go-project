package models

type ULBNearBy struct {
	Longitude         float64  `json:"longitude" bson:"longitude,omitempty"`
	Latitude          float64  `json:"latitude" bson:"latitude,omitempty"`
	KM                float64  `json:"km" bson:"km,omitempty"`
	CertificateStatus []string `json:"certificateStatus" bson:"certificateStatus,omitempty"`
}

type ULBNearByResponse struct {
	Name    string  `json:"name" bson:"name,omitempty"`
	Address Address `json:"address" bson:"address,omitempty"`
	Ref     struct {
		Inventory ULBInventory `json:"inventory" bson:"inventory,omitempty"`
		Address   *RefAddress  `json:"address" bson:"address,omitempty"`
	} `json:"ref" bson:"ref,omitempty"`
}

type CompostInStock struct {
	Quantity float64 `json:"quantity" bson:"quantity,omitempty"`
}
