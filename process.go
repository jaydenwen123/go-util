package util

import (
	"os"
	"path/filepath"
	"strings"
)

//GetCurrentProcessName 获取当前进程的名字，支持绝对路径和相对路径两种
// /usr/local/hello
// ./hello
func GetCurrentProcessName() string {
	processName := os.Args[0]
	//去除相对路径
	if strings.HasPrefix(processName, "./") && len(processName) > 2 {
		processName = processName[2:]
	}
	//绝对路径
	if strings.HasPrefix(processName, "/") {
		_, name := filepath.Split(processName)
		processName = name
	}
	//去除配置文件
	tmpArr := strings.Split(processName, " ")
	if len(tmpArr) > 0 {
		processName = tmpArr[0]
	}
	return processName
}
