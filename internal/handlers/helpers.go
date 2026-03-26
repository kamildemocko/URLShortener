package handlers

import (
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}

	if ip == "" {
		ip = r.RemoteAddr
		if strings.Contains(ip, ":") {
			ip, _, _ = net.SplitHostPort(ip)
		}
	}
	return ip
}
