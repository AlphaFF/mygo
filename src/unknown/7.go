package main

import (
	"bytes"
	"fmt"
)

func main() {
	var arr1 [5]int
	var arr2 *[5]int
	var arr3 = new([5]int)
	fmt.Println(arr1, arr2, *arr3)
	var arr4 = [5]string{3: "abc", 4: "edf"}
	fmt.Println(arr4)
	// for i := range arr4 {
	//  fmt.Println(i, arr4[i])
	// }

	var buffer bytes.Buffer
	// buffer.WriteString("hello")
	buffer.Write([]byte("Hello "))
	fmt.Println(buffer)
}
