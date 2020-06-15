package util

import (
	"github.com/astaxie/beego/logs"
	"os"
)

//初始化目录,目录不存在则创建，如果存在则直接跳过
func InitDir(path string) error {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logs.Error("init directory <", path, ">error.", err.Error())
			return err
		}
		logs.Info("init directory<", path, "> success.")
		return err
	}
	return nil
}

//IsExist 判断文件或者目录存在
func IsExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

//AppendData2File 追加数据到文件中
func AppendLine2File(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		logs.Error("open file error:%s", err.Error())
		return err
	}
	_, err = file.WriteString(data + "\n")
	if err != nil {
		logs.Error("write string error:%s", err.Error())
		return err
	}
	return nil
}
