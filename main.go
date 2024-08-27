package main

import "fmt"

func main() {
	defer fmt.Println(" world")
	fmt.Println("Hello")
	/*s1, s2 := test()
	fmt.Println(s1, s2)*/

	sum := 0
	//for i := 0; i < 10; i++ {
	for sum < 1000 { //while
		sum += 10
	}
	fmt.Println(sum)

	a := 0
	for a < 200 {
		//if a == 100 {
		//	fmt.Println("a is 100")
		//}

		switch i := isTest(a); i {
		case 1:
			fmt.Println("a = 2")
		case 2:
			fmt.Println("a = 3")
		default:
			fmt.Printf("unknown a=%d\n", a)
		}

		/*if i := isTest(a); i == 1 {
			fmt.Println("a = 2")
		} else if i == 2 {
			fmt.Println("a = 3")
		} else {
			fmt.Printf("unknown a=%d\n", a)
		}*/
		a++
	}
}

/*func test() (a, b string) {
	a = "hello"
	b = "world"
	return a, b
}*/

func isTest(a int) int {
	if a == 2 {
		return 1
	} else if a == 3 {
		return 2
	}
	return 3
}
