package handlers

import (
	"net/http"
	"fmt"
	"../utils"
	"../global"
	"../logger"
)

// 返回首页UI数据
// 从redis中获取页面的布局信息，将布局信息返回给客户端
func HomeHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	logger.Info.Printf("%s %s", request.Method, request.URL)
	isValid, err := utils.IsParamsValid(request)
	if !isValid {
		ResponseParamInvalid(responseWriter, err)
		return
	}
	level := request.Form.Get("level")
	platform := request.Form.Get("ptf")
	redisKey := fmt.Sprintf(utils.LeanKey, platform, level)
	result, redisErr := global.RedisClient.Get(redisKey).Result()
	if redisErr != nil {
		ResponseServerErr(responseWriter, redisErr.Error())
		return
	}
	ResponseAsJson(responseWriter, result)
}

// 返回详情页UI数据
// 从redis中获取详情页布局信息，将布局信息返回给客户端
func DetailHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	logger.Info.Printf("%s %s", request.Method, request.URL)
	isValid, err := utils.IsParamsValid(request)
	if !isValid {
		ResponseParamInvalid(responseWriter, err)
		return
	}
	level := request.Form.Get("level")
	platform := request.Form.Get("ptf")
	redisKey := fmt.Sprintf(utils.DetailKey, platform, level)
	result, redisErr := global.RedisClient.Get(redisKey).Result()
	if redisErr != nil {
		ResponseServerErr(responseWriter, redisErr.Error())
		return
	}
	ResponseAsJson(responseWriter, result)
}
