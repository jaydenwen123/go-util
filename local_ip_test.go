package util

import "testing"

func TestGetLocalAddress(t *testing.T) {
	ips, err := GetAllInterfaceIps()
	if err != nil {
		t.Logf("GetAllInterfaceIps failed:%s",err.Error())
	}else {
		t.Log(ips)
	}
}

func TestGetLocalAddressByIterface(t *testing.T) {
	address, err := GetLocalAddressByIterface("en0")
	if err != nil {
		t.Logf("GetAllInterfaceIps failed:%s",err.Error())
	}else {
		t.Log("address:",address)
	}
}
