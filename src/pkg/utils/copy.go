package utils

import (
	"encoding/json"
)

func Copy(dest interface{}, src interface{}) {
	data, _ := json.Marshal(src)
	_ = json.Unmarshal(data, dest)
}

func UniqueValues(arr []int) []int {
	unique := make(map[int]bool)
	var result []int

	for _, v := range arr {
		if !unique[v] {
			unique[v] = true
			result = append(result, v)
		}
	}

	return result
}
