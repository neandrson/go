package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type User struct {
	ID    uint64
	Name  string
	Email string
	Age   uint8
}

type Report struct {
	u        User
	ID       uint64
	Name     string
	Email    string
	Age      uint8
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	t, err := time.Parse("2006-01-02", reportDate)
	if err != nil {
		log.Fatal(err)
	}
	i, _ := strconv.Atoi(GenerateTimestampID())

	return Report{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Age:      user.Age,
		ReportID: i,
		Date:     t.Format("2006-01-02"),
	}
}

func PrintReport(report Report) {
	fmt.Println("ReportID  :", report.ReportID)
	fmt.Println("ReportDate:", report.Date)
	fmt.Println("ID        :", report.ID)
	fmt.Println("Name      :", report.Name)
	fmt.Println("Age       :", report.Age)
	fmt.Println("Email     :", report.Email)
}

func GenerateTimestampID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func GenerateUserReports(users []User, reportDate string) []Report {
	resultReport := make([]Report, 0, len(users))

	for _, user := range users {
		resultReport = append(resultReport, CreateReport(user, reportDate))
	}
	return resultReport
}

func main() {
	users := []User{
		{0, "Bill", "bill@yandex.ru", 33},
		{1, "Jane", "jane@mail.ru", 24},
		{2, "Tommy", "tommy@gmail.com", 37},
	}

	reports := GenerateUserReports(users, "2023-11-07")

	for _, r := range reports {
		PrintReport(r)
		fmt.Println()
	}
}
