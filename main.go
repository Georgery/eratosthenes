package main

import (
	"fmt"
	"time"
)


func main() {
	start := time.Now()
	primes := sieve_for_loop(100000)

	fmt.Println(len(primes))
	fmt.Println(time.Since(start))
}
