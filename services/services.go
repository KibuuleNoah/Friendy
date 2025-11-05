package services

import "encoding/json"

func StructToMap(src any, dst *map[string]any){
    data, _ := json.Marshal(src)
    json.Unmarshal(data, dst)
}
