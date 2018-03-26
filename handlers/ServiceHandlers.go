package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"../logger"
	"../utils"
	"../models"
)

// 查询接口，从solr中查询数据
func SearchHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	logger.Info.Printf("%s %s", request.Method, request.URL)
	isValid, err := utils.IsParamsValid(request)
	response := utils.Response{}
	if !isValid {
		logger.Warning.Println(response.ParamInvalid(err))
		fmt.Fprintf(responseWriter, response.ParamInvalid(err))
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
		logger.Error.Println("marshal musicData to json Error, on ServiceHandlers.go line 29: %s", jsonErr.Error())
		panic(jsonErr)
	}
	fmt.Fprintf(responseWriter, response.Success(string(jsonResult)))
}
