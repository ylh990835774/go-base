package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Data 一个测试的结构体
type Data struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	// init Data Array
	var marshalDataArray []Data
	for i := 0; i < 5; i++ {
		tmpData := Data{
			ID:   i,
			Name: "Hello-" + strconv.Itoa(i),
		}
		marshalDataArray = append(marshalDataArray, tmpData)
	}

	// [{0 Hello-0} {1 Hello-1} {2 Hello-2} {3 Hello-3} {4 Hello-4}]
	fmt.Println("\n---marshalDataArray---")
	fmt.Println(marshalDataArray)

	marshalDataJsonByte, err := json.Marshal(marshalDataArray)
	if err != nil {
		fmt.Printf("json.Marshal error: %v\n", err)
		return
	}
	marshalDataJsonStr := string(marshalDataJsonByte)
	// [{"id":0,"name":"Hello-0"},{"id":1,"name":"Hello-1"},{"id":2,"name":"Hello-2"},{"id":3,"name":"Hello-3"},{"id":4,"name":"Hello-4"}]
	fmt.Println("\n---marshalData---")
	fmt.Println(marshalDataJsonStr)

	var unmarshalData []Data // 这里的类型要和 marshalDataArray 一致
	err = json.Unmarshal(marshalDataJsonByte, &unmarshalData)
	if err != nil {
		fmt.Printf("json.Unmarshal error: %v\n", err)
	}
	// [{0 Hello-0} {1 Hello-1} {2 Hello-2} {3 Hello-3} {4 Hello-4}]
	fmt.Println("\n---unmarshalData---")
	fmt.Println(unmarshalData)
}
