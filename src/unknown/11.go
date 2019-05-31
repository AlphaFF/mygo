package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	raduis float32
}

type Shaper interface {
	Area() float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func (c *Circle) Area() float32 {
	return c.raduis * c.raduis * math.Pi
}

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool{
    n := data.Len()
    for i:=n-1:i>0;i--{
        if data.Less(i, i-1){
            return false
        }
    }
    return true
}

type IntArray []int

func (i IntArray) Len() int {return len(p)}
func (i IntArray) Less(i, j int) {return p[i] < p[j]}
func (i IntArray) Swap(i, j int) {p[i], p[j] = p[j], p[i]}


func main() {
	var sh Shaper
	s1 := new(Square)
	s1.side = 5

	sh = s1

	if t, ok := sh.(*Square); ok {
		fmt.Println("this is a square.", t, *t)
	}

	s2 := new(Circle)
	s2.raduis = 3

	sh = s2
	if u, ok := sh.(*Circle); ok {
		fmt.Println("this is a circle.", u, *u)
	} else {
		fmt.Println("this is nothing.")
	}
}
