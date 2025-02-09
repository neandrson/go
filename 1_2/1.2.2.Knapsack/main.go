package main

import (
	"fmt"
	"slices"
)

type Chest struct {
	val        []int
	wt         []int
	itemsToPut []int
	cap        int
	cost       int
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

func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	// рассчет стоимости
	n := len(chest.val)
	m := maxWeight

	// create dp data structure
	dp := make([][]int, n+1)
	dp1 := make([]int, len(chest.wt))
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < len(chest.wt); i++ {
		for j := 0; j <= maxWeight; j++ {
			if chest.wt[i] > j {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = max(dp[i][j-chest.wt[i]]+chest.val[i], dp[i][j])
				if i > 0 {
					if chest.wt[i] < maxWeight-chest.wt[i-1] {
						dp1[i] = i
					}
				} else {
					dp1[i] = i
				}
			}
		}
	}
	slices.Sort(dp1)
	//fmt.Println(dp, dp1[:]) //len(chest.wt)-1
	return dp[n][m], dp1[len(chest.wt)-2:]
}

/*func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}*/

func main() {
	//w := 5 // Грузоподъёмность рюкзака
	/*chest := Chest{
		val: []int{5, 3, 4}, // Стоимость
		wt:  []int{3, 2, 1}, // Масса
	}*/

	Cases := []Chest{
		{
			val:        []int{100, 400, 300, 500},
			wt:         []int{5, 4, 6, 3},
			itemsToPut: []int{1, 3},
			cap:        10,
			cost:       900,
		},
		{
			val:        []int{5, 3, 4},
			wt:         []int{3, 2, 1},
			itemsToPut: []int{0, 2},
			cap:        5,
			cost:       9,
		},
	}
	for _, c := range Cases {

		chest := &Chest{
			val: c.val,
			wt:  c.wt,
		}
		cost, items := Knapsack(chest, c.cap)

		if cost != c.cost {
			fmt.Printf("Expected cost: %d, got cost: %d\n", c.cost, cost)
		} else {
			fmt.Printf("got cost: %d\n", cost)
		}

		slices.Sort(items)
		slices.Sort(c.itemsToPut)
		if !slices.Equal(items, c.itemsToPut) {
			fmt.Printf("Expected items: %v, got items: %v\n", c.itemsToPut, items)
		} else {
			fmt.Printf("got items: %v\n", items)
		}
	}
}
