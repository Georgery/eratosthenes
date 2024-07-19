package main

import (
	"fmt"
	"time"
)

func sieve(numbers []int, primes *[]int) []int {
	prime := numbers[0]
	multiple := prime
	*primes = append(*primes, prime)
	numbers = numbers[1:]
	remaining_numbers := make([]int, 0)

	for _, v := range numbers {

		if multiple < v {
			multiple += prime
		}

		if multiple > v {
			remaining_numbers = append(remaining_numbers, v)
		}
	}
	return remaining_numbers
}

func main() {
	start := time.Now()
	max := 1000000
	primes := make([]int, 0)

	x := make([]int, 0)
	for i := 2; i <= max; i++ {
		x = append(x, i)
	}

	for len(x) > 0 {
		x = sieve(x, &primes)

	}
	// fmt.Println(primes)
	fmt.Println(time.Since(start))
}
