package main

func sieve_for_loop(max int) []int {

	primes := make([]int, 0)

	x := make([]int, 0)
	for i := 2; i <= max; i++ {
		x = append(x, i)
	}

	for len(x) > 0 {
		x = sieve_for_loop_runthrough(x, &primes)

	}

	return primes
}

func sieve_for_loop_runthrough(numbers []int, primes *[]int) []int {
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
