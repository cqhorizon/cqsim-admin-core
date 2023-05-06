package pkg

import (
	"fmt"
	"net"
)

// GetLocalHost 获取局域网ip地址
func GetLocalHost() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			adders, _ := netInterfaces[i].Addrs()

			for _, address := range adders {
				if aspnet, ok := address.(*net.IPNet); ok && !aspnet.IP.IsLoopback() {
					if aspnet.IP.To4() != nil {
						return aspnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}
