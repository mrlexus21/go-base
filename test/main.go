package main

import "fmt"

func main() {
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	k, ok := m[3]
	fmt.Println(k)
	fmt.Println(ok)
}
