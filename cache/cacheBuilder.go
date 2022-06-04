package cache

type RateLimitCacheType int8

const (
	RedisBackend RateLimitCacheType = iota
)

// GetCacheHandler cache builder ( factory method )
func GetCacheHandler(name RateLimitCacheType) ICache {
	switch name {
	case RedisBackend:
		return Redis{"127.0.0.1", 0, 6379}
	default:
		// set this temp
		return Redis{"127.0.0.1", 0, 6379}
	}
}
