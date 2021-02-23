package cache

import (
	"MacroBuffalo/conf"
	_ "encoding/json"
	"gopkg.in/redis.v4"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

// Wraps the Redis client to meet the Cache interface.
type RedisClient struct {
	pool *redis.Client
}

// until redigo supports sharding/clustering, only one host will be in hostList

func Config() *conf.Yaml {
	confInfo := new(conf.Yaml)
	var yamlFile, err = ioutil.ReadFile("../../conf/resources.yaml")
	if err != nil {
		log.Printf("resources file get err %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, confInfo)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return confInfo
}

func NewRedisCache() *RedisClient {
	redisConfig := Config().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         redisConfig.Addr,
		Password:     redisConfig.Password,
		MaxRetries:   redisConfig.MaxRetries,
		DialTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     redisConfig.PoolSize,
		PoolTimeout:  0,
		IdleTimeout:  0,
		DB:           redisConfig.DB,
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)

	c := &RedisClient{client}
	return c

}

/* Queue Function */
func (c *RedisClient) LPush(key string, value string) {
	cmd := c.pool.LPush(key, value)
	log.Println(cmd)
}

func (c *RedisClient) BRpop(key string) []string {
	data := c.pool.BRPop(0, key)
	return data.Val()
}
