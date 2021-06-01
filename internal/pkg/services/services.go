package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/rahulvramesh/ocha-url/internal/app/ocha/models"
	"github.com/rahulvramesh/ocha-url/internal/pkg/bigcache"
	"time"
)



// OchaService - Service  Function Encapsulation
type OchaService struct {
	Data models.Item
}

// CheckIfKeyExists - check if the key already added to the DS
func (o *OchaService) CheckIfKeyExists(key string) error {
	_, err := bigcache.GetDS().Get(key)
	if err != nil {
		return errors.New("the the desired shortcode is already in use")
	}

	return nil
}

// GetByKey - Check if the provided key already exists or not, if exists return the data
func (o *OchaService) GetByKey(key string) error {
	var (
		err error
	)
	itemObj, err := bigcache.GetDS().Get(key)
	if err != nil {
		return err
	}
	// Byte To Struct
	if err := json.Unmarshal(itemObj, &o.Data); err != nil {
		return err
	}

	// Update Last Seen & Counter
	err = o.CounterAndLastSeen(key)
	if err != nil {
		return err
	}

	return nil
}

// StoreLink - Store Key Value To Cache
func (o *OchaService) StoreLink(key, url string) error {

	// Create Item
	item := models.Item{
		URL:     url,
		Key:     key,
		StartDate: time.Now(),
		Counter: 0,
	}

	// Convert Struct to byte type for storage
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(item)
	if err != nil {
		return err
	}

	// Store to Big cache
	err = bigcache.GetDS().Set(key, reqBodyBytes.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// CounterAndLastSeen - function to increment counter every time there is a request
func (o *OchaService) CounterAndLastSeen(key string) error{

	o.Data.Counter++
	o.Data.LastSeenDate = time.Now()

	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(o.Data)
	if err != nil {
		return err
	}

	// Store to Big cache
	err = bigcache.GetDS().Set(key, reqBodyBytes.Bytes())
	if err != nil {
		return err
	}

	return nil
}

