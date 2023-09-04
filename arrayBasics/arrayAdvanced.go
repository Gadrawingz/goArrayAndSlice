package arrayBasics

import "fmt"

// Similarly when arrays are passed to functions as parameters,
// they are passed by value and the original array is unchanged.

func modifyLocalVariable(myNumber [8]int) {
	myNumber[0] = 34
	fmt.Println("INSIDE FUNCTION: ", myNumber)
}

func main() {

	// Variable Usage!
	newNumber := [...]int{10, 12, 16, 25, 43, 77, 88, 500}
	fmt.Println("B4ORE PASSING TO FUNCTION: ", newNumber)
	modifyLocalVariable(newNumber)
	fmt.Println("A4TER PASSING TO FUNCTION: ", newNumber)

}
