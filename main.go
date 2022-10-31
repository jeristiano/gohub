package main

import "fmt"

func main() {
	fmt.Println("hello world 您好世界 ")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	foo := make(map[string]string)
	foo["name"] = "长泽雅美"
	foo["age"] = "34"
	foo["nationality"] = "日本国"
	foo["occupation"] = "演员"

	foo1 := make(map[string]string)
	foo1["name"] = "新垣结衣"
	foo1["age"] = "34"
	foo1["nationality"] = "日本国"
	foo1["occupation"] = "演员"

	for k, v := range foo {
		fmt.Println(k, v)
	}

	items := make([]map[string]string, 5)

	items = append(items, foo)
	items = append(items, foo1)

	for key, value := range items {
		fmt.Println(key, value)
	}

}
