package main

import (
	"fmt"

	"github.com/positive-builder2/scannerN/scanner"
	"github.com/positive-builder2/scannerN/types"
)

func main() {
	host := "localhost"

	pool := scanner.NewPool(100)

	// resp := scanner.ProbeHTTP("localhost", 8081)
	// fmt.Println(resp)

	go func() {
		for res := range pool.Results {
			fmt.Printf(
				"%d/tcp open %s %q\n",
				res.Port,
				res.Service,
				res.Banner,
			)
		}
	}()

	//fmt.Println(scanner.ParsePorts("22,80,100-103"))

	for port := 1; port <= 1000; port++ {
		pool.Wg.Add(1)

		pool.Jobs <- types.Job{
			Host: host,
			Port: port,
		}
	}

	pool.Wg.Wait()

	close(pool.Jobs)
	close(pool.Results)
}
