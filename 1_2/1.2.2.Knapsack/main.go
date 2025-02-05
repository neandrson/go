package main

import (
	"fmt"
)

type Chest struct {
	val []int
	wt  []int
}

func Knapsack(chest *Chest, maxWeight int) int {
	n := len(chest.val) // Количество драгоценностей
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, maxWeight+1)
	}

	for item := 1; item <= n; item++ { // Переберём все предметы из сундука
		for capacity := 1; capacity <= maxWeight; capacity++ {
			// Всё ниже — о рюкзаке вместимостью capacity
			maxcostWithoutCurrent := matrix[item-1][capacity] //  Максимальная стоимость предыдущих предметов
			maxcostWithCurrent := 0                           //  Для хранения максимальной стоимости, если положим текущий предмет

			weightOfCurrent := chest.wt[item-1] // Масса текущего
			if capacity >= weightOfCurrent {    // Проверяем, влезет ли текущий предмет в рюкзак
				// Если текущий влез, смотрим, что ещё взять
				maxcostWithCurrent = chest.val[item-1]                  // Сначала положим текущий предмет
				remainingCapacity := capacity - weightOfCurrent         // Проверим, осталось ли место
				maxcostWithCurrent += matrix[item-1][remainingCapacity] // Максимальная стоимость оставшегося места
			}

			matrix[item][capacity] = max(maxcostWithoutCurrent, maxcostWithCurrent) // Выбираем, нужно ли класть текущий
		}
	}
	fmt.Println(matrix)
	item := GetSolution(maxWeight)
	return matrix[n][maxWeight]
}

func GetSolution(maxWeight Weight, items []int, getWeight func(*T) Weight, getValue func(*T) Value) (selection []T) {
	// Perform the dynamic programming 0-1 knapsack algorithm.
	maxValue := make([]int, maxWeight+1)
	for m := range items {
		itemWeight := getWeight(&items[m])
		itemValue := getValue(&items[m])
		for weight := maxWeight; weight >= itemWeight; weight-- {
			maxValueWithItem := itemValue + maxValue[weight-itemWeight].maxValue
			if maxValueWithItem > maxValue[weight].maxValue {
				maxValue[weight].maxValue = maxValueWithItem
				maxValue[weight].selectedItems.Set(&maxValue[weight-itemWeight].selectedItems).Add(m)
			}
		}
	}
	selection = make([]T, maxValue[maxWeight].selectedItems.Size())[:0]
	maxValue[maxWeight].selectedItems.Visit(func(index int) bool {
		selection = append(selection, items[index])
		return false
	})
	return
}

/*func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/

func main() {
	w := 10 // Грузоподъёмность рюкзака
	chest := Chest{
		val: []int{100, 400, 300, 500}, // Стоимость
		wt:  []int{5, 4, 6, 3},         // Масса
	}
	wt := Knapsack(&chest, w)
	fmt.Println(wt)
}

/*func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	// рассчет стоимости
	n := len(chest.val)
	m := maxWeight
	dp1 := make([]int, n)
	// create dp data structure
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < len(chest.wt); i++ {
		for j := 0; j <= maxWeight; j++ {
			if chest.wt[i] > j {
				dp[i+1][j] = dp[i][j]
				select {
				case dp[i][j] == 100:
					dp1[i] = 0
				case dp[i][j] == 400:
					dp1[i] = 1
				case dp[i][j] == 300:
					dp1[i] = 2
				default:
					dp1[i] = 3
				}
				//dp1[i] = i

			} else {
				dp[i+1][j] = Max(dp[i][j-chest.wt[i]]+chest.val[i], dp[i][j])
				//dp1[i] = i
			}
		}
		//dp1[i] = i
	}
	fmt.Println(dp, dp1)
	return dp[n][m], dp1
}*/

/*func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}*/

/*func main() {
	w := 10 // Грузоподъёмность рюкзака
	chest := Chest{
		val: []int{100, 400, 300, 500}, // Стоимость
		wt:  []int{5, 4, 6, 3},         // Масса
	}
	maxProfit, item := Knapsack(&chest, w)
	fmt.Println(maxProfit, item)
}*/
