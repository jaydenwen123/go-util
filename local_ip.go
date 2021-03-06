package util

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
	"strings"
)

const defaultNetInterfaceName = "eth1"

//GetLocalAddress 获取当前机器的ip地址,默认取得是eth1网卡
func GetLocalAddress() (string, error) {
	return GetLocalAddressByIterface(defaultNetInterfaceName)
}

//GetLocalAddressByIterface 根据网卡名字获取ip
func GetLocalAddressByIterface(interfaceName string) (string, error) {
	ipMap, err := GetAllInterfaceIps()
	if err != nil {
		logs.Error("GetAllInterfaceIps error:%s", err.Error())
		return "", err
	}
	addrs, ok := ipMap[interfaceName]
	if !ok {
		return "", fmt.Errorf("the iterface is not existed")
	}
	if len(addrs) >= 1 {
		for _, ad := range addrs {
			if len(strings.Split(ad, ".")) == 4 {
				return ad, nil
			}
		}
		return "", fmt.Errorf("there is not valid address")
	}
	return "", nil
}

//GetAllInterfaceIps 获取所有网卡的ip信息
func GetAllInterfaceIps() (map[string][]string, error) {
	ips := make(map[string][]string)
	interfaces, err := net.Interfaces()
	if err != nil {
		logs.Error("get Interfaces error:%s", err.Error())
		return nil, err
	}

	for _, iter := range interfaces {
		addrs, err := iter.Addrs()
		if err != nil {
			logs.Error("get Addrs error:%s", err.Error())
			continue
		}
		allAddrs := getAllAddrs(addrs)
		ips[iter.Name] = allAddrs
	}
	return ips, nil
}

func getAllAddrs(addrs []net.Addr) []string {
	realAddrs := make([]string, 0)
	for _, a := range addrs {
		tmpArr := strings.Split(a.String(), "/")
		if len(tmpArr) == 2 {
			realAddrs = append(realAddrs, tmpArr[0])
		}
	}
	return realAddrs
}
