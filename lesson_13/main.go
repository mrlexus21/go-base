package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//writeToFile()
	//appendFile()
	//c := exec.Command("systeminfo")
	//c := exec.Command("ssh", "-t", "kart", "top")
	c := exec.Command("cat", "test.txt")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	//c.Stderr = os.Stderr
	c.Run()
}

func appendFile() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.WriteString(" Qwerty")
	if err != nil {
		panic(err)
	}
}

func readFile() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func writeToFile() {
	data := []byte("Arthur")
	err := os.WriteFile("test.txt", data, 0777)
	if err != nil {
		return
	}
}
