/*
* @Author: wangfeng1
* @Date:   2019-03-01 16:33:18
* @Last Modified by:   wangfeng1
* @Last Modified time: 2019-03-04 14:16:48
 */

package main

import (
	"fmt"
)

func main() {
	greeting("wang", "feng", "zhang", "qi")
	greeting(person())
	fmt.Println("this is fmt print.")
	n := 0
	reply := &n
	multiply(10, 5, reply)
	fmt.Println("multiply:", *reply)
	fmt.Println("n:", n)
	var p *int
	i := 42
	p = &i
	fmt.Println(*p)
	*p = 12
	fmt.Println(i)
	// defer fmt.Println("finally...")
	fmt.Println("output before defer...")
	for i := 0; i < 5; i++ {
		fmt.Println("this is number: ", i)
	}
	var b = [10]int{1, 2, 3, 4}
	fmt.Println(b)
	var c = [...]string{"1", "2"}
	fmt.Println(c)
	var d [3]int
	fmt.Println(d)
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
	s := []int{1, 2, 3, 4}
	fmt.Println(s, len(s[:2]), cap(s))
	s1 := make([]int, 10)
	s2 := make([]string, 4, 6)
	fmt.Println(s1, s2)
	var s3 []int
	if s3 == nil {
		fmt.Println("this is a nil.", s3)
	}
	s4 := make([]int, 5)
	fmt.Println(s4)
	s5 := new([]int)
	fmt.Println(s5)

	data := make(chan int)
	go func() {
		// for d := range data {
		//  fmt.Println("this is now:", d)
		// }
		for {
			if d, ok := <-data; ok {
				fmt.Println("now is:", d)
			}
		}
	}()
	data <- 1
	data <- 2
	data <- 5
	close(data)
	s6 := make([]int, 6)
	s6 = append(s6, 7, 8)
	fmt.Println(s6)
}

func greeting(s string, t ...string) {
	fmt.Println("greeting:", s)
	for _, v := range t {
		fmt.Println("greeting:", v)
	}
}

func person() string {
	return "Handle"
}

func multiply(a, b int, reply *int) {
	*reply = a * b
}

func fp(a *[3]int) { fmt.Println(*a) }
