/*
* @Author: AlphaFF
* @Date:   2018-09-20 16:30:19
* @Last Modified by:   AlphaFF
* @Last Modified time: 2018-11-06 17:37:02
 */

// package main

// import (
//  "fmt"
//  "log"
//  "net/http"
//  "strings"
// )

// func sayHelloName(w http.ResponseWriter, r *http.Request) {
//  r.ParseForm()
//  fmt.Println(r.Form)
//  fmt.Println("path", r.URL.Path)
//  fmt.Println("scheme", r.URL.Scheme)
//  // fmt.Println(r.Form("url_long"))
//  for k, v := range r.Form {
//      fmt.Println("key:", k)
//      fmt.Println("val:", strings.Join(v, ""))
//  }
//  fmt.Fprintf(w, "Hello world...")
// }

// func main() {
//  http.HandleFunc("/", sayHelloName)
//  err := http.ListenAndServe(":9090", nil)
//  if err != nil {
//      log.Fatal("ListenAndServer:", err)
//  }
// }

package main

import (
	"fmt"
	"math"
)

// type person struct {
// 	name string
// 	age  int
// }

// func Older(p1, p2 person) (person, int) {
// 	if p1.age > p2.age {
// 		return p1, p1.age - p2.age
// 	}
// 	return p2, p2.age - p1.age
// }

// type Skills []string

// type Human struct {
// 	name   string
// 	age    int
// 	weight int
// }

// type Student struct {
// 	Human
// 	Skills
// 	int
// 	school string
// }

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	// for i := 0; i < 5; i++ {
	//  fmt.Println(i, &i)
	// }
	// var P person
	// P.name = "zhangsan"
	// P.age = 24
	// fmt.Println(P.name, P.age)

	// T := person{"tom", 18}
	// fmt.Println(T.name, T.age)

	// S := person{name: "lisi", age: 10}
	// fmt.Println(S.name)

	// Q := new(person)
	// fmt.Println(Q)

	// bp_older, bp_diff := Older(P, T)
	// fmt.Println(bp_older, bp_diff)

	// mark := Student{Human{"mark", 24, 170}, "peking university"}
	// fmt.Println(mark, mark.name)
	fmt.Println("hello world")
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	// areaIntf = sq1
	// if t, ok := areaIntf.(*Square); ok {
	// 	fmt.Println("t is:", t)
	// }
	// if u, ok := areaIntf.(*Circle); ok {
	// 	fmt.Println("u is:", u)
	// } else {
	// 	fmt.Println("don't contain type circle")
	// }
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Println("square is:", t)
	case *Circle:
		fmt.Println("circle is:", c)
	case nil:
		fmt.Println("nothing")
	default:
		fmt.Println("unexpected type")
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}
