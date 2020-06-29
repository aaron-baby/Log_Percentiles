package tdigest

import (
	"github.com/influxdata/tdigest"
	"log"
	"testing"
)

func TestInfluxdata(t *testing.T) {
	td := tdigest.NewWithCompression(1000)
	values := make(chan float64)
	go func() {
		for i := 1; i <= 10000; i++ {
			values <- float64(i)
		}
		close(values)
	}()
	for x := range values {
		td.Add(x, 1)
	}

	// Compute Quantiles
	log.Println("50th", td.Quantile(0.5))
	log.Println("75th", td.Quantile(0.75))
	log.Println("90th", td.Quantile(0.9))
	log.Println("99th", td.Quantile(0.99))
}
