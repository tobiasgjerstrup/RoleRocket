package main

import "rolerocket/internal/benchmark"

func main() {
	benchmark.RunBenchmark("http://localhost:8080/users", 500, 5)
}
