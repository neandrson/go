package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
	// Пример New World
	/*var neighbors int

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if i < 0 || j < 0 {
				continue
			} else if w.Alive(j, i) {
				neighbors++
			}
		}
	}
	return neighbors*/

	// Пример 2
	var n int
	y_prev := (y - 1 + w.Height) % w.Height
	x_prev := (x - 1 + w.Width) % w.Width
	y_next := (y + 1) % w.Height
	x_next := (x + 1) % w.Width

	if w.Cells[y_prev][x_prev] {
		n++
	}
	if w.Cells[y_prev][x] {
		n++
	}
	if w.Cells[y_prev][x_next] {
		n++
	}
	if w.Cells[y][x_prev] {
		n++
	}
	if w.Cells[y][x_next] {
		n++
	}
	if w.Cells[y_next][x_prev] {
		n++
	}
	if w.Cells[y_next][x] {
		n++
	}
	if w.Cells[y_next][x_next] {
		n++
	}

	return n
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

func (w *World) String() string {
	var result string

	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"

	for i := range w.Cells {
		for _, col := range w.Cells[i] {
			if col {
				result += greenSquare
			} else {
				result += brownSquare
			}
		}
		result += "\n"
	}

	return result
}

func (w *World) SaveState(filename string) error {
	/*inputFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inputFile.Close()*/

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем высоту и ширину сетки в файл
	/*err = binary.Write(file, binary.LittleEndian, int32(w.Height))
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, int32(w.Width))
	if err != nil {
		return err
	}*/

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := range w.Cells {
		arr, end := []string{}, "\n"

		for j := range w.Cells[i] {
			if w.Cells[i][j] == true {
				arr = append(arr, "1")
			} else {
				arr = append(arr, "0")
			}
		}

		row := strings.Join(arr, "")

		if i == len(w.Cells)-1 {
			end = ""
		}

		fmt.Fprint(writer, row+end)
	}

	return nil
}

func (w *World) LoadState(filename string) error {
	// Загрузка состояния игры из файла
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	new_cells := [][]bool{}

	file_scanner := bufio.NewScanner(file)
	for file_scanner.Scan() {
		col := []bool{}
		for _, let := range file_scanner.Text() {
			if string(let) == "1" {
				col = append(col, true)
			} else {
				col = append(col, false)
			}
		}

		new_cells = append(new_cells, col)

		index := len(new_cells) - 1
		if index > 0 && len(new_cells[index]) != len(new_cells[index-1]) {
			return fmt.Errorf("Different count")
		}
	}

	w.Cells = new_cells
	w.Height, w.Width = len(new_cells), len(new_cells[0])

	return nil
}

func main() {
	// Зададим размеры сетки
	height := 3
	width := 3
	i := 0 // Количество циклов
	// Объект для хранения текущего состояния сетки
	currentWorld := NewWorld(height, width)
	// Объект для хранения следующего состояния сетки
	nextWorld := NewWorld(height, width)
	// Установим начальное состояние
	currentWorld.Seed()
	// сохранение текущего состояния сетки в файл
	err := currentWorld.SaveState("error.log")
	if err != nil {
		panic(err)
	}
	for { // Цикл для вывода каждого состояния
		if i > 1 {
			break
		}
		// Загрузка состояния из файла
		err := currentWorld.LoadState("error.log")
		if err != nil {
			panic(err)
		}
		// Выведем текущее состояние на экран
		//fmt.Println(currentWorld)
		fmt.Print(currentWorld.String())

		// Рассчитываем следующее состояние
		NextState(currentWorld, nextWorld)
		// Изменяем текущее состояние
		currentWorld = nextWorld

		// сохранение текущего состояния сетки в файл
		err = currentWorld.SaveState("error.log")
		if err != nil {
			panic(err)
		}
		// Делаем паузу
		time.Sleep(1000 * time.Millisecond)
		// Специальная последовательность для очистки экрана после каждого шага
		fmt.Print("\033[H\033[2J")
		i++
	}
}
