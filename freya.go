package freya

import (
	"github.com/coci/freya/cache"
	"net/http"
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

// HandlerMuxMiddleware wrap http mux handler ( middleware )
func (f *Freya) HandlerMuxMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		// user ip address
		ip := getIPAddress(request)

		// if user didn't hit rate
		if f.Cache.IsAllowed(ip, int(f.requestNumber)) {

			// incr user rate
			f.Cache.AddCounter(ip, f.Duration)

			// serve func
			next.ServeHTTP(w, request)

		} else {
			// throw an error
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

	})
}
