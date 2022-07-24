package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float32
}

func main() {
	gobook := Book{
		Title:       "",
		Authors:     []string{"xushiwei", "hughlv", "pandamen", "haotuo"},
		Publisher:   "ituring.com.cn",
		IsPublished: true,
		Price:       9.99,
	}
	bytes, _ := json.Marshal(gobook)
	// 序列化对象为json字符串
	fmt.Println(string(bytes))
	// json字符串反序列为对象
	var book Book
	str := `{
    "Title": "",
    "Authors": [
        "xushiwei",
        "hughlv",
        "pandamen",
        "haotuo"
    ],
    "Publisher": "ituring.com.cn",
    "IsPublished": true,
    "Price": 9.99
}`
	// string to []byte
	data := []byte(str)
	json.Unmarshal(data, &book)
	fmt.Println(book)

	// 在不确定类型时，用interface{}解码json字符串
	var r interface{}
	json.Unmarshal(data, &r)
	fmt.Println(reflect.TypeOf(r))
	// 判断模板结构是否预期
	unknowBook, ok := r.(map[string]interface{})
	if ok {
		for k, v := range unknowBook {
			// 对 interface{} 断言类型
			switch v2 := v.(type) {
			case string:
				fmt.Println("is string", v2)
			case int:
				fmt.Println("is int", v2)
			case bool:
				fmt.Println("is bool", v2)
			case []interface{}:
				fmt.Println("is interface", v2)
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "not found handler")

			}
		}
	}
}
