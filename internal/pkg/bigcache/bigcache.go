package bigcache

import (
	"github.com/allegro/bigcache/v3"
	"log"
	"time"
)

var DS *bigcache.BigCache
var err error

// InitializeDataSource - global initialization of connection
func InitializeDataSource(){
	DS, err = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err !=nil {
		log.Fatal(err)
	}
}

// GetDS - Get shared connection
func GetDS() *bigcache.BigCache{
	return DS
}