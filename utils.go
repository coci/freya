package freya

import (
	"net"
	"net/http"
	"strings"
)

func getIPAddress(request *http.Request) string {
	//Get IP from the X-REAL-IP header
	ip := request.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip
	}

	//Get IP from X-FORWARDED-FOR header
	ips := request.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		return ""
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip
	}
	return ""
}
