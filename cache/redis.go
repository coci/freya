package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// Redis struct that implement ICache interface
type Redis struct {
	Host string
	Db   int8
	Port int16
}

var ctx = context.Background()

// redis client
func (r Redis) getRedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(int(r.Port)),
		Password: "",        // no password set
		DB:       int(r.Db), // use default DB
	})

	return client
}

// AddCounter create or increment user rate
func (r Redis) AddCounter(ip string, duration time.Duration) {
	// redis conn
	redisConn := r.getRedisConnection()

	// get rate via ip
	val := redisConn.Get(ctx, ip)

	// check existence of rate key
	if val.Val() == "" {
		// set rate key to 1
		redisConn.Set(ctx, ip, 1, duration)
	} else {
		// incr rate key
		redisConn.Incr(ctx, ip)
	}
}

// IsAllowed check user rate
func (r Redis) IsAllowed(ip string, requestNumber int) bool {
	// redis conn
	redisConn := r.getRedisConnection()

	// rate key via ip
	val := redisConn.Get(ctx, ip)

	// check existence of rate key
	if val.Val() != "" {
		// get amount of user rate via ip
		userRate, _ := val.Int()
		if userRate >= requestNumber {
			return false
		} else {
			return true
		}
	} else {
		// if key not exist so it's allowed
		return true
	}
}
