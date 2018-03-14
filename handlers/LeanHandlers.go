package handlers

import (
	"net/http"
	"fmt"
	"../utils"
)

// 返回首页数据，主要是页面的UI布局数据
func HomeHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	isValid, err := utils.IsParamsValid(request)
	if !isValid {
		response := utils.Response{Message:err}
		fmt.Fprintf(responseWriter, response.ParamInvalid())
		return
	}
	level := request.Form.Get("level")
	platform := request.Form.Get("ptf")
	redisKey := fmt.Sprintf(utils.LeanKey, platform, level)
}
