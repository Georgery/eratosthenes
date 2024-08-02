package main

import (
	"fmt"
	"time"
)


func main() {
	primes_up_to := 200000

	start := time.Now()
	primes := sieve_for_loop(primes_up_to)
	fmt.Println("Sieve in for-loop")
	fmt.Println("==========================================")
	fmt.Println("Number of primes below", primes_up_to, ":", len(primes))
	fmt.Println("Calculation took:", time.Since(start))
	fmt.Println()


	start = time.Now()
	primes = sieve_channels_in_slice(primes_up_to)
	fmt.Println("Sieve in channel slice")
	fmt.Println("==========================================")
	fmt.Println("Number of primes below", primes_up_to, ":", len(primes))
	fmt.Println("Calculation took:", time.Since(start))
	fmt.Println()
}
