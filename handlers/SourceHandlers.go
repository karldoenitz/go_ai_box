package handlers

import (
	"fmt"
	"net/http"
	"../logger"
	"../utils"
)

// 音乐查询接口，从solr中查询Music数据
func MusicHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	logger.Info.Printf("%s %s", request.Method, request.URL)
	isValid, err := utils.IsParamsValid(request)
	response := utils.Response{}
	if !isValid {
		logger.Warning.Println(response.ParamInvalid(err))
		fmt.Fprintf(responseWriter, response.ParamInvalid(err))
		return
	}
	// 此处是业务逻辑
}

// 播单查询接口，从solr中查询Playbill数据
func PlaybillHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	logger.Info.Printf("%s %s", request.Method, request.URL)
	isValid, err := utils.IsParamsValid(request)
	response := utils.Response{}
	if !isValid {
		logger.Warning.Println(response.ParamInvalid(err))
		fmt.Fprintf(responseWriter, response.ParamInvalid(err))
		return
	}
	// 此处是业务逻辑
}

// 歌手查询接口，从solr中查询Singer数据
func SingerHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	logger.Info.Printf("%s %s", request.Method, request.URL)
	isValid, err := utils.IsParamsValid(request)
	response := utils.Response{}
	if !isValid {
		logger.Warning.Println(response.ParamInvalid(err))
		fmt.Fprintf(responseWriter, response.ParamInvalid(err))
		return
	}
	// 此处是业务逻辑
}
