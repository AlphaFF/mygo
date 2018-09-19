/*
* @Author: alpha
* @Date:   2018-09-17 10:31:44
* @Last Modified by:   alpha
* @Last Modified time: 2018-09-18 17:20:05
 */

package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = iota
)

func main() {
	fmt.Println("world...")
	fmt.Println("hello world!")
	fmt.Println(a, b, c, d)

	if a := 1; a > 0 {
		fmt.Println("a>0")
	} else {
		fmt.Println("a<0")
	}

	// i := 0
	// Here:
	//  if i < 10 {
	//      i++
	//      fmt.Println(i)
	//  } else {
	//      fmt.Println("i > 10")
	//  }
	//  goto Here

	x := 3
	fmt.Println(x)
	x1 := add(&x)
	fmt.Println(x1)
	fmt.Println(x)

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	odd := filter(slice, isOdd)
}

// Hello test
type Hello interface {
}

func test(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func test1(arg ...int) {}

func add(a *int) int {
	*a = *a + 1
	return *a
}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
