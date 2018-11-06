/*
* @Author: alpha
* @Date:   2018-09-17 10:31:44
* @Last Modified by:   AlphaFF
* @Last Modified time: 2018-09-20 16:39:16
 */

// package main

// import "fmt"

// const (
// 	a = iota
// 	b
// 	c
// )

// const (
// 	d = iota
// )

// func main() {
// 	fmt.Println("world...")
// 	fmt.Println("hello world!")
// 	fmt.Println(a, b, c, d)

// 	if a := 1; a > 0 {
// 		fmt.Println("a>0")
// 	} else {
// 		fmt.Println("a<0")
// 	}

// 	// i := 0
// 	// Here:
// 	//  if i < 10 {
// 	//      i++
// 	//      fmt.Println(i)
// 	//  } else {
// 	//      fmt.Println("i > 10")
// 	//  }
// 	//  goto Here

// 	x := 3
// 	fmt.Println(x)
// 	x1 := add(&x)
// 	fmt.Println(x1)
// 	fmt.Println(x)

// 	slice := []int{1, 2, 3, 4, 5, 6, 7}
// 	odd := filter(slice, isOdd)
// 	fmt.Println(odd)

// 	var inter interface{}
// 	i := 5
// 	s := "app"
// 	inter = i
// 	fmt.Println(inter)
// 	inter = s
// 	fmt.Println(inter)
// }

// // Hello test
// type Hello interface {
// }

// func test(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func test1(arg ...int) {}

// func add(a *int) int {
// 	*a = *a + 1
// 	return *a
// }

// type testInt func(int) bool

// func isOdd(integer int) bool {
// 	if integer%2 == 0 {
// 		return false
// 	}
// 	return true
// }

// func isEven(integer int) bool {
// 	if integer%2 == 0 {
// 		return true
// 	}
// 	return false
// }

// func filter(slice []int, f testInt) []int {
// 	var result []int
// 	for _, value := range slice {
// 		if f(value) {
// 			result = append(result, value)
// 		}
// 	}
// 	return result
// }

package test

import (
	"fmt"
	"runtime"
	// "reflect"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name:)" + p.name + "-age:" + strconv.Itoa(p.age) + "years"
}

func main() {
	// list := make(List, 3)
	// list[0] = 1
	// list[1] = "hello"
	// list[2] = Person{"Jhon", 70}
	// for index, element := range list {
	// 	if value, ok := element.(int); ok {
	// 		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	// 	} else if value, ok := element.(string); ok {
	// 		fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
	// 	} else if value, ok := element.(Person); ok {
	// 		fmt.Printf("list[%d] is an Person and its value is %s\n", index, value)
	// 	} else {
	// 		fmt.Printf("list[%d] is a different type\n", index)
	// 	}
	// }
	// for index, ele := range list {
	// 	switch value := ele.(type) {
	// 	case int:
	// 		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	// 	case string:
	// 		fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
	// 	case Person:
	// 		fmt.Printf("list[%d] is an Person and its value is %s\n", index, value)
	// 	default:
	// 		fmt.Printf("list[%d] is a different type\n", index)
	// 	}
	// }
	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Println(v.Kind() == reflect.Float64)
	// fmt.Println(v.Float())

	// p := reflect.Valueof(&x)
	// v := p.Elem()
	// v.setFloat(7.1)
	// fmt.Println(v, x)

	// go say("world")
	// say("hello")

	// a := []int{7, 2, 8, -9, 4, 0}
	// c := make(chan int, 2)
	// go sum(a[:len(a)/2], c)
	// go sum(a[len(a)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)

	// c := make(chan int, 2)
	// c <- 1
	// c <- 2
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// c := make(chan int)
	// go fibonacci(10, c)
	// for i := range c {
	// 	fmt.Println(i)
	// }

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

// func fibonacci(n int, c chan int) {
// 	x, y := 1, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
