package handlers

import (
	"net/http"
	"fmt"
	"../utils"
	"../global"
)

// 返回首页UI数据
func HomeHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	response := utils.Response{}
	isValid, err := utils.IsParamsValid(request)
	if !isValid {
		fmt.Fprintf(responseWriter, response.ParamInvalid(err))
		return
	}
	level := request.Form.Get("level")
	platform := request.Form.Get("ptf")
	redisKey := fmt.Sprintf(utils.LeanKey, platform, level)
	result, redisErr := global.RedisClient.Get(redisKey).Result()
	if redisErr != nil {
		fmt.Fprintf(responseWriter, response.ServerError(redisErr.Error()))
		return
	}
	fmt.Fprintf(responseWriter, response.Success(result))
}

// 返回详情页UI数据
func DetailHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	response := utils.Response{}
	isValid, err := utils.IsParamsValid(request)
	if !isValid {
		fmt.Fprintf(responseWriter, response.ParamInvalid(err))
		return
	}
	level := request.Form.Get("level")
	platform := request.Form.Get("ptf")
	redisKey := fmt.Sprintf(utils.DetailKey, platform, level)
	result, redisErr := global.RedisClient.Get(redisKey).Result()
	if redisErr != nil {
		fmt.Fprintf(responseWriter, response.ServerError(redisErr.Error()))
		return
	}
	fmt.Fprintf(responseWriter, response.Success(result))
}
