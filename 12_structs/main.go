package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int
	firstName, lastName, city, gender string
	age int

}

//Greeting method (value reciever)

func (p Person) greet() string{
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer reciever)

func (p *Person) hasBirthday(){
	p.age++
}

// getMarried (pointer reciever)

func(p *Person) getMarried(spouseLastName string){
	if p.gender == "m"{
		return
	}else {
		p.lastName = spouseLastName
	}
}

func main()  {
	// init person using struct
	// person1 := Person{firstName: "Mario", lastName: "Bushra", city: "Sohag", gender: "M", age: 23}
	person2 := Person{"Mario", "Bushra", "Sohag","m",  23}

	// fmt.Println(person2)

	// fmt.Println(person2.firstName)
	person2.getMarried("Williams")
	person2.hasBirthday()
	fmt.Println(person2.greet())
}