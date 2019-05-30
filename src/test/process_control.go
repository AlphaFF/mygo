/*
* @Author: wangfeng1
* @Date:   2019-01-12 09:57:33
* @Last Modified by:   wangfeng1
* @Last Modified time: 2019-01-12 10:00:27
 */

package main

import "fmt"

func main() {
	if i := 3; i > 1 {
		fmt.Println("this is older 1")
	} else {
		fmt.Println("this is lower 1")
	}
	switch i := 3; i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
}
