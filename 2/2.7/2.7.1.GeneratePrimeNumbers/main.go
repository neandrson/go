package main

import "time"

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	primes := make(map[int]struct{})

	time.AfterFunc(10*time.Millisecond, func() {
		stop <- struct{}{}
	})
	for n := 2; n < N; n++ {
		isPrime := true
		for k, _ := range primes {
			if n%k == 0 {
				isPrime = false
				break
			}
		}
		select {
		case <-stop:
			return
		default:
			if isPrime {
				primes[n] = struct{}{}
				prime_nums <- n
			}
		}
	}
	close(prime_nums)
}

func main() {

}
