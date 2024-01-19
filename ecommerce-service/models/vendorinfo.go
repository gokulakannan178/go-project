package models

type VendorInfo struct {
	UniqueID     string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	VendorID     string     `json:"vendorId" bson:"vendorId,omitempty"`
	GSTNo        string     `json:"gstNo" bson:"gstNo,omitempty"`
	PanNo        string     `json:"panNo" bson:"panNo,omitempty"`
	TaxNo        string     `json:"taxNo" bson:"taxNo,omitempty"`
	AadhaarNo    string     `json:"aadhaarNo" bson:"aadhaarNo,omitempty"`
	AadhaarProof string     `json:"aadhaarProof" bson:"aadhaarProof,omitempty"`
	Status       string     `json:"status" bson:"status,omitempty"`
	Created      *CreatedV2 `json:"created" bson:"created,omitempty"`
}

type VendorInfoFilter struct {
	UniqueID       []string      `json:"uniqueId" bson:"uniqueId,omitempty"`
	VendorID       []string      `json:"vendorId" bson:"vendorId,omitempty"`
	Status         []string      `json:"status" bson:"status,omitempty"`
	Address        AddressSearch `json:"address" bson:"address,omitempty"`
	WantVendorInfo bool          `json:"wantVendorInfo" bson:"wantVendorInfo,omitempty"`
	SearchText     struct {
		UniqueID  string `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
		VendorID  string `json:"vendorId,omitempty" bson:"vendorId,omitempty"`
		GSTNo     string `json:"gstNo" bson:"gstNo,omitempty"`
		PanNo     string `json:"panNo" bson:"panNo,omitempty"`
		TaxNo     string `json:"taxNo" bson:"taxNo,omitempty"`
		AadhaarNo string `json:"aadhaarNo" bson:"aadhaarNo,omitempty"`
	} `json:"searchText"`
}

type RefVendorInfo struct {
	VendorInfo `bson:",inline"`
	Ref        struct {
		Vendor Vendor `json:"vendor" bson:"vendor,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
