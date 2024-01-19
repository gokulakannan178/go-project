package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TCLedger : ""
type TCLedger struct {
	ID             primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Date           *time.Time         `json:"date"  bson:"date,omitempty"`
	DateStr        string             `json:"dateStr"  bson:"dateStr,omitempty"`
	OpeningBalance float64
	ClosingBalance float64
	CashInHand     float64
	Record         []struct {
		Type   string
		Amount float64
		Txns   struct {
			Txn    string
			Amount float64
		}
	}
}
