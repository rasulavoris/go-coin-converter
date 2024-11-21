package cacher

import (
	"context"
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type CacheService struct {
	goCache    *cache.Cache
	redisCache *redis.Client
	context    context.Context
}

func NewCacheService(goCache *cache.Cache, redisCache *redis.Client) *CacheService {
	return &CacheService{
		goCache:    goCache,
		redisCache: redisCache,
		context:    context.Background(),
	}
}

func (s *CacheService) SaveData(key string, value interface{}) {
	s.goCache.Set(key, value, cache.DefaultExpiration)
}

func (s *CacheService) GetData(key string) (interface{}, bool) {
	return s.goCache.Get(key)
}

func (s *CacheService) SaveIntoRedis(key string, value interface{}, ttlInSeconds int) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("Error serializing model: %v", err)
	}

	err = s.redisCache.Set(s.context, key, jsonData, time.Duration(ttlInSeconds)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}
