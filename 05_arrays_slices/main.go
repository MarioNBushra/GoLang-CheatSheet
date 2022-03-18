package main

import "fmt"

func main()  {
	//Arrays
	 var fruitArr [2]string

	 fruitArr[0]= "Apple"
	 fruitArr[1]= "Orange"

	 //Declare and assign
	 studentsArr := [2]string{"Mario", "Bushra"}
	 fmt.Println(fruitArr)
	 
	 fmt.Println(studentsArr)

	 fruitSlice := []string{"Banana", "Grape", "Orange"}

	 fmt.Println(len(fruitSlice)) //get length of array
	 fmt.Println(fruitSlice[1:2]) //return Grape
	 fmt.Println(fruitSlice[1:3]) //return Grape and Orange

}