package scanner

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func ProbeHTTP(host string, port int) string {
	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout(
		"tcp",
		address,
		3*time.Second,
	)

	if err != nil {
		return ""
	}

	defer conn.Close()

	request :=
		"GET / HTTP/1.1\r\n" +
			"Host: " + host + "\r\n" +
			"Connection:close\r\n" +
			"\r\n"
	_, err = conn.Write([]byte(request))

	if err != nil {
		return ""
	}

	conn.SetReadDeadline(
		time.Now().Add(3 * time.Second),
	)

	var response strings.Builder

	buff := make([]byte, 4096)

	for {
		n, err := conn.Read(buff)

		if n > 0 {
			response.Write(buff[:n])
		}

		if err != nil {
			break
		}
	}
	return response.String()
}
