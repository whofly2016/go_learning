package main

func main() {
	var b Animal

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
