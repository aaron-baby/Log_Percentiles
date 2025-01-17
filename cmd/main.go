package main

import (
	"Log_Percentiles/pkg"
	"flag"
	"fmt"
	"github.com/influxdata/tdigest"
	"io/ioutil"
	"log"
	"regexp"
	"sync"
)

func main() {
	dir := flag.String("dir", ".", "log directory")
	flag.Parse()
	var wg sync.WaitGroup

	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	td := tdigest.NewWithCompression(1000)
	format := `$remote_addr [$time_local] "$request" $status $response_time`

	for _, file := range files {
		matched, _ := regexp.MatchString(`.log`, file.Name())
		if !matched {
			continue
		}

		wg.Add(1)

		l := pkg.NewLogParser(format)
		// Launch a goroutine parse log and add values
		go func(fn string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			l.AddValues(*dir + "/" + fn)

		}(file.Name())
		for x := range l.GetValues() {
			td.Add(x, 1)
		}
	}
	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()
	// Compute Quantiles
	fmt.Printf("90%% of requests return a response within %.f ms\n", td.Quantile(0.9))
	fmt.Printf("95%% of requests return a response within %.f ms\n", td.Quantile(0.95))
	fmt.Printf("99%% of requests return a response within %.f ms\n", td.Quantile(0.99))
}
