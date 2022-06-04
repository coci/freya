package cache

// GetCacheHandler cache builder ( factory method )
func GetCacheHandler(name string) ICache {
	switch name {
	case "redis":
		return Redis{"127.0.0.1", 0, 6379}
	default:
		// set this temp
		return Redis{"127.0.0.1", 0, 6379}
	}
}
