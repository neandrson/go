package main

import (
	"testing"
)

func TestEstimateTravelTime(t *testing.T) {
	car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
	motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}

	vehicles := []Vehicle{car, motorcycle}

	distance := 200.0

	travelTimes := EstimateTravelTime(vehicles, distance)
	//fmt.Println(travelTimes)
	expectedCarTime := distance / car.Speed
	if travelTimes["main.Car"] != expectedCarTime {
		t.Errorf("Ожидается время для автомобиля %.2f часа, получено %.2f", expectedCarTime, travelTimes["main.Car"])
	}

	expectedMotorcycleTime := distance / motorcycle.Speed
	if travelTimes["main.Motorcycle"] != expectedMotorcycleTime {
		t.Errorf("Ожидается время для мотоцикла %.2f часа, получено %.2f", expectedMotorcycleTime, travelTimes["main.Motorcycle"])
	}
}
