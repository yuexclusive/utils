package net

import (
	"log"
	"net"

	"github.com/ahmetb/go-linq"
)

func HostIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	var result []*net.IPNet
	linq.From(addrs).Where(func(c interface{}) bool {
		if ipnet, ok := c.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return true
		}
		return false
	}).ToSlice(&result)

	if len(result) <= 0 {
		log.Fatal("无法找到非回环IP地址")
	}

	return result[0].IP.String()
}
