package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//fmt.Println("hello world")
	go sayHello(ch)
	for i := range ch {
		fmt.Println(i)
	}
	//fmt.Println(<-ch)
	//time.Sleep(2 * time.Second)
}

func say(word string) {
	//time.Sleep(1 * time.Second)
	fmt.Println(word)

}

func sayHello(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		say("hello")
		exit <- i
	}
	close(exit)
}
