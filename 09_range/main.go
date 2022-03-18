package main

import "fmt"

func main()  {
	ids := []int{33, 76, 54, 23, 11, 2}

	//loop through ids

	for i, id := range ids{
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not Using Index
	for _, id := range ids{
		fmt.Printf("ID: %d\n", id)
	}

	//add ids together
	sum := 0

	for _, id := range ids{
		sum += id
	}

	fmt.Println("SUM", sum)

	//Range With Map
	accounts := map[string]string{"bob": "bob@gmail.com", "mario": "mario@gmail.com"}

	for k, v := range accounts{
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range accounts {
		fmt.Println("Name " + k )
	}
}