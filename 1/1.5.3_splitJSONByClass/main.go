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
	var data []map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	result := make(map[string][]byte)
	for _, item := range data {
		classObj, ok := item["class"]
		if !ok {
			return nil, fmt.Errorf(`параметр "class" отсутствует`)
		}
		class, ok := classObj.(string)
		if !ok {
			return nil, fmt.Errorf(`параметр "class" не является строкой`)
		}

		jsonBytes, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("ошибка создания JSON: %w", err)
		}

		if existingJSON, found := result[class]; found {
			result[class] = append(existingJSON, []byte(",")...)
			result[class] = append(result[class], jsonBytes...)
		} else {
			result[class] = append([]byte{'['}, jsonBytes...)
			continue
		}
		result[class] = append(result[class], ']')
	}

	return result, nil
}

func main() {
	jsonData := []byte(`[{
            "name": "Oleg",
            "class": "9B"
        },
        {
            "name": "Ivan",
            "class": "9A"
        },
        {
            "name": "Maria",
            "class": "9B"
        },
        {
            "name": "John",
            "class": "9A"
        }
	]`)

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
