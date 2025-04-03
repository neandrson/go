package main

func MultiplyPipeline(inputNums ...[]int) int {
	var multiply int = 1
	for _, input := range inputNums {
		numbers := NumbersGen(input...)
		filterNumbers := Filter(numbers)
		multiply *= Multiply(filterNumbers)
	}
	return multiply
}

func NumbersGen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()
	return out
}

func Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			if num > 0 {
				out <- num
			}
		}
	}()
	return out
}

func Multiply(in <-chan int) int {
	multiply := 1
	for num := range in {
		multiply *= num
	}
	return multiply
}
