package main

import "fmt"

// Определяем интерфейс Vehicle с полем Type
type Vehicle interface {
	CalculateTravelTime(distance float64) float64
	GetType() string
}

// Создаем структуру Car (автомобиль)
type Car struct {
	Speed    float64
	Type     string
	FuelType string
}

// Создаем структуру Motorcycle (мотоцикл)
type Motorcycle struct {
	Speed float64
	Type  string
}

// Реализуем метод CalculateTravelTime() для Car
func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

// Реализуем метод CalculateTravelTime() для Motorcycle
func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

// Реализуем метод GetType() для Car
func (c Car) GetType() string {
	return "main.Car" //c.Type
}

// Реализуем метод GetType() для Motorcycle
func (m Motorcycle) GetType() string {
	return "main.Motorcycle" //m.Type
}

// Функция EstimateTravelTime
func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	travelTimes := make(map[string]float64)

	for _, vehicle := range vehicles {
		time := vehicle.CalculateTravelTime(distance)
		travelTimes[vehicle.GetType()] = time
	}

	return travelTimes
}

func main() {
	// Создаем несколько транспортных средств
	car := Car{Speed: 60.0, Type: "Седан", FuelType: "Бензин"}
	motorcycle := Motorcycle{Speed: 80.0, Type: "Мотоцикл"}

	// Создаем список транспортных средств
	vehicles := []Vehicle{car, motorcycle}

	// Расстояние до пункта назначения
	distance := 200.0

	// Получаем карту с временем в пути для каждого транспортного средства
	travelTimes := EstimateTravelTime(vehicles, distance)

	// Выводим результат
	for vehicleType, time := range travelTimes {
		fmt.Printf("%s: Время в пути %.2f часа\n", vehicleType, time)
	}
}
