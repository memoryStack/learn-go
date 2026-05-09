package cache

import (
	"errors"

	"lru-cache/cache/strategies"
)

type CacheType int

type Cache interface {
	Get(key string) (int, error)
	Put(key string, value int) error
	Print()
}

const (
	LRU CacheType = iota
	LFU
	FIFO
)

func New(size int, _type CacheType) (Cache, error) {
	switch _type {
		case LRU:
			return strategies.NewLRUCache(size), nil
		// case LFU:
		// 	return &NewLFUCache(size)
		// case FIFO:
		// 	return &NewFIFOCache(size)
		default:
			return nil, errors.New("invalid cache type")
	}
}
