package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 10
var y int = 20

const pi = 3.14

func Calc(a, b *int) {
	*a = *a * *b
	*b = *a + *b
}

type User struct {
	Username string
	Email    string
	Password string
}

type Model interface {
	getData() (string, string, string)
	setData(a, b, c string)
}

func newUser(username, email, password string) *User {
	return &User{Username: username, Email: email, Password: password}
}

func (u User) printUsername() {
	fmt.Println(u.Username)
}

func (u User) getEmail() string {
	return u.Email
}

func (u User) getData() (string, string, string) {
	return u.Username, u.Email, u.Password
}

func (u *User) setData(username, email, password string) {
	u.Username = username
	u.Email = email
	u.Password = password
}

func printModel(mod Model) {
	a, b, c := mod.getData()
	fmt.Println(a, b, c)
}

func deferDemo() {
	fmt.Println("GABU")

	defer fmt.Println("aku 1")
	defer fmt.Println("aku 2")
	defer fmt.Println("aku 3")

	fmt.Println("BUGA")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker Starting")
	time.Sleep(time.Duration(id) * time.Millisecond * 1000)
	fmt.Printf("Worker %d Done", id)
}

func Tutorial() {
	// if x > 10 {
	// 	fmt.Println("it is bigger than 10")
	// } else if x <= 10 && y > 10 {
	// 	fmt.Println("It is smaller than 10 and y is bigger than 10")
	// } else {
	// 	fmt.Println("WOW")
	// }

	// if result := x * y * 2; result > 10 {
	// 	fmt.Println("WOW")
	// }

	// new := []string{"appel", "banana", "cherry"}
	// for index, fruit := range new {
	// 	fmt.Println(index, "#", fruit)
	// }

	// for i := 0; i < len(new); i++ {
	// 	fmt.Println(new[i])
	// }

	// value := 100
	// ptr := &value

	// *ptr = 10

	// fmt.Println(value)

	// orang := newUser("mama", "mama@mail.com", "1234")
	// printModel(orang)
	var wg sync.WaitGroup
	for i:=0;i<10;i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	
}
