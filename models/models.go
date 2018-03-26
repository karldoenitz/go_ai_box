package models

import (
	"fmt"
	"github.com/rtt/Go-Solr"
	"../global"
	"../logger"
)

type MusicData struct {
	Data []Music
	Length int
}

type PlaybillData struct {
	Data []Playbill
	Length int
}

type SingerData struct {
	Data []Singer
	Length int
}

type Music struct {
	Id int64
	Name string
	Singer []int64
	SearchWorld []string
	Url string
}

type Playbill struct {
	Id int64
	Name string
	SearchWorld []string
	Music []int64
}

type Singer struct {
	Id int64
	Name string
	SearchWorld []string
	Music []int64
	Playbill []int64
}

func GetMusicByName(name string) (result []Music) {
	solrMusicConn := global.SolrClientMusic
	fq := fmt.Sprintf("name:%s", name)
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"*:*"},
			"fq": []string{fq},
			"wt": []string{"json"},
			"fl": []string{"*"},
		},
		Rows: 10,
		Sort: "id+ASC",
	}
	queryString := q.String()
	fmt.Println(queryString)
	res, selectErr := solrMusicConn.Select(&q)
	if selectErr != nil {
		logger.Error.Printf("select ERROR in model.go line 63: %s", selectErr.Error())
		panic(selectErr)
	}
	solrResults := res.Results
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

func GetMusicById(id int) (result []Music) {
	solrMusicConn := global.SolrClientMusic
	fq := fmt.Sprintf("id:%d", id)
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"*:*"},
			"fq": []string{fq},
			"wt": []string{"json"},
			"fl": []string{"*"},
		},
		Rows: 10,
		Sort: "id+ASC",
	}
	queryString := q.String()
	fmt.Println(queryString)
	res, selectErr := solrMusicConn.Select(&q)
	if selectErr != nil {
		logger.Error.Printf("select ERROR in model.go line 111: %s", selectErr.Error())
		panic(selectErr)
	}
	solrResults := res.Results
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

func GetSingerById(id int) (result []Music) {
	solrMusicConn := global.SolrClientSinger
	fq := fmt.Sprintf("id:%d", id)
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"*:*"},
			"fq": []string{fq},
			"wt": []string{"json"},
			"fl": []string{"*"},
		},
		Rows: 10,
		Sort: "id+ASC",
	}
	queryString := q.String()
	fmt.Println(queryString)
	res, selectErr := solrMusicConn.Select(&q)
	if selectErr != nil {
		logger.Error.Printf("select ERROR in model.go line 159: %s", selectErr.Error())
		panic(selectErr)
	}
	solrResults := res.Results
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

func GetPlaybillById(id int) (result []Music) {
	solrMusicConn := global.SolrClientMusic
	fq := fmt.Sprintf("id:%d", id)
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"*:*"},
			"fq": []string{fq},
			"wt": []string{"json"},
			"fl": []string{"*"},
		},
		Rows: 10,
		Sort: "id+ASC",
	}
	queryString := q.String()
	fmt.Println(queryString)
	res, selectErr := solrMusicConn.Select(&q)
	if selectErr != nil {
		logger.Error.Printf("select ERROR in model.go line 63: %s", selectErr.Error())
		panic(selectErr)
	}
	solrResults := res.Results
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
