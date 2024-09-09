package main

import "fmt"

func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Println(s)
	return result
}

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

type Point struct {
	X, Y int
}

func (p Point) movePoint(x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y
	fmt.Println(*p)
}

func main() {
	//callbacks()
	//orderPrice := totalPrice(1)
	//fmt.Println(orderPrice(3))

	p := Point{1, 1}
	//fmt.Println(p.movePoint(1, 1))
	//p.movePointPtr(1, 1)

	fmt.Println(p)
	v := &p
	fmt.Println(v)
	v.movePoint(1, 1)
	fmt.Println(v)
	v.movePointPtr(1, 1)
	fmt.Println(v)
	fmt.Println(p)
}

func callbacks() {
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}
	result := doSomething(sumCallback, "plus")
	fmt.Println(result)

	multipleCallback := func(n1, n2 int) int {
		return n1 * n2
	}
	result = doSomething(multipleCallback, "multiple")
	fmt.Println(result)
}
