package handlers

import (
	"fmt"
	"net/http"
	"../utils"
	"../logger"
)

// 返回json数据，状态成功
func ResponseAsJson(responseWriter http.ResponseWriter, result string) {
	response := utils.Response{}
	responseWriter.Header().Set("Content-Type", "application/json")
	logger.Info.Println("Response Success ...")
	fmt.Fprintf(responseWriter, response.Success(result))
}

// 返回json数据，参数校验失败
func ResponseParamInvalid(responseWriter http.ResponseWriter, err string)  {
	response := utils.Response{}
	responseWriter.Header().Set("Content-Type", "application/json")
	logger.Warning.Println(err)
	fmt.Fprintf(responseWriter, response.ParamInvalid(err))
}

// 返回json数据，服务器端错误
func ResponseServerErr(responseWriter http.ResponseWriter, err string)  {
	response := utils.Response{}
	responseWriter.Header().Set("Content-Type", "application/json")
	logger.Error.Println(err)
	fmt.Fprintf(responseWriter, response.ServerError(err))
}