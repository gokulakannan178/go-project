package models

//Address : ""
type Address struct {
	No                string `json:"no" bson:"no,omitempty"`
	StateCode         string `json:"stateCode" bson:"stateCode,omitempty"`
	PinCode           string `json:"pinCode" bson:"pinCode"`
	DistrictCode      string `json:"districtCode" bson:"districtCode,omitempty"`
	BlockCode         string `json:"blockCode" bson:"blockCode,omitempty"`
	GramPanchayatCode string `json:"gramPanchayatCode" bson:"gramPanchayatCode,omitempty"`
	VillageCode       string `json:"villageCode" bson:"villageCode,omitempty"`
	// ZoneCode     string   `json:"zoneCode" bson:"zoneCode,omitempty"`
	// WardCode     string   `json:"wardCode" bson:"wardCode,omitempty"`
	AL1        string   `json:"al1" bson:"al1,omitempty"`
	Al2        string   `json:"al2" bson:"al2,omitempty"`
	PlotNo     string   `json:"plotNo" bson:"plotNo,omitempty"`
	KhataNo    string   `json:"khataNo" bson:"khataNo,omitempty"`
	PostalCode string   `json:"postalCode" bson:"postalCode,omitempty"`
	Landmark   string   `json:"landmark" bson:"landmark,omitempty"`
	Location   Location `json:"location" bson:"location,omitempty"`
}

//AddressV4 : ""
type AddressV4 struct {
	Name              string   `json:"name" bson:"name,omitempty"`
	Line1             string   `json:"line1" bson:"line1,omitempty"`
	Line2             string   `json:"line2" bson:"line2,omitempty"`
	Zipcode           string   `json:"postalcode" bson:"postalcode,omitempty"`
	StateCode         string   `json:"stateCode" bson:"stateCode,omitempty"`
	DistrictCode      string   `json:"districtCode" bson:"districtCode,omitempty"`
	BlockCode         string   `json:"blockCode" bson:"blockCode,omitempty"`
	GramPanchayatCode string   `json:"gramPanchayatCode" bson:"gramPanchayatCode,omitempty"`
	VillageCode       string   `json:"villageCode" bson:"villageCode,omitempty"`
	Country           string   `json:"country" bson:"country,omitempty"`
	Location          Location `json:"location" bson:"location,omitempty"`
	Type              string   `json:"type" bson:"type,omitempty"`
	DoorNo            string   `json:"doorNo" bson:"doorNo,omitempty"`
	Landmark          string   `json:"landmark" bson:"landmark,omitempty"`
	Ph                string   `json:"ph" bson:"ph,omitempty"`
	Fax               string   `json:"fax" bson:"fax,omitempty"`
}

//Location : ""
type Location struct {
	Type        string    `json:"type" bson:"type,omitempty"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates,omitempty"`
}

//RefAddress :""
type RefAddress struct {
	State         *State         `json:"state" bson:"state,omitempty"`
	District      *District      `json:"district" bson:"district,omitempty"`
	Village       *Village       `json:"village" bson:"village,omitempty"`
	Block         *Block         `json:"block" bson:"block,omitempty"`
	GramPanchayat *GramPanchayat `json:"gramPanchayat" bson:"gramPanchayat,omitempty"`
}

//AddressSearch : ""
type AddressSearch struct {
	StateCode         []string `json:"stateCode" bson:"stateCode,omitempty"`
	DistrictCode      []string `json:"districtCode" bson:"districtCode,omitempty"`
	VillageCode       []string `json:"villageCode" bson:"villageCode,omitempty"`
	BlockCode         []string `json:"blockCode" bson:"blockCode,omitempty"`
	GramPanchayatCode []string `json:"gramPanchayatCode" bson:"gramPanchayatCode,omitempty"`
	//ZoneCode     []string `json:"zoneCode" bson:"zoneCode,omitempty"`
	//WardCode     []string `json:"wardCode" bson:"wardCode,omitempty"`
	Country  []string `json:"country" bson:"country,omitempty"`
	Location Location `json:"location" bson:"location,omitempty"`
	Type     string   `json:"type" bson:"type,omitempty"`
}
