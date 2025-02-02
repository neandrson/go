package main

import (
	"fmt"
	"reflect"
	"time"
)

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
	stop := make(chan struct{})
	primeChan := make(chan int)

	// test for generating primes up to 10
	expectedPrimesUpTo10 := []int{2, 3, 5, 7}
	go GeneratePrimeNumbers(stop, primeChan, 10)

	receivedPrimes := make([]int, 0)
	for prime := range primeChan {
		receivedPrimes = append(receivedPrimes, prime)
	}

	// close stop channel to terminate generatePrimeNumbers goroutine
	close(stop)

	// check if received primes match the expected primes up to 10
	if !reflect.DeepEqual(receivedPrimes, expectedPrimesUpTo10) {
		fmt.Printf("Generated primes mismatch for N=10. Expected %v, got %v\n", expectedPrimesUpTo10, receivedPrimes)
	}
}
