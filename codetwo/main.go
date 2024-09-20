package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type JSONValue map[string]interface{}
type JSONList []interface{}
type JSONMap map[string]interface{}

func sanitizeString(s string) string {
	return strings.TrimSpace(s)
}

func transformValue(value JSONValue) (interface{}, bool) {
	for keyType, valueData := range value {
		switch keyType {
		case "S":
			strVal := sanitizeString(valueData.(string))
			if strVal == "" {
				return nil, false
			}
			if t, err := time.Parse(time.RFC3339, strVal); err == nil {
				return t.Unix(), true
			}
			return strVal, true
		case "N":
			strVal := sanitizeString(valueData.(string))
			strVal = strings.TrimLeft(strVal, "0")
			if strVal == "" {
				return nil, false
			}
			if numVal, err := strconv.ParseFloat(strVal, 64); err == nil {
				return numVal, true
			}
			return nil, false
		case "BOOL":
			strVal := sanitizeString(valueData.(string))
			switch strings.ToLower(strVal) {
			case "1", "t", "true":
				return true, true
			case "0", "f", "false":
				return false, true
			}
			return nil, false
		case "NULL":
			strVal := sanitizeString(valueData.(string))
			switch strings.ToLower(strVal) {
			case "1", "t", "true":
				return nil, true
			case "0", "f", "false":
				return nil, false
			}
			return nil, false
		case "L":
			list, ok := valueData.([]interface{})
			if !ok {
				return nil, false
			}
			result := make([]interface{}, 0)
			for _, item := range list {
				itemValue, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if val, valid := transformValue(itemValue); valid {
					result = append(result, val)
				}
			}
			if len(result) == 0 {
				return nil, false
			}
			return result, true
		case "M":
			mapData, ok := valueData.(map[string]interface{})
			if !ok {
				return nil, false
			}
			result := make(JSONMap)
			for k, v := range mapData {
				k = sanitizeString(k)
				if k == "" {
					continue
				}
				mapValue, ok := v.(map[string]interface{})
				if !ok {
					continue
				}
				if val, valid := transformValue(mapValue); valid {
					result[k] = val
				}
			}
			if len(result) == 0 {
				return nil, false
			}
			return result, true
		}
	}
	return nil, false
}

func main() {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading data.json:", err)
		return
	}

	var input map[string]JSONValue
	if err := json.Unmarshal(fileData, &input); err != nil {
		fmt.Println("Error parsing input JSON:", err)
		return
	}

	output := make([]JSONMap, 0)
	transformed := make(JSONMap)
	for key, value := range input {
		key = sanitizeString(key)
		if key == "" {
			continue
		}
		if val, valid := transformValue(value); valid {
			transformed[key] = val
		}
	}
	output = append(output, transformed)

	result, _ := json.MarshalIndent(output, "", "  ")
	fmt.Println(string(result))
}
