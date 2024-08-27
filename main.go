package main

import (
	"log"
	"net/http"
)

func main() {
	/* var a int8 = 127
	var b int16 = 32767
	var c int32 = 30
	var d int64 = 20
	var e uint8 = 10
	var f uint16
	var g uint32
	var h uint64
	var i byte //uint8
	var j rune //int32
	var k int  // int32 or int64
	var l uint // uint32 or uint64
	var a1 float32 = 1.8
	var b1 float64 = 1.8
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 4 + 90i

	var b23 bool = true

	var name string = "Hello World"
	fmt.Println(name)

	var (
		name1 = "Arthur"
		age1  = 18
	)
	fmt.Println(name1)
	fmt.Println(age1)

	var name2, age2 = "Arthur", 32
	fmt.Println(name2)
	fmt.Println(age2)

	cc := fmt.Sprintf("My name is %s and age %d", name2, age2)
	fmt.Println(cc)*/
	http.HandleFunc("/", Handler)
	log.Println("Start HTTP server on port 8001")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, another world!\n"))

}
