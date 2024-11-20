package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var students []Student
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return nil, err
	}

	classMap := make(map[string][]Student)
	result := make(map[string][]byte)

	for _, student := range students {
		fmt.Println(student)
		classMap[student.Class] = append(classMap[student.Class], student)
		fmt.Println(classMap[student.Class])
	}

	for class, students := range classMap {
		jsonStudents, err := json.Marshal(students)
		if err != nil {
			return nil, err
		}
		result[class] = jsonStudents
	}

	return result, nil
}

func main() {
	jsonData := []byte(`[{"class": "A", "name": "Alice"}, {"class": "B", "name": "Bob"}, {"class": "A", "name": "Alex"}]`)

	result, err := splitJSONByClass(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for class, jsonList := range result {
		fmt.Printf("Class: %s", class)
		for _, jsonData := range jsonList {
			fmt.Print(string(jsonData))
		}
	}
}
