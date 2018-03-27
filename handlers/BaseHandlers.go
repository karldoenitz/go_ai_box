package handlers

import (
	"fmt"
	"net/http"
	"../utils"
)

func ResponseAsJson(responseWriter http.ResponseWriter, result string) {
	response := utils.Response{}
	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(responseWriter, response.Success(result))
}

func ResponseParamInvalid(responseWriter http.ResponseWriter, err string)  {
	response := utils.Response{}
	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(responseWriter, response.ParamInvalid(err))
}

func ResponseServerErr(responseWriter http.ResponseWriter, err string)  {
	response := utils.Response{}
	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(responseWriter, response.ServerError(err))
}