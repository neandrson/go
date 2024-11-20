package main

import (
	"encoding/json"
	"fmt"
)

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var data []map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]byte)

	for _, item := range data {
		class := item["class"].(string)
		jsonBytes, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		result[class] = append(result[class], jsonBytes)
	}

return result, nil
}

func main() { jsonData := []byte([ {"class": "A", "name": "Alice"}, {"class": "B", "name": "Bob"}, {"class": "A", "name": "Alex"} ])

result, err := splitJSONByClass(jsonData)
if err != nil {
	fmt.Println("Error:", err)
	return
}

for class, jsonList := range result {
	fmt.Printf("Class: %s\n", class)
	for _, jsonData := range jsonList {
		fmt.Println(string(jsonData))
	}
}
}