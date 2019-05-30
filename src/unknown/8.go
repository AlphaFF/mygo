package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{"1": 10, "2": 20}
	fmt.Println(m)
	m["1"] = 30
	fmt.Println(m)
	if v, ok := m["1"]; ok {
		fmt.Println(v, m["1"])
	}
}
