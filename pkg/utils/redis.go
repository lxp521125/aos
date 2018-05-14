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

var RedisHandle = initRedisClient()

func initRedisClient() *RedisClient {
	var redisStuct = new(RedisClient)
	redisStuct.RedisClientHandle = getClient()
	return redisStuct
}

func getClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     consul.GdConsul["PUBLIC_REDIS_HOST"] + ":" + consul.GdConsul["PUBLIC_REDIS_PORT"],
		Password: consul.GdConsul["PUBLIC_REDIS_PASSWD"], // no password set
		DB:       0,                                      // use default DB
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

/*

client := redis.NewClient(&redis.Options{
	Addr:     "dev.redis.gaodunwangxiao.com:6379",
	Password: "gaodun.com", // no password set
	DB:       0,            // use default DB
})

pong, err := client.Ping().Result()
fmt.Println("client pong：", pong, err)

err = client.Set("key", "valueStone", 0).Err()
if err != nil {
	panic(err)
}

val, err := client.Get("key").Result()
if err != nil {
	panic(err)
}
fmt.Println("key", val)
*/
