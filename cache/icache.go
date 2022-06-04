package cache

import "time"

type ICache interface {
	AddCounter(ip string, duration time.Duration)
	getCounter(ip string) int32
	IsAllowed(ip string, requestNumber int) bool
}
