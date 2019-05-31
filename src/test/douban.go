/*
* @Author: wangfeng1
* @Date:   2019-01-11 19:45:16
* @Last Modified by:   wangfeng1
* @Last Modified time: 2019-01-11 20:14:25
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Swoop struct {
	url    string
	header map[string]string
}

func (swoop Swoop) get_html_header() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", swoop.url, nil)
	if err != nil {
		log.Fatalln("new Request err", err)
	}

	for key, value := range swoop.header {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("do client err", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read resp err", err)
	}
	return string(body)
}

func Convert() {
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}
	header["User-Agent"] = GetRandomUserAgent()
	f, err := os.Create("./douban.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("电影名称" + "\t" + "评分" + "\t" + "评价数量" + "\t" + "\r\n")
	for i := 0; i < 3; i++ {
		log.Println("正在爬取第" + strconv.Itoa(i) + "页数据...")
		url := "https://movie.douban.com/top250?start=" + strconv.Itoa(i*25)
		swoop := &Swoop{url, header}
		html := swoop.get_html_header()

		commentCount := `<span>(.*?)评价</span>`
		rp2 := regexp.MustCompile(commentCount)
		txt2 := rp2.FindAllStringSubmatch(html, -1)
		fmt.Println(i, txt2)
	}
}

var userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}

func main() {
	Convert()
}
