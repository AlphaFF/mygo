/*
* @Author: wangfeng1
* @Date:   2019-01-11 19:32:03
* @Last Modified by:   wangfeng1
* @Last Modified time: 2019-01-12 09:55:19
 */

package main

import "fmt"

func badCall() {
	e := "something is wrong."
	panic(e)
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panic: %s ==== \r\n", e)
		}
	}()
	badCall()
	fmt.Println("after badCall")
}

func main() {
	// fmt.Println("start")
	// panic("a server error occurred.")
	// fmt.Println("ending the program.")
	fmt.Println("calling test")
	test()
	fmt.Println("test completed")
}
