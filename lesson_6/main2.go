package main

import "fmt"

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
}

type User interface {
	ChangeName(newFio string)
	ChangeAddress(newAddress string)
}

func (u *user) ChangeName(newFio string) {
	u.FIO = newFio
}

func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress
}

func NewUser(fio, address, phone string) User {
	u := user{
		FIO:     fio,
		Address: address,
		Phone:   phone,
	}
	return &u
}

func main() {
	us := user{"Ivanov Ivan Ivanovich", "Moscow", "123456789"}

	fmt.Println(us)
	us.ChangeName("Petrov Petr Petrovich")
	us.ChangeAddress("Novgorod")
	fmt.Println(us)
}
