package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"hrms-services/models"

	"time"
)

func (s *Service) SetValueCacheMemory(key string, value interface{}, Duration time.Duration) error {
	//cachememory := new(models.CacheMemory)
	//t := time.Now()
	timein := time.Now().Local().Add(time.Second * time.Duration(Duration))
	bye, err := json.Marshal(value)
	if err != nil {
		return err
	}
	fmt.Println("key====>", key)
	fmt.Println("bye====>", bye)
	s.Cache.Mu.Lock()
	defer s.Cache.Mu.Unlock()
	var data models.CacheData
	data.Data = bye
	data.ExpireTime = &timein
	s.Cache.Data[key] = data
	fmt.Println("counters====>", s.Cache)

	return nil

}
func (s *Service) GetValueCacheMemory(key string, value interface{}) error {
	data := s.Cache.Data[key]
	// s.Cache.ExpireTime
	if &data == nil {
		return errors.New("data Invaild")
	}
	t := time.Now()
	if t.After(*data.ExpireTime) {
		delete(s.Cache.Data, key)
		return errors.New("Expired")
	}
	err := json.Unmarshal(data.Data, &value)
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) CacheGC() {
	for k, v := range s.Cache.Data {
		t := time.Now()
		if t.After(*v.ExpireTime) {
			delete(s.Cache.Data, k)
		}
	}
}
