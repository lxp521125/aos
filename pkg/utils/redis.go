package utils

import (
	"aos/pkg/consul"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	RedisClientHandle *redis.Client
}

func GetRedisHandle() *RedisClient {
	//  RedisHandle = initRedisClient()
	return initRedisClient()
}

func initRedisClient() *RedisClient {
	var redisStuct = new(RedisClient)
	redisStuct.RedisClientHandle = getClient()
	return redisStuct
}

func getClient() *redis.Client {
	consulInfo, _ := consul.InitConfig()
	fmt.Println("redis.go")
	client := redis.NewClient(&redis.Options{
		Addr:     consulInfo["PUBLIC_REDIS_HOST"] + ":" + consulInfo["PUBLIC_REDIS_PORT"],
		Password: consulInfo["PUBLIC_REDIS_PASSWD"], // no password set
		DB:       0,                                 // use default DB
	})
	pong, _ := client.Ping().Result()

	fmt.Println(pong)
	return client
}

func (rc *RedisClient) SetData(key string, value interface{}, expiration time.Duration) {
	err := rc.RedisClientHandle.Set(key, value, expiration).Err()
	if err != nil {
		// panic(err)
	}
}

func (rc *RedisClient) GetData(key string) interface{} {

	val, err := rc.RedisClientHandle.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		// panic(err)
	}
	return val
}
