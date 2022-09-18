package utils

import "github.com/dgraph-io/ristretto"

func NewRCache() (cache *ristretto.Cache, err error) {
	cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	return cache, err
}
