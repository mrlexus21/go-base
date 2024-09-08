package main

import "fmt"

type Point struct {
	X int
	Y int
	S string
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func main() {
	//pointers()
	//structs()
	//structMethod()
	//arrays()
	slices()
}

func slices() {
	/*letters := []string{"a", "b", "c"}
	fmt.Println(letters)
	letters = append(letters, "d")
	letters = append(letters, "e", "f")
	fmt.Println(letters)*/

	/*createSlice := make([]string, 3)
	createSlice[0] = "1"
	createSlice[1] = "2"
	createSlice[2] = "3"
	createSlice = append(createSlice, "4")
	fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	createSlice = append(createSlice, "5")
	fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	createSlice = append(createSlice, "6")
	fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	createSlice = append(createSlice, "7")
	fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	fmt.Println(createSlice)*/

	/*animalsArr := [4]string{
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}

	var a []string = animalsArr[0:2]
	//fmt.Println(a)

	var b []string = animalsArr[1:3]
	//fmt.Println(b)

	b[0] = "123"

	fmt.Println(animalsArr)
	fmt.Println(a)
	fmt.Println(b)*/

	/*s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := s[:6]
	fmt.Println(t)
	e := s[6:]
	fmt.Println(e)*/

	var snil []int
	fmt.Println(snil, len(snil), cap(snil))
	if snil == nil {
		fmt.Println("slice is nil")
	}
}

func arrays() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)

	numbers := [...]int{1, 2, 3}
	fmt.Println(numbers)
}

func structMethod() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "12345",
	}
	p2 := &p1
	p2.method()
}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "AAA",
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	p1.X = 123
	fmt.Println(p1)

	p2 := Point{X: 123}
	fmt.Println(p2)

	g := &p1
	fmt.Println(g)
	fmt.Println(*g)
	fmt.Println(&g)
}

func pointers() {
	a := "Hello world"
	fmt.Println(a)
	p := &a
	fmt.Println(*p)
	*p = "oh my"
	fmt.Println(a)

	b := 42
	fmt.Println(b)
	g := &b
	*g = *g / 2
	fmt.Println(*g)
}
