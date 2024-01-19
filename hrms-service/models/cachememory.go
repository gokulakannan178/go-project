package models

import (
	"sync"
	"time"
)

type CacheMemory struct {
	Mu         sync.Mutex           `json:"mu"  bson:"mu,omitempty"`
	Data       map[string]CacheData `json:"data"  bson:"data,omitempty"`
	ExpireTime *time.Time           `json:"expireTime"  bson:"expireTime,omitempty"`
}
type CacheData struct {
	Data       []byte
	ExpireTime *time.Time `json:"expireTime"  bson:"expireTime,omitempty"`
}
type Otp struct {
	Otp string `json:"otp"  bson:"otp,omitempty"`
}
