package redis

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

// SetKeyWithExpiration will simply set a key that will expire
// after exp millisecond
func SetKeyWithExpiration(rp *redis.Pool, key string, exp int64) error {
	red := rp.Get()
	defer func() {
		if err := red.Close(); err != nil {
			log.Fatalf("SetKeyWithExpiration(%s, %d) Close: %v", key, exp, err)
		}
	}()

	if _, err := red.Do("PSETEX", key, exp, ""); err != nil {
		return fmt.Errorf(
			"SetKeyWithExpiration(%s, %d) Do: %v",
			key, exp, err,
		)
	}
	return nil
}
