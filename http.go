package util

import (
	"net/http"
)

//SetHeader 解决跨域
func SetHeader(resp http.ResponseWriter) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Methods, Access-Control-Allow-Credentials, Origin, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	resp.Header().Set("content-type", "application/json")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
}

