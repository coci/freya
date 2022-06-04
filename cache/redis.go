package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type Redis struct {
	Host string
	Db   int8
	Port int16
}

var ctx = context.Background()

func (r Redis) getRedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(int(r.Port)),
		Password: "",        // no password set
		DB:       int(r.Db), // use default DB
	})

	return client
}

func (r Redis) AddCounter(ip string, duration time.Duration) {
	// incr rate key
}

func (r Redis) IsAllowed(ip string, requestNumber int) bool {
	// check allowed rate key
	return true
}
