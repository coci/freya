package cache

import (
	"time"
)

type Redis struct {
	Host string
	Db   int8
	Port int16
}

func (r Redis) AddCounter(ip string, duration time.Duration) {
	// incr rate key
}

func (r Redis) IsAllowed(ip string, requestNumber int) bool {
	// check allowed rate key
	return true
}
