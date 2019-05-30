package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	uid        int
	username   string
	department string
	created    string
}

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return 3.14 * c.radius * c.radius
}

type Simpler interface {
	Get()
	Set()
}

type Simple struct {
	value int
}

func (s *Simple) Get() int {
	return s.value
}

func (s *Simple) Set(b int) {
	s.value = b
}

func test(ser Simpler) {
	fmt.Println(ser)
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	fmt.Println(db)
	checkErr(err)

	stmt, _ := db.Prepare("insert into userinfo (username, department, created) values (?, ?, ?)")
	defer stmt.Close()
	stmt.Exec("test", "test", "2019-01-02")

	// stmt, _ = db.Prepare("delete from userinfo where uid = ?")
	// stmt.Exec(2)

	stmt, _ = db.Prepare("update userinfo set username = ? where uid = ?")
	stmt.Exec("hahahah", 4)

	rows, _ := db.Query("select * from userinfo;")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err := rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1
	fmt.Println("first:", areaIntf.Area())
	sq2 := new(Circle)
	sq2.radius = 2
	areaIntf = sq2
	fmt.Println("second:", areaIntf.Area())

	r := Rectangle{5, 3}
	q := &Square{5}

	shapes := []Shaper{r, q}
	for n := range shapes {
		fmt.Println(shapes[n])
		fmt.Println(shapes[n].Area())
	}

	var ser Simpler = &Simple{1000}
	// test(ser)
	fmt.Println(Simpler)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
