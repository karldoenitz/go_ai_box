package global

import (
	"github.com/go-redis/redis"
	"github.com/rtt/Go-Solr"
)

var RedisClient *redis.Client
var SolrClientMusic *solr.Connection
var SolrClientPlaybill *solr.Connection
var SolrClientSinger *solr.Connection
