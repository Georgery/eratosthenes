// The difference in this version is that we are noting down the non-prime numbers and then remove them from the vector.

package main

func sieve_for_loop_remove(max int) []int {
	primes := make([]int, 0)
	x := make([]int, 0)
	for i := 2; i <= max; i++ {
		x = append(x, i)
	}

	for len(x) > 0 {
		sieve_for_loop_runthrough_remove(&x, &primes)
	}

	return primes
}

func sieve_for_loop_runthrough_remove(numbers *[]int, primes *[]int) {
	prime := (*numbers)[0]
	multiple := prime
	*primes = append(*primes, prime)
	*numbers = (*numbers)[1:]

	indices_to_remove := make([]int, 0)

	for i, v := range *numbers {
		for multiple < v {
			multiple += prime
		}

		if multiple == v {
			indices_to_remove = append(indices_to_remove, i)
		}
	}

	for i := len(indices_to_remove) - 1; i >= 0; i-- {
		*numbers = append((*numbers)[:indices_to_remove[i]], (*numbers)[indices_to_remove[i]+1:]...)
	}


}