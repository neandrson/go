package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if len(line) < 10 {
			continue
		}
		timestampStr := line[:10] // Предполагаем, что дата всегда занимает первые 10 символов строки
		timestamp, err := time.Parse("02.01.2006", timestampStr)
		if err != nil {
			continue
		}

		if (timestamp.After(start) || timestamp.Equal(start)) && (timestamp.Before(end) || timestamp.Equal(end)) {
			logs = append(logs, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func main() {
	start := time.Date(2022, 12, 13, 0, 0, 0, 0, time.UTC)
	end := time.Date(2022, 12, 15, 0, 0, 0, 0, time.UTC)

	logs, err := ExtractLog("logfile.txt", start, end)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, log := range logs {
		fmt.Println(log)
	}
}
