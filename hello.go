package main

import (
	"fmt"
)

// Human 定义一个人
type Human struct {
	name  string
	age   int
	phont string
}

// Student 继承人
type Student struct {
	Human
	school string
	loan   float32
}

// Employee 工作者,也继承Human
type Employee struct {
	Human
	company string
	money   float32
}

// SayHi Human的方法
func (h *Human) SayHi() {
	fmt.Println("human say %s", h.name)
}

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	width, height float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

func main() {
	fmt.Println("hahaha..")
	switch a := 2; {
	case a > 0:
		fmt.Println("a==0")
		fallthrough
	case a > 1:
		fmt.Println("a==1")
	default:
		fmt.Println("None")
	}
	var a int
	for i := 0; i < 10; i++ {
		a++
		fmt.Println(a)
	}
	sq1 := new(Square)
	sq1.side = 5

	var areaInfo Shaper
	areaInfo = sq1
	fmt.Println(areaInfo.Area())

	r := Rectangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r, q}
	for n := range shapes {
		fmt.Println("shapes detail:", shapes[n])
		fmt.Println("shape is :", shapes[n].Area())
	}
}
