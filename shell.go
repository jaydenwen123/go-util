package util

import (
	"bytes"
	"os/exec"

	"github.com/astaxie/beego/logs"
)

//ExecShell 执行shell命令
func ExecShell(shell string, printLog bool) (string, error) {
	logs.Debug("the exec shell script:<%s>", shell)
	cmd := exec.Command("/bin/bash", "-c", shell)
	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Run()
	if printLog {
		logs.Debug("Stdout: [%s]", string(outbuf.Bytes()))
		logs.Debug("Stderr: [%s]", string(errbuf.Bytes()))
	}
	return outbuf.String(), err
}
