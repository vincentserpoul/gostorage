package redis

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

// Consume will check if a key exists and if it exists, will destroy it
func Consume(rp *redis.Pool, key string) error {
	red := rp.Get()
	defer func() {
		if err := red.Close(); err != nil {
			log.Fatalf("Exists(%v) Close: %v", key, err)
		}
	}()

	rawBytes, err := red.Do("GET", key)
	if err != nil {
		return fmt.Errorf("Exists GET(%s): %v", key, err)
	}
	if rawBytes == nil {
		return fmt.Errorf("Exists content(%s): %v", key, err)
	}

	if _, err := red.Do("DEL", key); err != nil {
		return fmt.Errorf("Exists DEL(%s): %v", key, err)
	}

	return nil
}
