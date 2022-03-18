package main

import "fmt"

func main()  {
	x := 5
	y := 10

	//IF Else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y) 
	}else{
		fmt.Printf("%d is less than %d\n", y, x) 
	}

	//else if
	color := "green"

	if color == "red" {
		fmt.Println("Colour is red")
	}else if color == "blue"{
		fmt.Println("Colour is blue")
	}else {
		fmt.Println("Colour is not blue or red")
	}

	//Switch
	switch color {
	case "red":
		fmt.Println("Colour is red")
	case "blue":
		fmt.Println("Colour is blue")
	default:
		fmt.Println("Colour is not red or blue")
	}
}