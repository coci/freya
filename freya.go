package freya

import (
	"github.com/coci/freya/cache"
	"sync"
	"time"
)

type Freya struct {
	requestNumber int16
	Cache         cache.ICache
	Duration      time.Duration
}

var lock = &sync.Mutex{}

// hold single instance of Freya
var singleFreyaInstance *Freya

// NewLimiter create or return Freya instance ( using singleton )
func NewLimiter(request int16, duration time.Duration, cacheType string) *Freya {
	if singleFreyaInstance != nil {
		return singleFreyaInstance
	} else {

		lock.Lock()
		defer lock.Unlock()

		// get proper cache handler with user input
		cacheType := cache.GetCacheHandler(cacheType)

		singleFreyaInstance := &Freya{requestNumber: request, Cache: cacheType, Duration: duration}

		return singleFreyaInstance
	}
}
