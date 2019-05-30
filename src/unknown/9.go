package main

import (
	"fmt"
	"regexp"
	// "strconv"
)

// 正则替换

func main() {
	searchIn := "123wo45"
	pat := "[0-9]+"
	if ok, v := regexp.MatchString(pat, searchIn); ok {
		fmt.Println("v", v)
	}
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##")
	fmt.Println(str)
}
