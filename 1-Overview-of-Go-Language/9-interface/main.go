package main

import "log"

type Rectangle interface {
	Area() int
	Range() int
}

type myValue struct {
	length int
	width  int
}

func (m myValue) Area() int {
	return m.length * m.width
}

func (m myValue) Range() int {
	return 2 * (m.length + m.width)
}

func main() {
	var r Rectangle
	r = myValue{10, 5}

	log.Println("The area of Rectangle is : ", r.Area())
	log.Println("The Range of Rectangle is : ", r.Range())

}
