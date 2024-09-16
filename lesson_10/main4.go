package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	done := make(chan string)
	for i := 0; i < 30; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("the once only")
			})
			fmt.Println("dd")
			done <- "ff"
		}()
	}
	for i := 0; i < 30; i++ {
		<-done
	}
}
