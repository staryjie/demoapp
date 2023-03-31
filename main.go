package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"time"
)

func getLocalIP() string {
	var ipStr []string
	interfaces, _ := net.Interfaces()
	for i := 0; i < len(interfaces); i++ {
		if (interfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := interfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//获取IPv6
					/*if ipnet.IP.To16() != nil {
					    fmt.Println(ipnet.IP.String())
					    ipStr = append(ipStr, ipnet.IP.String())
					}*/
					//获取IPv4
					if ipnet.IP.To4() != nil {
						//fmt.Println(ipnet.IP.String())
						ipStr = append(ipStr, ipnet.IP.String())
					}
				}
			}
		}
	}
	return ipStr[0]
}

func main() {
	started := time.Now()
	hostname, _ := os.Hostname()
	// 设置为生产模式
	gin.SetMode("release")
	r := gin.Default()
	//r.TrustedPlatform = "Client-IP"
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message":  "OK",
			"HostName": hostname,
			"ServerIP": getLocalIP(),
			"ClientIP": c.ClientIP(),
			"Version":  "v1.3",
			//"client-ip": c.Request.RemoteAddr,
		})
	})
	r.GET("/started", func(c *gin.Context) {
		duration := time.Since(started)
		if duration.Seconds() > 10 {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "InternalServerError", "data": fmt.Sprintf("error: %v", duration.Seconds())})
		} else {
			c.JSON(http.StatusOK, "OK")
		}
	})
	r.GET("/readyz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.GET("/livez", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.GET("/hostname", func(c *gin.Context) {
		c.String(http.StatusOK, hostname)
	})
	r.GET("/servername", func(c *gin.Context) {
		c.String(http.StatusOK, "demoapp")
	})
	r.Run("0.0.0.0:8080")
}
