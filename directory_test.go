package util

import "testing"

func TestInitDir(t *testing.T) {
	err := InitDir("test/hello/world")
	if err != nil {
		t.Log("Init dir error:",err.Error())
	}else{
		t.Log("init dir success...")
	}
}
