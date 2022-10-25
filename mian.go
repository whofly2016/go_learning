package main

import "fmt"

type Animal interface {
	Eat()
	Run()
}

type Cat struct {
	Name string
	Sex  bool
}

type Dog struct {
	Name string
}

func (d Dog) Run() {
	fmt.Println(d.Name, "run()")
}

func (d Dog) Eat() {
	fmt.Println(d.Name, "eat()")
}

func (c Cat) Run() {
	fmt.Println(c.Name, "run()")
}

func (c Cat) Eat() {
	fmt.Println(c.Name, "eat()")
}

func main() {
	var a Animal
	c := Cat{
		Name: "tom",
		Sex:  false,
	}

	a = c

	a.Eat()
	a.Run()
	// a.Name

}
