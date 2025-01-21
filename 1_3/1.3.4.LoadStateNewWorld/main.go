package main

import (
	"fmt"
	"io"
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
	var neighbors int

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

func (w *World) String() {
	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"
	for _, row := range w.Cells {
		for _, cell := range row {
			switch {
			case cell:
				//fmt.Printf("true")
				fmt.Printf(greenSquare)
				//fmt.Sprint("true")
			default:
				//fmt.Printf("false")
				fmt.Printf(brownSquare)
				//fmt.Sprint("false")
			}
			//fmt.Printf("%t", cell)
		}
		fmt.Printf("\n")
	}
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
	i := 0
	for _, row := range w.Cells {
		for _, cell := range row {
			var value string
			if cell {
				value = "1"
			} else {
				value = "0"
			}
			//err = binary.Write(file, binary.LittleEndian, value)
			//if err != nil {
			//	return err
			//}
			file.WriteString(value)
		}
		i++
		if i < len(w.Cells) {
			file.WriteString("\n")
		}
	}
	return nil
}

func (w *World) LoadState(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Читаеем высоту и ширину сетки в файл
	data := make([]byte, 64)
	width := 0
	for {
		_, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		//fmt.Print(string(data[:_]))
	}
	// количество строк
	height := strings.Count(string(data), "\n")

	// количество столбцов
	widths := 0
	for _, row := range data {
		if string(row) != "\n" {
			width++
		} else {
			widths = width
		}
	}
	if width != widths {
		return err
	}

	// создание сетки
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // Создаём новый слайс в каждой строке
	}
	w.Cells = cells

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
	//currentWorld.Seed()
	/*err := currentWorld.SaveState("error.log")
	if err != nil {
		panic(err)
	}*/
	for { // Цикл для вывода каждого состояния
		if i > 1 {
			break
		}
		// Загрузка состояния из файла
		err = currentWorld.LoadState("error.log")
		if err != nil {
			panic(err)
		}
		// Выведем текущее состояние на экран
		//fmt.Println(currentWorld)
		currentWorld.String()
		// сохранение текущего состояния сетки в файл

		// Рассчитываем следующее состояние
		NextState(currentWorld, nextWorld)
		// Изменяем текущее состояние
		currentWorld = nextWorld
		err := currentWorld.SaveState("error.log")
		if err != nil {
			panic(err)
		}
		// Делаем паузу
		time.Sleep(10000 * time.Millisecond)
		// Специальная последовательность для очистки экрана после каждого шага
		//fmt.Print("\033[H\033[2J")
		i++
	}
}
