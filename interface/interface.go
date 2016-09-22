package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) sayHi() {
	fmt.Printf("Hi I m %s and you can call me on %s\n", h.name, h.phone)
}

func (h Human) sing(lyrics string) {
	fmt.Println("I am singing", lyrics)
}

func (e Employee) sayHi() {
	fmt.Printf("Hi I am %s, I work at %s and you can call me at %s\n", e.name, e.company, e.phone)
}

type Men interface {
	sayHi()
	sing(lyrics string)
}

func (h Human) String() string {
	return "Name:" + h.name + ", Age:" + strconv.Itoa(h.age) + " years, Contact:" + h.phone
}

func main() {
	mike := Student{Human{"Mike", 24, "676-789-7878"}, "ABD School", 0}
	sam := Student{Human{"Sam", 22, "678-7877"}, "Harvard", 100}
	paul := Employee{Human{"Paul", 27, "676-4545"}, "Google", 1000}
	tom := Employee{Human{"Tom", 32, "454-678-5677"}, "Twitter", 300}
	siya := Human{"Siya", 21, "676-434-6564"}

	var i Men
	i = mike
	fmt.Println("This is Mike a Student")
	i.sayHi()
	i.sing("lalala laa")

	i = tom
	fmt.Println("This is an Employee")
	i.sayHi()
	i.sing("Tada tada ta da")

	fmt.Println("Lets have a slice of Men")
	x := make([]Men, 3)
	x[0], x[1], x[2] = sam, paul, mike

	for _, t := range x {
		t.sayHi()
	}
	fmt.Println("This is a Human", siya)
}
