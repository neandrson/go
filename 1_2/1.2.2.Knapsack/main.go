package main

import (
	"fmt"
)

type Chest struct {
	val []int
	wt  []int
}

/*func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	n := len(chest.val) // Количество драгоценностей
	matrix := make([][]int, n+1)
	dp := make([]int, n+1)
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
				maxcostWithCurrent = chest.val[item-1] // Сначала положим текущий предмет

				remainingCapacity := capacity - weightOfCurrent         // Проверим, осталось ли место
				maxcostWithCurrent += matrix[item-1][remainingCapacity] // Максимальная стоимость оставшегося места

			}

			matrix[item][capacity] = max(maxcostWithoutCurrent, maxcostWithCurrent) // Выбираем, нужно ли класть текущий
			//	dp[item] = remainingCapacity
		}
	}
	fmt.Println(matrix)

	return matrix[n][maxWeight], dp
}*/

/*func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/

func Knapsack(chest *Chest, maxWeight int) int {
	// рассчет стоимости
	n := len(chest.val)
	m := maxWeight

	// create dp data structure
	dp := make([][]int, n+1)
	dp1 := make([]int, len(chest.wt)+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < len(chest.wt); i++ {
		for j := 0; j <= maxWeight; j++ {
			if chest.wt[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i+1][j] = max(dp[i][j-chest.wt[i]]+chest.val[i], dp[i][j])
				if dp1[i] < maxWeight-chest.wt[i-1] {
					dp1[i] = i
				}
			}
		}
	}
	fmt.Println(dp, dp1)
	/*	num := len(chest.wt)
		dp1 := make([]int, maxWeight+1)

		for i := 0; i < num; i++ {
			for j := chest.wt[i]; j <= maxWeight; j++ {
				if s := dp1[j-chest.wt[i]] + chest.val[i]; dp1[j] < s {
					dp1[j] = s
				}
			}
		}*/

	//return dp[capacity]
	return dp[n][m]
}

/*func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}*/

func main() {
	w := 10 // Грузоподъёмность рюкзака
	chest := Chest{
		val: []int{100, 400, 300, 500}, // Стоимость
		wt:  []int{5, 4, 6, 3},         // Масса
	}
	maxProfit := Knapsack(&chest, w)
	fmt.Println(maxProfit)
}
