package main

import (
	"fmt"

	"github.com/positive-builder2/scannerN/scanner"
	"github.com/positive-builder2/scannerN/types"
)

func main() {
	host := "localhost"

	pool := scanner.NewPool(100)

	go func() {
		for res := range pool.Results {
			fmt.Printf(
				"%d/tcp open %s\n",
				res.Port,
				res.Service,
			)
		}
	}()

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
