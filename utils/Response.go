package utils

import "fmt"

type Response struct {
	Status int
	Message string
	Version int
	Data string
}

const(
	Success = iota
	ParamInvalid
	ServerError
)

// 将response实例转换为json字符串
func (response *Response)ToJsonString() (res string) {
	formatRes := `{"status": %d, "msg": "%s", "ver": %d, "data": %s}`
	return fmt.Sprintf(
		formatRes,
		response.Status,
		response.Message,
		response.Version,
		response.Data,
	)
}

// 返回Success结果
func (response *Response)Success(data string) (res string) {
	response.Status = Success
	response.Version = 0
	response.Message = "OK"
	response.Data = data
	res = response.ToJsonString()
	return res
}

// 参数校验失败
func (response *Response)ParamInvalid() (res string) {
	response.Status = ParamInvalid
	response.Version = 0
	response.Message = "Params Invalid"
	response.Data = `""`
	res = response.ToJsonString()
	return res
}

// 服务器错误
func (response *Response)ServerError(errorInfo string) (res string) {
	response.Status = ServerError
	response.Version = 0
	response.Message = errorInfo
	response.Data = `""`
	res = response.ToJsonString()
	return res
}
