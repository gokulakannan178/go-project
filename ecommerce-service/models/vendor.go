package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vendor struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UniqueID     string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	StoreName    string             `json:"storeName" bson:"storeName,omitempty"`
	EmailID      string             `json:"emailId" bson:"emailId,omitempty"`
	GSTNo        string             `json:"gstNo" bson:"gstNo,omitempty"`
	PANNo        string             `json:"panNo" bson:"panNo,omitempty"`
	TaxNo        string             `json:"taxNo" bson:"taxNo,omitempty"`
	Location     Location           `json:"location" bson:"location,omitempty"`
	Address      Address            `json:"address" bson:"address,omitempty"`
	Logo         string             `json:"logo" bson:"logo,omitempty"`
	AddressProof string             `json:"addressProof" bson:"addressProof,omitempty"`
	AadhaarProof string             `json:"aadhaarProof" bson:"aadhaarProof,omitempty"`
	GSTProof     string             `json:"gstProof" bson:"gstProof,omitempty"`
	MobileNo     string             `json:"mobileNo" bson:"mobileNo,omitempty"`
	Status       string             `json:"status" bson:"status,omitempty"`
	Created      *CreatedV2         `json:"created" bson:"created,omitempty"`
	Token        string             `json:"token" bson:"-"`
}

type VendorFilter struct {
	UniqueID       []string      `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status         []string      `json:"status" bson:"status,omitempty"`
	Address        AddressSearch `json:"address" bson:"address,omitempty"`
	WantVendorInfo bool          `json:"wantVendorInfo" bson:"wantVendorInfo,omitempty"`
	SearchText     struct {
		UniqueID  string `json:"uniqueId" bson:"uniqueId,omitempty"`
		MobileNo  string `json:"mobileNo" bson:"mobileNo,omitempty"`
		Name      string `json:"name" bson:"name,omitempty"`
		StoreName string `json:"storeName" bson:"storeName,omitempty"`
	} `json:"searchText"`
}

type RefVendor struct {
	Vendor `bson:",inline"`
	Ref    struct {
		VendorInfo *VendorInfo `json:"vendorInfo" bson:"vendorInfo,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
