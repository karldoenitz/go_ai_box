package handlers

import (
	"fmt"
	"net/http"
	"../logger"
	"../utils"
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
	// 此处是业务逻辑
}
