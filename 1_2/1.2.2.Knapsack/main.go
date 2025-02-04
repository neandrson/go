package main

import (
	"fmt"
	"math"
)

type Chest struct {
	val []int
	wt  []int
}

/*func Knapsack(chest *Chest, maxWeight int) (int, int) {
	n := len(chest.val) // Количество драгоценностей
	//k := 0
	// Выделим память под слайсы
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, maxWeight+1)
	}

	for i := chest.wt[0]; i <= maxWeight; i++ {
		matrix[0][i] = chest.val[0]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j <= maxWeight; j++ {
			for n := 0; n*chest.wt[i] <= j; n++ {
				if chest.wt[i-1] <= j {
					matrix[i][j] = max(matrix[i-1][j-n*chest.wt[i]]+n*chest.val[i], matrix[i][j])
				} else {
					matrix[i][j] = matrix[i-1][j]
				}
			}
		}
	}
	fmt.Println("the dp array is", matrix)
	return chest.wt[len(chest.wt)-1], matrix[len(chest.wt)][maxWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	w := 10 // Грузоподъёмность рюкзака
	chest := Chest{
		val: []int{100, 400, 300, 500}, // Стоимость
		wt:  []int{5, 4, 6, 3},         // Масса
	}
	wt, les := Knapsack(&chest, w)
	fmt.Println(wt, les)
}*/

func Knapsack(chest *Chest, maxWeight int) (int, []int) {
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
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	w := 10 // Грузоподъёмность рюкзака
	chest := Chest{
		val: []int{100, 400, 300, 500}, // Стоимость
		wt:  []int{5, 4, 6, 3},         // Масса
	}
	maxProfit, item := Knapsack(&chest, w)
	fmt.Println(maxProfit, item)
}
