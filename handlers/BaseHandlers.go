package handlers

import (
	"fmt"
	"net/http"
	"../utils"
	"../logger"
)

func ResponseAsJson(responseWriter http.ResponseWriter, result string) {
	response := utils.Response{}
	responseWriter.Header().Set("Content-Type", "application/json")
	logger.Info.Println("Response Success ...")
	fmt.Fprintf(responseWriter, response.Success(result))
}

func ResponseParamInvalid(responseWriter http.ResponseWriter, err string)  {
	response := utils.Response{}
	responseWriter.Header().Set("Content-Type", "application/json")
	logger.Warning.Println(err)
	fmt.Fprintf(responseWriter, response.ParamInvalid(err))
}

func ResponseServerErr(responseWriter http.ResponseWriter, err string)  {
	response := utils.Response{}
	responseWriter.Header().Set("Content-Type", "application/json")
	logger.Error.Println(err)
	fmt.Fprintf(responseWriter, response.ServerError(err))
}