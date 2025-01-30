package main

import "fmt"

type Chest struct {
	val []int
	wt  []int
}

func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	n := len(chest.val) // Количество драгоценностей
	k := 0
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
				matrix[i][j] = max(matrix[i][j], matrix[i-1][j-n*chest.wt[i]]+n*chest.val[i])
			}
		}
	}
	fmt.Println("the dp array is", matrix)
	return k, matrix[len(chest.wt)-1][maxWeight]
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
	cost, items := Knapsack(&chest, w)
	fmt.Println(cost, items)
}
