package main

import "rolerocket/internal/benchmark"

func main() {
	benchmark.RunBenchmark("http://localhost:8080/users?username=convoke", 10000, 5)
}

/*
commands to get list of all active things. Failures will start appearing at around 40-50k open connections on windows by default

netsh int ipv4 show dynamicport tcp
netstat -an | Select-String ":8080" | Group-Object { ($_ -split '\s+')[-1] } | Format-Table Count, Name
*/
