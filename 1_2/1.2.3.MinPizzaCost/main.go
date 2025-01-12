package main

import "fmt"

func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
	for _, coin := range coins {
		// Переберём все значения слайса с количеством способов
		for j := range ways {
			// Если монета меньше значения, которое нам нужно получить
			if coin <= j {
				// Обновим значение с учётом предыдущих расчётов
				ways[j] += ways[j-coin]
			}
		}
	}
	fmt.Println(ways[N])
}

func main() {

}
