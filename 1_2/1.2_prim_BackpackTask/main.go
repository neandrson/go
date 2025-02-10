package main

import "fmt"

// Сундук с драгоценностями
type Chest struct {
	cost []int // Стоимость предметов
	mass []int // Масса предметов
}

func Knapsack(chest *Chest, maxWeight int) int {
	n := len(chest.cost) // Количество драгоценностей
	// Выделим память под слайсы
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, maxWeight+1)
	}

	for item := 1; item <= n; item++ { // Переберём все предметы из сундука
		for capacity := 1; capacity <= maxWeight; capacity++ {
			// Всё ниже — о рюкзаке вместимостью capacity
			maxcostWithoutCurrent := matrix[item-1][capacity] //  Максимальная стоимость предыдущих предметов
			maxcostWithCurrent := 0                           //  Для хранения максимальной стоимости, если положим текущий предмет

			weightOfCurrent := chest.mass[item-1] // Масса текущего
			if capacity >= weightOfCurrent {      // Проверяем, влезет ли текущий предмет в рюкзак
				// Если текущий влез, смотрим, что ещё взять
				maxcostWithCurrent = chest.cost[item-1]                 // Сначала положим текущий предмет
				remainingCapacity := capacity - weightOfCurrent         // Проверим, осталось ли место
				maxcostWithCurrent += matrix[item-1][remainingCapacity] // Максимальная стоимость оставшегося места
			}

			matrix[item][capacity] = max(maxcostWithoutCurrent, maxcostWithCurrent) // Выбираем, нужно ли класть текущий
		}
	}
	return matrix[n][maxWeight]
}

func main() {
	w := 10 // Грузоподъёмность рюкзака
	chest := Chest{
		cost: []int{100, 400, 300, 500}, // Стоимость
		mass: []int{5, 4, 6, 3},         // Масса
	}
	a := Knapsack(&chest, w)
	fmt.Println(a)
}
