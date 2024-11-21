package cacher

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gocoinconverter/cmd/config"
	"gopkg.in/yaml.v3"
	"os"
)

func InitRedisConnection() *redis.Client {
	actualConfig := readConfiguration()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", actualConfig.Redis.Host, actualConfig.Redis.Port),
		Password: actualConfig.Redis.Password,
		DB:       actualConfig.Redis.DB,
	})

	fmt.Println("Redis started")
	return rdb
}

func readConfiguration() config.Config {
	file, err := os.Open("/Users/rasulabdulaev/Documents/personal/go-coin-converter/cmd/config/config.yaml")
	if err != nil {
		fmt.Printf("Error opening configuration file: %s\n", err)
	}
	defer file.Close()

	var actualConfig config.Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&actualConfig); err != nil {
		fmt.Printf("YAML file can not be decoded: %s\n", err)
	}

	return actualConfig
}
