package utils

import (
	"errors"
	"net"
)

// GetLocalIPv4 获取本地IPV4地址
func GetLocalIPv4() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range interfaces {
		// 获取非回环地址的接口和已启用的接口
		if iface.Flags&net.FlagLoopback != net.FlagLoopback && iface.Flags&net.FlagUp != 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}

			for _, addr := range addrs {
				// 获取非回环地址的IPV4地址
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil
				}
			}
		}
	}
	return "", errors.New("get local IP error")
}

// MustGetLocalIPv4 获取本地IPV4地址对外函数
func MustGetLocalIPv4() string {
	ipv4, err := GetLocalIPv4()
	if err != nil {
		panic("get local IP error")
	}
	return ipv4
}
