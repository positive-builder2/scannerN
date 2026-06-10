package scanner

import (
	"net"
	"strconv"
	"time"

	"github.com/positive-builder2/scannerN/types"
)

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
