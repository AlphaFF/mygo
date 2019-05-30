package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "sync"
)

type Crawl struct {
	url    string
	header map[string]string
}

func (crawl Crawl) getHtml() string {
	resp, err := http.Get(crawl.url)
	if err != nil {
		log.Fatalln("request err:", err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("response err:", err)
	}
	return string(html)

}

func testPanic() {
	f, err := os.Open("text.txt")
	if err != nil {
		panic("文件不存在.")
	}
	fmt.Println(f)
}

func main() {
	header := map[string]string{
		"Host": "baidu.com",
	}
	url := "http://www.baidu.com"
	crawl := &Crawl{url, header}
	fmt.Println(crawl)
	// result := crawl.getHtml()
	// fmt.Println(result)

	// wg := sync.WaitGroup{}
	// wg.Add(100)
	// for i := 0; i < 100; i++ {
	//  go func(i int) {
	//      fmt.Println(`number is `, i)
	//      wg.Done()
	//  }(i)
	// }
	// wg.Wait()
	d := test(4, 5)
	fmt.Println(`d is`, d)
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("error is", e)
		}
	}()
	testPanic()
}

// 测试参数定义返回值
func test(a, b int) (c int) {
	c = a + b
	fmt.Println(c)
	return
}
