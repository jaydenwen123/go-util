package util

import "testing"

func TestInitDir(t *testing.T) {
	err := InitDir("test/hello/world")
	if err != nil {
		t.Log("Init dir error:", err.Error())
	} else {
		t.Log("init dir success...")
	}
}

func TestAppendLine2File(t *testing.T) {
	AppendLine2File("test.txt", "hello")
	AppendLine2File("test.txt", "hello1")
	AppendLine2File("test.txt", "hello2")
	AppendLine2File("test.txt", "hello3")
	AppendLine2File("test.txt", `{"hello":"world","usernmae":"jaydenwen"}`)
	AppendLine2File("test.txt", `{"hello":"world","usernmae":"jaydenwen"}`)
	AppendLine2File("test.txt", `{"hello":"world","usernmae":"jaydenwen"}`)
}
