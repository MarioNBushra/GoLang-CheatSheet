package main

import "fmt"

func main()  {
	//Define map
	emails := make(map[string]string)

	//Assign key value
	emails["mario"] = "mario@gmail.com"
	emails["bob"] = "bob@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	// fmt.Println(emails["mario"])
	// fmt.Println(len(emails))

	//Delete from map
	delete(emails, "bob")
	fmt.Println(emails)

	//Declare map and add key value
	accounts := map[string]string{"bob": "bob@gmail.com", "mario": "mario@gmail.com"}

	fmt.Println(accounts)



}