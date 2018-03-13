package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
	"net/http"
)

func IsParamsValid(request *http.Request) (result bool, msg string) {
	request.ParseForm()
	token := request.Form.Get("token")
	level := request.Form.Get("level")
	signTime := request.Form.Get("signTime")
	deviceId := request.Form.Get("deviceId")
	if token == "" || level == "" || signTime == "" || deviceId == "" {
		return false, "not enough token"
	}
	now := time.Now().Unix()
	signTimeInt, _ := strconv.Atoi(signTime)
	if int64(signTimeInt) > now || now - int64(signTimeInt) > 60 * 30 {
		return false, "signTime expired or invalid"
	}
	paramsInput := deviceId + signTime + level
	w := md5.New()
	io.WriteString(w, paramsInput)
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	if token != md5str2 {
		return false, "token invalid"
	}
	result = true
	msg = "OK"
	return result, msg
}
