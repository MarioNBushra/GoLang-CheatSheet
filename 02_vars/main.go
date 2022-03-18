package main

import "fmt"

func main()  {
	// var name  = "Mario"
	var age  = 37
	const isCool = true

	//ShortHand
	addres := "Sohag"
	size := 1.3

	name, email := "Mario", "mario@gmail.com"

	fmt.Println(name, age, isCool, addres, size, email)
	fmt.Printf("%T\n", size) //get type of name
}