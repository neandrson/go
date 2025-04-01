package main

import (
	"errors"
	"sync"
)

func ProcessSum(
	summer func(arr []int, result chan<- int),
	nums []int,
	chunkSize int,
) (int, error) {
	if chunkSize <= 0 {
		return 0, errors.New("non-positive chunk's size")
	}

	size := len(nums) / chunkSize
	if len(nums)%chunkSize != 0 {
		size++
	}

	wg := sync.WaitGroup{}
	wg.Add(size)
	result := make(chan int, size)
	defer close(result)

	var first, last int = 0, chunkSize
	if last > len(nums) {
		last = len(nums)
	}

	for first <= len(nums) {
		go func(arr []int) {
			defer wg.Done()
			summer(arr, result)
		}(nums[first:last])

		first += chunkSize
		last += chunkSize
		if last > len(nums) {
			last = len(nums)
		}
	}

	wg.Wait()

	var sum int
	for i := 0; i < size; i++ {
		select {
		case v, ok := <-result:
			if ok {
				sum += v
			} else {
				break
			}
		default:
			break
		}
	}

	return sum, nil
}

func SumChunk(arr []int, result chan<- int) {
	var sum int
	for _, v := range arr {
		sum += v
	}
	result <- sum
}
