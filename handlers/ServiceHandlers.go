package handlers

import (
	"net/http"
	"encoding/json"
	"../logger"
	"../utils"
	"../models"
)

// 查询接口，从solr中查询数据
// 根据客户端传的searchWord参数，到solr中进行相关查询，然后将数据返回给客户端
func SearchHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	logger.Info.Printf("%s %s", request.Method, request.URL)
	isValid, err := utils.IsParamsValid(request)
	if !isValid {
		ResponseParamInvalid(responseWriter, err)
		return
	}
	searchWord := request.Form.Get("searchWord")
	musicSearchResults := models.GetMusicByName(searchWord)
	musicData := models.MusicData{
		Data: musicSearchResults,
		Length: len(musicSearchResults),
	}
	jsonResult, jsonErr := json.Marshal(musicData)
	if jsonErr != nil {
		ResponseServerErr(responseWriter, jsonErr.Error())
	}
	ResponseAsJson(responseWriter, string(jsonResult))
}
