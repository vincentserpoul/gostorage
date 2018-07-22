package redis

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

// Exists return true if key is present
func Exists(rp *redis.Pool, key string) (bool, error) {
	red := rp.Get()
	defer func() {
		if err := red.Close(); err != nil {
			log.Fatalf("Exists(%v) Close: %v", key, err)
		}
	}()

	rawBytes, err := red.Do("GET", key)
	if err != nil {
		return false, fmt.Errorf("Exists GET(%s): %v", key, err)
	}
	if rawBytes == nil {
		return false, fmt.Errorf("Exists content(%s): %v", key, err)
	}

	return true, nil
}
