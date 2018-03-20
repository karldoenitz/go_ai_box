package main

import (
	"net/http"
	"github.com/go-redis/redis"
	"../handlers"
	"../utils"
	"../global"
	"../logger"
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
	logger.Info.Printf("redis connectting ...\n")
	client, redisClientErr := RedisFactory()
	if redisClientErr != nil {
		logger.Error.Printf("redis connect failed: %s", client)
		return
	}
	global.RedisClient = client
	logger.Info.Printf("Server is running ...\n")
	// url配置
	http.HandleFunc("/lean/home", handlers.HomeHandler)
	http.HandleFunc("/lean/detail", handlers.DetailHandler)
	// 端口监听
	httpServerErr := http.ListenAndServe(utils.ServerIPPort, nil)
	if httpServerErr != nil {
		logger.Error.Printf("error is: %s", httpServerErr)
	}
}
