package main

import (
	"fmt"
	"time"
)

func testVariableArgs(a ...int) {
	fmt.Println(a)
}

func testVariableArgs1(b ...interface{}) {
	for _, value := range b {
		switch v := value.(type) {
		case int:
			fmt.Println("this is int:", v, value)
		case string:
			fmt.Println("this is string:", v, value)
		case bool:
			fmt.Println("this is bool:", v, value)
		default:
			fmt.Println("this is no.", v, value)
		}
	}
}

func fibonacci(n int) (res uint64) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// 闭包, 没看懂.
func fibonacci1() func() int {
	a, b := -1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

const LIM = 41

var fibs [LIM]uint64

// 通过内存缓存来替身性能
func fibonacci2(n int) (res uint64) {
	fmt.Println("fibs:", fibs)
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res
	return
}

func main() {
	fmt.Println("hello world.")
	t := []int{3, 4, 5}
	testVariableArgs(t...)
	testVariableArgs(12, 24, 25, 46)
	testVariableArgs1(12, "13", true)
	fmt.Println(fibonacci(4))
	s := fibonacci1()
	start := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println(s())
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)

	var result uint64 = 0
	start = time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci2(i)
		fmt.Println("result:", result)
	}
	end = time.Now()
	delta = end.Sub(start)
	print("delta:", delta)
}
