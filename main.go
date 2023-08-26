package main

import (
	"fmt"
	"helloworld/users"
)

func main() {
	fmt.Println("Welcome to Golang Rest API")
	users.UserHandler()

	// var Firstname string = "Mahendra"
	// age := 29
	// fmt.Printf("Firstname type is %T \n", Firstname)
	// fmt.Printf("Age type is %T\n", age)
	// fmt.Printf("My Firstname is %v and age is %v", Firstname, age)

	// fruitArray := [3]string{}
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Mango"
	// fruitArray[2] = "Banana"
	// fmt.Println("\nFruit array is :", fruitArray)

	// for key, value := range fruitArray {
	// 	fmt.Printf("\nKey is %v, Value is %v", key, value)
	// }
}
