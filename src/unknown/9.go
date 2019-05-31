package main

import (
	"fmt"
	"regexp"
	// "strconv"
	"sync"
)

type info struct {
	mu   sync.Mutex
	name string
}

func Update(info *info) {
	info.mu.Lock()
	info.name = "hahhaa"
	info.mu.Unlock()
}

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

	// 没看懂
	// var mu sync.Mutex
	// info := new(info{mu, "wahaha"})
	// Update(info)
	// fmt.Println(info)
}
