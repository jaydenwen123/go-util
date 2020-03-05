package util

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

//从json文件中恢复出golang对象
func LoadObjectFromJsonFile(filePath string, obj interface{}) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logs.Error("read json file error.", err.Error())
		panic(err.Error())
	}
	LoadObjFromJsonStr(string(data), obj)
}

//LoadObjFromJsonStr 将jsonStr字符串转为go对象
func LoadObjFromJsonStr(jsonStr string, obj interface{}) {

	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		logs.Error("json to object error.", err.Error())
		panic(err.Error())
	}
}
