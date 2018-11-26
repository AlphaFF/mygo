package main

import (
	"fmt"
	"math"
	"reflect"
	// "time"
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

type any interface{}

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Println("bool", v)
		case int:
			fmt.Println("int", v)
		case float32:
			fmt.Println("float32", v)
		case string:
			fmt.Println("string", v)
		case specialString:
			fmt.Println("specialString", v)
		default:
			fmt.Println("unknown type", v)
		}
	}
	testFunc(whatIsThis)
}

type T struct {
	A int
	B string
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
		fmt.Println("circle is:", t)
	case nil:
		fmt.Println("nothing")
	default:
		fmt.Println("unexpected type")
	}

	TypeSwitch()
	fmt.Println(reflect.TypeOf(sq1))
	fmt.Println(reflect.ValueOf(sq1))
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v = v.Elem()
	fmt.Println(v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v)

	t := T{23, "skido"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Println(i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("sunset strip")
	fmt.Println("t is now", t)

	ch := make(chan string, 5)
	go sendData(ch)
	// go getData(ch)
	for {
		fmt.Println(<-ch)
	}

	// time.Sleep(1e9)
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func sendData(ch chan string) {
	ch <- "washington"
	ch <- "Tripoli"
	ch <- "london"
	ch <- "beijing"
	ch <- "tokio"
}

func getData(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
}
