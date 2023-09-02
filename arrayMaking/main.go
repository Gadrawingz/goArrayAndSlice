package main

import "fmt"

func main() {
	fmt.Println("Array time!")

	// 1. var arrayName[length] datatype
	var players [11]string

	// 2. var arrayName = [length] datatype{ values }
	var fruits = [5]string{"Mango", "Orange", "Apple"}

	// 3. var arrayName = [...] datatype{ values }
	var oddNumbers = [...]int{1, 5, 7, 77}

	// Empty arrays
	var emptyArray [10]int

	fmt.Println(players)
	fmt.Println(fruits)
	fmt.Println(oddNumbers)
	fmt.Println(emptyArray)

}
