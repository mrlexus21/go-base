package main

import "fmt"

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type InterfaceHere interface {
	Sum() int
}

func main() {
	//var a InterfaceHere
	//sh := structHere{1, 2}
	//os := otherStruct{2, 3}

	//a = &sh
	//fmt.Println(a.Sum())
	//a = os
	//fmt.Println(a.Sum())
	//var os *structHere
	//var i InterfaceHere = otherStruct{3, 2}
	//var i InterfaceHere
	//i = os
	//i.Sum()
	//fmt.Printf("%v, %T", i, i)
	var a interface{} = "jelly"
	//fmt.Printf("%v %T\n", a, a)
	//a = 42
	//fmt.Printf("%v %T", a, a)
	//s := a.(string)

	//s, ok := a.(string)
	//fmt.Println(s)
	//fmt.Println(ok)
	//g := a.(float64)
	switch a.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	default:
		fmt.Printf("unknown type %T", a)
	}
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}
