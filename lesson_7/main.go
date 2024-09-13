package main

import (
	"fmt"
	"io"
	"strings"
)

type appError struct {
	Err    error
	Custom string
	Field  int
}

func (e *appError) Error() string {
	//err := fmt.Errorf("new error", e.Custom)
	//return err.Error()
	//return e.Err.Error()
	return fmt.Sprintf("%s error", e.Custom)
}

func (e *appError) Unwrap() error {
	return e.Err
}

func main() {
	/*err := method1()
	if err != nil {
		fmt.Println(err.Unwrap())
	}*/
	r := strings.NewReader("hello world")
	arr := make([]byte, 4)
	for {
		n, err := r.Read(arr)
		fmt.Printf("n= %d err=%v arr= %v\n", n, err, arr)
		fmt.Printf("arr n bytes: %q\n", arr[:n])
		if err == io.EOF {
			break
		}
	}
}

func method1() *appError {
	/*return &AppError{
		Err:    fmt.Errorf("my error"),
		Custom: "value here",
		Field:  89,
	}*/
	return method2()
}

func method2() *appError {
	return method3()
}

func method3() *appError {
	//return fmt.Errorf("some error")
	return &appError{
		Err:    fmt.Errorf("internal error"),
		Custom: "something goes wrong",
	}
}
