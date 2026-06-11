package scanner

import (
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/positive-builder2/scannerN/types"
)

func ParsePorts(s string) []int {
	var ports []int

	parts := strings.Split(s, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if strings.Contains(part, "-") {
			//
			bounds := strings.Split(part, "-")

			if len(bounds) != 2 {
				continue
			}

			start, err1 := strconv.Atoi(bounds[0])
			end, err2 := strconv.Atoi(bounds[1])

			if err1 != nil || err2 != nil {
				continue
			}

			for p := start; p <= end; p++ {
				ports = append(ports, p)
			}
		} else {
			port, err := strconv.Atoi(part)

			if err == nil {
				ports = append(ports, port)
			}
		}
	}
	return ports
}

func ScanPort(job types.Job) types.Result {
	address := job.Host + ":" + strconv.Itoa(job.Port)

	conn, err := net.DialTimeout(
		"tcp",
		address,
		500*time.Millisecond,
	)

	if err != nil {
		return types.Result{
			Host: job.Host,
			Port: job.Port,
			Open: false,
		}
	}

	defer conn.Close()

	conn.SetReadDeadline(
		time.Now().Add(2 * time.Second),
	)

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	banner := ""

	if n > 0 {
		banner = string(buf[:n])
	}

	return types.Result{
		Host:    job.Host,
		Port:    job.Port,
		Open:    true,
		Service: GuessService(job.Port),
		Banner:  banner,
	}
}
