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
	/*var students []map[string]interface{}
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]byte)

	for _, student := range students {
		class := student["class"].(string)
		if _, ok := result[class]; !ok {
			result[class] = []byte(string(jsonData))
		}
		//fmt.Println(string(result[class]))
		//else {
		//	result[class] = append(result[class], ',')
		//	result[class] = append(result[class], jsonData...)
		//}
	}
	for class, data := range result {
		result[class] = []byte(string(data))
		fmt.Println(string(result[class]))
	}*/

	var data []map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]byte)
	result[class] = append(result[class], "[")
	for _, item := range data {
		class := item["class"].(string)
		jsonBytes, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		//fmt.Println(string(jsonBytes))
		result[class] = append(result[class], jsonBytes...)
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
		fmt.Printf("%s:", class)
		for _, jsonData := range jsonList {
			fmt.Print(string(jsonData))
		}
	}
}
