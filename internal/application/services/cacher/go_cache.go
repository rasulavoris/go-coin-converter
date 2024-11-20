package cacher

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type GoCache struct{}

func DefineCache(ttlInSeconds int, clearInSeconds int) *cache.Cache {
	return cache.New(time.Duration(ttlInSeconds)*time.Second, time.Duration(clearInSeconds)*time.Second)
}
