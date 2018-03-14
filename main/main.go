package main

import (
	"net/http"
	"fmt"
	"github.com/go-redis/redis"
	"../handlers"
	"../utils"
	"../global"
)

func RedisFactory() (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:       utils.RedisServer,
		Password:   utils.RedisPwd,
		DB:         utils.DB,
	})
	_, err = client.Ping().Result()
	return client, err
}

func main()  {
	fmt.Printf("redis connectting ...")
	client, redisClientErr := RedisFactory()
	if redisClientErr != nil {
		fmt.Printf("redis connect failed: %s", client)
		return
	}
	global.RedisClient = client
	fmt.Printf("Server is running ...")
	http.HandleFunc("/lean/home", handlers.HomeHandler)
	httpServerErr := http.ListenAndServe("127.0.0.1:8908", nil)
	if httpServerErr != nil {
		fmt.Printf("error is: %s", httpServerErr)
	}
}
