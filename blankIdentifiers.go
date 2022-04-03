package main

import "fmt"

func main() {
	stringA := "This is the message to print!"

	//count, err := fmt.Println(stringA)
	//fmt.Println(count, err)

	_, err := fmt.Println(stringA)
	fmt.Println(err)

	// assignment mismatch: 1 variable but fmt.Println returns 2 values
	// obj := fmt.Println(stringA)
	// fmt.Println(obj)
}
