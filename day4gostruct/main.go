package day4gostruct

import "fmt"

type Person struct { 
	Name string
	Age int
	Address string
}

func main() { 
	var person Person
	person.Name = "Mike"
	person.Age = 25
	person.Address = "New York"

	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.Address)
}
