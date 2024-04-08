While:

```bash
go build -gcflags '-m -l'
```

Outputs:

```
./main.go:9:2: moved to heap: f
```

The Benchmark, does not show this allocation:

```
goos: darwin
goarch: arm64
pkg: github.com/svenliebig/go-verify-alocations
BenchmarkAllocations-8   	1000000000	         0.2932 ns/op	       0 B/op	       0 allocs/op
```

Based on this article:

https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d
