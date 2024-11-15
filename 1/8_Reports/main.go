package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uint64
	Name  string
	Email string
	Age   uint8
}

type Report struct {
	u        User
	ReportID uuid.UUID
	Date     time.Time
}

func CreateReport(user User, reportDate string) Report {
	t, err := time.Parse("02-01-2006", reportDate)
	if err != nil {
		log.Fatal(err)
	}

	return Report{
		u:        user,
		ReportID: uuid.New(),
		Date:     t,
	}
}

func PrintReport(report Report) {
	fmt.Println("ReportID  :", report.ReportID.String())
	fmt.Println("ReportDate:", report.Date.Format(time.RFC1123))
	fmt.Println("ID        :", report.u.ID)
	fmt.Println("Name      :", report.u.Name)
	fmt.Println("Age       :", report.u.Age)
	fmt.Println("Email     :", report.u.Email)
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

	reports := GenerateUserReports(users, "07-11-2023")

	for _, r := range reports {
		PrintReport(r)
		fmt.Println()
	}
}
