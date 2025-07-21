package benchmark

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func RunBenchmark(url string, totalRequests int, concurrency int) {
	var wg sync.WaitGroup
	start := time.Now()

	requestsPerWorker := totalRequests / concurrency
	results := make(chan time.Duration, totalRequests)
	var failureCount int64

	// Reusable HTTP client with keep-alives and timeouts
	client := &http.Client{
		Timeout: 2 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        concurrency,
			MaxIdleConnsPerHost: concurrency,
		},
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				reqStart := time.Now()
				resp, err := client.Get(url)
				if err != nil {
					atomic.AddInt64(&failureCount, 1)
				} else {
					resp.Body.Close()
				}
				results <- time.Since(reqStart)
			}
		}()
	}

	wg.Wait()
	close(results)

	// Analyze results
	var total time.Duration
	var fastest, slowest time.Duration
	fastest = time.Hour

	for r := range results {
		total += r
		if r < fastest {
			fastest = r
		}
		if r > slowest {
			slowest = r
		}
	}

	fmt.Printf("Benchmark results for %s\n", url)
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("Concurrency: %d\n", concurrency)
	fmt.Printf("Total time: %v\n", time.Since(start))
	fmt.Printf("Average latency: %v\n", total/time.Duration(totalRequests))
	fmt.Printf("Fastest: %v\n", fastest)
	fmt.Printf("Slowest: %v\n", slowest)
	fmt.Printf("Failures: %d\n", failureCount)
}
