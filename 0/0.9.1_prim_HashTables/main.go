package main

import "fmt"

func containsDuplicate(nums []int) bool {
	/*arrLen := len(nums)
	for i := 0; i < arrLen; i++ { // перебираем все индексы исходного массива
		for j := i + 1; j < arrLen; j++ { // перебираем все индексы после i (чтобы не перебирать повторяющиеся варианты)
			if nums[i] == nums[j] {
				return true // дубликат найден
			}
		}
	}*/

	set := make(map[int]bool)
    for _, num := range nums {
        if _, ok := set[num]; ok // если такой ключ существует, переходим к return 
        {
            return true
        }
        set[num] = true
    }

	return false // после полной проверки дубликат не найден
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 1} // Пример входных данных
	fmt.Println(containsDuplicate(numbers))
}
