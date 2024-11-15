package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateReport(t *testing.T) {
	user := User{ID: 1, Name: "Иван", Email: "ivan@example.com", Age: 30}
	reportDate := time.Now().Format("2006-01-02")

	report := CreateReport(user, reportDate)
	fmt.Println(report)
	// Проверяем, что поля в отчете заполнены корректно
	if report.ID != user.ID {
		t.Errorf("Ожидается ID пользователя %d, получено %d", user.ID, report.ID)
	}
	if report.Name != user.Name {
		t.Errorf("Ожидается имя пользователя %s, получено %s", user.Name, report.Name)
	}
	if report.Email != user.Email {
		t.Errorf("Ожидается Email пользователя %s, получено %s", user.Email, report.Email)
	}
	if report.Age != user.Age {
		t.Errorf("Ожидается возраст пользователя %d, получено %d", user.Age, report.Age)
	}
	if report.ReportID.String() == "" {
		t.Errorf("Ожидается положительное значение, получено %d", report.ReportID.String())
	}
	_, err := time.Parse("02-01-2006", reportDate)
	if err != nil {
		t.Errorf("Неверная дата", err)
	}
}
