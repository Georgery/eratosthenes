package main


func generate_channel(max int) <-chan int {
	c := make(chan int, 1000000)

	go func() {
		for i := 2; i <= max; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func new_channel_after(from <-chan int, primes *[]int) (<-chan int, bool) {
	to := make(chan int, 20)

	prime, ok := <-from // first one in the channel must be a prime

	if !ok {
		close(to)
		return to, true
	}

	*primes = append(*primes, prime) // ...and therefore goes here
	prime_multiple := prime

	go func() {
		for number := range from {
			if prime_multiple < number {
				prime_multiple += prime
			}
			if prime_multiple > number {
				to <- number
			}
		}

		close(to)
	}()

	return to, false
}

func sieve_channels_in_slice(max int) []int {
	primes := make([]int, 0)
	from := generate_channel(max)

	channels := []<-chan int{from}

	for {

		to, channel_empty := new_channel_after(channels[len(channels)-1], &primes)
		if channel_empty {
			break
		}

		channels = append(channels, to)
	}

	return primes
}
