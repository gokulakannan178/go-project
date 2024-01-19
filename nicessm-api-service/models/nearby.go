package models

type NearBy struct {
	Longitude         float64  `json:"longitude" bson:"longitude,omitempty"`
	Latitude          float64  `json:"latitude" bson:"latitude,omitempty"`
	KM                float64  `json:"km" bson:"km,omitempty"`
	CertificateStatus []string `json:"certificateStatus" bson:"certificateStatus,omitempty"`
}
