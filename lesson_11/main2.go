package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type AppError struct {
	Message string
	Err     error
}

func (ae *AppError) Error() string {
	return ae.Message
}

func main() {
	/*defer func() {
		fmt.Println("ok")
	}()*/
	//panic("something goes wrong")
	divide(4, 0)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("after panic")
}

func divide(a, b int) {
	defer func() {
		var appErr *AppError
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appErr) {
					fmt.Println("This is AppError panic!", err)
				} else {
					fmt.Println("This is other error panic")
				}
			default:
				panic("This is default go panic")
			}
			log.Println("panic happened:", err)
		}
	}()
	fmt.Println(div(a, b))
}

func div(a, b int) int {
	if b == 0 {
		//panic(fmt.Errorf("qwerty"))
		panic(&AppError{
			Message: "this divided by zero",
			Err:     nil,
		})
		//panic("123")
	}
	return a / b
}
