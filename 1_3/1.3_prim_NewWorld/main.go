package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

func NewWorld(height, width int) *World {
	// Создаём тип World с количеством слайсов hight (количество строк)
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // Создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

func (w *World) Next(x, y int) bool {
	n := w.Neighbors(x, y)       // Получим количество живых соседей
	alive := w.Cells[y][x]       // Текущее состояние клетки
	if n < 4 && n > 1 && alive { // Если соседей двое или трое, а клетка жива,
		return true // то следующее её состояние — жива
	}
	if n == 3 && !alive { // Если клетка мертва, но у неё трое соседей,
		return true // клетка оживает
	}

	return false // В любых других случаях — клетка мертва
}

func (w *World) Neighbors(x, y int) int {
	var neighbors int

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if w.Alive(j, i) {
				neighbors++
			}
		}
	}
	return neighbors
}

func (w World) Alive(x, y int) bool {
	y = (w.Height + y) % w.Height
	x = (w.Width + x) % w.Width
	return w.Cells[y][x]
}
func NextState(oldWorld, newWorld *World) {
	// Переберём все клетки, чтобы понять, в каком они состоянии
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// Для каждой клетки получим новое состояние
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

func (w *World) Seed() {
	// Снова переберём все клетки
	for _, row := range w.Cells {
		for i := range row {
			//rand.Intn(10) возвращает случайное число из диапазона	от 0 до 9
			if rand.Intn(10) == 1 {
				row[i] = true
			}
		}
	}
}

func main() {
	// Зададим размеры сетки
	height := 10
	width := 10
	// Объект для хранения текущего состояния сетки
	currentWorld := NewWorld(height, width)
	// Объект для хранения следующего состояния сетки
	nextWorld := NewWorld(height, width)
	// Установим начальное состояние
	currentWorld.Seed()
	for { // Цикл для вывода каждого состояния
		// Выведем текущее состояние на экран
		fmt.Println(currentWorld)
		// Рассчитываем следующее состояние
		NextState(currentWorld, nextWorld)
		// Изменяем текущее состояние
		currentWorld = nextWorld
		// Делаем паузу
		time.Sleep(100 * time.Millisecond)
		// Специальная последовательность для очистки экрана после каждого шага
		fmt.Print("\033[H\033[2J")
	}
}
