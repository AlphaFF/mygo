/*
* @Author: AlphaFF
* @Date:   2018-09-20 16:30:19
* @Last Modified by:   AlphaFF
* @Last Modified time: 2018-10-23 18:06:32
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
)

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Skills
	int
	school string
}

func main() {
	// for i := 0; i < 5; i++ {
	//  fmt.Println(i, &i)
	// }
	var P person
	P.name = "zhangsan"
	P.age = 24
	fmt.Println(P.name, P.age)

	T := person{"tom", 18}
	fmt.Println(T.name, T.age)

	S := person{name: "lisi", age: 10}
	fmt.Println(S.name)

	Q := new(person)
	fmt.Println(Q)

	bp_older, bp_diff := Older(P, T)
	fmt.Println(bp_older, bp_diff)

	mark := Student{Human{"mark", 24, 170}, "peking university"}
	fmt.Println(mark, mark.name)
}
