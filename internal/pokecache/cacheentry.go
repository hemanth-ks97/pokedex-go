package pokecache

import "time"

type cacheentry struct {
	created_at time.Time
	val        []byte
}
