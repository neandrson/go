package main

import (
	"encoding/json"
	"fmt"
)

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	mergedData := make([]map[string]interface{}, 0)

	for _, jsonData := range jsonDataList {
		var tempData []map[string]interface{}
		err := json.Unmarshal(jsonData, &tempData)
		if err != nil {
			return nil, err
		}

		/*for key, value := range tempData {
			mergedData[key] = value
		}*/
		mergedData = append(mergedData, tempData...)
	}

	mergedJSONData, err := json.Marshal(mergedData)
	if err != nil {
		return nil, err
	}

	return mergedJSONData, nil
}

func main() {
	jsonData1 := []byte(`[{"name": "Alice", "age": 30}]`)
	jsonData2 := []byte(`[{"city": "New York", "country": "USA"}]`)

	mergedJSON, err := mergeJSONData(jsonData1, jsonData2)
	if err != nil {
		fmt.Println("Error merging JSON data:", err)
		return
	}

	fmt.Println(string(mergedJSON))
}
