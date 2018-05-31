package models

import (
	"fmt"
	"github.com/rtt/Go-Solr"
	"../global"
	"../logger"
)

// 获取solr查询语句
// 初始化并返回一个solr查询Query实例
func getSolrQuery(fq string) (q solr.Query) {
	q = solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"*:*"},
			"fq": []string{fq},
			"wt": []string{"json"},
			"fl": []string{"*"},
		},
		Rows: 10,
		Sort: "id+ASC",
	}
	return q
}

// 根据solr链接及solr查询语句查询结果
// 失败则抛出异常
func getSolrQueryResult(conn *solr.Connection, fq string) (result *solr.DocumentCollection) {
	q := getSolrQuery(fq)
	queryString := q.String()
	logger.Info.Println(queryString)
	res, selectErr := conn.Select(&q)
	if selectErr != nil {
		logger.Error.Printf("solr query ERROR: %s", selectErr.Error())
		panic(selectErr)
	}
	solrResults := res.Results
	return solrResults
}

// 根据音乐名称获取音乐model数据
// 返回值为音乐model类型的数组
func GetMusicByName(name string) (result []Music) {
	solrMusicConn := global.SolrClientMusic
	fq := fmt.Sprintf("name:%s", name)
	solrResults := getSolrQueryResult(solrMusicConn, fq)
	for i := 0; i < solrResults.NumFound; i++ {
		collection := solrResults.Collection[i]
		id := int64(collection.Fields["id"].(float64))
		name := collection.Fields["name"].(string)
		url := collection.Fields["url"].(string)
		singerInterface := collection.Fields["singer"].([]interface{})
		searchWordInterface := collection.Fields["searchWord"].([]interface{})
		singer := make([]int64, len(singerInterface))
		for i := range singerInterface {
			singer[i] = int64(singerInterface[i].(float64))
		}
		searchWord := make([]string, len(searchWordInterface))
		for i := range searchWordInterface {
			searchWord[i] = searchWordInterface[i].(string)
		}
		musicObj := Music{
			Id: id,
			Name: name,
			Url: url,
			Singer: singer,
			SearchWorld: searchWord,
		}
		result = append(result, musicObj)
	}
	return result
}

//根据音乐id获取音乐model数据
//返回值为音乐model类型的数组
func GetMusicById(id int) (result []Music) {
	solrMusicConn := global.SolrClientMusic
	fq := fmt.Sprintf("id:%d", id)
	solrResults := getSolrQueryResult(solrMusicConn, fq)
	for i := 0; i < solrResults.NumFound; i++ {
		collection := solrResults.Collection[i]
		id := int64(collection.Fields["id"].(float64))
		name := collection.Fields["name"].(string)
		url := collection.Fields["url"].(string)
		singerInterface := collection.Fields["singer"].([]interface{})
		searchWordInterface := collection.Fields["searchWord"].([]interface{})
		singer := make([]int64, len(singerInterface))
		for i := range singerInterface {
			singer[i] = int64(singerInterface[i].(float64))
		}
		searchWord := make([]string, len(searchWordInterface))
		for i := range searchWordInterface {
			searchWord[i] = searchWordInterface[i].(string)
		}
		musicObj := Music{
			Id: id,
			Name: name,
			Url: url,
			Singer: singer,
			SearchWorld: searchWord,
		}
		result = append(result, musicObj)
	}
	return result
}

//根据id获取歌手model数据
//返回值为歌手model类型的数组
func GetSingerById(id int) (result []Singer) {
	solrSingerConn := global.SolrClientSinger
	fq := fmt.Sprintf("id:%d", id)
	solrResults := getSolrQueryResult(solrSingerConn, fq)
	for i := 0; i < solrResults.NumFound; i++ {
		collection := solrResults.Collection[i]
		id := int64(collection.Fields["id"].(float64))
		name := collection.Fields["name"].(string)
		playbillInterface := collection.Fields["playbill"].([]interface{})
		musicInterface := collection.Fields["music"].([]interface{})
		searchWordInterface := collection.Fields["searchWord"].([]interface{})
		playbill := make([]int64, len(playbillInterface))
		for i := range playbillInterface {
			playbill[i] = int64(playbillInterface[i].(float64))
		}
		music := make([]int64, len(musicInterface))
		for i := range musicInterface {
			music[i] = int64(musicInterface[i].(float64))
		}
		searchWord := make([]string, len(searchWordInterface))
		for i := range searchWordInterface {
			searchWord[i] = searchWordInterface[i].(string)
		}
		singerObj := Singer{
			Id: id,
			Name: name,
			Playbill: playbill,
			Music: music,
			SearchWorld: searchWord,
		}
		result = append(result, singerObj)
	}
	return result
}

//根据id获取播单model数据
//返回值为播单model类型的数组
func GetPlaybillById(id int) (result []Playbill) {
	solrPlaybillConn := global.SolrClientPlaybill
	fq := fmt.Sprintf("id:%d", id)
	solrResults := getSolrQueryResult(solrPlaybillConn, fq)
	for i := 0; i < solrResults.NumFound; i++ {
		collection := solrResults.Collection[i]
		id := int64(collection.Fields["id"].(float64))
		name := collection.Fields["name"].(string)
		musicInterface := collection.Fields["music"].([]interface{})
		searchWordInterface := collection.Fields["searchWord"].([]interface{})
		music := make([]int64, len(musicInterface))
		for i := range musicInterface {
			music[i] = int64(musicInterface[i].(float64))
		}
		searchWord := make([]string, len(searchWordInterface))
		for i := range searchWordInterface {
			searchWord[i] = searchWordInterface[i].(string)
		}
		musicObj := Playbill{
			Id: id,
			Name: name,
			Music: music,
			SearchWorld: searchWord,
		}
		result = append(result, musicObj)
	}
	return result
}
