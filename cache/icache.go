package cache

import "time"

type ICache interface {
	AddCounter(ip string, duration time.Duration)
	IsAllowed(ip string, requestNumber int) bool
}
