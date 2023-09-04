package main

import "fmt"

// Go Program that uses a function to print a multidimensional array!
// Array has a fixed size at [4][2]
func printMyArray(a [4][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {

	//  It is possible to create multidimensional arrays
	a := [4][2]string{
		{"Lion", "Koala"},
		{"Zebra", "Tiger"},
		{"Panda", "Python"},
		{"Donkey", "Cheetah"},
		/* The compiler will complain if you omit this comma, Because,
		this comma is necessary. */
	}
	printMyArray(a)

	var newArray [4][2]string
	newArray[0][0] = "Coca Cola"
	newArray[0][1] = "Mango"
	newArray[1][0] = "Apple"
	newArray[1][1] = "Avocado"
	newArray[2][0] = "Apple"
	newArray[2][1] = "Lime"
	newArray[3][0] = "Yellow Fish"
	newArray[3][1] = "Brown Dish"
	fmt.Println()
	printMyArray(newArray)

}
