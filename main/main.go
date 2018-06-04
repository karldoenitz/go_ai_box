package main

import (
	"net/http"
	"github.com/go-redis/redis"
	"github.com/rtt/Go-Solr"
	"../handlers"
	"../utils"
	"../global"
	"../logger"
)

// 初始化redis数据工厂
// 失败则返回的err不为nil
func RedisFactory() (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:       utils.RedisServer,
		Password:   utils.RedisPwd,
		DB:         utils.DB,
	})
	_, err = client.Ping().Result()
	return client, err
}

// 初始化solr数据工厂
// 若失败则返回err
func SolrFactory() (err error) {
	musicCore, musicErr := solr.Init(utils.SolrHost, utils.SolrPort, utils.SolrMusicCore)
	if musicErr != nil {
		return musicErr
	}
	global.SolrClientMusic = musicCore
	playbillCore, playbillErr := solr.Init(utils.SolrHost, utils.SolrPort, utils.SolrPlaybillCore)
	if playbillErr != nil {
		return playbillErr
	}
	global.SolrClientPlaybill = playbillCore
	singerCore, singerCoreErr := solr.Init(utils.SolrHost, utils.SolrPort, utils.SolrSingerCore)
	if singerCoreErr != nil {
		return singerCoreErr
	}
	global.SolrClientSinger = singerCore
	return nil
}

func main()  {
	logger.Info.Printf("redis connectting ...\n")
	// redis链接初始化
	client, redisClientErr := RedisFactory()
	if redisClientErr != nil {
		logger.Error.Printf("redis connect failed: %s", redisClientErr.Error())
		return
	}
	global.RedisClient = client
	// solr链接初始化
	solrErr := SolrFactory()
	if solrErr != nil {
		logger.Error.Printf("solr connect failed: %s", solrErr.Error())
		return
	}
	logger.Info.Printf("Server is running ...\n")
	// url配置
	http.HandleFunc("/lean/home", handlers.HomeHandler)
	http.HandleFunc("/lean/detail", handlers.DetailHandler)
	http.HandleFunc("/service/search", handlers.SearchHandler)
	http.HandleFunc("/source/music", handlers.MusicHandler)
	http.HandleFunc("/source/singer", handlers.SingerHandler)
	http.HandleFunc("/source/playbill", handlers.PlaybillHandler)
	// 端口监听
	httpServerErr := http.ListenAndServe(utils.ServerIPPort, nil)
	if httpServerErr != nil {
		logger.Error.Printf("error is: %s", httpServerErr)
	}
}
