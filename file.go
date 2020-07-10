package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/astaxie/beego/logs"
)

var cmp = regexp.MustCompile(`([\/:*?"<>|]+)`)

//TrimIllegalStr 去除文件名中的非法字符
func TrimIllegalStr(filename string, replaceStr string) string {
	newStr := cmp.ReplaceAllString(filename, replaceStr)
	return newStr
}

//WriteData2File 将数据写入到文件
func WriteData2File(basepath, filename, content string, perm os.FileMode) error {
	var err error

	if ! IsExist(basepath) {
		err = InitDir(basepath)
		if err != nil {
			logs.Error("Init the dir<%s> conf error:%s", basepath, err.Error())
			return err
		}
	}
	shellFileName := filepath.Join(basepath, filename)
	//err = util.AppendLine2File(shellFileName, content)
	err = ioutil.WriteFile(shellFileName, []byte(content), perm)
	if err != nil {
		logs.Error("init the <%s> error:%s", shellFileName, err.Error())
		return err
	}
	return nil
}
