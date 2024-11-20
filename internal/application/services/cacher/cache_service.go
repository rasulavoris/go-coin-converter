package cacher

import (
	"github.com/patrickmn/go-cache"
)

type CacheService struct {
	cache *cache.Cache
}

func NewCacheService(cache *cache.Cache) *CacheService {
	return &CacheService{cache: cache}
}

func (s *CacheService) SaveData(key string, value interface{}) {
	s.cache.Set(key, value, cache.DefaultExpiration)
}

func (s *CacheService) GetData(key string) (interface{}, bool) {
	return s.cache.Get(key)
}
