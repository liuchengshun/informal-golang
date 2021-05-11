package main

import "fmt"

type students struct {
	person1 person
	person2 person
}

type person struct {
	name string
	age  int
}

func (s *students) addPerson(p ...person) {
	s.person1 = p
}

func main() {
	s := students{}
	person1 := person{name: "liuchengshun", age: 25}
	person2 := person{name: "liuchengyi", age: 26}

	s.addPerson(person1)
	s.addPerson(person2)
	fmt.Println("students:", s)
}
