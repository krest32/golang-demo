package entity

import "fmt"

type Person struct {
	name string
	year int
}

var person Person

func ParsePerson() *Person {
	person :=&Person{}
	person.name = "我"
	return person
}

func GetPerson() Person {
	fmt.Printf(person.name)
	return person
}

func SetPerson(name string, year int) Person {
	person = Person{
		name: name,
		year: year,
	}
	return person
}

