package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return nil, err
	}

	for i := range students {
		students[i].Grade++
	}

	updatedData, err := json.Marshal(students)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}

func main() {
	jsonData := []byte(`[{"name": "Alice", "grade": 85}]`)
	modifiedData, err := modifyJSON(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(modifiedData))
}
