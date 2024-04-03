package utils

import "encoding/json"

func Copy(src interface{}, dest interface{}) {
	data, _ := json.Marshal(src)
	_ = json.Unmarshal(data, dest)
}
