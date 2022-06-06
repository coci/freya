package freya

import (
	"net/http"
	"sync"
	"time"

	"github.com/coci/freya/cache"
)

type Freya struct {
	requestNumber int16
	Cache         cache.ICache
	Duration      time.Duration
}

var lock = &sync.Mutex{}

// hold single instance of Freya
var singleFreyaInstance *Freya = nil

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

// HandlerMuxMiddleware wrap http handler func ( middleware )
func (f *Freya) HandlerFuncMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		ip := getIPAddress(request)

		if f.Cache.IsAllowed(ip, int(f.requestNumber)) {
			f.Cache.AddCounter(ip, f.Duration)

			next.ServeHTTP(w, request)

		} else {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

	})
}

// NewLimiter create or return Freya instance ( using singleton )
func NewLimiter(request int16, duration time.Duration, cacheType cache.RateLimitCacheType) *Freya {
	if singleFreyaInstance != nil {

		return singleFreyaInstance

	} else {

		lock.Lock()
		defer lock.Unlock()

		// get proper cache handler with user input
		cacheType := cache.GetCacheHandler(cacheType)

		singleFreyaInstance = &Freya{requestNumber: request, Cache: cacheType, Duration: duration}

		return singleFreyaInstance
	}
}
