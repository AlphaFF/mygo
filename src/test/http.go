// /*
// * @Author: AlphaFF
// * @Date:   2018-11-26 20:16:27
// * @Last Modified by:   wangfeng1
// * @Last Modified time: 2018-12-25 19:30:42
//  */

// package main

// import (
// 	"fmt"
// 	"math"
// 	"runtime"
// 	"strconv"
// 	// "time"
// 	"strings"
// )

// // type Student struct {
// // 	name string
// // 	age  int
// // }

// const (
// 	B float64 = 1 << (iota * 10)
// 	KB
// 	MB
// 	GB
// )

// // type S struct{}

// type Human struct {
// 	name  string
// 	age   int
// 	phone string
// }

// type Student struct {
// 	Human
// 	school string
// 	loan   float32
// }

// type Employee struct {
// 	Human
// 	company string
// 	money   float32
// }

// func (h *Human) SayHi() {
// 	fmt.Println("name is: %s, and age is:  %s", h.name, h.age)
// }

// func (h *Human) Sing(lyric string) {
// 	fmt.Println("lalala", lyric)
// }

// func (h *Human) Guzzle(beerStein string) {
// 	fmt.Println("Guzzle", beerStein)
// }

// func (e *Employee) SayHi() {
// 	fmt.Println("company is %s, money is %s", e.company, e.money)
// }

// func (s *Student) BorrowMoney(amount float32) {
// 	s.loan += amount
// }

// func (e *Employee) SpendSalary(amount float32) {
// 	e.money -= amount
// }

// type Men interface {
// 	SayHi()
// 	Sing(lyrics string)
// 	Guzzle(beerStein string)
// }

// type YoungChap interface {
// 	SayHi()
// 	Sing(song string)
// 	BorrowMoney(amount float32)
// }

// type ElderlyGent interface {
// 	SayHi()
// 	Sing(song string)
// 	SpendSalary(amount float32)
// }

// type Element interface{}
// type List []Element

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) String() string {
// 	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
// }

// func main() {
// 	// fmt.Println("hello world..")
// 	// jack := Student{"张三", 24}
// 	// fmt.Println(jack)
// 	// fmt.Println(KB)
// 	fmt.Println(math.Pi)
// 	a := 1
// 	a++
// 	var p *int
// 	p = &a
// 	fmt.Println(*p)
// 	// b := 10
// 	// if b := 5; b > 2 {
// 	// 	fmt.Println(b)
// 	// }
// 	// fmt.Println(b)
// 	// for {
// 	// 	b--
// 	// 	if b < 3 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(b)
// 	// }
// 	// for i := 0; i < 3; i++ {
// 	// 	b++
// 	// 	fmt.Println(b)
// 	// }
// 	// fmt.Println("over...")

// 	switch c := 1; {
// 	case c >= 0:
// 		fmt.Println("c==0")
// 		fallthrough
// 	case c >= 1:
// 		fmt.Println("c==1")
// 	default:
// 		fmt.Println("c not exist.")
// 	}
// 	d := make(map[string]int)
// 	fmt.Println(d)
// 	d["1"] = 2
// 	d["2"] = 4
// 	d["3"] = 9
// 	fmt.Println(d)
// 	for k := range d {
// 		fmt.Println(k, d[k])
// 	}
// 	delete(d, "3")
// 	fmt.Println(d)
// 	eyExist, ok := d["4"]
// 	if ok {
// 		fmt.Println("4 is ", eyExist)
// 	} else {
// 		fmt.Println("4 in not in d")
// 	}
// 	// pa := Student{"lisi", 25}
// 	// var pt *Student
// 	// pt = &pa
// 	// fmt.Println(pt.name, pt.age)
// 	list := make(List, 3)
// 	fmt.Println(list)
// 	list[0] = 1
// 	list[1] = "hello"
// 	list[2] = Person{"devi", 70}
// 	for index, element := range list {
// 		// if value, ok := element.(int); ok {
// 		// 	fmt.Println("this element is int", index, value)
// 		// } else if value, ok := element.(string); ok {
// 		// 	fmt.Println("this element is string", index, value)
// 		// } else if value, ok := element.(Person); ok {
// 		// 	fmt.Println("this element is Person", index, value)
// 		// } else {
// 		// 	fmt.Println("this element is not any of int, string, Person")
// 		// }
// 		switch value := element.(type) {
// 		case int:
// 			fmt.Println("this element is int", index, value)
// 		case string:
// 			fmt.Println("this element is string", index, value)
// 		case Person:
// 			fmt.Println("this element is Person", index, value)
// 		default:
// 			fmt.Println("this element is not any of int, string, Person")
// 		}
// 	}
// 	// go say("world")
// 	say("hello")
// 	// ch1 := make(chan int)
// 	// // go sendData(ch1)
// 	// // getData(ch1)
// 	// ch2 := make(chan int)
// 	// go pump1(ch1)
// 	// go pump2(ch2)
// 	// go suck(ch1, ch2)
// 	// time.Sleep(1 * 1e9)
// 	f()
// 	b()
// 	s := "This is a test"
// 	fmt.Println(strings.HasPrefix(s, "TH"))
// 	fmt.Println(strings.HasSuffix(s, "st"))
// 	fmt.Println(strings.Contains(s, "st"))
// 	fmt.Println(strings.Index(s, "s"))
// 	strings.Replace(s, "is", "as", 1)
// 	fmt.Println(strings.Replace(s, "is", "as", -1))
// 	fmt.Println(`this is a raw string \n`)
// }

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		runtime.Gosched()
// 		fmt.Println(s)
// 	}
// }

// func pump(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		ch <- i
// 	}
// }

// func sendData(ch chan string) {
// 	ch <- "w"
// 	ch <- "t"
// 	ch <- "l"
// 	ch <- "b"
// 	ch <- "t"
// 	close(ch)
// }

// func getData(ch chan string) {
// 	for {
// 		input, open := <-ch
// 		if !open {
// 			break
// 		}
// 		fmt.Println("%s", input)
// 	}
// }

// func pump1(ch chan int) {
// 	for i := 0; ; i++ {
// 		ch <- i * 2
// 	}
// }

// func pump2(ch chan int) {
// 	for i := 0; ; i++ {
// 		ch <- i + 5
// 	}
// }

// func suck(ch1, ch2 chan int) {
// 	for {
// 		select {
// 		case v := <-ch1:
// 			fmt.Println("receive from channel 1", v)
// 		case v := <-ch2:
// 			fmt.Println("receive from channel 2", v)
// 		}
// 	}
// }

// func f() {
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println("defer:", i)
// 	}
// }

// func trace(s string)   { fmt.Println("entering:", s) }
// func untrace(s string) { fmt.Println("leaving:", s) }

// func a() {
// 	trace("a")
// 	defer untrace("a")
// 	fmt.Println("in a")
// }

// func b() {
// 	trace("b")
// 	defer untrace("b")
// 	fmt.Println("in b")
// 	a()
// }
